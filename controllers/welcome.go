package controllers

import (
	"github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/jssdk"
	"math/rand"
	"regexp"
	"time"
)

type WelcomeController struct {
	beego.Controller
}

func (this *WelcomeController) Get() {
	var AccessTokenServer = mp.NewDefaultAccessTokenServer("wxa2dc73d6e114eb30", "a058552dcb7da6dc4c0f646eecf3a92a", nil)
	var mpClient = mp.NewClient(AccessTokenServer, nil)
	var TicketServer = jssdk.NewDefaultTicketServer(mpClient)
	tick, err := TicketServer.Ticket()
    TicketServer.
	if err != nil {
		beego.Error(err)
	}
	this.Ctx.WriteString(tick)
    jssdk.WXConfigSign(tick, nonceStr, timestamp, url)
	// name := this.GetString("name")
	// if IsChineseName(name) {
	// 	this.Ctx.WriteString(name + "is")
	// } else {
	// 	this.Ctx.WriteString(name + "no")
	// }
}

func (this *WelcomeController) Submit() {
	name := this.GetString("name")
	mobile := this.GetString("mobile")

	if name == "" || mobile == "" {
		this.Data["json"] = map[string]string{"code": "error", "info": "姓名或密码为空"}
		this.ServeJson()
		return
	}
}

func (this *WelcomeController) Luck() {
	if !this.Ctx.Input.IsAjax() {
		this.Ctx.WriteString("invalid request")
		return
	}
	rand.Seed(time.Now().Unix())
	number := rand.Intn(100)
	if number > 0 && number <= 10 {
		this.Data["json"] = map[string]string{"code ": "success", "info": "恭喜您中奖!"}
	} else {
		this.Data["json"] = map[string]string{"code": "error", "info": "很遗憾!没有中奖!"}
	}
	this.ServeJson()
	return
}

//检测是否是电话号码
func IsMobile(mobile string) bool {
	pattern := "^1(3[0-9]|4[57]|5[0-35-9]|8[0-9]|70)\\d{8}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}

func IsChineseName(name string) bool {
	pattern := "^([\u4e00-\u9fa5]{2,4})$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}
