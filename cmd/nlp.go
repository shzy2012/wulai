package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
)

//nlpEntitiesExtract 实体抽取
func nlpEntitiesExtract(content string) string {

	botResp, err := WulaiClient.NLPEntitiesExtract(content)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			return cliErr.Message()
		} else if serErr, ok := err.(*errors.ServerError); ok {
			return serErr.Message()
		} else {
			return err.Error()
		}
	}

	bytes, _ := json.Marshal(botResp)
	return fmt.Sprintf("%s\n", bytes)
}

//nlpTokenize 分词&词性标注
func nlpTokenize(content string) string {

	botResp, err := WulaiClient.NLPTokenize(content)
	if err != nil {
		if cliErr, ok := err.(*errors.ClientError); ok {
			return cliErr.Message()
		} else if serErr, ok := err.(*errors.ServerError); ok {
			return serErr.Message()
		} else {
			return err.Error()
		}
	}

	bytes, _ := json.Marshal(botResp)
	return fmt.Sprintf("%s\n", bytes)
}

/****************
 user cmd
****************/

//NlpEntitiesExtract 实体抽取
var NlpEntitiesExtract = &ishell.Cmd{
	Name:     "nlpEntitiesExtract",
	Aliases:  []string{"nlpentitiesextract", "nlp_entities_extract", "nee"},
	Help:     "实体抽取",
	LongHelp: "实体抽取 nlpEntitiesExtract query  \n\n参数 query:待实体抽取query",
	Func: func(c *ishell.Context) {
		frontFunc(c, 1, func() {
			c.Println("操作结果:", nlpEntitiesExtract(c.Args[0]))
		})
	},
}

//NlpTokenize 分词&词性标注
var NlpTokenize = &ishell.Cmd{
	Name:     "nlpTokenize",
	Aliases:  []string{"nlptokenize", "nlp_tokenize", "nt"},
	Help:     "实体抽取",
	LongHelp: "实体抽取 nlpTokenize query  \n\n参数 query:分词&词性标注",
	Func: func(c *ishell.Context) {
		frontFunc(c, 1, func() {
			c.Println("操作结果:", nlpTokenize(c.Args[0]))
		})
	},
}
