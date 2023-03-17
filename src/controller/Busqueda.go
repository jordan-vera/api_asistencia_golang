package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarBusqueda(c *gin.Context) {
	var data models.Busqueda

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO busqueda (identificacion1, identificacion2, fechahora) VALUES (?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	sqlQ.Exec(data.Identificacion1, data.Identificacion2, data.Fechahora)
	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}
