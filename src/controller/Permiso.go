package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarPermiso(c *gin.Context) {
	var data models.Permisos

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO permisos (idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado) VALUES (?,?,?,?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	res, errorr := sqlQ.Exec(data.Idtipopermiso, data.Identificacion, data.Desde, data.Hasta, data.Motivo, data.Estadojefe, data.Fechasolicitud, data.Tiempoestimado)
	if errorr != nil {
		panic(errorr)
	}

	idpermiso, errId := res.LastInsertId()
	if errId != nil {

	}

	c.JSON(http.StatusCreated, gin.H{"response": idpermiso})
}

func GetCountPermiso(c *gin.Context) {
	identificacion := c.Param("identificacion")
	var contador int = 0

	query := `select count(*) from permisos where identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
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
		c.JSON(http.StatusCreated, gin.H{"response": contador})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func GetAllPermisos(c *gin.Context) {
	identificacion := c.Param("identificacion")
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo from permisos
	            inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	            where identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo)
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

func GetTipoPermiso(c *gin.Context) {
	var contador int = 0
	var d models.Tipopermisos
	var datos []models.Tipopermisos

	query := `select * from tipopermiso`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idtipopermiso, &d.Tipo)
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

func EliminarPermiso(c *gin.Context) {
	idpermiso := c.Param("idpermiso")

	query, err := conexion.SessionMysql.Prepare("DELETE FROM permisos WHERE idpermiso = ?")
	if err != nil {
		panic(err)
	}

	query.Exec(idpermiso)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "registro eliminado!"})
}

func GetAllPermisoFecha(c *gin.Context) {
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos
	identificacion := c.Param("identificacion")
	mes := c.Param("mes")
	anio := c.Param("anio")

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo  from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where identificacion = ? and mes = ? and anio = ?
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, identificacion, mes, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo)
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

func GetAllPermisosadmin(c *gin.Context) {
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo  from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where estadojefe = 'SOLICITADO'
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo)
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

func AutorizarPermiso(c *gin.Context) {
	idpermiso := c.Param("idpermiso")

	query, err2 := conexion.SessionMysql.Prepare("update permisos set estadojefe = 'AUTORIZADO' where idpermiso = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(idpermiso)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func GetAllPorestadoPermiso(c *gin.Context) {
	estadojefe := c.Param("estadojefe")
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo  from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where estadojefe = ?
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, estadojefe)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo)
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

func GetPermisosFiltro(c *gin.Context) {
	estadojefe := c.Param("estadojefe")
	identificacion := c.Param("identificacion")

	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos
	var query = ""

	if estadojefe != "vacio" && identificacion != "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo  from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = ? and identificacion = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, estadojefe, identificacion)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else if estadojefe != "vacio" && identificacion == "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo  from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, estadojefe)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else if estadojefe == "vacio" && identificacion != "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo  from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where identificacion = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, identificacion)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo)
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
