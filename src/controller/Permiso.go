package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/global"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarPermiso(c *gin.Context) {
	var data models.Permisos

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO permisos (idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso, anio) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	if data.Idtipopermiso == 19 {
		data.Escargovacaciones = 1
	}

	anio := strings.Split(data.Hasta, "-")[2]

	res, errorr := sqlQ.Exec(data.Idtipopermiso, data.Identificacion, data.Desde, data.Hasta, data.Motivo, data.Estadojefe, data.Fechasolicitud, data.Tiempoestimado, "", 0, data.Escargovacaciones, data.Horainiciopermiso, data.Horafinpermiso, anio)
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

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso, permisos.anio from permisos
	            inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	            where identificacion = ?
				order by permisos.idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso, &d.Anio)
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

func GetPermisoPorAnioMesDia(c *gin.Context) {
	var contador int = 0
	var d models.Permisos

	var data models.PermisosAnioMesDia

	fmt.Println(data)

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	query := `	SELECT 
					permisos.idpermiso, 
					permisos.idtipopermiso, 
					permisos.identificacion, 
					permisos.desde, 
					permisos.hasta, 
					permisos.motivo, 
					permisos.tiempoestimado, 
					permisos.horainiciopermiso, 
					permisos.horafinpermiso,
					detallepermiso.iddetallepermiso
				FROM permisos
				inner join detallepermiso on detallepermiso.idpermiso = permisos.idpermiso
				where numerodia = ? and mes = ? and detallepermiso.anio = ? and permisos.estadojefe = 'AUTORIZADO' and permisos.identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, data.Numerodia, data.Mes, data.Anio, data.Identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Tiempoestimado, &d.Horainiciopermiso, &d.Horafinpermiso, &d.Iddetallepermiso)
		if errsql != nil {
			panic(err)
		}
	}

	if contador > 0 {
		var esUltimoDiaPermiso bool = false
		horasTiempoEstimado := global.ConvertirAhInt(strings.Split(d.Tiempoestimado, " ")[3])
		diasTiempoEstimado := global.ConvertirAhInt(strings.Split(d.Tiempoestimado, " ")[0])
		if diasTiempoEstimado > 0 && horasTiempoEstimado > 0 {
			// consultar si es el ultimo el el primero
			esUltimoDiaPermiso = verificarPosicionDetallePermisos(d.Idpermiso, d.Iddetallepermiso)
		}
		c.JSON(http.StatusCreated, gin.H{"response": d, "esultimodiapermiso": esUltimoDiaPermiso})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func verificarPosicionDetallePermisos(idpermiso int, iddetallepermiso int) bool {
	var data models.DetallepermisosVerificar
	var datas []models.DetallepermisosVerificar
	query := `SELECT iddetallepermiso FROM detallepermiso where idpermiso =?
	order by numerodia asc`
	filas, err := conexion.SessionMysql.Query(query, idpermiso)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&data.Iddetallepermiso)
		if errsql != nil {
			panic(err)
		}
		datas = append(datas, data)
	}

	cantidadFilas := len(datas)
	var contadorFilas int = 0
	for i := 0; i < cantidadFilas; i++ {
		contadorFilas++
		if datas[i].Iddetallepermiso == iddetallepermiso {
			break
		}
	}
	if contadorFilas == cantidadFilas {
		return true
	} else {
		return false
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

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where identificacion = ?
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
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

	anio := c.Param("anio")

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where estadojefe = 'SOLICITADO' and permisos.anio = ?
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
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

func EditEsCargoVacaciones(c *gin.Context) {
	idpermiso := c.Param("idpermiso")
	escargovacaciones := c.Param("escargovacaciones")

	query, err2 := conexion.SessionMysql.Prepare("update permisos set escargovacaciones = ? where idpermiso = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(escargovacaciones, idpermiso)

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

	query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
	inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
	where estadojefe = ?
	order by idpermiso desc`

	filas, err := conexion.SessionMysql.Query(query, estadojefe)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
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
	anio := c.Param("anio")

	var contador int = 0
	var d models.Permisos
	var datos []models.Permisos
	var query = ""

	if estadojefe != "vacio" && identificacion != "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = ? and identificacion = ? and permisos.anio = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, estadojefe, identificacion, anio)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else if estadojefe != "vacio" && identificacion == "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = ? and permisos.anio = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, estadojefe, anio)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else if estadojefe == "vacio" && identificacion != "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where identificacion = ? and permisos.anio = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, identificacion, anio)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
			if errsql != nil {
				panic(err)
			}
			datos = append(datos, d)
		}
	} else if estadojefe == "vacio" && identificacion == "vacio" {
		query = `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where permisos.anio = ?
		order by idpermiso desc`
		filas, err := conexion.SessionMysql.Query(query, anio)
		if err != nil {
			panic(err)
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
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
		query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = 'AUTORIZADO' and identificacion = ? and str_to_date(left(desde,10), '%d-%m-%Y') > str_to_date(left(?,10), '%d-%m-%Y')`
		filas, err := conexion.SessionMysql.Query(query, identificacion, obtenerLaFechaUltimaVacaciones(identificacion))
		if err != nil {
			errorGeneral = err
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
			if errsql != nil {
				errorGeneral = err
			}
			datos = append(datos, d)
		}
	} else {
		query := `select idpermiso, permisos.idtipopermiso, identificacion, desde, hasta, motivo, estadojefe, fechasolicitud, tiempoestimado, tipo, autorizador, calculadoenvacaciones, escargovacaciones, horainiciopermiso, horafinpermiso from permisos 
		inner join tipopermiso on tipopermiso.idtipopermiso = permisos.idtipopermiso
		where estadojefe = 'AUTORIZADO' and identificacion = ?`
		filas, err := conexion.SessionMysql.Query(query, identificacion)
		if err != nil {
			errorGeneral = err
		}
		for filas.Next() {
			contador++
			errsql := filas.Scan(&d.Idpermiso, &d.Idtipopermiso, &d.Identificacion, &d.Desde, &d.Hasta, &d.Motivo, &d.Estadojefe, &d.Fechasolicitud, &d.Tiempoestimado, &d.Tipo, &d.Autorizador, &d.Calculadoenvacaciones, &d.Escargovacaciones, &d.Horainiciopermiso, &d.Horafinpermiso)
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
	var d models.VacacionesDetalle
	query := `
			SELECT vacacionesdetalles.numerodia, vacacionesdetalles.mes, vacacionesdetalles.anio FROM vacacionesdetalles 
		    inner join vacaciones on vacaciones.idvacaciones = vacacionesdetalles.idvacaciones
		    WHERE vacaciones.identificacion = ? LIMIT 1
	`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Numerodia, &d.Mes, &d.Anio)
		if errsql != nil {
			panic(err)
		}
	}
	if contador > 0 {
		return fmt.Sprintf("%d-%d-%d", d.Numerodia, d.Mes, d.Anio)
	} else {
		return ""
	}
}
