package http

// RequestSchema
type RequestSchema struct {
	Url          string
	Method       string
	Headers      map[string][]string //头部
	ContentType  string
	Values       map[string][]string //get请求参数
	Body         any
	FileToUpload string              //上传文件
	OutputFile   string              //输出文件
}
