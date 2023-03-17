package controller

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/global"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarMarcacion(c *gin.Context) {
	var data models.Marcaciones

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	if verificarSiMarcacionEsSeguida(data.IDASISTENCIA) == false {
		saveimage(data.FILE, data.IMAGEN)

		sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO marcaciones (IDASISTENCIA, HORA, TIPO, IDSUCURSAL, IMAGEN) VALUES (?,?,?,?,?)")
		if err2 != nil {
			panic(err2)
		}

		sqlQ.Exec(data.IDASISTENCIA, global.HoraActual(), data.TIPO, data.IDSUCURSAL, data.IMAGEN)
	}

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func verificarSiMarcacionEsSeguida(idasistencia int) bool {
	var respuesta bool = false
	var hora string = ""
	var contador int = 0
	query := `SELECT HORA FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC LIMIT 1`
	filas, err := conexion.SessionMysql.Query(query, idasistencia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&hora)
		if errsql != nil {
			log.Fatal(errsql)
		}
	}

	if contador > 0 {
		partesHora := strings.Split(hora, ":")
		partesHoraActual := strings.Split(global.HoraActual(), ":")

		if partesHora[0] == partesHoraActual[0] {
			if partesHora[1] == partesHoraActual[1] {
				respuesta = true
			} else {
				respuesta = false
			}
		} else {
			respuesta = false
		}

	} else {
		respuesta = false
	}

	return respuesta
}

func saveimage(archivo string, foto string) {
	file, err := base64.StdEncoding.DecodeString(archivo)
	if err != nil {
		panic(err)
	}

	err2 := ioutil.WriteFile(global.UrlImagenesMarcaciones+foto, file, 0644)
	if err2 != nil {
		panic(err2)
	}
}

func Getultimamarcacion(c *gin.Context) {
	idasistencia := c.Param("idasistencia")
	var contador int = 0
	var d models.Marcaciones

	query := `SELECT TIPO FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC LIMIT 1`
	filas, err := conexion.SessionMysql.Query(query, idasistencia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.TIPO)
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

func GetMarcacionesHoy(c *gin.Context) {
	idasistencia := c.Param("idasistencia")
	var contador int = 0
	var d models.Marcaciones
	var datos []models.Marcaciones

	query := `SELECT * FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC`

	filas, err := conexion.SessionMysql.Query(query, idasistencia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDMARCACION, &d.IDASISTENCIA, &d.HORA, &d.TIPO, &d.IDSUCURSAL, &d.IMAGEN)
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
