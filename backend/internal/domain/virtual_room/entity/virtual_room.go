package entity

type VirtualRoom struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	Capacity int    `json:"capacity"`
}

func NewVirtualRoom(roomName string, capacity int) *VirtualRoom {
	return &VirtualRoom{
		RoomName: roomName,
		Capacity: capacity,
	}
}
