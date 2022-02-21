package cognito

import (
	"errors"
	"go_sample/app/domain/model"
	"go_sample/app/infrastructure/config"
	util_error "go_sample/app/utility/error"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

const CUSTOM_USER_ACTIVATION_DATE = "custom:activation_date"

type CognitoSession *string

type AccountRepository struct {
	client   *cognitoidentityprovider.CognitoIdentityProvider
	poolID   string
	clientID string
}

func NewAccountRepository(config config.Config) *AccountRepository {
	mySession := session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(config.AwsDefaultRegion),
		}))

	return &AccountRepository{
		client:   cognitoidentityprovider.New(mySession),
		poolID:   config.AwsCognitoPoolID,
		clientID: config.AwsCognitoClientID,
	}
}

func (repo *AccountRepository) FindByID(id model.UserID) (*model.Account, error) {
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: &repo.poolID,
		Filter:     aws.String("sub=" + string(id)),
	}

	output, err := repo.client.ListUsers(input)
	if err != nil {
		return nil, err
	}

	// subはIDなので、一意が保証されている
	user := output.Users[0]
	userID, err := repo.getUserIDFromAttributes(user.Attributes)
	if err != nil {
		return nil, err
	}

	account := model.Account{
		ID:        model.UserID(*userID),
		Email:     model.Email(*user.Username),
		CreatedAt: *user.UserCreateDate,
		UpdatedAt: *user.UserLastModifiedDate,
		Enabled:   *user.Enabled,
	}
	return &account, nil
}

func (repo *AccountRepository) FindByEmail(email model.Email) (*model.Account, error) {
	input := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: &repo.poolID,
		Username:   aws.String(string(email)),
	}

	output, err := repo.client.AdminGetUser(input)
	if err != nil {
		return nil, err
	}

	id, err := repo.getUserIDFromAttributes(output.UserAttributes)
	if err != nil {
		return nil, err
	}

	account := model.Account{
		ID:        model.UserID(*id),
		Email:     model.Email(*output.Username),
		CreatedAt: *output.UserCreateDate,
		UpdatedAt: *output.UserLastModifiedDate,
		Enabled:   *output.Enabled,
	}
	return &account, nil
}

func (repo *AccountRepository) Store(object model.Account) (*model.UserID, error) {
	var attrs []*cognitoidentityprovider.AttributeType
	attrs = append(attrs, repo.createUserAttribute(cognitoidentityprovider.UsernameAttributeTypeEmail, string(object.Email)))

	input := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:     &repo.poolID,
		Username:       aws.String(string(object.Email)),
		UserAttributes: attrs,
	}

	_, err := repo.client.AdminCreateUser(input)
	if err != nil {
		switch err.(type) {
		case *cognitoidentityprovider.UsernameExistsException:
			return nil, util_error.NewErrEmailAlreadyExists(err.Error())
		default:
			return nil, err
		}
	}

	if account, err := repo.FindByEmail(object.Email); err != nil {
		return nil, err
	} else {
		return &account.ID, nil
	}
}

func (repo *AccountRepository) Authenticate(email model.Email, password model.Password) error {
	if _, err := repo.authenticate(email, password); err != nil {
		return err
	}
	if err := repo.validateUserStatus(email); err != nil {
		return err
	}
	return nil
}

func (repo *AccountRepository) authenticate(email model.Email, password model.Password) (CognitoSession, error) {
	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		UserPoolId: &repo.poolID,
		ClientId:   &repo.clientID,
		AuthFlow:   aws.String("ADMIN_NO_SRP_AUTH"),
		AuthParameters: aws.StringMap(
			map[string]string{
				"USERNAME": string(email),
				"PASSWORD": string(password),
			},
		),
	}
	output, err := repo.client.AdminInitiateAuth(input)

	if err != nil {
		switch err.(type) {
		case *cognitoidentityprovider.NotAuthorizedException:
			return nil, util_error.NewErrAuthenticationFailed()
		default:
			return nil, err
		}
	} else {
		return output.Session, nil
	}
}

func (repo *AccountRepository) validateUserStatus(email model.Email) error {
	user, err := repo.getCognitoUserByEmail(email)
	if err != nil {
		return err
	}

	if !*user.Enabled {
		return util_error.NewErrEntityNotExists("enabled cognito user")
	}

	if *user.UserStatus == "FORCE_CHANGE_PASSWORD" {
		if activationDate, _ := repo.getActivationDateFromAttributes(user.Attributes); activationDate != nil {
			// 強制的にパスワード変更が求められる場合
			return util_error.NewErrChangePasswordRequired()
		} else {
			// Activationが未実施の場合
			return util_error.NewErrActivationRequired()
		}
	}
	return nil
}

