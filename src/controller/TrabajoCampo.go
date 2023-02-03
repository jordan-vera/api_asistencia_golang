package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarTrabajoCampo(c *gin.Context) {
	var data models.Trabajocampo

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO trabajocampo (identificacion, fecha, comentario, dia, mes, anio) VALUES (?,?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	res, errorr := sqlQ.Exec(data.Identificacion, data.Fecha, data.Comentario, data.Dia, data.Mes, data.Anio)
	if errorr != nil {
		panic(errorr)
	}

	idcampo, errId := res.LastInsertId()
	if errId != nil {

	}

	c.JSON(http.StatusCreated, gin.H{"response": idcampo})
}

func GetAllTrabajoCampoPorIdentificacion(c *gin.Context) {
	var contador int = 0
	var d models.Trabajocampo
	var datos []models.Trabajocampo
	identificacion := c.Param("identificacion")
	mes := c.Param("mes")
	anio := c.Param("anio")

	query := `select * from trabajocampo where identificacion = ? and mes = ? and anio = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion, mes, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idcampo, &d.Identificacion, &d.Fecha, &d.Comentario, &d.Dia, &d.Mes, &d.Anio)
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

func GetAllTrabajoCampoFiltro(c *gin.Context) {
	var contador int = 0
	var d models.Trabajocampo
	var datos []models.Trabajocampo
	var query string = ""
	identificacion := c.Param("identificacion")
	mes := c.Param("mes")
	anio := c.Param("anio")

	if identificacion == "vacio" {
		query = `select * from trabajocampo where mes = ? and anio = ? order by idcampo desc`
		filas, err := conexion.SessionMysql.Query(query, mes, anio)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idcampo, &d.Identificacion, &d.Fecha, &d.Comentario, &d.Dia, &d.Mes, &d.Anio)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else {
		query = `select * from trabajocampo where identificacion = ? and mes = ? and anio = ? order by idcampo desc`
		filas, err := conexion.SessionMysql.Query(query, identificacion, mes, anio)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idcampo, &d.Identificacion, &d.Fecha, &d.Comentario, &d.Dia, &d.Mes, &d.Anio)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}
