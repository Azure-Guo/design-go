package service

import (
	"design-go/handler"
	"design-go/pojo"
)

func Suggest(userName string) []string {
	userInfo := getUserInfo(userName)
	result := handler.ProcessHandler(userInfo)
	return result
}

// 因为需要查询缓存，如果缓存没有，需要查库.本应在dao层，为了简便写这里
func getUserInfo(username string) pojo.UserInfo {
	return pojo.UserInfo{}
}
