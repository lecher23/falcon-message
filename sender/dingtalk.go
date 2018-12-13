package sender

import (
	"errors"
	"log"
	"net/http"
	"os"
	"github.com/labstack/echo"
)

var cfgGateway string

func init() {
	val, has := os.LookupEnv("DING_API")
	cfgGateway = `https://oapi.dingtalk.com/robot/send?access_token=`
	if has {
		cfgGateway = val
		log.Println("use dd gateway from env config", cfgGateway)
	}
}

type DingTalk struct {
}

func (d *DingTalk) Send(token string, content string) error {
	if token == "" {
		return errors.New("need dingding token")
	}

	// 发送钉钉
	ding := NewDing(token)
	ding.Gateway = cfgGateway
	result := ding.SendMessage(Message{Content: content})
	log.Println(result)
	if !result.Success {
		log.Println("token:", token)
		return echo.NewHTTPError(http.StatusBadRequest, result.ErrMsg)
	}

	return nil
}

func NewDingTalk() *DingTalk {
	return &DingTalk{}
}
