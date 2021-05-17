package errno

type Code uint16

const (
	CodeOk       Code = 0
	CodeNotFound Code = 404

	// 1000～1999 区间表示参数错误
	ErrInvalidParams     Code = 1000
	CodeInvalidSignature Code = 1001

	// 2000～2999 区间表示用户错误
	CodePermissionDenied    Code = 2000
	CodeTokenNotRequired    Code = 2001
	CodeTokenAuthExpired    Code = 2002
	CodeTokenRecordExist    Code = 2003
	CodeFromHashNoExist     Code = 2004
	CodeFileNameIllegal     Code = 2005
	CodFileNameRecordExist  Code = 2006
	CodeFileInitializeError Code = 2007
	CodeFileHashNotFound    Code = 2008
	CodeFileSizeExceedLimit Code = 2009
	CodeThumbUploadError    Code = 2010
	CodeChunkUploadError    Code = 2011
	CodeInvalidChunkHash    Code = 2012
	CodeCreateChunkError    Code = 2013
	CodeFileUnCompleted     Code = 2014
	CodeFileIsCompleted     Code = 2015
	CodeChunkReadError      Code = 2016
	CodeFileDeleteError     Code = 2017
	CodeFolderIdNotFound    Code = 2018
	CodeFolderIsExisted     Code = 2019
	CodeFileCopyFailed      Code = 2020
	CodeFileMoveFailed      Code = 2021
	CodeFileRenameFailed    Code = 2022
	CodeThumbnailNotFound   Code = 2023

	// 3000～3999 区间表示接口异常
	CodeUnknownError       Code = 3000
	CodeRequestExceedLimit Code = 3001
	CodeInvalidOperation   Code = 3002
	CodeDevHardDiskLoss    Code = 3003
)

var (
	codeMap = map[Code]string{
		CodeOk:       "该操作已完成",
		CodeNotFound: "该请求路径无效或不存在 :(",

		// 1000～1999 区间表示参数错误
		ErrInvalidParams:     "缺少参数或参数错误",
		CodeInvalidSignature: "SIGN 参数验证失败",

		// 2000～2999 区间表示用户错误
		CodePermissionDenied:    "您没有权限访问",
		CodeTokenNotRequired:    "TOKEN 不存在或者错误",
		CodeTokenAuthExpired:    "TOKEN 已过期，请重新获取",
		CodeTokenRecordExist:    "TOKEN 已存在，请重新绑定",
		CodeFromHashNoExist:     "参数列表中包含不存在的文件Hash",
		CodeFileNameIllegal:     "该名称含有非法字符，请重命名后上传",
		CodFileNameRecordExist:  "该名称在当前文件夹已存在，请重命名",
		CodeFileInitializeError: "初始化文件失败，请重新操作",
		CodeFileHashNotFound:    "该文件不存在，请重新操作",
		CodeFileSizeExceedLimit: "您上传的文件大小超出限制，默认是 10M，请重新上传",
		CodeThumbUploadError:    "文件缩略图 Thumb 上传失败",
		CodeChunkUploadError:    "分片文件 Chunk 上传失败",
		CodeInvalidChunkHash:    "分片文件 SHA256 检验失败",
		CodeCreateChunkError:    "分片文件创建元数据失败",
		CodeFileUnCompleted:     "文件未合并或缺少分片文件",
		CodeFileIsCompleted:     "文件合并已完成，请勿重复操作",
		CodeChunkReadError:      "文件无法读取，主从数据已损坏",
		CodeFileDeleteError:     "文件删除失败，请重新操作",
		CodeFolderIdNotFound:    "文件夹ID不存在，请重新操作",
		CodeFolderIsExisted:     "文件夹已存在，请重命名后创建",
		CodeFileCopyFailed:      "文件复制时出错，请检查源和目标文件夹是否正确",
		CodeFileMoveFailed:      "文件移动时出错，请检查源和目标文件夹是否正确",
		CodeFileRenameFailed:    "文件重命名时出错，请重新操作",
		CodeThumbnailNotFound:   "文件缩略图未上传，请重新上传",

		// 3000～3999 区间表示接口异常
		CodeUnknownError:       "发生未知错误",
		CodeRequestExceedLimit: "您请求得太频繁",
		CodeInvalidOperation:   "非法的操作",
		CodeDevHardDiskLoss:    "设备挂载盘丢失，NAS系统已无法正常工作 (请联系工作人员操作)",
	}
)

func (e Code) Error() string {
	return codeMap[e]
}

func (e Code) Message() string {
	return codeMap[e]
}
