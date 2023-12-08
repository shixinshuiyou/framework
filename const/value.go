package _const

const (
	TraceID = "X-Request-Trace-ID"
)

// 架构设计中：需要携带的应用信息
const (
	ServiceName    = "X-Service-Name"    // 服务标志/名称
	ServiceVersion = "X-Service_Version" // 服务版本信息
)

// 架构设计中：需要携带的账号信息
const (
	UserID   = "X-App-User-ID"
	UserName = "X-App-User-Name"
)