func (repo *AccountRepository) Activate(email model.Email, currentPassword model.Password, newPassword model.Password) error {
	session, err := repo.authenticate(email, currentPassword)
	if err != nil {
		return err
	}

	err = repo.validateUserStatus(email)
	if err == nil {
		// UserStatusに問題が無い場合はActivateの必要がないのでエラーとする
		return util_error.NewErrActivationNotRequired()
	}

	if errors.As(err, &util_error.ErrActivationRequired{}) {
		if err = repo.changePassword(session, email, newPassword); err != nil {
			return err
		}
		if err = repo.setActivationDate(email); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func (repo *AccountRepository) setActivationDate(email model.Email) error {
	var attrs []*cognitoidentityprovider.AttributeType
	attr := repo.createUserAttribute(CUSTOM_USER_ACTIVATION_DATE, time.Now().Format("2006-01-02T15:04:05+09:00"))
	attrs = append(attrs, attr)

	input := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserPoolId:     &repo.poolID,
		Username:       aws.String(string(email)),
		UserAttributes: attrs,
	}

	if _, err := repo.client.AdminUpdateUserAttributes(input); err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *AccountRepository) ChangePassword(email model.Email, currentPassword model.Password, newPassword model.Password) error {
	session, err := repo.authenticate(email, currentPassword)
	if err != nil {
		return err
	}

	err = repo.validateUserStatus(email)
	if err == nil || errors.As(err, &util_error.ErrChangePasswordRequired{}) {
		if err = repo.changePassword(session, email, newPassword); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func (repo *AccountRepository) changePassword(session CognitoSession, email model.Email, newPassword model.Password) error {
	input := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		UserPoolId:    &repo.poolID,
		ClientId:      &repo.clientID,
		ChallengeName: aws.String("NEW_PASSWORD_REQUIRED"),
		ChallengeResponses: aws.StringMap(
			map[string]string{
				"USERNAME":     string(email),
				"NEW_PASSWORD": string(newPassword),
			},
		),
		Session: session,
	}

	_, err := repo.client.AdminRespondToAuthChallenge(input)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *AccountRepository) getCognitoUserByEmail(email model.Email) (*cognitoidentityprovider.UserType, error) {
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: &repo.poolID,
		Filter:     aws.String("email='" + string(email) + "'"),
	}

	output, err := repo.client.ListUsers(input)
	if err != nil {
		return nil, err
	}

	users := output.Users
	if len(users) > 0 {
		// Cognito User Pool で email によるログインを有効にしていれば、emailはユニークであることが保証される。
		return users[0], nil

	} else {
		return nil, util_error.NewErrEntityNotExists("cognito user")
	}
}

func (repo *AccountRepository) Update(object model.Account) (*model.UserID, error) {
	userID := model.UserID("xxxx")
	return &userID, nil
}

func (repo *AccountRepository) Enable(id model.UserID) error {
	account, err := repo.FindByID(id)
	if err != nil {
		return err
	}

	input := &cognitoidentityprovider.AdminEnableUserInput{
		UserPoolId: &repo.poolID,
		Username:   aws.String(string(account.Email)),
	}

	if _, err = repo.client.AdminEnableUser(input); err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *AccountRepository) Disable(id model.UserID) error {
	account, err := repo.FindByID(id)
	if err != nil {
		return err
	}

	input := &cognitoidentityprovider.AdminDisableUserInput{
		UserPoolId: &repo.poolID,
		Username:   aws.String(string(account.Email)),
	}

	if _, err = repo.client.AdminDisableUser(input); err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *AccountRepository) Delete(id model.UserID) error {
	account, err := repo.FindByID(id)
	if err != nil {
		return err
	}

	input := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: &repo.poolID,
		Username:   aws.String(string(account.Email)),
	}

	if _, err = repo.client.AdminDeleteUser(input); err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *AccountRepository) createUserAttribute(name string, value string) *cognitoidentityprovider.AttributeType {
	return &cognitoidentityprovider.AttributeType{
		Name:  aws.String(name),
		Value: aws.String(value),
	}
}

func (repo *AccountRepository) getActivationDateFromAttributes(attributes []*cognitoidentityprovider.AttributeType) (*string, error) {
	for _, v := range attributes {
		if *v.Name == CUSTOM_USER_ACTIVATION_DATE {
			return v.Value, nil
		}
	}
	return nil, util_error.NewErrUnexpected("'" + CUSTOM_USER_ACTIVATION_DATE + "' does not exist in cognito user attributes")
}

func (repo *AccountRepository) getUserIDFromAttributes(attributes []*cognitoidentityprovider.AttributeType) (*string, error) {
	// UserAttributesのsubの値をUserIDとして取得する
	for _, v := range attributes {
		if *v.Name == "sub" {
			return v.Value, nil
		}
	}
	return nil, util_error.NewErrUnexpected("'sub' does not exist in cognito user attributes")
}
