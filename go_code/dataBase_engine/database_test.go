package databaseengine

import (
	"fmt"
	"testing"
)

func TestInsertConnection(t *testing.T) {
	engine, err := CreateDBEngine()
	if err != nil {
		t.Error(err)
	}
	pc := Pc{
		Id: 2,
		OS: "Windows",
		PO: []string{
			"Hola", "muchachos",
		},
		Critical_vulns: []string{"no"},
	}
	po := Po_on_pc{
		Id:      10,
		Version: "1.8.0",
		CriticalVuln: []string{
			"cve-20102111",
		},
	}
	wu := WebUser{
		Id:    10,
		Login: "futurame",
		Hash:  "11111",
	}

	u := User{
		Id:       10,
		Username: "Hola",
		Email:    "asdas",
		Password: "dasdasdas",
	}

	su := SystemUser{
		Id:    1,
		Login: "adasdasd",
		Hash:  "Sdasdasda",
	}

	s := make([]interface{}, 6)
	s[0] = pc
	s[1] = po
	s[2] = wu
	s[3] = u
	s[4] = su
	err = AddInterfaceToDatabase(engine, s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("test passed")
}

func TestGetConnection(t *testing.T) {
	engine, err := CreateDBEngine()
	if err != nil {
		t.Error(err)
	}
	user, _ := GetUserFromDatabase(engine, "Hola")
	fmt.Println(user)
}

func TestSysGet(t *testing.T) {
	engine, err := CreateDBEngine()
	if err != nil {
		t.Error(err)
	}
	user, _ := GetSystemUserFromDatabase(engine, "adasdasd")
	fmt.Println(user)
}
