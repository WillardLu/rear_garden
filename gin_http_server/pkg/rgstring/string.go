// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language go 1.18.1
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package rgstring

import "strings"

// 读取两个特定字符串之间的字符串
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
