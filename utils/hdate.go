package utils

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func DataInterval(start string, end string) int {
	var y1, m1, d1 int
	var y2, m2, d2 int

	year_start, _ := strconv.Atoi(start[0:4])
	month_start, _ := strconv.Atoi(start[5:7])
	day_start, _ := strconv.Atoi(start[8:10])

	year_end, _ := strconv.Atoi(end[0:4])
	month_end, _ := strconv.Atoi(end[5:7])
	day_end, _ := strconv.Atoi(end[8:10])

	m1 = (month_start + 9) % 12
	y1 = year_start - m1/10
	d1 = 365*y1 + y1/4 - y1/100 + y1/400 + (m1*306+5)/10 + (day_start - 1)

	m2 = (month_end + 9) % 12
	y2 = year_end - m2/10
	d2 = 365*y2 + y2/4 - y2/100 + y2/400 + (m2*306+5)/10 + (day_end - 1)

	return (d2 - d1)
}

func MonthInterval(start string, pmt_freq int) (int, error) {
	lt, err := AddMonths(start, pmt_freq)
	if err != nil {
		fmt.Errorf("%v", err)
		return 0, err
	}
	return DataInterval(lt, start), nil
}

func AddMonths(start string, num int) (string, error) {
	var oldmths = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var mths = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(start[0:4])
	month, _ := strconv.Atoi(start[5:7])
	if month > 12 || month < 1 {
		fmt.Errorf("月份数错误，请输入月份在1-12之间的数字")
		return ",", errors.New("月份数错误，请输入月份在1-12之间的数字")
	}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		oldmths[2] = 29
	}

	day, _ := strconv.Atoi(start[8:10])
	if day > oldmths[month] || day < 1 {
		fmt.Errorf("%v", strconv.Itoa(day)+"不在该夜的正确天数范围内")
		return ",", errors.New(strconv.Itoa(day) + "不在该夜的正确天数范围内")
	}
	new_day := day

	if num > 0 {

		new_months := (month + num) % 12
		if new_months == 0 {
			new_months = 12
		}
		new_year := year + (month+num)/12

		if (month+num)/12 > 0 && (month+num)%12 == 0 {
			new_year = new_year - 1
		}

		//		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		//			oldmths[2] = 29
		//		}
		if (new_year%4 == 0 && new_year%100 != 0) || new_year%400 == 0 {
			mths[2] = 29
		}
		if oldmths[month] == day {
			new_day = mths[new_months]
		} else {
			if day > mths[new_months] {
				new_day = mths[new_months]
			} else {
				new_day = day
			}
		}
		return FormatIntTodate(new_year, new_months, new_day)
	} else if num == 0 {
		return start, nil
	} else {

		Y := int(math.Ceil((float64(-num-month) + 0.000000000000001) / 12))
		new_months := month + 12*Y + num
		new_year := year - Y

		//		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		//			oldmths[2] = 29
		//		}
		if (new_year%4 == 0 && new_year%100 != 0) || new_year%400 == 0 {
			mths[2] = 29
		}
		if oldmths[month] == day {
			new_day = mths[new_months]
		} else {
			if day > mths[new_months] {
				new_day = mths[new_months]
			} else {
				new_day = day
			}
		}
		return FormatIntTodate(new_year, new_months, new_day)
	}

}

func AddDays(start string, num int) (string, error) {
	var mths = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(start[0:4])
	month, _ := strconv.Atoi(start[5:7])
	if month > 12 || month < 1 {
		fmt.Errorf("月份数错误，请输入月份在1-12之间的数字")
		return ",", errors.New("月份数错误，请输入月份在1-12之间的数字")
	}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		mths[2] = 29
	}

	day, _ := strconv.Atoi(start[8:10])
	if day > mths[month] || day < 1 {
		fmt.Errorf("%v", strconv.Itoa(day)+"不在该夜的正确天数范围内")
		return ",", errors.New(strconv.Itoa(day) + "不在该夜的正确天数范围内")
	}

	if num > 0 {
		for num > 0 {
			if num+day > mths[month] {
				num = num - (mths[month] - day + 1)
				day = 1
				month = month + 1
				if month > 12 {
					month = 1
					year = year + 1
					if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
						mths[2] = 29
					} else {
						mths[2] = 28
					}
				}
			} else {
				day = num + day
				num = 0
			}
		}
	}
	if num < 0 {
		num = -num
		for num > 0 {
			if num-day >= 0 {
				num = num - day
				month = month - 1
				if month < 1 {
					month = 12
					year = year - 1
					if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
						mths[2] = 29
					} else {
						mths[2] = 28
					}
				}
				day = mths[month]
			} else {
				day = day - num
				num = 0
			}
		}
	}
	return FormatIntTodate(year, month, day)
}

