package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

//msgBotResponse 获取机器人回复
func msgBotResponse(userID, content string) string {

	//消息类型[文本消息]
	textMsg := &wulai.Text{
		Content: content,
	}

	botResp, err := WulaiClient.MsgBotResponse(userID, textMsg, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			return cliErr.Error()
		} else if serErr, ok := err.(*errors.ServerError); ok {
			return serErr.Error()
		} else {
			return err.Error()
		}
	}

	bytes, _ := json.Marshal(botResp)
	return fmt.Sprintf("%s\n", bytes)
}

//msgBotResponseTask 获取任务机器人回复
func msgBotResponseTask(userID, content string) string {

	//消息类型[文本消息]
	textMsg := &wulai.Text{
		Content: content,
	}

	botResp, err := WulaiClient.MsgBotResponseTask(userID, textMsg, "")
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			return cliErr.Error()
		} else if serErr, ok := err.(*errors.ServerError); ok {
			return serErr.Error()
		} else {
			return err.Error()
		}
	}

	bytes, _ := json.Marshal(botResp)
	return fmt.Sprintf("%s\n", bytes)
}

/****************
 bot cmd
****************/

//MsgBotResponse 获取机器人回复
var MsgBotResponse = &ishell.Cmd{
	Name:     "msgBotResponse",
	Aliases:  []string{"msgbotresponse", "msg_bot_response", "mbr"},
	Help:     "获取机器人回复",
	LongHelp: "获取机器人回复 msgBotResponse userID context\n\n\n\n参数 userID:用户ID context:问题",
	Func: func(c *ishell.Context) {
		frontFunc(c, 2, func() {
			userID := c.Args[0]
			context := c.Args[1]
			fmt.Println("操作结果:", msgBotResponse(userID, context))
		})
	},
}

//MsgBotResponseTask 获取任务机器人回复
var MsgBotResponseTask = &ishell.Cmd{
	Name:     "msgBotResponseTask",
	Aliases:  []string{"msgbotresponsetask", "msg_bot_response_task", "mbrt"},
	Help:     "获取任务机器人回复",
	LongHelp: "获取任务机器人回复 msgBotResponseTask userID context\n\n\n\n参数 userID:用户ID context:问题",
	Func: func(c *ishell.Context) {
		frontFunc(c, 2, func() {
			userID := c.Args[0]
			context := c.Args[1]
			fmt.Println("操作结果:", msgBotResponseTask(userID, context))
		})
	},
}
