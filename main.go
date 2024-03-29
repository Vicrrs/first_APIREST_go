package main

import "fmt"

// Estrutura struct(sao como formularios) para album
type album struct {
	ID     string `json: "id"`
	Title  string `json: "title"`
	Artist string `json: "artist"`
	Year   string `json: "year"`
}

func main() {
	fmt.Println("Hello, GO!")
}
