package controllers

import (
	"customer/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var arrayBooks []models.Book

func CheckListoBooks(ctx *gin.Context) {
	for {
		response, _ := http.Get("http://localhost:8080/books/")

		if response.StatusCode == http.StatusOK {

			body, err := io.ReadAll(response.Body)

			defer response.Body.Close()

			if err != nil {
				fmt.Println("Error al obtener el cuerpo de la respuesta que nos da el servidor")
				return
			}

			var result models.ResponseBook

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error", err)
				return
			}

			if len(arrayBooks) != len(result.Book) {
				fmt.Println("Hay cambios")

				arrayBooks = result.Book
			} else {
				fmt.Println("No hay cambios")
			}

		} else {
			fmt.Println("Error con código de estado: ", response.StatusCode)
		}

		time.Sleep(10 * time.Second)

	}

}
