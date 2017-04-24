package i18n

import "net/http"

func NoAuth(req *http.Request) string {
	return Get(req, "as_of_date_no_auth")
}

// 执行成功
func Success(req *http.Request) string {
	return Get(req, "success")
}

// 用户连接已断开
func Disconnect(req *http.Request) string {
	return Get(req, "as_of_date_disconnect")
}

// 页面不存在
func PageNotFound(req *http.Request) string {
	return Get(req, "page_not_found")
}

// 没有权限读取域中的信息
func ReadDomain(req *http.Request, domain_id string) string {
	return Get(req, "read_domain_insufficient", domain_id)
}

// 没有分隔符,格式不正确
func NoSeparator(req *http.Request, id string) string {
	return Get(req, "as_of_date_no_separator", id)
}

// 没有权限向这个域中写入信息
func WriteDomain(req *http.Request, domain_id string) string {
	return Get(req, "write_domain_insufficient", domain_id)
}
