package model

type FileType uint8

const (
	FileTypeFile   FileType = 1
	FileTypeFolder FileType = 2
)

type File struct {
	ID        uint32       `gorm:"primaryKey"                                                       json:"-"`
	UID       uint32       `gorm:"column:uid;      type:int(10) unsigned;    not null; index"       json:"-"`
	PID       uint32       `gorm:"column:pid;      type:int(10) unsigned;    not null; index"       json:"pid"`
	Hash      string       `gorm:"column:hash;     type:char(32);            not null; uniqueIndex" json:"hash"`
	Name      string       `gorm:"column:name;     type:varchar(255);        not null; index"       json:"name"`
	Size      uint64       `gorm:"column:size;     type:bigint(20) unsigned; not null"              json:"size"`
	Ext       FileExt      `gorm:"column:ext;      type:varchar(8);          not null"              json:"ext"`
	Type      FileType     `gorm:"column:type;     type:tinyint(1);          not null"              json:"type"`
	Category  FileCategory `gorm:"column:category; type:tinyint(1);          not null"              json:"category"`
	Status    uint8        `gorm:"column:status;   type:tinyint(1);          not null"              json:"status"`
	Objects   []*Object    `gorm:"foreignKey:FID"                                                   json:"-"`
	Thumbnail *Thumbnail   `gorm:"foreignKey:FID"                                                   json:"-"`
	BaseTime
}
