package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarDetallePermiso(c *gin.Context) {
	var data models.Detallepermisos

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO detallepermiso (idpermiso, numerodia, mes, anio) VALUES (?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	sqlQ.Exec(data.Idpermiso, data.Numerodia, data.Mes, data.Anio)
	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func GetDetallePermisos(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos

	query := `select permisos.idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, numerodia, tipopermiso.tipo from permisos
	          inner join detallepermiso on detallepermiso.idpermiso = permisos.idpermiso
			  inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
			  where detallepermiso.mes = ? and detallepermiso.anio = ? and estadojefe = ? `

	filas, err := conexion.SessionMysql.Query(query, mes, anio, "AUTORIZADO")
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Numerodia, &d.Tipo)
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

func GetDetallePermisosFecha(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	dia := c.Param("dia")
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos

	query := `select permisos.idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, numerodia, tipopermiso.tipo from permisos
	          inner join detallepermiso on detallepermiso.idpermiso = permisos.idpermiso
			  inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
			  where detallepermiso.mes = ? and detallepermiso.anio = ? and numerodia = ?`

	filas, err := conexion.SessionMysql.Query(query, mes, anio, dia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Numerodia, &d.Tipo)
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
