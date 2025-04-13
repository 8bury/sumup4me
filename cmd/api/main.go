package main

import (
	"log"

	"github.com/8bury/sumup4me/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Iniciando aplicação sumup4me...")

	api := echo.New()

	log.Println("Configurando API e dependências...")
	config.ConfigureApi(api)

	port := ":8080"
	log.Printf("Servidor iniciado na porta %s", port)
	api.Start(port)
}
