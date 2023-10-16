package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarVacaciones(c *gin.Context) {
	var data models.Vacaciones
	var errorGeneral error = nil

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO vacaciones (identificacion, cantidaddias, estado, anio) VALUES (?,?,?,?)")
	if err2 != nil {
		errorGeneral = err2
	}

	res, errorr := sqlQ.Exec(data.Identificacion, data.Cantidaddias, data.Estado, data.Anio)
	if errorr != nil {
		errorGeneral = errorr
	}

	idvacaciones, errId := res.LastInsertId()
	if errId != nil {
		errorGeneral = errId
	}

	if errorGeneral != nil {
		c.JSON(http.StatusCreated, gin.H{"error": errorGeneral})
	} else {
		c.JSON(http.StatusCreated, gin.H{"response": idvacaciones})
	}
}

func verificarSiTieneVacacionesEsteAnio(identificacion string, anio int) bool {
	var contador int = 0

	query := `select count(*) from vacaciones where identificacion = ? and anio = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&contador)
		if errsql != nil {
			panic(err)
		}
	}

	if contador > 0 {
		return true
	} else {
		return false
	}
}

func GetVacacionesAll(c *gin.Context) {
	var contador int = 0
	var d models.Vacaciones
	var datos []models.Vacaciones

	anio := c.Param("anio")

	query := `SELECT idvacaciones, identificacion, cantidaddias, estado, anio FROM vacaciones WHERE anio = ?`

	filas, err := conexion.SessionMysql.Query(query, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idvacaciones, &d.Identificacion, &d.Cantidaddias, &d.Estado, &d.Anio)
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

func GetVacacionesPorIdentificacion(c *gin.Context) {
	var contador int = 0
	var d models.VacacionesDetalleFilter
	var datos []models.VacacionesDetalleFilter

	identificacion := c.Param("identificacion")

	query := `
	            SELECT 
				    idvacaciones, identificacion, cantidaddias, estado, anio 
				FROM vacaciones 
				WHERE identificacion = ?
				ORDER BY idvacaciones DESC`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idvacaciones, &d.Identificacion, &d.Cantidaddias, &d.Estado, &d.Anio)
		if errsql != nil {
			panic(err)
		}
		d.Detalle = obtenerDetalleVacaciones(d.Idvacaciones)
		datos = append(datos, d)
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func getnombreEmpleado(identificacion string) string {
	var nombre string = ""
	query := `SELECT nombreUnido FROM Personas.Persona WHERE identificacion = @identificacion`
	filas, err := conexion.Session.Query(query, sql.Named("identificacion", identificacion))
	if err != nil {
		panic(err)
	}
	for filas.Next() {
		errsql := filas.Scan(&nombre)
		if errsql != nil {
			panic(err)
		}
	}
	return nombre
}

func Updatevacaciones(c *gin.Context) {
	var data models.Vacaciones

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	query, err2 := conexion.SessionMysql.Prepare("update vacaciones set cantidaddias = ? where idvacaciones = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(data.Cantidaddias, data.Idvacaciones)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}
