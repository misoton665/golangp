package main

import (
	"time"
)

type IbukiUser struct {
	Id        int32  `db:"pk"`
	Name      string `db:"unique" size:"255"`
	Email     string `db:"unique" size:"255"`
	CreatedAt time.Time
}

type Password struct {
	Id        int32 `db:"pk"`
	UserId    int32
	Password  string `size:"255"`
	CreatedAt time.Time
}
