package mapping

import (
	"WebProject/ginEssential/utils/ContactDb"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	Id        string
	Name      string
	Age       int
	Telephone string
	db        *sqlx.DB
}

var user *User

func GetUser() *User {

	if user != nil {
		return user
	}
	user = &User{db: ContactDb.DB}
	return user
}

func (this User) TelephonIsexist(tele string) User {
	sqlStr := "select id,name,age,telephone from user where telephone = ?"

	this.db.Get(&this, sqlStr, tele)

	return this
}

func (this User) Login(pwd, tele string) *User {

	sqlStr := "select id,name,age,telephone from user where pwd = ? and telephone = ? "

	this.db.Get(&this, sqlStr, pwd, tele)

	return &this
}

func (this User) AddUser(name, pwd, telephone string) bool {

	sqlStr := "insert into user(name , pwd , telephone) values (?,?,?)"

	_, err := this.db.Exec(sqlStr, name, pwd, telephone)

	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func (this *User) FindById(id string) (*User, bool) {

	sqlStr := "select id,name,age,telephone from user where id = ?"
	this.db.Get(this, sqlStr, id)
	if len(this.Id) == 0 {
		return this, false
	}
	return this, true
}
