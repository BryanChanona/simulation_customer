package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckListBookLong(ctx *gin.Context) {
	for {
		response, err := http.Get("http://localhost:8080/books/longPolling")

		if err != nil {
			fmt.Println("Error al hacer la petición:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		defer response.Body.Close()

		// Verificar el código de estado
		if response.StatusCode != http.StatusOK {
			fmt.Println("Error, código de estado:", response.StatusCode)
			time.Sleep(5 * time.Second)
			continue
		}

		// Leer el cuerpo de la respuesta
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error al leer el cuerpo:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Verificar si la respuesta está vacía
		if len(body) == 0 {
			fmt.Println("Respuesta vacía del servidor")
			time.Sleep(5 * time.Second)
			continue
		}

		// Definir la estructura esperada del JSON
		type Response struct {
			Message string `json:"message"`
			Total   int    `json:"total"`
		}

		var result Response

		// Intentar deserializar la respuesta JSON
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Error al parsear JSON:", err, "Respuesta recibida:", string(body))
			time.Sleep(5 * time.Second)
			continue
		}

		// Imprimir el mensaje recibido del servidor
		fmt.Println("Esto es lo que nos dice el server:", result.Message)

		// Esperar antes de hacer otra solicitud
		time.Sleep(10 * time.Second)
	}
}
