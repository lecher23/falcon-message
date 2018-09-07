package util

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

// EncodeJSON json序列化(禁止 html 符号转义)
func EncodeJSON(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//StringToInt string 类型转 int
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("agent 类型转换失败, 请检查配置文件中 agentid 配置是否为纯数字(%v)", err)
		return 0
	}
	return n
}

// HandleContent [P2][PROBLEM][10-13-33-153][][测试 all(#1) net.port.listen port=2 0==0][O3 2017-06-06 16:46:00]
func HandleContent(content string) string {
	parts := strings.Split(content, "][")
	if len(parts) != 6 {
		return "[错误]: 解析以下通知出错:\n" + content
	}
	alertLevel := parts[0][1:]
	alertStatus := parts[1]
	alertEnv := parts[2]
	msgParts := strings.Split(parts[4], " ")
	if len(msgParts) < 2 {
		return "[错误]: 解析通知正文出错:\n" + parts[4]
	}

	alertText := msgParts[0]
	alertDetail := strings.Join(msgParts[1:], ",")
	alertTime := parts[5][:len(parts[5])-1]

	content = fmt.Sprintf("Falcon报警\n\n"+
		"严重等级: %s\n"+
		"当前状态: %s\n"+
		"报警环境: %s\n"+
		"报警内容: %s\n"+
		"报警参数: %s\n"+
		"产生时间: %s", alertLevel, alertStatus, alertEnv, alertText, alertDetail, alertTime)

	return content
}

