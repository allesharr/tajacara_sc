package databaseengine

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// var engine *xorm.Engine = nil

type DbEngine struct{}

type SystemUser struct {
	Id         int    `json:"Id" xorm:"int  not null unique autoincr 'id'"`
	Login      string `json:"login" xorm:"text not null 'login'"`
	Hash       string `json:"-" xorm:"text not null 'hash'"`
	LastCookie string `json:"last_cookie" xorm:"text 'last_cookie'"`
}

const (
	DB_HOST     = "10.47.75.135"
	DB_PORT     = 5432
	DB_USER     = "karim"
	DB_PASSWORD = "karim"
	DB_NAME     = "tajacara_base"
)

type User struct {
	Id       int    `json:"Id" xorm:"int  not null unique autoincr 'id'"`
	Username string `json:"username" xorm:"text not null 'username'"`
	Email    string `json:"email" xorm:"text not null unique 'email'"`
	Password string `json:"-" xorm:"text not null 'password'"`
}

type WebUser struct {
	Id    int    `json:"Id" xorm:"int not null unique autoincr 'id'"`
	Login string `json:"Login" xorm:"text not null unique 'login'"`
	Hash  string `json:"Hash" xorm:"text not null 'hash'"`
}
type Po_on_pc struct {
	Id           int      `json:"Id" xorm:"int not null unique autoincr 'id'"`
	Version      string   `json:"Version" xorm:"text 'version'"`
	CriticalVuln []string `json:"CriticalVuln" xorm:"json'critical_vuln'"`
}
type Pc struct {
	Id             int      `json:"Id" xorm:"int not null unique autoincr 'id'"`
	OS             string   `json:"OS" xorm:"text 'os'"`
	PO             []string `json:"PO" xorm:"json 'po'"`
	Critical_vulns []string `json:"CriticalVuln" xorm:"json 'critical_vulns'"`
}

func Check_if_not_Table(x *xorm.Engine, d interface{}) {
	t, err := x.IsTableExist(d)
	if err != nil {
		log.Fatal(err)
	}
	if !t {
		x.CreateTables(d)
	}
}

func CreateDBEngine() (*xorm.Engine, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	engine, err := xorm.NewEngine("postgres", dbinfo)
	if err != nil {

		return nil, err
	}
	if err := engine.Ping(); err != nil {

		return nil, err
	}
	if err := engine.Sync(new(User)); err != nil {

		return nil, err
	}
	engine.ShowSQL(true)
	return engine, nil
}

func CleanPostgres(x *xorm.Engine) error {
	var e error = nil
	if x.DriverName() == "postgres" {
		sess := x.NewSession()
		defer sess.Close()

		if _, err := sess.Exec("DROP DATABASE postgres"); err != nil {
			e = errors.New("failed to drop postgres_database")
		}

		if _, err := sess.Exec("CREATE DATABASE postgres"); err != nil {
			e = errors.New("failed to create database postgres")
		}
	}
	return e
}

// Control with whiteList of what shold be in datatabase
func AddInterfaceToDatabase(x *xorm.Engine, f []interface{}) error {
	var err error = nil
	fmt.Println(x.DriverName())
	if x.DriverName() == "postgres" {
		sess := x.NewSession()
		defer sess.Close()

		for _, e := range f {

			switch v := e.(type) {
			case User:
				b, _ := x.IsTableExist(v)
				if !b {
					x.CreateTables(v)
				}
				_, err = x.Insert(v)
			case Pc:
				b, _ := x.IsTableExist(v)
				if !b {
					x.CreateTables(v)
				}
				_, err = x.Insert(v)
			case Po_on_pc:
				b, _ := x.IsTableExist(v)
				if !b {
					x.CreateTables(v)
				}
				_, err = x.Insert(v)
			case WebUser:
				b, _ := x.IsTableExist(v)
				if !b {
					x.CreateTables(v)
				}
				_, err = x.Insert(v)
			case SystemUser:
				b, _ := x.IsTableExist(v)
				if !b {
					x.CreateTables(v)
				}
				_, err = x.Insert(v)
			default:
				fmt.Println("wrong Type does not Insert to database.")
			}
			if err != nil {
				return err
			}
		}
		// pc.Critical_vulns = json.Marshal(pc.Critical_vulns)

	} else {
		err = errors.New("cannot write to databse, wrong driver")
	}
	return err
}

func AddUserToDatabase(x *xorm.Engine, user User) error {
	var err error = nil
	if x.DriverName() == "postgres" {
		// pc.Critical_vulns = json.Marshal(pc.Critical_vulns)
		sess := x.NewSession()
		defer sess.Close()
		_, err := x.Insert(&user)
		if err != nil {
			return err
		}

	} else {
		err = errors.New("cannot write to databse, wrong driver")
	}
	return err
}
func GetSystemUserFromDatabase(x *xorm.Engine, login string) (SystemUser, error) {
	su := SystemUser{
		Login: login,
	}
	has, err := x.Get(&su)
	if err != nil {
		return SystemUser{}, err
	}
	if !has {
		return su, nil
	} else {
		return su, nil
	}

}
func GetUserFromDatabase(x *xorm.Engine, username string) (User, error) {
	u := User{
		Username: username,
	}
	has, err := x.Get(&u)
	if err != nil {
		return User{}, err
	}
	if !has {
		return u, nil
	} else {
		return u, nil
	}
}
