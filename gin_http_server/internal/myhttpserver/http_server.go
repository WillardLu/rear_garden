package myhttpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
)

type TemplateList struct {
	name string
	file string
}

// 读取模板文件
func readTemplates(templates []*TemplateList) int {
	// 读取模板列表配置文件
	tree, err := toml.LoadFile("config/templates_list.toml")
	if err != nil {
		fmt.Println("载入TOML文件时发生错误：", err)
		return -1
	}
	route := tree.Get("templates").([]interface{})

	// 把模板列表中的名称与实际的文件内容逐一读入结构数组中
	for i := 0; i < len(route); i = i + 2 {
		templates[i/2].name = route[i].(string)
		contents, err := os.ReadFile(route[i+1].(string))
		if err != nil {
			fmt.Println("读取模板文件时发生错误：", err)
			return -1
		}
		templates[i/2].file = string(contents)
	}
	return 0
}

// 创建http server
func CreateHttpServer() {
	// 从TOML配置文件中读取http服务器参数
	config, err := toml.LoadFile("http_server_config.toml")
	if err != nil {
		fmt.Println("载入TOML文件时发生错误：", err)
		return
	}
	// 设置运行模式
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// 设置路由
	if SetupRouter(router) == -1 {
		return
	}

	srv := &http.Server{
		Addr:              config.Get("address").(string) + ":" + config.Get("port").(string),
		Handler:           router,
		ReadTimeout:       time.Duration(config.Get("ReadTimeout").(int64)) * time.Second,
		WriteTimeout:      time.Duration(config.Get("WriteTimeout").(int64)) * time.Second,
		ReadHeaderTimeout: time.Duration(config.Get("ReadHeaderTimeout").(int64)) * time.Second,
		IdleTimeout:       time.Duration(config.Get("IdleTimeout").(int64)) * time.Second,
	}
	// 声明一个匿名函数，并创建一个goroutine（有的翻译为协程）
	go func() {
		// 监听请求
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 关闭（或重启）服务
	// 1）创建通道，用来接收信号
	quit := make(chan os.Signal, 1)
	// 2）监听和捕获信号
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("\n>> 开始关闭 http server……")

	// 3）创建一个子节点的context,5秒后自动超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("\n>> http server 关闭时出错：", err)
	}
	select {
	case <-ctx.Done():
	default:
	}
	log.Println("\n>> http server 退出。")
	return
}

// 设置路由
//  gin router
//  返回值：0 成功，-1 失败
func SetupRouter(router *gin.Engine) int {
	// 设置CSS样式的位置
	router.StaticFS("/css", http.Dir("./static/css"))
	// 设置静态图片的位置
	router.StaticFS("/images", http.Dir("./static/images"))

	// 读取模板文件内容
	// 通过这样的方式把模板文件内容读入内存，以减少磁盘读取
	template, error := os.ReadFile("./templates/template1.html")
	if error != nil {
		log.Println("Error reading file:", error)
		return -1
	}

	// 这里不能直接调用多参数的函数，
	// 需要使用func(c *gin.Context)作为中转来调用多参数的函数
	router.GET("/", func(c *gin.Context) {
		homePage(c, string(template))
	})
	router.GET("/task-list", func(c *gin.Context) {
		taskList(c, string(template))
	})

	return 0
}

// 生成主页
//  c - 指向gin.Context的指针
//  page - 模板文件内容
func homePage(c *gin.Context, template string) {
	var msg []byte
	str := "<a href=\"#\" onclick=\"location.reload()\">主页</a>"
	msg = []byte(strings.Replace(string(template), "{{.sub-dir}}", str, -1))
	c.Writer.Write(msg)
	// c.File("./static/index.html") // 这种方式可以直接读取文件内容并显示在浏览器上
}

// 生成任务列表页
//  c - 指向gin.Context的指针
//  page - 模板文件内容
func taskList(c *gin.Context, template string) {
	var msg []byte
	str := "<a href=\"\">任务列表</a>"
	msg = []byte(strings.Replace(template, "{{.sub-dir}}", str, -1))
	c.Writer.Write(msg)
}
