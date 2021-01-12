package docx

var (
	chineseStr = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}
)

// Num2Str 只能转换二位数
func Num2ChineseStr(num int) string {
	if num > 100 {
		return ""
	}
	var res string
	if num >= 10 {
		var prefix string
		if num/10-1 != 0 {
			prefix = chineseStr[num/10-1]
		}
		res = prefix + "十"
	}
	if num%10 > 0 {
		res += chineseStr[num%10-1]
	}

	return res
}
