package service

import (
	"design-go/pay/facade"
	"design-go/pay/pojo"
)

func Pay(body pojo.PayBody) bool {
	success := facade.Pay(body)
	if success {
		saveToDb(body)
	}
	return success
}

func saveToDb(body pojo.PayBody) {

}
