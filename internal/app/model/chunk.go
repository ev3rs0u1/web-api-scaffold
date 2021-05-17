package model

type Chunk struct {
	ID    uint32 `gorm:"primaryKey"                                                   json:"-"`
	Hash  string `gorm:"column:hash; type:char(64);            not null; uniqueIndex" json:"hash"`
	Size  uint64 `gorm:"column:size; type:bigint(20) unsigned; not null"              json:"size"`
	Index uint32 `gorm:"-"                                                            json:"index,omitempty"`
	BaseTime
}
