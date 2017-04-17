package i18n

import (
	"os"
	"path/filepath"
	"net/http"
	"strings"
	"github.com/nicksnyder/go-i18n/i18n"
)

// 从i18n配置文件中获取id对应的翻译值.
func Get(req *http.Request,translationID string, args ...interface{}) string {
	lang:=strings.Split(req.Header.Get("accept-language"),",")
	if len(lang)==0{
		return translationID
	}
	T,err:=i18n.Tfunc(lang[0])
	if err != nil {
		T,_ = i18n.Tfunc("zh-cn")
	}
	return T(translationID,args...)

}

func GetSuccess(req *http.Request) string {
	return Get(req,"success")
}

func GetDisconnect(req *http.Request) string {
	return Get(req,"as_of_date_disconnect")
}

func GetPageNotFound(req *http.Request) string {
	return Get(req,"page_not_found")
}

// 初始化i18n文件
func init(){
	HOME:=os.Getenv("HBIGDATA_HOME")
	i18n.LoadTranslationFile(filepath.Join(HOME,"i18n/zh-cn.yaml"))
	i18n.LoadTranslationFile(filepath.Join(HOME,"i18n/en-us.yaml"))
}
