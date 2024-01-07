// @copyright Copyright 2024 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func main() {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		fmt.Println("载入TOML文件时发生错误：", err)
		return
	}
	// 1. 字符串
	str1 := config.Get("host").(string)
	fmt.Println("1、字符串\nhost: ", str1)
	// 2. 整数
	int1 := config.Get("strength").(int64)
	fmt.Println("2、整数\nstrength: ", int1)
	// 3. 浮点数
	float1 := config.Get("pi").(float64)
	fmt.Println("3、浮点数\npi: ", float1)
	// 4. 布尔型
	bool1 := config.Get("is_true").(bool)
	fmt.Println("4、布尔值\nis_true: ", bool1)
	// 5. 本地日期时间
	ldt := config.Get("ldt").(toml.LocalDateTime)
	fmt.Println("5、本地日期时间\nldt: ", ldt)
	// 注意：在TOML配置文件中，对于本地日期类型的数据，在最后一定要有一个空格，
	// 否则在 LoadFile 时就会报错。
	ld1 := config.Get("ld1").(toml.LocalDate)
	fmt.Println("5、本地日期\nld1: ", ld1)
	lt1 := config.Get("lt1").(toml.LocalTime)
	fmt.Println("5、本地时间\nlt1: ", lt1)
	// 6. 数组
	arr1 := config.Get("arr1").([]interface{})
	fmt.Println("6、数组\narr1: ", arr1[0], arr1[1], arr1[2])
	arr2 := config.Get("arr2").([]interface{})
	fmt.Println("6、数组\narr2: ", arr2[0], arr2[1], arr2[2])
	// 7. 表
	// 7.1 直接获取
	directAddr := config.Get("server.address").(string)
	directPort := config.Get("server.port").(int64)
	fmt.Println("7.1、表：直接获取\naddress: ", directAddr, "port: ", directPort)
	// 7.,2 间接获取
	server := config.Get("server").(*toml.Tree)
	addr := server.Get("address").(string)
	port := server.Get("port").(int64)
	fmt.Println("7.2、表：间接获取\naddress: ", addr, "port: ", port)
	// 8. 内联表
	server1 := config.Get("server1").(*toml.Tree)
	addr1 := server1.Get("address").(string)
	port1 := server1.Get("port").(int64)
	fmt.Println("8、内联表\naddress: ", addr1, "port: ", port1)
	// 9. 表数组
	users := config.Get("users").([]*toml.Tree)
	for _, user := range users {
		name := user.Get("name").(string)
		email := user.Get("email").(string)
		fmt.Printf("9、表数组\nname: %s, email: %s\n", name, email)
	}
}
