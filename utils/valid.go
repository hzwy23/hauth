package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var MuYezi = []string{"0", "1"}
var IsValue = []string{"0", "1"}
var ValueMethod = []string{"101", "102", "103", "104", "105", "106"}

type valid struct {
}

//检验汉子，单词，括号，逗号，问好，句号,除号
//后边变参做多接受两个
//第一个标识校验最小长度，第二个标识校验最大长度
//正确匹配，返回true，错误匹配返回false
func ValidHanAndWord(str string, args ...int) bool {
	return vd.checkHanAndWord(str, args...)
}

//描述: 日期校验
//
//输入参数:
//
//	date: 校验的字符串
//
//输出参数:
//
//	bool:  true：字符串为日期格式  false：字符串不为日期格式
//
func ValidDate(date string) bool {
	return vd.checkdate(date)
}

//描述: 字符串内容为数字校验
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为数字）或2个 args[0]为长度下限  args[1]为长度下限
//
//输出参数:
//
//	bool:  true：字符串全由数字字符组成(1234)  false：字符串不是全由数字字符组成(654r)
//
func ValidNumber(num string, args ...int) bool {
	return vd.checkallnum(num, args...)
}

//描述: 字符串全为字母校验
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为字母）或2个 args[0]为长度下限  args[1]为长度下限
//
//输出参数:
//
//	bool:  true：全为字母  false：不全是字母
//
func ValidAlpha(s string, args ...int) bool {
	return vd.checkalpha(s, args...)
}

//描述: 字符串全为数字字母下划线
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为数字字母下划线）或2个 args[0]为字符串长度下限  args[1]为字符串长度下限
//
//输出参数:
//
//	bool:  true：满足描述  false：不满足描述
//
func ValidWord(word string, args ...int) bool {
	return vd.checkWord(word, args...)
}

//描述: 字符串全为数字字母
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为数字字母）或2个 args[0]为字符串长度下限  args[1]为字符串长度下限
//
//输出参数:
//
//	bool:  true：满足描述  false：不满足描述
//
func ValidAlphaNumber(s string, args ...int) bool {
	return vd.checkwordnumber(s, args...)
}

//描述: 字符串全为数字，包含小数（－123，＋132）
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为数字小数）或2个 args[0]为字符串长度下限  args[1]为字符串长度下限
//
//输出参数:
//
//	bool:  true：满足描述  false：不满足描述
//
func ValidBalance(s string, args ...int) bool {
	return vd.checkfloat(s, args...)
}

//描述: 字符串全为中文汉字
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为汉字）或2个 args[0]为汉字长度下限  args[1]为汉字长度下限
//
//输出参数:
//
//	bool:  true：字符串全为中文汉字 false：字符串不全为中文汉字
//
func ValidHan(s string, args ...int) bool {
	return vd.checkhan(s, args...)
}

//描述: 字符串全为中文汉字加数字，必须以中文汉字开头
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为汉字数字）或2个 args[0]为字符串长度下限  args[1]为字符串长度下限
//
//输出参数:
//
//	bool:  true：字符串全为中文汉字加数字 false：字符串不全为中文汉字加数字
//
func ValidOrderHanNumber(s string, args ...int) bool {
	return vd.checkhanumber(s, args...)
}

//描述: 字符串全为字母加数字，开头必须为字母
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验是否为字母加数字）或2个 args[0]为字符串长度下限  args[1]为字符串长度下限
//
//输出参数:
//
//	bool:  true：字符串全为字母数字 false：字符串不全为字母数字
//
func ValidOrderAlpNumber(s string, args ...int) bool {
	return vd.checkOrderAlpNumber(s, args...)
}

//描述: 字符串全为中文汉字或字母或括号，句号，逗号，省略号
//
//输入参数:
//
//	s: 校验的字符串
//
//	args:   限0个（只校验满足描述）或2个 args[0]为字符长度下限  args[1]为字符长度下限
//
//输出参数:
//
//	bool:  true：满足描述 false：不满足描述
//
func ValidHanWord(s string, args ...int) bool {
	return vd.checkHanAndWord(s, args...)
}

func ValidHanAndAlnum(s string, args ...int) bool {
	if strings.TrimSpace(s) == "" {
		fmt.Println("false")
		return false
	}
	return vd.checkHanAndAlnum(s, args...)
}

//描述: 字符串s在字符串数组ss中
//
//输入参数:
//
//	s: 校验的字符串
//
//	ss: 对比的数组切片
//
//输出参数:
//
//	bool:  true：满足描述 false：不满足描述
//
func ValidStandard(s string, ss []string) bool {
	return vd.checkstand(s, ss)
}

//描述: 字符串是否是IP地址
//
//输入参数:
//
//	s: 校验的字符串
//
//输出参数:
//
//	bool:  true：输入字符串为ip地址格式 false：输入字符串不为IP地址格式
//
func ValidIP(s string) bool {
	return vd.checkip(s)
}

//描述: 字符串是否是email
//
//输入参数:
//
//	s: 校验的字符串
//
//输出参数:
//
//	bool:  true：输入字符串为email格式 false：输入字符串不为email格式
//
func ValidEmail(s string) bool {
	return vd.checkemail(s)
}

