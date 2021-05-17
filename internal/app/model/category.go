package model

import "web-api-scaffold/internal/pkg/fileutil"

type FileCategory uint8

const (
	NoneCategory  FileCategory = 0
	ImageCategory FileCategory = 1
	VideoCategory FileCategory = 2
	AudioCategory FileCategory = 3
	DocCategory   FileCategory = 4
	OtherCategory FileCategory = 5
)

var (
	ExtCategoryMap = map[FileExt]FileCategory{
		"":           NoneCategory,
		ImageExtJPG:  ImageCategory,
		ImageExtPEG:  ImageCategory,
		ImageExtPNG:  ImageCategory,
		ImageExtBMP:  ImageCategory,
		ImageExtGIF:  ImageCategory,
		ImageExtTIF:  ImageCategory,
		ImageExtICO:  ImageCategory,
		ImageExtICON: ImageCategory,
		ImageExtWebP: ImageCategory,
		ImageExtHEIC: ImageCategory,
		ImageExtHEIF: ImageCategory,

		VideoExtMP4:  VideoCategory,
		VideoExtM4V:  VideoCategory,
		VideoExtMKV:  VideoCategory,
		VideoExtAVI:  VideoCategory,
		VideoExtWMV:  VideoCategory,
		VideoExtFLV:  VideoCategory,
		VideoExtMOV:  VideoCategory,
		VideoExtMPG:  VideoCategory,
		VideoExtMPEG: VideoCategory,
		VideoExt3GP:  VideoCategory,
		VideoExtWebM: VideoCategory,
		VideoExtSWF:  VideoCategory,
		VideoExtF4V:  VideoCategory,

		AudioExtMP3:  AudioCategory,
		AudioExtM4A:  AudioCategory,
		AudioExtMID:  AudioCategory,
		AudioExtMIDI: AudioCategory,
		AudioExtWAV:  AudioCategory,
		AudioExtOGG:  AudioCategory,
		AudioExtAAC:  AudioCategory,
		AudioExtAMR:  AudioCategory,
		AudioExtAPE:  AudioCategory,
		AudioExtFLAC: AudioCategory,

		DocExtTXT:  DocCategory,
		DocExtPDF:  DocCategory,
		DocExtDOC:  DocCategory,
		DocExtDOCX: DocCategory,
		DocExtXLS:  DocCategory,
		DocExtXLSX: DocCategory,
		DocExtPPT:  DocCategory,
		DocExtPPTX: DocCategory,
	}
)

func IsMediaTypeByName(s string) bool {
	typ, has := ExtCategoryMap[FileExt(fileutil.Ext(s))]
	return typ == ImageCategory || typ == VideoCategory && has
}
