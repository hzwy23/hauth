package i18n

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"os"
	"path/filepath"
)


// 从i18n配置文件中获取id对应的翻译值.
func Get(translationID string, args ...interface{}) string{
	T,_:=i18n.Tfunc("zh-cn")
	return T(translationID,args)
}

func Success()string{
	return Get("success")
}

func Disconnect()string{
	return Get("as_of_date_disconnect")
}

// 初始化i18n文件
func init(){
	HOME:=os.Getenv("HBIGDATA_HOME")
	i18n.MustLoadTranslationFile(filepath.Join(HOME,"i18n/zh-cn.yaml"))
	i18n.LoadTranslationFile(filepath.Join(HOME,"i18n/en-us.yaml"))
}
