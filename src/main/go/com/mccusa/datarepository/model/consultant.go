package model

// Consultant maps to the 'consultants' table
type Consultant struct {
	ID       uint           `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	Area     *string        `gorm:"column:AREA" json:"area,omitempty"`
	UserName *string        `gorm:"column:USER_NAME" json:"userName,omitempty"`
	Mail     *string        `gorm:"column:MAIL" json:"mail,omitempty"`
	Role     []JsonMetadata `gorm:"column:ROLE;type:json" json:"role,omitempty"`
	Password *string        `gorm:"column:PASSWORD;size:100" json:"password,omitempty"`
	Access   []JsonMetadata `gorm:"column:ACCESS;type:json" json:"access,omitempty"`
}
