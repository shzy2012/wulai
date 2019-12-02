package cmd

import (
	"fmt"
	"os"

	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

const (
	notInitClient = "未初始化吾来客户端，请使用init初始化"
	initClientFromEnv = "已通过环境变量，完成初始化"
	missedParam   = "缺少参数"
)

//WulaiClient 吾来客户端
var WulaiClient *wulai.Client
var pubkey string
var secret string

//输出颜色
var yellow func(a ...interface{}) string

func init() {
	yellow = color.New(color.FgYellow).SprintFunc()
	pubkey, _ = os.LookupEnv("pubkey")
	secret, _ = os.LookupEnv("secret")
	if pubkey != "" && secret != "" {
		initClient(secret, pubkey)
		fmt.Println(yellow(initClientFromEnv))
	}
}

func initClient(secret, pubkey string) {
	WulaiClient = wulai.NewClient(secret, pubkey)
	WulaiClient.Version = "v2"
	WulaiClient.SetDebug(false)
}

func envIsOk() bool {
	if WulaiClient != nil {
		return true
	}

	return false
}

//frontFunc 前置函数，检查环境变量
func frontFunc(c *ishell.Context, paramCount int, method func()) {
	if envIsOk() {
		if len(c.Args) < paramCount {
			c.Println(yellow(missedParam))
		} else {
			method()
		}
	} else {
		c.Println(yellow(notInitClient))
	}
}

/****************
 初始化环境 cmd
****************/

//InitEnv 初始化环境
var InitEnv = &ishell.Cmd{
	Name:     "init",
	Help:     "初始化吾来客户端",
	LongHelp: "初始化吾来客户端 init pubkey secret",
	Func: func(c *ishell.Context) {
		if len(c.Args) < 2 {
			fmt.Println(`缺少参数： 参数pubkey,secret不能为空.`)
			return
		}

		pubkey = c.Args[0]
		secret = c.Args[1]
		initClient(secret, pubkey)
		fmt.Println("操作成功: pubkey:"+pubkey+" secret:", secret)
		return
	},
}

//PrintEnv 打印环境变量
var PrintEnv = &ishell.Cmd{
	Name:     "print",
	Help:     "打印 pubkey secret",
	LongHelp: "打印当前的pubkey secret",
	Func: func(c *ishell.Context) {
		fmt.Println("当前变量: pubkey:"+pubkey+" secret:", secret)
	},
}
