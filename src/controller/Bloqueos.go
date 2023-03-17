package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/global"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarBlqueo(c *gin.Context) {
	var data models.Bloqueos

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	if verificarSiYaTieneBloqueo(data.Anio, data.Mes, data.Dia, data.Identificacion) == false {
		sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO bloqueo (identificacion, dia, mes, anio, estado, hora) VALUES (?,?,?,?,?,?)")
		if err2 != nil {
			panic(err2)
		}

		sqlQ.Exec(data.Identificacion, data.Dia, data.Mes, data.Anio, 0, global.HoraActual())
	}

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func GetBloqueosAll(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")

	var d models.Bloqueos
	var datos []models.Bloqueos

	query := `select * from bloqueo where mes = ? and anio = ?`

	filas, err := conexion.SessionMysql.Query(query, mes, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueo, &d.Identificacion, &d.Dia, &d.Mes, &d.Anio, &d.Estado, &d.Hora)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetBloqueosAllPorEstado(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	estado := c.Param("estado")

	var d models.Bloqueos
	var datos []models.Bloqueos

	query := `select * from bloqueo where mes = ? and anio = ? and estado = ? order by idbloqueo desc`

	filas, err := conexion.SessionMysql.Query(query, mes, anio, estado)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueo, &d.Identificacion, &d.Dia, &d.Mes, &d.Anio, &d.Estado, &d.Hora)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetBloqueosAllPorFecha(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	dia := c.Param("dia")

	var d models.Bloqueos
	var datos []models.Bloqueos

	query := `select * from bloqueo where mes = ? and anio = ? and dia = ?`

	filas, err := conexion.SessionMysql.Query(query, mes, anio, dia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueo, &d.Identificacion, &d.Dia, &d.Mes, &d.Anio, &d.Estado, &d.Hora)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetBloqueosIdentificacionMesAnio(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	identificacion := c.Param("identificacion")

	var d models.Bloqueos
	var datos []models.Bloqueos

	query := `select * from bloqueo where mes = ? and anio = ? and identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, mes, anio, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueo, &d.Identificacion, &d.Dia, &d.Mes, &d.Anio, &d.Estado, &d.Hora)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func verificarSiYaTieneBloqueo(anio int, mes int, dia int, identificacion string) bool {
	var contador int = 0

	query := `select * from bloqueo where dia = ? and mes = ? and anio = ? and identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, dia, mes, anio, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
	}

	if contador > 0 {
		return true
	} else {
		return false
	}
}

func verificarSiElBloqueoEstaAutorizado(anio int, mes int, dia int, identificacion string) bool {
	var contador int = 0

	query := `select * from bloqueo where dia = ? and mes = ? and anio = ? and estado = 1 and identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, dia, mes, anio, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
	}

	if contador > 0 {
		return true
	} else {
		return false
	}
}

func verificarSiEstaPuntual() bool {
	horaSistema := global.NumHoraActual()
	minutoSistema := global.NumMinutoActual()
	horaBD := 0
	minutosBD := 0

	query := `SELECT hora, minuto FROM horaentrada LIMIT 1`
	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&horaBD, &minutosBD)
		if errsql != nil {
			panic(err)
		}
	}

	if horaSistema <= horaBD {
		if horaSistema == horaBD {
			if minutoSistema <= minutosBD {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func verificarSiEsPrimeraAsistenciaMarcacion(identificacion string) bool {
	var resultado bool = false
	var contador int = 0

	query := `SELECT IDASISTENCIA FROM asistencias WHERE IDENTIFICACION = ? AND FECHA = ?`
	filas, err := conexion.SessionMysql.Query(query, identificacion, global.FechaActual())
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
	}

	if contador > 0 {
		resultado = true
	}

	return resultado
}

func VerificarsiPuedeMarcarAsistencia(c *gin.Context) {
	var resultado string = ""
	identificacion := c.Param("identificacion")

	if verificarSiEsPrimeraAsistenciaMarcacion(identificacion) == true {
		resultado = "AUTORIZADO"
	} else {
		if verificarSiEstaPuntual() == true {
			resultado = "AUTORIZADO"
		} else {
			if verificarSiYaTieneBloqueo(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
				if verificarSiElBloqueoEstaAutorizado(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
					resultado = "AUTORIZADO"
				} else {
					resultado = "BLOQUEADO"
				}
			} else {
				sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO bloqueo (identificacion, dia, mes, anio, estado, hora) VALUES (?,?,?,?,?,?)")
				if err2 != nil {
					panic(err2)
				}

				sqlQ.Exec(identificacion, global.NumDiaActual(), global.NumMesActual(), global.NumAnioActual(), 0, global.HoraActual())
				resultado = "BLOQUEADO"
			}
		}
	}
	c.JSON(http.StatusCreated, gin.H{"response": resultado})
}

func AutorizarBloqueos(c *gin.Context) {
	idbloqueo := c.Param("idbloqueo")

	query, err2 := conexion.SessionMysql.Prepare("update bloqueo set estado = 1 where idbloqueo = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(idbloqueo)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}
