package main

import "gorm.io/gorm"

type Todo struct {
	ID   int    `json:"id"`
	Note string `json:"note"`
	Done bool   `json:"done"`
	gorm.Model
}

func (t *Todo) TableName() string {
	return "m_todo"
}
