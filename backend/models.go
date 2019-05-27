package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Llave struct {
	gorm.Model
	IdLlave      int64 `gorm:"primary_key"`
	NombreLlave  string
	LlavePublica []byte
	LlavePrivada []byte
}

type Entrada struct {
	IdLlave    string
	TextoPlano string
}
