package response

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseGetTransferedCustomerList struct {
	response.ResponseWX
	ExternalContactList []*ResponseCustomer `json:"customer"`
}

type ResponseCustomer struct {
	ExternalUserID string `json:"external_userid"` // "woAJ2GCAAAd4uL12hdfsdasassdDmAAAAA",
	Status         string `json:"Status"`          //
	TakeoverTime   int    `json:"takeover_time"`   //
	ErrCode        string `json:"errcode"`
}
