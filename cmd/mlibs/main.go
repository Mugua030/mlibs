package main

import (
	"log"
	//	"mlibs/dao"
	//	"mlibs/services"
	//	"net/http"
	//	"mlibs/handlers"
)

func main() {
	log.Println("begin run ...")
	//defer dao.Db.Close()

	// error process
	/*
		uid := int64(3)
		userinfo, err := dao.GetUserInfo(uid)
		if err != nil {
			log.Fatalf("GetUserInfo fail: %v", err)
		}
		if userinfo == nil   {
			log.Print("no current user uid=", uid)
			return
		}
	*/

	// test v2
	/*
		orderid := int64(3)
		orderinfo, err := services.GetOrderInfo(orderid)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("result:: order=%#v",  orderinfo)
	*/

	app := NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("app run fail, error=%v", err)
	}

}
