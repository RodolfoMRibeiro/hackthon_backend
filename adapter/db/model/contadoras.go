package db_model

import (
	"hackthon/common/library"

	"gorm.io/gorm"
)

type Contadora struct {
	gorm.Model
	Cnpj        string
	RazaoSocial string
}

func (Contadora) TableName() string {
	return library.TB_CONTADORA
}
