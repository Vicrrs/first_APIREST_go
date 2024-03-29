package main

import "fmt"

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

func main() {
	fmt.Println("Hello, GO!")
}
