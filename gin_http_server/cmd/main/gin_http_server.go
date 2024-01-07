package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func main() {
	// myhttpserver.CreateHttpServer()
	tree, err := toml.LoadFile("config/templates_list.toml")
	if err != nil {
		fmt.Println("载入TOML文件时发生错误：", err)
		return
	}
	route := tree.Get("templates").([]interface{})

	for _, v := range route {
		fmt.Println(v)
	}

	return
}