//将字符串格式化未yyyy-mm-dd形式
func FormatIntTodate(year int, month int, day int) (string, error) {
	nyear := strconv.Itoa(year)
	nmonth := strconv.Itoa(month)
	nday := strconv.Itoa(day)
	for len(nyear) < 4 {
		nyear = "0" + nyear
	}
	for len(nmonth) < 2 {
		nmonth = "0" + nmonth
	}
	for len(nday) < 2 {
		nday = "0" + nday
	}
	rst := nyear + "-" + nmonth + "-" + nday
	return rst, nil
}

func FormatStringToDate(date string) (string, error) {
	r, err := regexp.Compile(`^[\d]{4}-[\d]{1,2}-[\d]{1,2}$`)
	if err != nil {
		fmt.Errorf("%v", err)
		return "", err
	}

	if str := r.FindString(date); str != "" {
		if len(str) == 10 {
			return str, nil
		} else {
			s := strings.Split(str, "-")
			y := s[0]
			m := s[1]
			d := s[2]
			for len(m) < 2 {
				m = "0" + m
			}
			for len(d) < 2 {
				d = "0" + d
			}
			if newD := y + "-" + m + "-" + d; ValidDate(newD) {
				return newD, nil
			} else {
				return "", errors.New("输入内容不合理，无法转换成正确的日期")
			}
		}
	} else {
		rc, err := regexp.Compile(`^[\d]{4}/[\d]{1,2}/[\d]{1,2}$`)
		if err != nil {
			fmt.Errorf("%v", err)
			return "", err
		}
		if str := rc.FindString(date); str != "" {
			if len(str) == 10 {
				return strings.Replace(str, "/", "-", -1), nil
			} else {
				s := strings.Split(str, "/")
				y := s[0]
				m := s[1]
				d := s[2]
				for len(m) < 2 {
					m = "0" + m
				}
				for len(d) < 2 {
					d = "0" + d
				}
				if newD := y + "-" + m + "-" + d; ValidDate(newD) {
					return newD, nil
				} else {
					return "", errors.New("输入内容不合理，无法转换成正确的日期")
				}
			}
		} else {
			return "", errors.New("输入参数不合理，无法转换成日期类型")
		}
	}

}

func AGTB(A string, B string) bool {
	return A > B
}

func AGTEB(A string, B string) bool {
	return A >= B
}

// 第二个参数大于第一个参数时，返回 1
// 第二个参数小于第一个参数时，返回 －1
// 第一个参数等于第二个参数时，返回 0
func CompareDate(start string, end string) int {

	year_start, _ := strconv.Atoi(start[0:4])
	month_start, _ := strconv.Atoi(start[5:7])
	day_start, _ := strconv.Atoi(start[8:10])

	year_end, _ := strconv.Atoi(end[0:4])
	month_end, _ := strconv.Atoi(end[5:7])
	day_end, _ := strconv.Atoi(end[8:10])

	if year_end > year_start {
		return 1
	} else if year_end == year_start {
		if month_end > month_start {
			return 1
		} else if month_end == month_start {
			if day_end > day_start {
				return 1
			} else if day_end == day_start {
				return 0
			} else {
				return -1
			}
		} else {
			return -1
		}
	} else {
		return -1
	}
}

//日期校验
func ValidDate(date string) bool {
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

func FormatToDate(str string) (string, error) {
	return DateFormat(str, "YYYY-MM-DD")
}

func FormatToTime(str string) (string, error) {
	return DateFormat(str, "YYYY-MM-DD HH24:MM:SS")
}

// 将字符串类型的时间转换成日期
func DateFormat(str string, model string) (string, error) {
	switch model {
	case "YYYY-MM-DD":
		pattern := `^[1-2]{1}[0-9]{3}[-|/][0-9]{1,2}[-|/][0-9]{1,2}`
		re, err := regexp.Compile(pattern)
		if err != nil {
			return str, err
		}

		rst := re.FindString(str)
		if rst == "" {
			return str, errors.New("no match")
		}
		return rst, nil
	case "YYYY-MM-DD HH24:MM:SS":
		pattern := `^[1-2]{1}[0-9]{3}(-|/)[0-9]{2}(-|/)[0-9]{2}(T)[0-9]{2}:[0-9]{2}:[0-9]{2}`
		re, err := regexp.Compile(pattern)
		if err != nil {
			return str, err
		}
		rst := re.FindString(str)
		if rst == "" {
			return str, errors.New("no match")
		}
		return strings.Replace(rst, "T", " ", 1), nil
	}
	return str, errors.New("model is unsupported.")
}
