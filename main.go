package main

import (
	"log"
	"mlibs/dao"
)

func main() {
	log.Println("begin run ...")
	defer dao.Db.Close()

	// error process
	uid := int64(3)
	userinfo, err := dao.GetUserInfo(uid)
	if err != nil {
		log.Fatal(err)
	}
	if userinfo == (dao.User{})   {
		log.Print("no current uid=", uid)
		return
	}
	log.Printf("result=%v", userinfo)
}
