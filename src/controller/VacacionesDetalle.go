package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarVacacionesDetalle(c *gin.Context) {
	var data models.VacacionesDetalle
	var errorGeneral error = nil

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO vacacionesdetalles (idvacaciones, numerodia, mes, anio) VALUES (?,?,?,?)")
	if err2 != nil {
		errorGeneral = err2
	}

	_, errorr := sqlQ.Exec(data.Idvacaciones, data.Numerodia, data.Mes, data.Anio)
	if errorr != nil {
		errorGeneral = errorr
	}

	if errorGeneral != nil {
		c.JSON(http.StatusCreated, gin.H{"error": errorGeneral})
	} else {
		c.JSON(http.StatusCreated, gin.H{"response": "Hecho"})
	}
}

func GetVacacionesDetalleAll(c *gin.Context) {
	var contador int = 0
	var d models.VacacionesDetalleMesAnio
	var datos []models.VacacionesDetalleMesAnio

	anio := c.Param("anio")
	mes := c.Param("mes")

	query := `
	            SELECT 
				    vacaciones.idvacaciones, vacaciones.identificacion, vacaciones.cantidaddias, vacaciones.estado, vacacionesdetalles.numerodia, vacacionesdetalles.mes, vacacionesdetalles.anio
				FROM vacaciones 
				INNER JOIN vacacionesdetalles ON vacacionesdetalles.idvacaciones = vacaciones.idvacaciones 
				WHERE vacacionesdetalles.anio = ? AND vacacionesdetalles.mes = ?;`

	filas, err := conexion.SessionMysql.Query(query, anio, mes)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idvacaciones, &d.Identificacion, &d.Cantidaddias, &d.Estado, &d.Numerodia, &d.Mes, &d.Anio)
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

func GetVacacionesDetalleAllIdentificacion(c *gin.Context) {
	var contador int = 0
	var d models.VacacionesDetalleMesAnio
	var datos []models.VacacionesDetalleMesAnio

	anio := c.Param("anio")
	mes := c.Param("mes")
	identificacion := c.Param("identificacion")

	query := `
	            SELECT 
				    vacaciones.idvacaciones, vacaciones.identificacion, vacaciones.cantidaddias, vacaciones.estado, vacacionesdetalles.numerodia, vacacionesdetalles.mes, vacacionesdetalles.anio
				FROM vacaciones 
				INNER JOIN vacacionesdetalles ON vacacionesdetalles.idvacaciones = vacaciones.idvacaciones 
				WHERE vacacionesdetalles.anio = ? AND vacacionesdetalles.mes = ? AND vacaciones.identificacion = ?;`

	filas, err := conexion.SessionMysql.Query(query, anio, mes, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idvacaciones, &d.Identificacion, &d.Cantidaddias, &d.Estado, &d.Numerodia, &d.Mes, &d.Anio)
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

func EliminarDetallesVacaciones(c *gin.Context) {
	idvacaciones := c.Param("idvacaciones")

	query, err := conexion.SessionMysql.Prepare("DELETE FROM vacacionesdetalles WHERE idvacaciones = ?")
	if err != nil {
		panic(err)
	}

	query.Exec(idvacaciones)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "registro eliminado!"})
}

func EliminarDetallesVacacionesOne(c *gin.Context) {
	iddetallevacaciones := c.Param("iddetallevacaciones")

	query, err := conexion.SessionMysql.Prepare("DELETE FROM vacacionesdetalles WHERE iddetallevacaciones = ?")
	if err != nil {
		panic(err)
	}

	query.Exec(iddetallevacaciones)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "registro eliminado!"})
}
