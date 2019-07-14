package main

import (
	"time"
	uuid "github.com/google/uuid"
)

func guuid() string {
	id := uuid.New()
	return id.String()

}

// User represents a blog user.
type User struct {
	ID        string    `json:"uuid"`
	Fname     string    `json:"fname"`
    Lname     string    `json:"lname"`
	Age       int       `json:"age"`
	Timestamp string `json:timestamp`
}

func NewUser(fname string, lname string, age int) User {
	return User{ID: guuid(), Fname: fname, Lname: lname, Age: age, Timestamp:time.Now().Format(time.RFC3339)}
}

type UserObject struct{
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Age       int       `json:"age"`
}

func UserMap(u *User) map[string]interface{} {
	var m = make(map[string]interface{})
	m["uuid"] = u.ID
	m["fname"] = u.Fname
	m["lname"] = u.Lname
	m["age"] = u.Age
	m["timestamp"] = u.Timestamp
	return m


}

// Users represents a list of Users of this blog.
type Users []User