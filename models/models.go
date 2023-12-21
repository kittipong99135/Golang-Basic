package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Employee_id string `json:"Employee_id" validate:"required,excludesall= !@#?.*"`
	Name        string `json:"Name" validate:"required"`
	Lastname    string `json:"Lastname" validate:"required"`
	Birthday    string `json:"Birthday" validate:"required"`
	Age         int    `json:"Age" validate:"required,numeric,excludesall= !_@#?.*"`
	Email       string `json:"Email" validate:"required,email,excludesall= !_@#?.*"`
	Tel         string `json:"Tel" validate:"required,len=10,numeric,excludesall= !_@#?.*"`
}

type UsersResult struct {
	Name      string `json:"Name" validate:"required"`
	Lastname  string `json:"Lastname" validate:"required"`
	Birthday  string `json:"Birthday" validate:"required"`
	Age       int    `json:"Age" validate:"required,numeric,excludesall= !_@#?.*"`
	Genertion string
}

type UsersReturn struct {
	Users          []UsersResult `json:"Users" validate:"required,excludesall= !@#?.*"`
	GenZ           int
	GenY           int
	GenX           int
	BabyBoomer     int
	G_I_Generation int
}
