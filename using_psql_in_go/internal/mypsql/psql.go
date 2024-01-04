// @copyright Copyright 2024 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package mypsql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml"
)

// 读取PostgreSQL数据库连接参数
//  file TOML配置文件名
//  返回值：配置内容，错误值
func GetPsqlConfig(file string) (string, int) {
	config, err := toml.LoadFile("pg_config.toml")
	if err != nil {
		fmt.Println("载入TOML文件时发生错误：", err)
		return "", -1
	}
	var str = ""
	var paramKey = [6]string{"host", "port", "user", "password", "dbname",
		"sslmode"}
	for i := 0; i < 6; i++ {
		// 检查参数是否存在
		err := config.Get(paramKey[i])
		if err == nil {
			fmt.Println("配置文件中缺少参数：", paramKey[i])
			return "", -1
		}
		str = str + paramKey[i] + "=" + config.Get(paramKey[i]).(string) + " "
	}
	return str, 0
}

// 连接PostgreSQL数据库
//  config 数据库连接参数
//  返回值：数据库连接句柄，连接信息
func ConnectPsql(config string) (*sql.DB, string) {
	// sql.Open()函数在判断连接是否成功时，显然只做了有限的检测，即只检测第一个参数是否正确。
	// 在第一个参数正确的情况下，无论第二个字符串参数的内容是什么（即使为空），都不会报错。所以
	// 需要做进一步的检验，才能准确判断是否连接成功。
	db, err := sql.Open("postgres", config)
	if err != nil {
		return nil, "连接数据库时发生错误：" + err.Error()
	}
	_, err = db.Query("SELECT 0")
	if err != nil {
		return nil, "查询数据库时发生错误：" + err.Error()
	}
	return db, ">>数据库连接成功！"
}

// 关闭数据库连接
//  db 数据库连接句柄
//  返回值：关闭结果信息，错误值
func ClosePsql(db *sql.DB) (string, int) {
	err := db.Close()
	if err != nil {
		return "关闭数据库连接时发生错误：" + err.Error(), -1
	}
	return "\n>>数据库连接已关闭！", 0
}

// 查询数据
//  db 数据库连接句柄
//  selectStatement 查询语句。语句结尾有没有分号都可以。
//  返回值：查询结果，错误值
func PsqlSelect(db *sql.DB, selectStatement string) (rows *sql.Rows) {
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println("查询数据库时发生错误：", err)
		return nil
	}
	return rows
}

// 插入、修改和删除数据
//  db 数据库连接句柄
//  execStatement sql语句。语句结尾有没有分号都可以。
//  返回值：0 成功，nil 失败
func PsqlExec(db *sql.DB, execStatement string) int {
	_, err := db.Exec(execStatement)
	if err != nil {
		fmt.Println("执行\""+execStatement+"\"时发生错误：", err)
		return -1
	}
	return 0
}
