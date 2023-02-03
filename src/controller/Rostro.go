package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func GetAllRostros(c *gin.Context) {
	var contador int = 0
	var d models.Rostros
	var datos []models.Rostros

	query := `select * from rostros`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDROSTROS, &d.IDENTIFICACION, &d.IMAGEN)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func Getonerostros(c *gin.Context) {
	identificacion := c.Param("identificacion")
	var contador int = 0
	var d models.Rostros

	query := `select * from rostros where IDENTIFICACION = ?`
	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDROSTROS, &d.IDENTIFICACION, &d.IMAGEN)
		if errsql != nil {
			log.Fatal(errsql)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": d})
	} else {
		c.JSON(http.StatusCreated, gin.H{"errors": "No hay datos"})
	}
}
