package entity

type Equipment struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

func NewEquipment(name string, userID int) *Equipment {
	return &Equipment{
		Name:   name,
		UserID: userID,
	}
}
