package main

import (
	"github.com/pequenojoohn/go-gin/database"
	"github.com/pequenojoohn/go-gin/models"
	"github.com/pequenojoohn/go-gin/routes"
)

func main() {
	database.ConnectedDataBase()
	models.Alunos = []models.Aluno{
		{Nome: "Jonathan Matos", CPF: "00000000000", RG: "4444444"},
		{Nome: "Ana", CPF: "00000000000", RG: "4444444"},
	}
	routes.HandleRequest()
}
