package model

type FileExt string

const (
	ImageExtJPG  FileExt = "jpg"
	ImageExtPEG  FileExt = "jpeg"
	ImageExtPNG  FileExt = "png"
	ImageExtBMP  FileExt = "bmp"
	ImageExtGIF  FileExt = "gif"
	ImageExtTIF  FileExt = "tif"
	ImageExtICO  FileExt = "ico"
	ImageExtICON FileExt = "icon"
	ImageExtWebP FileExt = "webp"
	ImageExtHEIC FileExt = "heic"
	ImageExtHEIF FileExt = "heif"

	VideoExtMP4  FileExt = "mp4"
	VideoExtM4V  FileExt = "m4v"
	VideoExtMKV  FileExt = "mkv"
	VideoExtAVI  FileExt = "avi"
	VideoExtWMV  FileExt = "wmv"
	VideoExtFLV  FileExt = "flv"
	VideoExtMOV  FileExt = "mov"
	VideoExtMPG  FileExt = "mpg"
	VideoExtMPEG FileExt = "mpeg"
	VideoExt3GP  FileExt = "3gp"
	VideoExtWebM FileExt = "webm"
	VideoExtSWF  FileExt = "swf"
	VideoExtF4V  FileExt = "f4v"
	VideoExtRMVB FileExt = "rmvb"

	AudioExtMP3  FileExt = "mp3"
	AudioExtM4A  FileExt = "m4a"
	AudioExtMID  FileExt = "mid"
	AudioExtMIDI FileExt = "midi"
	AudioExtWAV  FileExt = "wav"
	AudioExtOGG  FileExt = "ogg"
	AudioExtAAC  FileExt = "aac"
	AudioExtAMR  FileExt = "amr"
	AudioExtAPE  FileExt = "ape"
	AudioExtFLAC FileExt = "flac"

	DocExtTXT  FileExt = "txt"
	DocExtPDF  FileExt = "pdf"
	DocExtDOC  FileExt = "doc"
	DocExtDOCX FileExt = "docx"
	DocExtXLS  FileExt = "xls"
	DocExtXLSX FileExt = "xlsx"
	DocExtPPT  FileExt = "ppt"
	DocExtPPTX FileExt = "pptx"
)

var (
	ExtContentTypeMap = map[FileExt]string{
		ImageExtJPG:  "image/jpg",
		ImageExtPEG:  "image/jpeg",
		ImageExtPNG:  "image/png",
		ImageExtBMP:  "image/bmp",
		ImageExtGIF:  "image/gif",
		ImageExtTIF:  "image/tif",
		ImageExtICO:  "image/x-icon",
		ImageExtICON: "image/x-icon",
		ImageExtWebP: "image/webp",
		ImageExtHEIC: "image/heic",
		ImageExtHEIF: "image/heif",

		VideoExtMP4:  "video/mp4",
		VideoExtM4V:  "video/x-m4v",
		VideoExtMKV:  "video/x-matroska",
		VideoExtAVI:  "video/x-msvideo",
		VideoExtWMV:  "video/x-ms-wmv",
		VideoExtFLV:  "video/x-flv",
		VideoExtMOV:  "video/quicktime",
		VideoExtMPG:  "video/mpeg",
		VideoExtMPEG: "video/mpeg",
		VideoExt3GP:  "video/3gpp",
		VideoExtWebM: "video/webm",
		VideoExtRMVB: "video/rmvb",

		AudioExtMP3:  "audio/mpeg",
		AudioExtM4A:  "audio/x-m4a",
		AudioExtMID:  "audio/mid",
		AudioExtMIDI: "audio/midi",
		AudioExtWAV:  "audio/wav",
		AudioExtOGG:  "application/ogg",
		AudioExtAAC:  "audio/aac",
		AudioExtAMR:  "audio/amr",
		AudioExtAPE:  "audio/ape",
		AudioExtFLAC: "audio/flac",

		DocExtTXT:  "text/plain",
		DocExtPDF:  "application/pdf",
		DocExtDOC:  "application/msword",
		DocExtDOCX: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		DocExtXLS:  "application/vnd.ms-excel",
		DocExtXLSX: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		DocExtPPT:  "application/vnd.ms-powerpoint",
		DocExtPPTX: "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	}
)

func (e FileExt) String() string {
	return string(e)
}
