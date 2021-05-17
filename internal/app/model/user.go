package model

type User struct {
	ID     uint32 `gorm:"primaryKey"                                            json:"-"`
	Token  string `gorm:"column:token;  type:char(32);   not null; uniqueIndex" json:"token"`
	Owner  uint8  `gorm:"column:owner;  type:tinyint(1); not null"              json:"owner"`
	Active uint8  `gorm:"column:active; type:tinyint(1); not null"              json:"active"`
	BaseTime
}
