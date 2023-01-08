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

	user := "root"
	password := "@Murilo1202"
	host := "localhost"
	conectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/cursos?charset=utf8mb4&parseTime=True&loc=Local", user, password, host)
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
