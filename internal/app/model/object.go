package model

type Object struct {
	ID       uint32 `gorm:"primaryKey"                                              json:"-"`
	FID      uint32 `gorm:"column:fid;      type:int(10) unsigned; not null; index" json:"fid"`
	CID      uint32 `gorm:"column:cid;      type:int(10) unsigned; not null; index" json:"cid"`
	Position uint32 `gorm:"column:position; type:int(10) unsigned; not null"        json:"index"`
	Chunk    *Chunk `gorm:"foreignKey:CID"                                          json:"chunk,omitempty"`
	BaseTime
}
