package main

import (
	"wulai/cmd"

	"github.com/abiosoft/ishell"
)

func main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("吾来机器人shell,help帮助,exit退出")

	//shell.IgnoreCase(true)
	// register a function for "greet" command.
	shell.AddCmd(cmd.InitEnv)
	shell.AddCmd(cmd.PrintEnv)

	//用户类
	shell.AddCmd(cmd.UserCreate)
	shell.AddCmd(cmd.UserAttributeUpdate)
	shell.AddCmd(cmd.UserAttributeList)

	//bot类
	shell.AddCmd(cmd.MsgBotResponse)
	shell.AddCmd(cmd.MsgBotResponseTask)
	shell.AddCmd(cmd.MsgBotResponseKeyword)
	shell.AddCmd(cmd.MsgBotResponseQa)

	//nlp类
	shell.AddCmd(cmd.NlpEntitiesExtract)
	shell.AddCmd(cmd.NlpTokenize)

	// start shell
	shell.Run()
}
