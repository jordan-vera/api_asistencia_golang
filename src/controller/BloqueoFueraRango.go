package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/global"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarBlqueoFueraRango(c *gin.Context) {
	var data models.Bloqueofuerarango

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	if verificarSiYaTieneBloqueoFueraRango(data.Anio, data.Mes, data.Dia, data.Identificacion, data.Tipobloqueo) == false {
		sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO bloquefuerarango (identificacion, tipobloqueo, dia, mes, anio, horab, estado, justificacion, autorizadorb) VALUES (?,?,?,?,?,?,?,?,?)")
		if err2 != nil {
			panic(err2)
		}

		sqlQ.Exec(data.Identificacion, data.Tipobloqueo, data.Dia, data.Mes, data.Anio, global.HoraActual(), 0, data.Identificacion, data.Autorizadorb)
	}

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func verificarSiYaTieneBloqueoFueraRango(anio int, mes int, dia int, identificacion string, tipo string) bool {
	var contador int = 0

	query := `SELECT idbloqueofuerarando FROM bloquefuerarango WHERE dia = ? AND mes = ? AND anio = ? AND identificacion = ? AND tipobloqueo = ?`

	filas, err := conexion.SessionMysql.Query(query, dia, mes, anio, identificacion, tipo)
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

func GetBloqueosFueraRangoAll(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")

	var d models.Bloqueofuerarango
	var datos []models.Bloqueofuerarango

	query := `SELECT 
	                idbloqueofuerarando, identificacion, tipobloqueo, dia, mes, anio, horab, estado, justificacion, autorizadorb 
			    FROM bloquefuerarango 
				WHERE mes = ? AND anio = ?`

	filas, err := conexion.SessionMysql.Query(query, mes, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueofuerarando, &d.Identificacion, &d.Tipobloqueo, &d.Dia, &d.Mes, &d.Anio, &d.Horab, &d.Estado, &d.Justificacion, &d.Autorizadorb)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetBloqueosFueraRanfoAllPorEstado(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	estado := c.Param("estado")

	var d models.Bloqueofuerarango
	var datos []models.Bloqueofuerarango

	query := `  SELECT 
	                idbloqueofuerarando, identificacion, tipobloqueo, dia, mes, anio, horab, estado, justificacion, autorizadorb 
	            FROM bloquefuerarango 
				WHERE mes = ? and anio = ? and estado = ? order by idbloqueofuerarando desc`

	filas, err := conexion.SessionMysql.Query(query, mes, anio, estado)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueofuerarando, &d.Identificacion, &d.Tipobloqueo, &d.Dia, &d.Mes, &d.Anio, &d.Horab, &d.Estado, &d.Justificacion, &d.Autorizadorb)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetBloqueosFueraRangoAllPorFecha(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	dia := c.Param("dia")

	var d models.Bloqueofuerarango
	var datos []models.Bloqueofuerarango

	query := `  SELECT 
	                idbloqueofuerarando, identificacion, tipobloqueo, dia, mes, anio, horab, estado, justificacion, autorizadorb
	            FROM bloquefuerarango 
				WHERE mes = ? AND anio = ? AND dia = ?`

	filas, err := conexion.SessionMysql.Query(query, mes, anio, dia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueofuerarando, &d.Identificacion, &d.Tipobloqueo, &d.Dia, &d.Mes, &d.Anio, &d.Horab, &d.Estado, &d.Justificacion, &d.Autorizadorb)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetBloqueosFueraRangoIdentificacionMesAnio(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")
	identificacion := c.Param("identificacion")

	var d models.Bloqueofuerarango
	var datos []models.Bloqueofuerarango

	query := `  SELECT 
	                idbloqueofuerarando, identificacion, tipobloqueo, dia, mes, anio, horab, estado, justificacion, autorizadorb 
				FROM bloquefuerarango WHERE mes = ? AND anio = ? AND identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, mes, anio, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idbloqueofuerarando, &d.Identificacion, &d.Tipobloqueo, &d.Dia, &d.Mes, &d.Anio, &d.Horab, &d.Estado, &d.Justificacion, &d.Autorizadorb)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func verificarSiElBloqueoEstaAutorizadoFueraRango(anio int, mes int, dia int, identificacion string, tipo string) bool {
	var contador int = 0

	query := `SELECT idbloqueofuerarando FROM bloquefuerarango WHERE dia = ? AND mes = ? AND anio = ? AND estado = 1 AND identificacion = ? AND tipobloqueo=?`

	filas, err := conexion.SessionMysql.Query(query, dia, mes, anio, identificacion, tipo)
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

func VerificarRangoSalidaAlmuerzo(c *gin.Context) {

}
