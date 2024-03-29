package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura struct(sao como formularios) para album
type album struct {
	ID     string `json: "id"`
	Title  string `json: "title"`
	Artist string `json: "artist"`
	Year   string `json: "year"`
}

// Definindo slice( usadas para armazenar vários valores do mesmo tipo em uma única variável)
var albums = []album{
	{ID: "1", Title: "Sou frida", Artist: "Fabio Brazza", Year: "2023"},
	{ID: "2", Title: "Não é sério", Artist: "CBJR", Year: "2002"},
	{ID: "3", Title: "Manic", Artist: "Flash Dance", Year: "1980"},
}

/*
O contexto é a parte mais importante do gin. Ele nos permite passar variáveis (c) entre middleware,
gerenciar o fluxo, validar o JSON de uma solicitação e renderizar uma resposta JSON
*/
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
