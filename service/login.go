package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WXLoginResp  openid 和sessionkey
type WXLoginResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Login 用得到的code登录
func Login(code string) (WXLoginResp, error) {
	get := "https://api.weixin.qq.com/sns/jscode2session?appid=wxf2b2ebb728ece45e&secret=1c708c51001e9e33d20a40d2949455f5&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(get)
	if err != nil {
		return WXLoginResp{}, err
	}
	var u = WXLoginResp{}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return WXLoginResp{}, err
	}

	if u.ErrCode != 0 {
		return u, fmt.Errorf(fmt.Sprintf("ErrCode:%d  ErrMsg:%s", u.ErrCode, u.ErrMsg))
	}

	return u, nil
}
