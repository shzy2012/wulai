package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

//userCreate 创建用户
func userCreate(userID string) string {

	user, err := WulaiClient.UserCreate(userID, "", "")
	if err != nil {
		return err.Error()
	}

	return user.UserID
}

//userAttributeCreate 更新用户属性
func userAttributeCreate(userID, attrID, attrValue string) string {

	err := WulaiClient.UserAttributeCreate(userID, attrID, attrValue)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			return cliErr.Message()
		} else if serErr, ok := err.(*errors.ServerError); ok {
			return serErr.Message()
		} else {
			return err.Error()
		}
	}

	return "成功"
}

//userAttributeList 获取用户属性列表
func userAttributeList(isAttrGroup bool, page, pageSize int) string {

	resp, err := WulaiClient.UserAttributeList(isAttrGroup, page, pageSize)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			return cliErr.Message()
		} else if serErr, ok := err.(*errors.ServerError); ok {
			return serErr.Message()
		} else {
			return err.Error()
		}
	}

	bytes, _ := json.Marshal(resp)
	return fmt.Sprintf("%s\n", bytes)
}

/****************
 user cmd
****************/

//UserCreate 创建用户
var UserCreate = &ishell.Cmd{
	Name:     "userCreate",
	Aliases:  []string{"usercreate", "user_create", "uc"},
	Help:     "创建用户",
	LongHelp: "创建用户 userCreate userID  \n\n参数 userID:用户ID",
	Func: func(c *ishell.Context) {
		frontFunc(c, 1, func() {
			c.Println("操作结果:", userCreate(c.Args[0]))
		})
	},
}

//UserAttributeUpdate 更新用户属性
var UserAttributeUpdate = &ishell.Cmd{
	Name:     "userAttributeUpdate",
	Aliases:  []string{"userattributeupdate", "user_attribute_update", "uau"},
	Help:     "更新用户属性",
	LongHelp: "更新用户属性 userAttributeUpdate userID  \n\n参数 userID:用户ID attrID:属性id attrValue:属性值",
	Func: func(c *ishell.Context) {
		frontFunc(c, 3, func() {
			userID, attrID, attrValue := c.Args[0], c.Args[1], c.Args[2]
			c.Println("操作结果:", userAttributeCreate(userID, attrID, attrValue))
		})
	},
}

//UserAttributeList 获取用户属性列表
var UserAttributeList = &ishell.Cmd{
	Name:     "userAttributeList",
	Aliases:  []string{"userattributelist", "user_attribute_list", "ual"},
	Help:     "获取用户属性列表",
	LongHelp: "获取用户属性列表 userAttributeList page  \n\n参数 page:页码,从1开始 pageSize:每页数量",
	Func: func(c *ishell.Context) {
		frontFunc(c, 2, func() {

			page, pageSize := c.Args[0], c.Args[1]
			pageI, err := strconv.Atoi(page)
			if err != nil {
				c.Println(yellow(`参数错误: 参数page需要输入数字.`))
				return
			}

			pageSizeI, err := strconv.Atoi(pageSize)
			if err != nil {
				c.Println(yellow(`参数错误: 参数pageSize需要输入数字.`))
				return
			}
			c.Println("操作结果:", userAttributeList(true, pageI, pageSizeI))
		})
	},
}
