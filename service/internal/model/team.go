package model

type Team struct {
	ID   uint16 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:character(256)" json:"name"`
}

// TableName Rename
func (Team) TableName() string {
	return "Team"
}
