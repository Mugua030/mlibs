package services

import (
	"log"

	"mlibs/dao"
	"github.com/pkg/errors"
)

type Order interface {
	GetOrderInfo(orderID int64) (OrderDetail, error)
	GenOrder() (int64, error)
}

type OrderDetail struct {
	OrderID int64
	Title string
	OrderPrice float64
	UserName string
	UserID int64
	Addr *dao.Address
	CreateTs int64
	PayTs int64
}

type MiOrder struct {}

func GetOrderInfo(orderID int64) (*OrderDetail, error) {
	uid := int64(2)

	addr, err := dao.GetDefaultAddr(uid)
	if err != nil {
		return nil, err
	}
	if addr == nil {
		log.Print("user addr is empty")
	}
	// user info
	user, err := dao.GetUserInfo(uid)
	if err != nil {
		return nil, errors.WithMessage(err, "services: get user info fail")
	}
	username := ""
	if user != nil {
		username = user.Name
	}



	orderInfo := &OrderDetail{
		OrderID: int64(1),
		Title: "test",
		Addr: nil,
		UserName: username,
	}

	return orderInfo, nil
}
func (m *MiOrder) GenOrder() {}