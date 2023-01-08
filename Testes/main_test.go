package Testes

import (
	"bytes"
	"encoding/json"
	"github.com/MuriloRibg/API-Rest_Gin.git/Controllers"
	"github.com/MuriloRibg/API-Rest_Gin.git/Database"
	"github.com/MuriloRibg/API-Rest_Gin.git/Models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

var ID int

func CriaALunoMock() {
	aluno := Models.Aluno{CPF: "12312312312", RG: "123123123", Nome: "Murilo Ribeiro"}
	Database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno Models.Aluno
	Database.DB.Delete(&aluno, ID)
}

func TestVerificarStatusCodeListarAlunos(t *testing.T) {
	Database.ConectarBanco()
	CriaALunoMock()
	defer DeletaAlunoMock() //ao finalizar o programa

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", Controllers.Listar)

	//criando o request
	request, _ := http.NewRequest("GET", "/alunos", nil)

	//armazena o response
	response := httptest.NewRecorder()

	//Fazendo a requisição
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestRecuperarAlunoPorCpf(t *testing.T) {
	Database.ConectarBanco()
	CriaALunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/aluno/cpf/:cpf", Controllers.RecuperPorCPF)

	request, _ := http.NewRequest("GET", "/aluno/cpf/12312312312", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestRecuperarAlunoId(t *testing.T) {
	Database.ConectarBanco()
	CriaALunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/aluno/:id", Controllers.Recuperar)

	path := "/aluno/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var alunoMock Models.Aluno

	//converte os bytes para um json
	if err := json.Unmarshal(response.Body.Bytes(), &alunoMock); err != nil {
		panic(err.Error())
	}

	assert.Equal(t, "12312312312", alunoMock.CPF)
	assert.Equal(t, "123123123", alunoMock.RG)
	assert.Equal(t, "Murilo Ribeiro", alunoMock.Nome)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeletarALuno(t *testing.T) {
	Database.ConectarBanco()
	CriaALunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/aluno/:id", Controllers.Deletar)

	path := "/aluno/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditarAluno(t *testing.T) {
	Database.ConectarBanco()
	CriaALunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PUT("/aluno/:id", Controllers.Editar)

	aluno := Models.Aluno{CPF: "78945612374", RG: "123456789", Nome: "Murilo R"}
	varJson, _ := json.Marshal(aluno) //passando para json
	path := "/aluno/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("PUT", path, bytes.NewBuffer(varJson)) //passando para bytes
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var alunoMock Models.Aluno
	if err := json.Unmarshal(response.Body.Bytes(), &alunoMock); err != nil {
		panic(err.Error())
	}

	assert.Equal(t, "78945612374", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, "Murilo R", alunoMock.Nome)
	assert.Equal(t, http.StatusOK, response.Code)
}
