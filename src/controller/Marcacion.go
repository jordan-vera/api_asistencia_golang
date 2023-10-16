package controller

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/global"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func AgregarMarcacion(c *gin.Context) {
	var data models.Marcaciones

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	if verificarSiMarcacionEsSeguida(data.IDASISTENCIA) == false {
		saveimage(data.FILE, data.IMAGEN)

		sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO marcaciones (IDASISTENCIA, HORA, TIPO, IDSUCURSAL, IMAGEN) VALUES (?,?,?,?,?)")
		if err2 != nil {
			panic(err2)
		}

		sqlQ.Exec(data.IDASISTENCIA, global.HoraActual(), data.TIPO, data.IDSUCURSAL, data.IMAGEN)
	}

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func NuevaMarcacionParaEdit(c *gin.Context) {
	var data models.Marcaciones

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO marcaciones (IDASISTENCIA, HORA, TIPO, IDSUCURSAL, IMAGEN) VALUES (?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	sqlQ.Exec(data.IDASISTENCIA, data.HORA, data.TIPO, data.IDSUCURSAL, data.IMAGEN)

	c.JSON(http.StatusCreated, gin.H{"response": "hecho"})
}

func verificarSiMarcacionEsSeguida(idasistencia int) bool {
	var respuesta bool = false
	var hora string = ""
	var contador int = 0
	query := `SELECT HORA FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC LIMIT 1`
	filas, err := conexion.SessionMysql.Query(query, idasistencia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&hora)
		if errsql != nil {
			log.Fatal(errsql)
		}
	}

	if contador > 0 {
		partesHora := strings.Split(hora, ":")
		partesHoraActual := strings.Split(global.HoraActual(), ":")

		if partesHora[0] == partesHoraActual[0] {
			if partesHora[1] == partesHoraActual[1] {
				respuesta = true
			} else {
				respuesta = false
			}
		} else {
			respuesta = false
		}

	} else {
		respuesta = false
	}

	return respuesta
}

func saveimage(archivo string, foto string) {
	file, err := base64.StdEncoding.DecodeString(archivo)
	if err != nil {
		panic(err)
	}

	err2 := ioutil.WriteFile(global.UrlImagenesMarcaciones+foto, file, 0644)
	if err2 != nil {
		panic(err2)
	}
}

