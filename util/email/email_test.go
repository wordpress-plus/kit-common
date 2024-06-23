package email

import (
	"github.com/alice52/jasypt-go"
	"testing"
)

var secretQq string
var secretNet string
var secretGmail string

func init() {
	secretQq, _ = jasypt.New().DecryptWrapper("ENC(hlomTQKeIwivZYpT22kVC/oiRnewPAXza2LZo87/0PObwbdYVF/p5+NCb/069aZmP2D/p740TbMTl8W9uslWzg==)")
	secretNet, _ = jasypt.New().DecryptWrapper("ENC(mPbyo0f3VKO5kBGJaKOggTLbfqXR103iOfQhn548ff+EI1hMrj5q3YffhwMRKNdM)")
	secretGmail, _ = jasypt.New().DecryptWrapper("ENC(8YwWbLC7ZwTZZrE76TzTjtBq19M8NVq0AJRjzQ7jYpKQWOMQLxlka0foCqfdKbIN0Yrql8R8WUFaPhpnKg0tuw==)")
}
func TestDoSendQQ(t *testing.T) {
	err := DoSend("1252068782@qq.com", "123", "email.Bod",
		"zzhang_xz@163.com", "zack", secretNet)
	if err != nil {
		panic(err)
	}
}

func TestDoSendNet(t *testing.T) {

	err2 := DoSend("zzhang_xz@163.com", "123", "email.Bod",
		"1252068782@qq.com", "zack", secretQq)
	if err2 != nil {
		panic(err2)
	}
}

func TestDoSendGmail(t *testing.T) {
	err3 := DoSend("zzhang_xz@163.com", "123", "email.Bod",
		"danielzhang182@gmail.com", "dylan", secretGmail)
	if err3 != nil {
		panic(err3)
	}
}
