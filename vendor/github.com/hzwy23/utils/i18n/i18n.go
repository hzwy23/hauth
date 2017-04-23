package i18n

import (
	"github.com/hzwy23/utils/logs"
	"github.com/nicksnyder/go-i18n/i18n"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 从i18n配置文件中获取id对应的翻译值.
func Get(req *http.Request, translationID string, args ...interface{}) string {
	lang := strings.Split(req.Header.Get("accept-language"), ",")
	if len(lang) == 0 {
		return translationID
	}
	T, err := i18n.Tfunc(lang[0])
	if err != nil {
		T, _ = i18n.Tfunc("zh-cn")
	}
	return T(translationID, args...)
}

// 执行陈宫
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

// 注册i18n文件
func Register(file string) {
	err := i18n.LoadTranslationFile(file)
	if err != nil {
		logs.Error(err)
	}
}

// 初始化i18n文件
func init() {
	// register system default i18n file
	HOME := os.Getenv("HBIGDATA_HOME")
	i18n.LoadTranslationFile(filepath.Join(HOME, "views", "i18n", "zh-cn.yaml"))
	i18n.LoadTranslationFile(filepath.Join(HOME, "views", "i18n", "en-us.yaml"))
}
