package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/global"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarAsistencia(c *gin.Context) {
	var data models.Asistencias

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO asistencias (IDENTIFICACION, FECHA, MES, ANIO, DIA, NOMBREDIA, JUSTIFICACION, HORASJUSTIFICADAS) VALUES (?,?,?,?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	sqlQ.Exec(data.IDENTIFICACION, global.FechaActual(), data.MES, data.ANIO, data.DIA, data.NOMBREDIA, data.JUSTIFICACION, data.HORASJUSTIFICADAS)
	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func Verificarsiexisteasistencia(c *gin.Context) {
	identificacion := c.Param("identificacion")
	fecha := c.Param("fecha")
	var contador int = 0
	var d models.Asistencias

	query := `SELECT IDASISTENCIA FROM asistencias WHERE IDENTIFICACION = ? AND FECHA = ?`
	filas, err := conexion.SessionMysql.Query(query, identificacion, fecha)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDASISTENCIA)
		if errsql != nil {
			log.Fatal(errsql)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": d})
	} else {
		c.JSON(http.StatusCreated, gin.H{"errors": "No hay datos"})
	}
}

func Getfechaasistencia(c *gin.Context) {
	idempleado := c.Param("identificacion")
	fecha := c.Param("fecha")
	var contador int = 0
	var d models.Asistencias

	query := `SELECT * FROM asistencias WHERE IDENTIFICACION = ? AND FECHA = ?`
	filas, err := conexion.SessionMysql.Query(query, idempleado, fecha)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDASISTENCIA, &d.IDENTIFICACION, &d.FECHA, &d.MES, &d.ANIO, &d.DIA, &d.NOMBREDIA, &d.JUSTIFICACION, &d.HORASJUSTIFICADAS)
		if errsql != nil {
			log.Fatal(errsql)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": d})
	} else {
		c.JSON(http.StatusCreated, gin.H{"errors": "No hay datos"})
	}
}

func GetAsistenciasMarcacionesAllEmpleados(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")

	var d models.EmpleadosAsistencias
	var datos []models.EmpleadosAsistencias

	query := `select identificacion, nombreUnido from Personas.Persona
	inner join Nomina.EMPLEADO on EMPLEADO.SECUENCIALPERSONANATURAL = Persona.secuencial
	where EMPLEADO.CODIGOESTADOEMPLEADO = 'A' 
	order by nombreUnido asc`

	filas, err := conexion.Session.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Identificacion, &d.NombreUnido)
		if errsql != nil {
			panic(err)
		}
		d.Asistencias = getAsistenciasPorEmpleado(d.Identificacion, mes, anio)
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func GetAsistenciasMarcacionesAllServiciosProfecionales(c *gin.Context) {
	mes := c.Param("mes")
	anio := c.Param("anio")

	var d models.ServiciosProfesionalesAsistencia
	var datos []models.ServiciosProfesionalesAsistencia

	query := `select * from serviciosprofecionales where estado = 1`

	filas, err := conexion.SessionMysql.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.Idservicio, &d.Nombres, &d.Usuario, &d.Clave, &d.Identificacion, &d.Idsucursal, &d.Estado)
		if errsql != nil {
			panic(err)
		}
		d.Asistencias = getAsistenciasPorEmpleado(d.Identificacion, mes, anio)
		datos = append(datos, d)
	}

	c.JSON(http.StatusCreated, gin.H{"response": datos})
}

func getAsistenciasPorEmpleado(identificacion string, mes string, anio string) []models.AsistenciasMarcaciones {
	var d models.AsistenciasMarcaciones
	var datos []models.AsistenciasMarcaciones

	query := `SELECT * FROM asistencias 
	          WHERE IDENTIFICACION = ? AND MES = ? AND ANIO = ? 
			  ORDER BY IDASISTENCIA DESC`
	filas, err := conexion.SessionMysql.Query(query, identificacion, mes, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.IDASISTENCIA, &d.IDENTIFICACION, &d.FECHA, &d.MES, &d.ANIO, &d.DIA, &d.NOMBREDIA, &d.JUSTIFICACION, &d.HORASJUSTIFICADAS)
		if errsql != nil {
			log.Fatal(errsql)
		}
		d.MARCACIONES = obtenerMarcaciones(d.IDASISTENCIA)
		//d.JUSTIFICACION = calcularHorasMarcaciones(obtenerMarcaciones(d.IDASISTENCIA), d.ANIO, d.MES, d.DIA)
		datos = append(datos, d)
	}
	return datos
}

