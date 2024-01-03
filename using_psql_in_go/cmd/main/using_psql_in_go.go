package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml"
)

// 读取PostgreSQL数据库连接参数
//  file TOML配置文件名
//  返回值：配置内容，错误值
func GetPostgresConfig(file string) (string, int) {
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
func ConnectPostgres(config string) (*sql.DB, string) {
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
func ClosePostgres(db *sql.DB) (string, int) {
	err := db.Close()
	if err != nil {
		return "关闭数据库连接时发生错误：" + err.Error(), -1
	}
	return ">>数据库连接已关闭！", 0
}

func main() {
	// 获取数据库连接参数
	config, err := GetPostgresConfig("pg_config.toml")
	if err != 0 {
		fmt.Println("获取PostgreSQL数据库连接参数时发生错误")
		return
	}
	// 连接数据库
	db, msg := ConnectPostgres(config)
	fmt.Println(msg)
	if db == nil {
		return
	}
	// 关闭数据库连接
	msg, err = ClosePostgres(db)
	fmt.Println(msg)

	return
}
