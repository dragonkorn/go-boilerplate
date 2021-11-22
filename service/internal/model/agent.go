package model

type Agent struct {
	ID        uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `gorm:"type:character(256)" json:"firstName"`
	LastName  string `gorm:"type:character(256)" json:"lastName"`
}

// TableName Rename
func (Agent) TableName() string {
	return "Agent"
}
