package constant

const (
	BinFileVersion   = 2.15
	BinFileSererName = "bin-fs"

	MaxBufferSize        = 32 << 10
	MaxLogContentLength  = 1 << 10
	MaxUploadChunkSize   = 10 << 20
	MaxBodyLimit         = (10 << 20) + (1 << 10)
	DeviceSingleDiskSize = 3848290697216
	TaraXSystem          = "android"
	TaraWSSystem         = "linux"
	DeviceSecret         = "7ce1150025dac5d2598b2fc45c4b494c"
	HeaderKeyToken       = "token"
	CtxKeyRequestID      = "req-id"
	CtxKeyUserID         = "user-id"
	SyncClientUA         = "sync-client/1.0"
)
