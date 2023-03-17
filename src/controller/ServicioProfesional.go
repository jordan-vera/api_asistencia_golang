package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func LoginServicios(c *gin.Context) {
	var contador int = 0

	usuario := c.Param("usuario")
	clave := c.Param("clave")

	u := models.ServiciosProfesionales{}

	query := `select * from serviciosprofecionales where usuario = ? and clave = MD5(?)`
	filas, errsql := conexion.SessionMysql.Query(query, usuario, clave)
	if errsql != nil {
		panic(errsql)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&u.Idservicio, &u.Nombres, &u.Usuario, &u.Clave, &u.Identificacion, &u.Idsucursal, &u.Estado)
		if errsql != nil {
			panic(errsql)
		}
	}

	if contador > 0 {
		token, err := createToken(2)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusAccepted, gin.H{"servicioprofesional": u, "token": token})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"error": "Contraseña Incorrecta!!"})
	}
}

func GetAllServiciosProfesionales(c *gin.Context) {
	var contador int = 0
	var d models.ServiciosProfesionales
	var datos []models.ServiciosProfesionales

	query := `select idservicio, nombres, usuario, identificacion, idsucursal from serviciosprofecionales where estado = 1`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idservicio, &d.Nombres, &d.Usuario, &d.Identificacion, &d.Idsucursal)
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
