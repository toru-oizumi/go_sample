package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type RoomPresenter struct{}

func NewRoomPresenter() *RoomPresenter {
	return &RoomPresenter{}
}

func (p *RoomPresenter) BuildFindByIDResponse(object *model.Room) (*output.FindRoomByIDResponse, error) {
	return &output.FindRoomByIDResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
func (p *RoomPresenter) BuildFindAllResponse(objects model.Rooms) (output.FindAllRoomsResponse, error) {
	var result output.FindAllRoomsResponse
	for _, object := range objects {
		result = append(
			result,
			model.Room{
				ID:            object.ID,
				Name:          object.Name,
				OwnerUserID:   object.OwnerUserID,
				VisitorUserID: object.VisitorUserID,
				CreatedAt:     object.CreatedAt,
				UpdatedAt:     object.UpdatedAt,
			},
		)
	}

	return result, nil
}
func (p *RoomPresenter) BuildCreateResponse(object *model.Room) (*output.CreateRoomResponse, error) {
	return &output.CreateRoomResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
func (p *RoomPresenter) BuildUpdateResponse(object *model.Room) (*output.UpdateRoomResponse, error) {
	return &output.UpdateRoomResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
