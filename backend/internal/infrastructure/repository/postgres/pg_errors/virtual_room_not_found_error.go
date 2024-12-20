package pg_errors

type VirtualRoomNotFoundError struct {
	Name string
}

func NewVirtualRoomNotFoundError(name string) *VirtualRoomNotFoundError {
	return &VirtualRoomNotFoundError{Name: name}
}

func (e *VirtualRoomNotFoundError) Error() string { return "Virtual room not found " + e.Name }
