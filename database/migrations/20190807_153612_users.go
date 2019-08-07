package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20190807_153612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20190807_153612{}
	m.Created = "20190807_153612"

	migration.Register("Users_20190807_153612", m)
}

// Run the migrations
func (m *Users_20190807_153612) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(`id` int(11) DEFAULT NULL,`username` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,`password` varchar(128) NOT NULL,`avatar` varchar(128) NOT NULL,`address` varchar(128) NOT NULL)")
}

// Reverse the migrations
func (m *Users_20190807_153612) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
}
