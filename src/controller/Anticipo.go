package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarAnticipo(c *gin.Context) {
	var data models.Anticipos

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO solicitud_anticipos_sueldos (fecha, identificacion, cantidadanticipo, motivo_si_es_segundo, meses_a_deducir, anio, mes, dia, estodogerente) VALUES (?,?,?,?,?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	sqlQ.Exec(data.Fecha, data.Identificacion, data.Cantidadanticipo, data.Motivo_si_es_segundo, data.Meses_a_deducir, data.Anio, data.Mes, data.Dia, "PENDIENTE")
	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func GetAnticiposPendientes(c *gin.Context) {

	anio := c.Param("anio")

	var d models.Anticipos
	var datos []models.Anticipos

	query := `
	            select * from solicitud_anticipos_sueldos 
				where estodogerente = 'PENDIENTE' and anio = ?
				order by idanticipo desc`

	filas, err := conexion.SessionMysql.Query(query, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idanticipo, &d.Fecha, &d.Identificacion, &d.Cantidadanticipo, &d.Motivo_si_es_segundo, &d.Meses_a_deducir, &d.Anio, &d.Mes, &d.Dia, &d.Estodogerente)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetAnticiposPorEstadoGerente(c *gin.Context) {
	estado := c.Param("estado")
	anio := c.Param("anio")
	var d models.Anticipos
	var datos []models.Anticipos

	query := `
	            select * from solicitud_anticipos_sueldos 
				where estodogerente = ? and anio = ?
				order by idanticipo desc`

	filas, err := conexion.SessionMysql.Query(query, estado, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idanticipo, &d.Fecha, &d.Identificacion, &d.Cantidadanticipo, &d.Motivo_si_es_segundo, &d.Meses_a_deducir, &d.Anio, &d.Mes, &d.Dia, &d.Estodogerente)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetAnticiposPorIdentificacion(c *gin.Context) {

	var d models.Anticipos
	var datos []models.Anticipos
	identificacion := c.Param("identificacion")
	anio := c.Param("anio")

	query := `
	            select * from solicitud_anticipos_sueldos 
				where identificacion = ? and anio = ?
				order by idanticipo desc`

	filas, err := conexion.SessionMysql.Query(query, identificacion, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idanticipo, &d.Fecha, &d.Identificacion, &d.Cantidadanticipo, &d.Motivo_si_es_segundo, &d.Meses_a_deducir, &d.Anio, &d.Mes, &d.Dia, &d.Estodogerente)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetAnticiposPorIdentificacionMesAnio(c *gin.Context) {

	var d models.Anticipos
	var datos []models.Anticipos
	identificacion := c.Param("identificacion")
	mes := c.Param("mes")
	anio := c.Param("anio")

	query := `
	            select * from solicitud_anticipos_sueldos 
				where identificacion = ? and mes = ? and anio = ?
				order by idanticipo desc
			`

	filas, err := conexion.SessionMysql.Query(query, identificacion, mes, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idanticipo, &d.Fecha, &d.Identificacion, &d.Cantidadanticipo, &d.Motivo_si_es_segundo, &d.Meses_a_deducir, &d.Anio, &d.Mes, &d.Dia, &d.Estodogerente)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetNumerocliente(c *gin.Context) {

	var numeroCliente int
	identificacion := c.Param("identificacion")

	query := `
				SELECT numeroCliente FROM Personas.Persona
				INNER JOIN Clientes.Cliente ON Cliente.secuencialPersona = Persona.secuencial
				WHERE identificacion = @identificacion
			`

	filas, err := conexion.Session.Query(query, sql.Named("identificacion", identificacion))
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&numeroCliente)
		if errsql != nil {
			panic(err)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"response": numeroCliente})
}

func AutorizarAnticiposGerente(c *gin.Context) {
	idanticipo := c.Param("idanticipo")

	query, err2 := conexion.SessionMysql.Prepare("update solicitud_anticipos_sueldos set estodogerente = 'AUTORIZADO' where idanticipo = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(idanticipo)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func NegarAnticiposGerente(c *gin.Context) {
	idanticipo := c.Param("idanticipo")

	query, err2 := conexion.SessionMysql.Prepare("update solicitud_anticipos_sueldos set estodogerente = 'NEGADO' where idanticipo = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(idanticipo)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func EliminarAnticipo(c *gin.Context) {
	idanticipo := c.Param("idanticipo")

	query, err := conexion.SessionMysql.Prepare("DELETE FROM solicitud_anticipos_sueldos WHERE idanticipo = ?")
	if err != nil {
		panic(err)
	}

	query.Exec(idanticipo)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "registro eliminado!"})
}
