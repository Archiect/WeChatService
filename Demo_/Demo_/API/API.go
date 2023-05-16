package API

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	 "net/http"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "ID"
	appSecret       = "key"
)

func GetOpenId(code string) string {
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code) //拼接网址
	resp, err := http.Get(url)                                  //发起请求
	if err != nil {
		fmt.Println("请求失败")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json := gojsonq.New().FromString(string(body)).Find("openid")
	openId := json.(string)
	return openId
}