//描述: 字符串是否是大陆手机号码格式
//
//输入参数:
//
//	s: 校验的字符串
//
//输出参数:
//
//	bool:  true：输入字符串为大陆手机号码格式 false：输入字符串不为大陆手机号码格式
//
func ValidMobile(s string) bool {
	return vd.checkmobile(s)
}

//描述: 字符串是否是大陆固定电话号码格式
//
//输入参数:
//
//	s: 校验的字符串
//
//输出参数:
//
//	bool:  true：输入字符串为固定电话号码格式 false：输入字符串不为固定电话号码格式
//
func ValidTel(s string) bool {
	return vd.checktel(s)
}

//描述: 字符串内容是数字0或者1
//
//输入参数:
//
//	s: 校验的字符串
//
//输出参数:
//
//	bool:  true：字符串内容是数字0或者1   false：字符串内容不是数字0或者1
//
func ValidBool(s string) bool {
	return vd.checkbool(s)
}

var vd *valid

func newValid() *valid {
	v := new(valid)
	return v
}

func ValidAlnumAndSymbol(s string, args ...int) bool {
	return vd.checkwordnumberAndSymbol(s, args...)
}

func (v *valid) checkwordnumberAndSymbol(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkwordnumber  参数错误")
		return false
	}
	min := "1"
	max := "999"
	if len(args) == 2 {
		if args[1]-args[0] < 0 {
			fmt.Errorf("checkwordnumber  第二个数字参数应该大于等于第一个")
			return false
		}
		min = strconv.Itoa(args[0])
		max = strconv.Itoa(args[1])
	}

	r, err := regexp.Compile(`^[[:alnum:](\_)(\()(\))]{` + min + `,` + max + `}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	return r.MatchString(s)
}

func init() {
	vd = newValid()
}

//校验全位数字
func (v *valid) checkallnum(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkallnum  参数错误")
		return false
	}
	min := "1"
	max := "999"
	if len(args) == 2 {
		if args[1]-args[0] < 0 {
			fmt.Errorf("checkallnum  第二个数字参数应该大于等于第一个")
			return false
		}
		min = strconv.Itoa(args[0])
		max = strconv.Itoa(args[1])
	}

	r, err := regexp.Compile(`^[\d]{` + min + `,` + max + `}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	return r.MatchString(s)

}

//
func (v *valid) checkalpha(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkalpha  参数错误")
		return false
	}
	min := 1
	max := 999
	if len(args) == 2 {
		min = args[0]
		max = args[1]
	}
	for _, ss := range s {
		if ('Z' < ss || ss < 'A') && ('z' < ss || ss < 'a') {
			return false
		}
	}
	if len(s) <= max && len(s) >= min {
		return true
	}
	return false

}
func (this *valid) checkWord(word string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkWord  参数错误")
		return false
	}
	var min = "0"
	var max = "999"
	for index, val := range args {
		if index == 0 {
			min = strconv.Itoa(val)
		} else if index == 1 {
			max = strconv.Itoa(val)
		}
	}

	r, err := regexp.Compile(`^[[:word:]]{` + min + `,` + max + `}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	return r.MatchString(word)
}
func (v *valid) checkstand(s string, ss []string) bool {
	if len(s) == 0 || len(ss) == 0 {
		return false
	}
	for _, val := range ss {
		if s == val {
			return true
		}
	}
	return false
}
func (v *valid) checkhanumber(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("参数错误")
		return false
	}
	min := 1
	max := 999
	if len(args) == 2 {
		if args[1]-args[0] < 0 {
			fmt.Errorf("第二个数字参数应该大于等于第一个")
			return false
		}
		min = args[0]
		max = args[1]
	}
	r, err := regexp.Compile(`^[\p{Han}]+\d?\d*$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	if r.MatchString(s) {
		if len(s) >= min && len(s) <= max {
			return true
		}
	}
	return false
}

func (v *valid) checkOrderAlpNumber(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkOrderAlpNumber  参数错误")
		return false
	}
	min := 1
	max := 999
	if len(args) == 2 {
		if args[1]-args[0] < 0 {
			fmt.Errorf("checkOrderAlpNumber  第二个数字参数应该大于等于第一个")
			return false
		}
		min = args[0]
		max = args[1]
	}
	r, err := regexp.Compile(`^[a-zA-Z]+\d?$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	if r.MatchString(s) {
		if len(s) >= min && len(s) <= max {
			return true
		}
	}
	return false
}

//
func (v *valid) checkwordnumber(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkwordnumber  参数错误")
		return false
	}
	min := "1"
	max := "999"
	if len(args) == 2 {
		if args[1]-args[0] < 0 {
			fmt.Errorf("checkwordnumber  第二个数字参数应该大于等于第一个")
			return false
		}
		min = strconv.Itoa(args[0])
		max = strconv.Itoa(args[1])
	}

	r, err := regexp.Compile(`^[[:alnum:]]{` + min + `,` + max + `}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	return r.MatchString(s)
}

