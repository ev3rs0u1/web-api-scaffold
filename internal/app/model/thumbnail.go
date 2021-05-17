package model

type Thumbnail struct {
	ID   uint32 `gorm:"primaryKey"                                                   json:"-"`
	FID  uint32 `gorm:"column:fid;  type:int(10)    unsigned; not null; uniqueIndex" json:"fid"`
	Size uint64 `gorm:"column:size; type:bigint(20) unsigned; not null"              json:"size"`
}

func (Thumbnail) TableName() string {
	return "thumbs"
}
