package dao

import (
	"log"
	"database/sql"
)

type User struct {
	Uid  int64
	Name string
	Addr string
}
/*
func (u User) IsEmpty() bool {
	if u == (User{}) {
		return true
	}
	return false
}
*/

var (
	uid  int64
	name string
	addr string
)

func GetUserInfo(uid int64) (User, error) {
	err := Db.QueryRow("select user_id, name, address from user where user_id = ?", uid).Scan(&uid, &name, &addr)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
		log.Fatal(err)
	}
	userRow := User{
		Uid:  uid,
		Name: name,
		Addr: addr,
	}

	return userRow, nil
}
