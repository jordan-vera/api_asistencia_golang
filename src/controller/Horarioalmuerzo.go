package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarHorarioAlmuerzo(c *gin.Context) {
	var data models.Horarioalmuerzo
	var errorGeneral error = nil

	err := c.ShouldBindJSON(&data)
	if err != nil {
		errorGeneral = err
	}

	if verificarSiTieneHorarioRegistrado(data.Identificacion) == false {

		sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO horarioalmuerzo (entrada, salida, salida_a_casa, identificacion) VALUES (?,?,?,?)")
		if err2 != nil {
			errorGeneral = err2
		}

		sqlQ.Exec(data.Entrada, data.Salida, data.Salidaacasa, data.Identificacion)
	}

	if errorGeneral != nil {
		c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": errorGeneral})
	}
}

func GetHorasAlmuerzo(c *gin.Context) {
	var contador int = 0
	var d models.HorarioalmuerzoColaborador
	var datos []models.HorarioalmuerzoColaborador

	query := `SELECT idhorarioalmuerzo, entrada, salida, salida_a_casa, identificacion FROM horarioalmuerzo`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idhorarioalmuerzo, &d.Entrada, &d.Salida, &d.Salidaacasa, &d.Identificacion)
		if errsql != nil {
			panic(err)
		}
		if buscarEmpleado(d.Identificacion) != "" {
			contador++
			d.Nombres = buscarEmpleado(d.Identificacion)
			datos = append(datos, d)
		} else if buscarServicios(d.Identificacion) != "" {
			contador++
			d.Nombres = buscarServicios(d.Identificacion)
			datos = append(datos, d)
		}

	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func buscarEmpleado(identificacion string) string {
	var nombre string = ""
	query := `SELECT nombreUnido FROM Personas.Persona
	INNER JOIN Nomina.EMPLEADO ON EMPLEADO.SECUENCIALPERSONANATURAL = Persona.secuencial
	WHERE EMPLEADO.CODIGOESTADOEMPLEADO = 'A' AND identificacion = @identificacion`

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

func buscarServicios(identificacion string) string {
	var nombre string = ""
	query := `SELECT nombres FROM serviciosprofecionales WHERE identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
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

func GetHorasAlmuerzoPorIdentificacion(c *gin.Context) {
	var contador int = 0
	var d models.Horarioalmuerzo
	var datos []models.Horarioalmuerzo
	identificacion := c.Param("identificacion")

	query := `SELECT idhorarioalmuerzo, entrada, salida, salida_a_casa, identificacion FROM horarioalmuerzo WHERE identificacion = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idhorarioalmuerzo, &d.Entrada, &d.Salida, &d.Salidaacasa, &d.Identificacion)
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

func verificarSiTieneAsignadoHorarioAlmuerzo(identificacion string) bool {
	var contador int = 0
	query := `SELECT count(*) FROM horarioalmuerzo where identificacion = ?`
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
		return true
	} else {
		return false
	}
}

func GetHorasAlmuerzoPorEntrada(c *gin.Context) {
	var contador int = 0
	var d models.Horarioalmuerzo
	var datos []models.Horarioalmuerzo
	entrada := c.Param("entrada")

	query := `SELECT idhorarioalmuerzo, entrada, salida, salida_a_casa, identificacion FROM horarioalmuerzo WHERE entrada = ?`

	filas, err := conexion.SessionMysql.Query(query, entrada)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Idhorarioalmuerzo, &d.Entrada, &d.Salida, &d.Salidaacasa, &d.Identificacion)
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

func verificarSiTieneHorarioRegistrado(identificacion string) bool {

	var contador int = 0

	query := `SELECT count(*) FROM horarioalmuerzo WHERE identificacion = ?`

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
		return true
	} else {
		return false
	}

}

func ActualizarHorario(c *gin.Context) {
	var data models.Horarioalmuerzo
	var errorGeneral error = nil

	err := c.ShouldBindJSON(&data)
	if err != nil {
		errorGeneral = err
	}

	query, err2 := conexion.SessionMysql.Prepare("update horarioalmuerzo set entrada = ?, salida = ?, salida_a_casa = ? where idhorarioalmuerzo = ?")
	if err2 != nil {
		errorGeneral = err2
	}

	query.Exec(data.Entrada, data.Salida, data.Salidaacasa, data.Idhorarioalmuerzo)

	if errorGeneral != nil {
		c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": errorGeneral})
	}
}
