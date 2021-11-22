package model

type AgentReport struct {
	ID              uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	AsOfDate        string `gorm:"type:character(12)" json:"name"`
	CaseSubmit      uint16 `gorm:"type:smallint" json:"caseSubmit"`
	CaseApprove     uint16 `gorm:"type:smallint" json:"caseApprove"`
	FYPSubmit       uint16 `gorm:"type:smallint" json:"fypSubmit"`
	FYPApprove      uint16 `gorm:"type:smallint" json:"fypApprove"`
	FYCLifeSubmit   uint16 `gorm:"type:smallint" json:"fycLifeSubmit"`
	FYCLifeApprove  uint16 `gorm:"type:smallint" json:"fycLifeApprove"`
	FYCOtherSubmit  uint16 `gorm:"type:smallint" json:"fycOtherSubmit"`
	FYCOtherApprove uint16 `gorm:"type:smallint" json:"fycOtherApprove"`
}

// TableName Rename
func (AgentReport) TableName() string {
	return "AgentReport"
}
