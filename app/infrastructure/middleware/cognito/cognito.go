package cognito

import (
	"go_sample/app/domain/model"
	"go_sample/app/infrastructure/config"
	util_error "go_sample/app/utility/error"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

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
	userID, err := repo.getUserIDFromSub(user.Attributes)
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

	id, err := repo.getUserIDFromSub(output.UserAttributes)
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
	emailAttr := repo.createUserAttribute(cognitoidentityprovider.UsernameAttributeTypeEmail, string(object.Email))
	attrs = append(attrs, emailAttr)

	input := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:     &repo.poolID,
		Username:       aws.String(string(object.Email)),
		UserAttributes: attrs,
	}

	_, err := repo.client.AdminCreateUser(input)
	if err != nil {
		return nil, err
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
	} else {
		return nil
	}
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
		return nil, err
	} else {
		return output.Session, nil
	}
}

func (repo *AccountRepository) Activate(email model.Email, currentPassword model.Password, newPassword model.Password) error {
	session, err := repo.authenticate(email, currentPassword)
	if err != nil {
		return err
	}

	if err = repo.activate(session, email, newPassword); err != nil {
		return err
	} else {
		return nil
	}
}

func (repo *AccountRepository) activate(session CognitoSession, email model.Email, newPassword model.Password) error {
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
		return err
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

func (repo *AccountRepository) createUserAttribute(name, value string) *cognitoidentityprovider.AttributeType {
	return &cognitoidentityprovider.AttributeType{
		Name:  aws.String(name),
		Value: aws.String(value),
	}
}

func (repo *AccountRepository) getUserIDFromSub(attributes []*cognitoidentityprovider.AttributeType) (*string, error) {
	// UserAttributesのsubの値をUserIDとして取得する
	for _, v := range attributes {
		if *v.Name == "sub" {
			return v.Value, nil
		}
	}
	// TODO: エラーの返し方
	return nil, util_error.NewErrBadRequest("")
}
