package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func main() {
	config, err := toml.LoadFile("pg_config.toml")
	if err != nil {
		fmt.Println("载入TOML文件时发生错误：", err)
		return
	}

	pg_host := config.Get("host").(string)
	pg_port := config.Get("port").(string)
	pg_user := config.Get("user").(string)
	pg_password := config.Get("password").(string)
	pg_dbname := config.Get("dbname").(string)
	pg_sslmode := config.Get("sslmode").(string)
	fmt.Println("host: ", pg_host)
	fmt.Println("port: ", pg_port)
	fmt.Println("user: ", pg_user)
	fmt.Println("password: ", pg_password)
	fmt.Println("dbname: ", pg_dbname)
	fmt.Println("sslmode: ", pg_sslmode)
}
