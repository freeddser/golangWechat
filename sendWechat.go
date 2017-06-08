package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type WxToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type SentTemplete struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func GetAccessToken(corpid,corpsecret string) string {
	resp, err := http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid="+corpid+"&corpsecret="+corpsecret)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "get AccessToken faild!"
	}
	//fmt.Println(string(body))
	var Tokenjson WxToken

	err1 := json.Unmarshal(body, &Tokenjson)
	if err1 != nil {

		fmt.Println(err1)
	}
	//fmt.Println(Tokenjson.AccessToken)

	return Tokenjson.AccessToken
}



func Send_Msg_To_Group(access_token string,agentid int, tagidstr, msg string) {
	//fmt.Print(access_token)

	var SentContent SentTemplete
	SentContent.Agentid = agentid
	SentContent.Toparty = ""
	SentContent.Totag = tagidstr
	SentContent.Msgtype = "text"
	SentContent.Text.Content = msg

	if bs, err := json.Marshal(&SentContent); err != nil {
		panic(err)
	} else {
		resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="+access_token,
			"application/json; charset=UTF-8",
			strings.NewReader(string(bs)))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(resp)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
	}

}


func main() {
	//company id
	var  corpid string="sdfsdfsdfaads"
	//app key
	var corpsecret string="asdfsafdsafsdafsadfad"
	//app id
	var agentid int=1000002

	if len(os.Args) < 4 {
		fmt.Print("#####Please input secret, tagid,msg:100 'alert xxx '#####\n")
		fmt.Println("Example:go run sentWechat.go 123456 1 testmsg")
		os.Exit(1)
	}
  #if you don't need this can remove
	if os.Args[1] != "123456" {
		fmt.Print("#####secret error,exit#####")
		os.Exit(1)
	}
	Send_Msg_To_Group(GetAccessToken(corpid,corpsecret), agentid,os.Args[2], os.Args[3])

}
