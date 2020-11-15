package main

import (
	"log"
)

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) GetAge() int64 {
	return u.age
}

//GetGender abc
func (u *User) GetGender() bool {
	return u.gender
}
func (u *User) GetAddress() string {
	return u.address
}
func Bai7() {
	m := make(map[string]*User)
	u1 := User{name: "Phong", age: 22, gender: true, address: "HP"}
	u2 := User{name: "Kha", age: 22, gender: true, address: "ND"}
	u3 := User{name: "Tuan", age: 22, gender: true, address: "BG"}
	u4 := User{name: "Long", age: 22, gender: true, address: "HN"}
	u5 := User{name: "Huy", age: 22, gender: true, address: "HD"}
	u6 := User{name: "Huu", age: 22, gender: true, address: "NA"}
	u7 := User{name: "Quang", age: 22, gender: true, address: "HN"}
	u8 := User{name: "Kien", age: 22, gender: true, address: "TH"}
	u9 := User{name: "Chung", age: 22, gender: true, address: "HN"}
	u10 := User{name: "Tuan1", age: 22, gender: true, address: "PT"}
	m[u1.GetName()] = &u1
	m[u2.GetName()] = &u2
	m[u3.GetName()] = &u3
	m[u4.GetName()] = &u4
	m[u5.GetName()] = &u5
	m[u6.GetName()] = &u6
	m[u7.GetName()] = &u7
	m[u8.GetName()] = &u8
	m[u9.GetName()] = &u9
	m[u10.GetName()] = &u10

	slice := make([]User, 10)
	for i := range m {
		slice = append(slice, *m[i])
	}
	for _, v := range slice {
		log.Println(v)
	}
}
