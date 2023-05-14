package authhandler

import (
	"log"
	databaseengine "tajacara/go_code/dataBase_engine"

	"xorm.io/xorm"
)

func IsAuthSysytemUser(db *xorm.Engine, loging string, hash string) bool {

	user, err := databaseengine.GetSystemUserFromDatabase(db, loging)
	if err != nil {
		log.Print("Cannot get SystemUser from database with login", loging)
	}
	if user.Hash == hash {
		return true
	}
	return false
}

func CheckAuth(db *xorm.Engine, loging string, cookie string) bool {
	user, err := databaseengine.GetSystemUserFromDatabase(db, loging)
	if err != nil {
		log.Print("Cannot get SystemUser from database with login while checking ", loging)
	}
	if user.LastCookie == cookie && user.Login == loging {
		return true
	} else {
		return false
	}

}
