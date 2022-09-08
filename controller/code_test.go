package controller

import "testing"

func TestResCode_Msg(t *testing.T) {
	tests := []struct {
		name string
		c    ResCode
		want string
	}{
		{"1", 10000, "success"},
		{"2", 10001, "请求参数错误"},
		{"3", 10002, "用户已存在"},
		{"4", 10003, "用户名不存在"},
		{"5", 10004, "用户名或密码错误"},
		{"6", 10005, "服务繁忙"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Msg(); got != tt.want {
				t.Errorf("Msg() = %v, want %v", got, tt.want)
			}
		})
	}
}
