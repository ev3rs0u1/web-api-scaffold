package model

type Node struct {
	PID   uint32 `gorm:"column:pid;   type:int(10) unsigned; not null; primaryKey" json:"pid"`
	CID   uint32 `gorm:"column:cid;   type:int(10) unsigned; not null; primaryKey" json:"cid"`
	Depth uint32 `gorm:"column:depth; type:int(10) unsigned; not null; index"      json:"depth"`
}
