package rgstring

import "strings"

// 读取两个特定字符串之间的字符串
// 参数注释：
//  str    源字符串
//  start  开头字符串
//  end    结尾字符串
//  返回值：两个特定字符串之间的字符串
func ReadBetween(str string, start string, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return ""
	}
	return str[s:][:e]
}
