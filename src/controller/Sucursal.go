package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func GetSucursal(c *gin.Context) {
	var contador int = 0
	var d models.Sucursales
	var datos []models.Sucursales

	query := `SELECT * FROM sucursal`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDSUCURSAL, &d.SUCURSAL, &d.LONGITUD, &d.LATITUD)
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

func Getonesucursal(c *gin.Context) {
	idsucursal := c.Param("idsucursal")
	var contador int = 0
	var d models.Sucursales

	query := `SELECT * FROM SUCURSAL WHERE IDSUCURSAL = ?`
	filas, err := conexion.SessionMysql.Query(query, idsucursal)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDSUCURSAL, &d.SUCURSAL, &d.LONGITUD, &d.LATITUD)
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
