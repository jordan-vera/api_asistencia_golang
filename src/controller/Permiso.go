package controller

import (
	"net/http"
	"strings"

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

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO permisos (idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, autorizador) VALUES (?,?,?,?,?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	res, errorr := sqlQ.Exec(data.Idtipopermiso, data.Identificacion, data.Desde, data.Hasta, data.Motivo, data.Estadojefe, data.Fechasolicitud, data.Tiempoestimado, "")
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

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador from permisos
	            inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	            where identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
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

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where identificacion = ?
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
		if errsql != nil {
			panic(err)
		}
		if verificarSiMesAnio(mes, anio, d.Fechasolicitud) == true {
			datos = append(datos, d)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func verificarSiMesAnio(mes string, anio string, fechasolicitud string) bool {
	var resultado bool = false
	if mes == "1" {
		mes = "01"
	} else if mes == "2" {
		mes = "02"
	} else if mes == "3" {
		mes = "03"
	} else if mes == "4" {
		mes = "04"
	} else if mes == "5" {
		mes = "05"
	} else if mes == "6" {
		mes = "06"
	} else if mes == "7" {
		mes = "07"
	} else if mes == "8" {
		mes = "08"
	} else if mes == "9" {
		mes = "09"
	}
	parteFechaSolicitud := strings.Split(fechasolicitud, "-")

	if parteFechaSolicitud[1] == mes && parteFechaSolicitud[0] == anio {
		resultado = true
	}
	return resultado
}

func GetAllPermisosadmin(c *gin.Context) {
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where estadojefe = 'SOLICITADO'
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
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
	usuario := c.Param("usuario")

	query, err2 := conexion.SessionMysql.Prepare("update permisos set estadojefe = 'AUTORIZADO', autorizador = ? where idpermiso = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(usuario, idpermiso)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func NegarPermiso(c *gin.Context) {
	idpermiso := c.Param("idpermiso")
	usuario := c.Param("usuario")

	query, err2 := conexion.SessionMysql.Prepare("update permisos set estadojefe = 'NEGADO', autorizador = ? where idpermiso = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(usuario, idpermiso)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func GetAllPorestadoPermiso(c *gin.Context) {
	estadojefe := c.Param("estadojefe")
	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where estadojefe = ?
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, estadojefe)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
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
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = ? and identificacion = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, estadojefe, identificacion)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else if estadojefe != "vacio" && identificacion == "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador  from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, estadojefe)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else if estadojefe == "vacio" && identificacion != "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where identificacion = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, identificacion)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
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

func GetPermisosParaCalcularVacaciones(c *gin.Context) {
	var contador int = 0
	var errorGeneral error = nil
	var d models.Permisos
	var datos []models.Permisos
	identificacion := c.Param("identificacion")

	if obtenerLaFechaUltimaVacaciones(identificacion) != "" {
		query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = 'AUTORIZADO' and identificacion = ? and str_to_date(left(desde,10), '%d-%m-%Y') > str_to_date(left(?,10), '%d-%m-%Y')`
		filas, err := conexion.SessionMysql.Query(query, identificacion, obtenerLaFechaUltimaVacaciones(identificacion))
		if err != nil {
			errorGeneral = err
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
			if errsql != nil {
				errorGeneral = err
			}
			datos = append(datos, d)
		}
	} else {
		query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador  from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = 'AUTORIZADO' and identificacion = ?`
		filas, err := conexion.SessionMysql.Query(query, identificacion)
		if err != nil {
			errorGeneral = err
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador)
			if errsql != nil {
				errorGeneral = err
			}
			datos = append(datos, d)
		}
	}

	if errorGeneral != nil {
		c.JSON(http.StatusCreated, gin.H{"error": errorGeneral})
	} else if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func obtenerLaFechaUltimaVacaciones(identificacion string) string {
	var contador int = 0
	var d models.Vacaciones
	query := `SELECT idvacaciones, identificacion, cantidaddias, fechainicio, fechafin, estado, anio FROM vacaciones WHERE identificacion = ? ORDER BY idvacaciones desc LIMIT 1`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idvacaciones, &d.Identificacion, &d.Cantidaddias, &d.Fechainicio, &d.Fechafin, &d.Estado, &d.Anio)
		if errsql != nil {
			panic(err)
		}
	}
	if contador > 0 {
		return d.Fechafin
	} else {
		return ""
	}
}
