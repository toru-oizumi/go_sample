package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type RoomPresenter struct{}

func NewRoomPresenter() *RoomPresenter {
	return &RoomPresenter{}
}

func (p *RoomPresenter) BuildFindByIdResponse(object *model.Room) (*output.FindRoomByIdResponse, error) {
	return &output.FindRoomByIdResponse{
		Id:            object.Id,
		Name:          object.Name,
		OwnerUserId:   object.OwnerUserId,
		VisitorUserId: object.VisitorUserId,
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
				Id:            object.Id,
				Name:          object.Name,
				OwnerUserId:   object.OwnerUserId,
				VisitorUserId: object.VisitorUserId,
				CreatedAt:     object.CreatedAt,
				UpdatedAt:     object.UpdatedAt,
			},
		)
	}

	return result, nil
}
func (p *RoomPresenter) BuildCreateResponse(object *model.Room) (*output.CreateRoomResponse, error) {
	return &output.CreateRoomResponse{
		Id:            object.Id,
		Name:          object.Name,
		OwnerUserId:   object.OwnerUserId,
		VisitorUserId: object.VisitorUserId,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
func (p *RoomPresenter) BuildUpdateResponse(object *model.Room) (*output.UpdateRoomResponse, error) {
	return &output.UpdateRoomResponse{
		Id:            object.Id,
		Name:          object.Name,
		OwnerUserId:   object.OwnerUserId,
		VisitorUserId: object.VisitorUserId,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
