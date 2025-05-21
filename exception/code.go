package exception

var (
	UnknownError        = &CustomError{"UnknownError", "未知的错误"}
	FileNotFound        = &CustomError{"FileNotFound", "文件%s未找到"}
	FileFormatError     = &CustomError{"FileFormatError", "文件%s格式错误"}
	InvalidParam        = &CustomError{"InvalidParam", "无效的参数%s值：%v"}
	ParamIsBlank        = &CustomError{"ParamIsBlank", "参数%s值为空"}
	JsonLoadError       = &CustomError{"JsonLoadError", "JSON字符串序列化失败：%s"}
	JsonDumpError       = &CustomError{"JsonLoadError", "JSON字符串反序列化失败"}
	MethodNotImplement  = &CustomError{"MethodNotImplement", "方法%s暂未实现"}
	HttpDownloadError   = &CustomError{"HttpDownloadError", "Http下载失败(%s)"}
	CreateFileError     = &CustomError{"CreateFileError", "文件创建失败：%s"}
	CopyFileError       = &CustomError{"CopyFileError", "文件复制失败：%s"}
	OpenFileError       = &CustomError{"OpenFileError", "文件打开失败：%s"}
	HttpError           = &CustomError{"HttpError", "Http请求错误：%s"}
	GetUserHomeDirError = &CustomError{"GetUserHomeDirError", "获取用户家目录失败"}
)
