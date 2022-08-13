package Database

import (
	"fmt"
	"github.com/MuriloRibg/API-Rest_Gin.git/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectarBanco() {
	conectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/estudos?charset=utf8mb4&parseTime=True&loc=Local", "delta", "#Murilo1202", "aquarius-knight-db.mysql.database.azure.com")
	DB, err = gorm.Open(mysql.Open(conectString))
	if err != nil {
		panic(err.Error())
	}

	//Criando as tabelas
	err := DB.AutoMigrate(&Models.Aluno{})
	if err != nil {
		panic(err.Error())
	}
}