func GetAsistenciaPorMesAnioEmpleado(c *gin.Context) {
	identificacion := c.Param("identificacion")
	mes := c.Param("mes")
	anio := c.Param("anio")
	var contador int = 0
	var d models.AsistenciasMarcaciones
	var datos []models.AsistenciasMarcaciones

	query := `SELECT * FROM asistencias WHERE IDENTIFICACION = ? AND MES = ? AND ANIO = ? order by IDASISTENCIA DESC`
	filas, err := conexion.SessionMysql.Query(query, identificacion, mes, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDASISTENCIA, &d.IDENTIFICACION, &d.FECHA, &d.MES, &d.ANIO, &d.DIA, &d.NOMBREDIA, &d.JUSTIFICACION, &d.HORASJUSTIFICADAS)
		if errsql != nil {
			log.Fatal(errsql)
		}
		d.MARCACIONES = obtenerMarcaciones(d.IDASISTENCIA)
		datos = append(datos, d)
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"errors": "No hay datos"})
	}
}

func obtenerMarcaciones(idasistencia int) []models.Marcaciones {
	var d models.Marcaciones
	var datos []models.Marcaciones

	query := `SELECT * FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION ASC`

	filas, err := conexion.SessionMysql.Query(query, idasistencia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&d.IDMARCACION, &d.IDASISTENCIA, &d.HORA, &d.TIPO, &d.IDSUCURSAL, &d.IMAGEN)
		if errsql != nil {
			panic(err)
		}
		datos = append(datos, d)
	}

	return datos
}

func JustificarAsistencia(c *gin.Context) {
	var data models.AsistenciasJustificacion

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	query, err2 := conexion.SessionMysql.Prepare("update asistencias set JUSTIFICACION = ?, HORASJUSTIFICADAS = ? where IDASISTENCIA = ?")
	if err2 != nil {
		panic(err2)
	}

	query.Exec(data.JUSTIFICACION, data.HORASJUSTIFICADAS, data.IDASISTENCIA)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func GetAsistenciaAllPorFecha(c *gin.Context) {
	fecha := c.Param("fecha")
	var contador int = 0
	var d models.Asistencias
	var datos []models.Asistencias

	query := `SELECT * FROM asistencias WHERE FECHA = ?`
	filas, err := conexion.SessionMysql.Query(query, fecha)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDASISTENCIA, &d.IDENTIFICACION, &d.FECHA, &d.MES, &d.ANIO, &d.DIA, &d.NOMBREDIA, &d.JUSTIFICACION, &d.HORASJUSTIFICADAS)
		if errsql != nil {
			log.Fatal(errsql)
		}
		datos = append(datos, d)
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"errors": "No hay datos"})
	}
}

/*

func calcularHorasMarcaciones(marcaciones []models.Marcaciones, anio int, mes int, dia int) string {
	var entrada = ""
	var salida = ""
	var diferencias = 0.00

	for i := 0; i < len(marcaciones); i++ {
		if marcaciones[i].TIPO == "ENTRADA" {
			entrada = marcaciones[i].HORA
		} else {
			salida = marcaciones[i].HORA
		}

		if entrada != "" && salida != "" {
			horaEntrada, _ := strconv.Atoi(strings.Split(entrada, ":")[0])
			horaSalida, _ := strconv.Atoi(strings.Split(salida, ":")[0])

			minutoEntrada, _ := strconv.Atoi(strings.Split(entrada, ":")[1])
			minutoSalida, _ := strconv.Atoi(strings.Split(salida, ":")[1])

			dateentrada := time.Date(anio, time.Month(mes), dia, horaEntrada, minutoEntrada, 0, 0, time.UTC)
			datesalida := time.Date(anio, time.Month(mes), dia, horaSalida, minutoSalida, 0, 0, time.UTC)
			fmt.Println(dateentrada)
			fmt.Println(datesalida)

			diferencia := datesalida.Sub(dateentrada)
			fmt.Println(diferencia)

			diferencias, eror := time.Parse("2006-01-02", diferencia.Hours())

			entrada = ""
			salida = ""
		}
	}

	return fmt.Sprintf("%f", diferencias)
}
*/