func (v *valid) checkfloat(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkfloat  参数错误")
		return false
	}
	min := 1
	max := 999
	if len(args) == 2 {
		if args[1]-args[0] < 0 {
			fmt.Errorf("checkfloat  第二个数字参数应该大于等于第一个")
			return false
		}
		min = args[0] //strconv.Itoa(args[0])
		max = args[1] //strconv.Itoa(args[1])
	}
	cfloat := `^[(\-)|(+)]?(\d*\.)?\d+$`
	match, _ := regexp.MatchString(cfloat, s)
	if match {
		if len(s) >= min && len(s) <= max {
			return true
		}
		return false
	}
	return false
}
func (v *valid) checkhan(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkhan  参数错误")
		return false
	}
	min := "1"
	max := "999"
	if len(args) == 2 {
		if args[1]-args[0] < 0 {
			fmt.Errorf("checkhan  第二个数字参数应该大于等于第一个")
			return false
		}
		min = strconv.Itoa(args[0])
		max = strconv.Itoa(args[1])
	}
	r, err := regexp.Compile(`^[\p{Han}]{` + min + `,` + max + `}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	return r.MatchString(s)

}
func (v *valid) checkHanAndWord(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkHanAndWord 参数错误")
		return false
	}
	var min = "1"
	var max = "999"
	for index, val := range args {
		if index == 0 {
			min = strconv.Itoa(val)
		} else if index == 1 {
			max = strconv.Itoa(val)
		}
	}

	r, err := regexp.Compile(`^[\p{Han}[:word:](\()(\))(\{)(\})(,)(.)(?)(。)(……)(、)(/)]{` + min + `,` + max + `}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	return r.MatchString(s)
}
func (v *valid) checkHanAndAlnum(s string, args ...int) bool {
	if len(args) > 2 {
		fmt.Errorf("checkHanAndWord 参数错误")
		return false
	}
	var min = "1"
	var max = "999"
	for index, val := range args {
		if index == 0 {
			min = strconv.Itoa(val)
		} else if index == 1 {
			max = strconv.Itoa(val)
		}
	}

	r, err := regexp.Compile(`^[\p{Han}[:alnum:](\()(\))(\{)(\})(,)(.)(?)(。)(……)(、)(/)]{` + min + `,` + max + `}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}
	return r.MatchString(s)
}
func (v *valid) checkip(s string) bool {
	cip := `^(\d+)\.(\d+)\.(\d+)\.(\d+)$`
	match, _ := regexp.MatchString(cip, s)
	return match
}
func (v *valid) checkemail(s string) bool {
	cemail := "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?"
	match, _ := regexp.MatchString(cemail, s)
	return match
}
func (v *valid) checkmobile(s string) bool {
	cmobile := `^(\+86)?1[3|4|5|6|8]\d{9}$`
	match, _ := regexp.MatchString(cmobile, s)
	return match
}
func (v *valid) checktel(s string) bool {
	fixtel := `^(0\d{2,3}(-)?)?\d{7,8}$`
	match, _ := regexp.MatchString(fixtel, s)
	return match
}
func (v *valid) checkbool(s string) bool {
	if len(s) != 1 {
		return false
	}
	for _, ss := range s {
		if ss == '1' || ss == '0' {
			return true
		}
	}
	return false

}

//日期校验
func (v *valid) checkdate(date string) bool {
	r, err := regexp.Compile(`^[\d]{4}/[\d]{1,2}/[\d]{1,2}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return false
	}

	if r.MatchString(date) {
		d := strings.Split(date, "/")
		if len(d) != 3 {
			fmt.Errorf("日期格式错误")
			return false
		}
		year, err := strconv.Atoi(d[0])
		if err != nil {
			fmt.Errorf("%v", err)
			return false
		}
		mths := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			mths[2] = 29
		}
		month, err := strconv.Atoi(d[1])
		if err != nil {
			fmt.Errorf("%v", err)
			return false
		}
		day, err := strconv.Atoi(d[2])
		if err != nil {
			fmt.Errorf("%v", err)
			return false
		}
		if month > 12 || month < 0 || mths[month] < day {
			return false
		}
		return true
	} else {

		rc, err := regexp.Compile(`^[\d]{4}-[\d]{1,2}-[\d]{1,2}$`)
		if err != nil {
			fmt.Errorf("%v", err)
			return false
		}
		if rc.MatchString(date) {
			d := strings.Split(date, "-")
			if len(d) != 3 {
				fmt.Errorf("日期格式错误")
				return false
			}
			year, err := strconv.Atoi(d[0])
			if err != nil {
				fmt.Errorf("%v", err)
				return false
			}
			mths := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				mths[2] = 29
			}
			month, err := strconv.Atoi(d[1])
			if err != nil {
				fmt.Errorf("%v", err)
				return false
			}
			day, err := strconv.Atoi(d[2])
			if err != nil {
				fmt.Errorf("%v", err)
				return false
			}
			if month > 12 || month < 0 || mths[month] < day {
				return false
			}
			return true
		} else {
			return false
		}
	}
}