func Getultimamarcacion(c *gin.Context) {
	idasistencia := c.Param("idasistencia")
	var contador int = 0
	var d models.Marcaciones

	query := `SELECT TIPO FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC LIMIT 1`
	filas, err := conexion.SessionMysql.Query(query, idasistencia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.TIPO)
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

func GetMarcacionesHoy(c *gin.Context) {
	idasistencia := c.Param("idasistencia")
	var contador int = 0
	var d models.Marcaciones
	var datos []models.Marcaciones

	query := `SELECT * FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC`

	filas, err := conexion.SessionMysql.Query(query, idasistencia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.IDMARCACION, &d.IDASISTENCIA, &d.HORA, &d.TIPO, &d.IDSUCURSAL, &d.IMAGEN)
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

func RealizarMarcacionGeneral(c *gin.Context) {
	var resultadoHttp string = ""
	var data models.Marcaciones
	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	identificacion := c.Param("identificacion")

	// verifica si tiene creada una asistencia
	if verificarSiEsPrimeraAsistenciaMarcacion(identificacion) == true {
		idasistencia := obtenerIdAsistencia(identificacion)
		tipoMarcacion := verificarQueTipoMarcacionCorresponde(idasistencia)

		if global.NombreDia() == "sábado" || global.NombreDia() == "domingo" {
			if tipoMarcacion == "Error" {
				if verificarSiYaTieneBloqueo(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
					if verificarSiElBloqueoEstaAutorizado(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
						crearMarcacion(identificacion, "ENTRADA", data.IDSUCURSAL, data.IMAGEN, data.FILE)
						resultadoHttp = "HECHO"
					} else {
						resultadoHttp = "BLOQUEADO"
					}
				}
			} else {
				if tipoMarcacion == "SALIDA-DEL-ALMUERZO" {
					crearMarcacion(identificacion, "SALIDA", data.IDSUCURSAL, data.IMAGEN, data.FILE)
					resultadoHttp = "HECHO"
				} else {
					resultadoHttp = "YA TIENES TODAS LAS MARCACIONES"
				}

			}
		} else {
			if tipoMarcacion == "Error" {
				if verificarSiYaTieneBloqueo(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
					if verificarSiElBloqueoEstaAutorizado(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
						crearMarcacion(identificacion, "ENTRADA", data.IDSUCURSAL, data.IMAGEN, data.FILE)
						resultadoHttp = "HECHO"
					} else {
						resultadoHttp = "BLOQUEADO"
					}
				}
			} else {
				tipo := ""
				if tipoMarcacion == "SALIDA-DEL-ALMUERZO" {
					tipo = "SALIDA"
				} else if tipoMarcacion == "ENTRADA-DEL-ALMUERZO" {
					tipo = "ENTRADA"
				} else if tipoMarcacion == "SALIDA-A-CASA" {
					tipo = "SALIDA"
				}

				// si tiene horario de almuerzo

				if verificarSiTieneHorarioAlmuerzoParametrizado(identificacion) {
					horarioalmuerzoOne := obtenerHorarioDelAlmuerzo(identificacion)
					estaDentroDelRango := calcularSiEstaEnElRangoHorario(tipoMarcacion, horarioalmuerzoOne)
					if verificarQueTipoMarcacionCorresponde(idasistencia) == "MARCACIONES-COMPLETAS" {
						resultadoHttp = "YA TIENES TODAS LAS MARCACIONES"
					} else if estaDentroDelRango == true {
						crearMarcacion(identificacion, tipo, data.IDSUCURSAL, data.IMAGEN, data.FILE)
						resultadoHttp = "HECHO"
					} else {
						resultadoHttp = "MARCACIÓN FUERA DEL RANGO DEL HORARIO ESTABLECIDO"
					}
				} else {
					if tipo != "" {
						crearMarcacion(identificacion, tipo, data.IDSUCURSAL, data.IMAGEN, data.FILE)
						resultadoHttp = "HECHO"
					} else {
						resultadoHttp = "SOLO SE PERMITEN 4 MARCACIONES"
					}
				}
			}
		}
	} else {
		//crear asistencia y la primera marcacion, verificando que sea puntual
		crearAsistencia(identificacion)
		if verificarSiEstaPuntual() == true {
			crearMarcacion(identificacion, "ENTRADA", data.IDSUCURSAL, data.IMAGEN, data.FILE)
			resultadoHttp = "HECHO"
		} else {
			if verificarSiYaTieneBloqueo(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
				if verificarSiElBloqueoEstaAutorizado(global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual(), identificacion) == true {
					crearMarcacion(identificacion, "ENTRADA", data.IDSUCURSAL, data.IMAGEN, data.FILE)
					resultadoHttp = "HECHO"
				} else {
					resultadoHttp = "BLOQUEADO"
				}
			} else {
				sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO bloqueo (identificacion, dia, mes, anio, estado, hora, autorizador) VALUES (?,?,?,?,?,?,?)")
				if err2 != nil {
					panic(err2)
				}

				sqlQ.Exec(identificacion, global.NumDiaActual(), global.NumMesActual(), global.NumAnioActual(), 0, global.HoraActual(), "")
				resultadoHttp = "BLOQUEADO"
			}
		}

	}
	c.JSON(http.StatusCreated, gin.H{"response": resultadoHttp})
}

func crearAsistencia(identificacion string) {
	sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO asistencias (IDENTIFICACION, FECHA, MES, ANIO, DIA, NOMBREDIA, JUSTIFICACION, HORASJUSTIFICADAS) VALUES (?,?,?,?,?,?,?,?)")
	if err2 != nil {
		panic(err2)
	}

	sqlQ.Exec(identificacion, global.FechaActual(), global.NumMesActual(), global.NumAnioActual(), global.NumDiaActual(), global.NombreDia(), "", 0)
}

func crearMarcacion(identificacion string, tipo string, idsucursal int, imagen string, file string) {
	idasistencia := obtenerIdAsistencia(identificacion)
	if verificarSiMarcacionEsSeguida(idasistencia) == false {
		saveimage(file, imagen)

		sqlQ, err2 := conexion.SessionMysql.Prepare("INSERT INTO marcaciones (IDASISTENCIA, HORA, TIPO, IDSUCURSAL, IMAGEN) VALUES (?,?,?,?,?)")
		if err2 != nil {
			panic(err2)
		}

		sqlQ.Exec(idasistencia, global.HoraActual(), tipo, idsucursal, imagen)
	}
}

func calcularSiEstaEnElRangoHorario(tipo string, horarioalmuerzo models.Horarioalmuerzo) bool {
	var result bool = false
	if tipo == "SALIDA-DEL-ALMUERZO" {
		diferencia := global.CalcularHora(horarioalmuerzo.Salida)
		if global.EsPositivoNeutro(diferencia) == false {
			if math.Abs(float64(diferencia)) <= 5 {
				result = true
			} else {
				result = false
			}
		} else {
			result = false
		}

	} else if tipo == "ENTRADA-DEL-ALMUERZO" {
		diferencia := global.CalcularHora(horarioalmuerzo.Entrada)
		if global.EsPositivoNeutro(diferencia) == false {
			if math.Abs(float64(diferencia)) <= 5 {
				result = true
			} else {
				result = false
			}
		} else {
			if diferencia <= 5 {
				result = true
			} else {
				result = false
			}
		}
	} else if tipo == "SALIDA-A-CASA" {
		diferencia := global.CalcularHora(horarioalmuerzo.Salidaacasa)
		if global.EsPositivoNeutro(diferencia) == false {
			if math.Abs(float64(diferencia)) <= 15 {
				result = true
			} else {
				result = false
			}
		} else {
			result = false
		}
	}

	return result
}

func obtenerHorarioDelAlmuerzo(identificacion string) models.Horarioalmuerzo {
	var contador int = 0
	var d models.Horarioalmuerzo
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
	}
	return d
}

func verificarSiTieneHorarioAlmuerzoParametrizado(identificacion string) bool {
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

func verificarQueTipoMarcacionCorresponde(idasistencia int) string {
	var countMarcaciones int = 0
	var tipoUltimaMarcacion string = ""
	queryCount := `SELECT count(*) FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC`
	queryTipo := `SELECT TIPO FROM marcaciones WHERE IDASISTENCIA = ? ORDER BY IDMARCACION DESC LIMIT 1`

	filasCount, err := conexion.SessionMysql.Query(queryCount, idasistencia)
	if err != nil {
		panic(err)
	}

	for filasCount.Next() {
		errsql := filasCount.Scan(&countMarcaciones)
		if errsql != nil {
			panic(err)
		}
	}

	filasTipo, err := conexion.SessionMysql.Query(queryTipo, idasistencia)
	if err != nil {
		panic(err)
	}

	for filasTipo.Next() {
		errsql := filasTipo.Scan(&tipoUltimaMarcacion)
		if errsql != nil {
			panic(err)
		}
	}

	// primero verificamos si ya estan las 4 marcaciones
	if countMarcaciones == 4 && tipoUltimaMarcacion == "SALIDA" {
		return "MARCACIONES-COMPLETAS"
	} else if countMarcaciones == 3 && tipoUltimaMarcacion == "ENTRADA" {
		return "SALIDA-A-CASA"
	} else if countMarcaciones == 2 && tipoUltimaMarcacion == "SALIDA" {
		return "ENTRADA-DEL-ALMUERZO"
	} else if countMarcaciones == 1 && tipoUltimaMarcacion == "ENTRADA" {
		return "SALIDA-DEL-ALMUERZO"
	} else {
		return "Error"
	}
}

func obtenerIdAsistencia(identificacion string) int {
	var idasistencia int = 0
	query := `SELECT IDASISTENCIA FROM asistencias WHERE IDENTIFICACION = ? AND FECHA = ?`
	filas, err := conexion.SessionMysql.Query(query, identificacion, global.FechaActual())
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&idasistencia)
		if errsql != nil {
			panic(err)
		}
	}
	return idasistencia
}

func EliminarMarcacion(c *gin.Context) {
	idmarcacion := c.Param("idmarcacion")

	query, err := conexion.SessionMysql.Prepare("DELETE FROM marcaciones WHERE IDMARCACION = ?")
	if err != nil {
		panic(err)
	}

	query.Exec(idmarcacion)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "registro eliminado!"})
}
