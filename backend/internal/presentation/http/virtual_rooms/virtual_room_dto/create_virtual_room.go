package virtual_room_dto

type CreateVirtualRoom struct {
	RoomName string `json:"room_name"`
	Capacity int    `json:"capacity"`
}
