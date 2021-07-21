// 所有对外暴露的方法写在这里面
package logic

import (
	"errors"
	"strconv"
)

func (e *Engine) Cmd(event string, callback func(req *Request)) {
	Broadcaster.on(event, callback)
}

// SendMsgById 通过用户id发送消息
func (e *Engine) SendMsgByID(id int, event string, msg interface{}) error {
	user := Broadcaster.users[id]
	if user == nil {
		return errors.New("用户" + strconv.Itoa(id) + "不在线")
	}
	user.MessageChannel <- NormalMessage(event, msg)
	return nil
}

// CLoseConnById 通过id关闭连接
func (e *Engine) CloseConnByID(id int) {
	user := Broadcaster.users[id]
	if user != nil {
		user.close = true
	}
}

// 根据过滤条件关闭连接
func (e *Engine) CloseConnByFilter(callback func(user *User) bool) {
	// 游客和登录用户一起循环
	recipient := [2]map[int]*User{
		Broadcaster.users, Broadcaster.tourists,
	}
	for _, v := range recipient {
		for _, v1 := range v {
			ok := callback(v1)
			if ok {
				v1.close = true
			}
		}
	}
}

// SendMsgByFilter 遍历所有当前在线的用户返回的不是nil就发送消息
func (e *Engine) SendMsgByFilter(event string, callback func(user *User) interface{}) {
	// 游客和登录用户一起循环
	recipient := [2]map[int]*User{
		Broadcaster.users, Broadcaster.tourists,
	}
	for _, v := range recipient {
		for _, v1 := range v {
			msg := callback(v1)
			if msg == nil {
				continue
			}
			v1.MessageChannel <- NormalMessage(event, msg)
		}
	}
}
