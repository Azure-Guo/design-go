package handler

import (
	"design-go/pojo"
)

// 抽象处理方法
type AbstractSuggestRequirementHandler interface {
	processHandler(userInfo pojo.UserInfo, suggestList []string) []string
}

// 具体处理方法
type PersonalCheckHandler struct{}
type CityCheckHandler struct{}
type NewCheckHandler struct{}
type RecentCheckHandler struct{}

func (*PersonalCheckHandler) processHandler(userInfo pojo.UserInfo, suggestList []string) []string {
	// 通过个人资质的check,找到了4个可以投放的业务，放到suggestList 中

	suggestList = append(suggestList, "1")
	suggestList = append(suggestList, "2")
	suggestList = append(suggestList, "3")
	suggestList = append(suggestList, "4")
	return suggestList
}
func (*CityCheckHandler) processHandler(userInfo pojo.UserInfo, suggestList []string) []string {
	// 通过获取userinfo的city属性
	// userInfo.City
	// 通过city和之前保留的业务信息进行对比，筛选出剩余的业务投放
	suggestList = deleteSlice(suggestList, "1")
	return suggestList
}
func (*NewCheckHandler) processHandler(userInfo pojo.UserInfo, suggestList []string) []string {
	if userInfo.IsNewUser {
		// 特定的list
		suggestList = make([]string, 0)
	}
	return suggestList
}
func (*RecentCheckHandler) processHandler(userInfo pojo.UserInfo, suggestList []string) []string {
	// 近期购买产品,进行筛选
	// userInfo.BuyProducts
	suggestList = deleteSlice(suggestList, "2")
	return suggestList
}

func deleteSlice(suggestList []string, item string) []string {
	result := make([]string, 0)
	for i := 0; i < len(suggestList); i++ {
		if suggestList[i] != item {
			result = append(result, suggestList[i])
		}
	}
	return result
}
