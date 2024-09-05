package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Definir rota para abrir um chamado
	e.POST("/chamados", criarChamado)

	// Iniciar o servidor
	e.Logger.Fatal(e.Start(":8080"))
}

func criarChamado(c echo.Context) error {
	// Lógica para criar um chamado (via WhatsApp ou E-mail)
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Chamado criado com sucesso!",
	})
}
