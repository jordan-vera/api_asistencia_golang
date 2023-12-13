package controller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
	"github.com/jordan-vera/api_asistencia_golang/src/global"
	"github.com/jordan-vera/api_asistencia_golang/src/models"
)

func Login(c *gin.Context) {
	var contador int = 0
	var secuencialPersona int = 0

	usuario := c.Param("codigo")
	clave := c.Param("clave")

	u := models.Usuario{}

	query := `select codigo, nombre, Seguridades.Usuario.estaActivo, secuencialOficina, secuencialPersona, Persona.identificacion from Seguridades.Usuario
		INNER JOIN Seguridades.Usuario_Complemento ON Seguridades.Usuario_Complemento.codigoUsuario = Seguridades.Usuario.codigo
		INNER JOIN Personas.Persona ON Persona.secuencial = Usuario_Complemento.secuencialPersona
		WHERE Seguridades.Usuario.codigo = @user AND Seguridades.Usuario_Complemento.clave = UPPER(sys.fn_varbintohexsubstring(0, HashBytes('SHA1',CAST(Seguridades.Usuario.codigo + @clave AS VARCHAR)),1,0)) ;
		`
	filas, errsql := conexion.Session.Query(query, sql.Named("user", usuario), sql.Named("clave", clave))
	if errsql != nil {
		panic(errsql)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&u.Codigouser, &u.Nombre, &u.Estadoactivo, &u.Secuencialoficina, &secuencialPersona, &u.Identificacion)
		if errsql != nil {
			panic(errsql)
		}
	}

	if contador > 0 {
		if verificarSiServicioProfesionalEstaActivo(u.Identificacion) == "inactivo" {
			c.JSON(http.StatusAccepted, gin.H{"error": "No autorizado!!"})
		} else if verificarSiServicioProfesionalEstaActivo(u.Identificacion) == "noexiste" || verificarSiServicioProfesionalEstaActivo(u.Identificacion) == "activo" {
			if verificarSiEstaEnVacaciones(u.Identificacion, global.NumAnioActual(), global.NumMesActual(), global.NumDiaActual()) {
				c.JSON(http.StatusAccepted, gin.H{"error": "No puedes ingresar por que estas en vacaciones!!"})
			} else {
				token, err := createToken(2)
				if err != nil {
					panic(err)
				}
				c.JSON(http.StatusAccepted, gin.H{"empleado": u, "secuencialpersona": secuencialPersona, "token": token, "rol": verificarSiTieneRoles(u.Identificacion)})
			}
		}
	} else {
		c.JSON(http.StatusAccepted, gin.H{"error": "Contraseña Incorrecta!!"})
	}

}

func verificarSiEstaEnVacaciones(identificacion string, anio int, mes int, dia int) bool {

	var idvacaciones int = 0

	query := `SELECT idvacaciones FROM vacaciones WHERE identificacion = ? AND anio = ?`

	filas, err := conexion.SessionMysql.Query(query, identificacion, anio)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&idvacaciones)
		if errsql != nil {
			panic(err)
		}
	}

	if idvacaciones != 0 {
		if buscarEnDetalleVacaciones(idvacaciones, anio, mes, dia) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func buscarEnDetalleVacaciones(idvacaciones int, anio int, mes int, dia int) bool {
	var count int = 0
	query := `SELECT COUNT(*) FROM vacacionesdetalles WHERE idvacaciones = ? AND anio = ? AND mes = ? AND numerodia = ?`

	filas, err := conexion.SessionMysql.Query(query, idvacaciones, anio, mes, dia)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		errsql := filas.Scan(&count)
		if errsql != nil {
			panic(err)
		}
	}

	if count > 0 {
		return true
	} else {
		return false
	}
}

func verificarSiServicioProfesionalEstaActivo(identificacion string) string {
	var estado int = 0
	var contador int = 0

	query := `SELECT estado FROM serviciosprofecionales WHERE identificacion = ?`
	filas, errsql := conexion.SessionMysql.Query(query, identificacion)
	if errsql != nil {
		panic(errsql)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&estado)
		if errsql != nil {
			panic(errsql)
		}
	}

	if contador > 0 {
		if estado == 1 {
			return "activo"
		} else {
			return "inactivo"
		}
	} else {
		return "noexiste"
	}
}

func verificarSiTieneRoles(identificacion string) string {
	var contador int = 0
	var tipouser string = ""
	query := `  SELECT tipousuario.tipouser FROM usuarioroles 
	            INNER JOIN tipousuario on tipousuario.idtipousuario = usuarioroles.idtipousuario
	            WHERE identificacion = ? and tipousuario.estado = 1`
	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&tipouser)
		if errsql != nil {
			panic(err)
		}
	}
	if contador > 0 {
		return tipouser
	} else {
		return ""
	}
}

func GetEmpleados(c *gin.Context) {
	var contador int = 0
	var d models.Empleados
	var datos []models.Empleados

	query := `select identificacion, nombreUnido from Personas.Persona
	inner join Nomina.EMPLEADO on EMPLEADO.SECUENCIALPERSONANATURAL = Persona.secuencial
	where EMPLEADO.CODIGOESTADOEMPLEADO = 'A' 
	order by nombreUnido asc`

	filas, err := conexion.Session.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Identificacion, &d.NombreUnido)
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

func GetEmpleadosSinVacaciones(c *gin.Context) {
	var contador int = 0
	var d models.Empleados
	var datos []models.Empleados
	anio := c.Param("anio")

	query := `select identificacion, nombreUnido from Personas.Persona
	inner join Nomina.EMPLEADO on EMPLEADO.SECUENCIALPERSONANATURAL = Persona.secuencial
	where EMPLEADO.CODIGOESTADOEMPLEADO = 'A' 
	order by nombreUnido asc`

	filas, err := conexion.Session.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Identificacion, &d.NombreUnido)
		if errsql != nil {
			panic(err)
		}

		if verificarSiTieneAsignadaVacaciones(d.Identificacion, anio) == false {
			datos = append(datos, d)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func GetEmpleadosConHorarioAlmuerzo(c *gin.Context) {
	var contador int = 0
	var d models.Empleados
	var datos []models.Empleados

	query := `select identificacion, nombreUnido from Personas.Persona
	inner join Nomina.EMPLEADO on EMPLEADO.SECUENCIALPERSONANATURAL = Persona.secuencial
	where EMPLEADO.CODIGOESTADOEMPLEADO = 'A' 
	order by nombreUnido asc`

	filas, err := conexion.Session.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Identificacion, &d.NombreUnido)
		if errsql != nil {
			panic(err)
		}

		if verificarSiTieneAsignadoHorarioAlmuerzo(d.Identificacion) == false {
			datos = append(datos, d)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func GetEmpleadosConVacaciones(c *gin.Context) {
	var contador int = 0
	var d models.EmpleadosConVacacionesDetalle
	var datos []models.EmpleadosConVacacionesDetalle
	anio := c.Param("anio")

	query := `select identificacion, nombreUnido from Personas.Persona
	inner join Nomina.EMPLEADO on EMPLEADO.SECUENCIALPERSONANATURAL = Persona.secuencial
	where EMPLEADO.CODIGOESTADOEMPLEADO = 'A' 
	order by nombreUnido asc`

	filas, err := conexion.Session.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Identificacion, &d.NombreUnido)
		if errsql != nil {
			panic(err)
		}
		d.Vacaciones = obtenerVacacionesY_Detalle(d.Identificacion)
		if verificarSiTieneAsignadaVacaciones(d.Identificacion, anio) == true {
			datos = append(datos, d)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func GetEmpleadoConVacacionesIdentificacion(c *gin.Context) {
	var contador int = 0
	var d models.EmpleadosConVacacionesDetalle
	var datos []models.EmpleadosConVacacionesDetalle
	anio := c.Param("anio")
	identificacion := c.Param("identificacion")

	query := `select identificacion, nombreUnido from Personas.Persona
	inner join Nomina.EMPLEADO on EMPLEADO.SECUENCIALPERSONANATURAL = Persona.secuencial
	where EMPLEADO.CODIGOESTADOEMPLEADO = 'A' and identificacion = @identificacion
	order by nombreUnido asc`

	filas, err := conexion.Session.Query(query, sql.Named("identificacion", identificacion))
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Identificacion, &d.NombreUnido)
		if errsql != nil {
			panic(err)
		}
		d.Vacaciones = obtenerVacacionesY_Detalle(d.Identificacion)
		if verificarSiTieneAsignadaVacaciones(d.Identificacion, anio) == true {
			datos = append(datos, d)
		}
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": datos})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func obtenerVacacionesY_Detalle(identificacion string) []models.VacacionesDetalleFilter {
	var data models.VacacionesDetalleFilter
	var datos []models.VacacionesDetalleFilter
	query := `SELECT idvacaciones, identificacion, cantidaddias, estado, anio FROM vacaciones where identificacion = ?`
	filas, err := conexion.SessionMysql.Query(query, identificacion)
	if err != nil {
		panic(err)
	}
	for filas.Next() {
		errsql := filas.Scan(&data.Idvacaciones, &data.Identificacion, &data.Cantidaddias, &data.Estado, &data.Anio)
		if errsql != nil {
			panic(err)
		}
		data.Detalle = obtenerDetalleVacaciones(data.Idvacaciones)
		datos = append(datos, data)
	}
	return datos
}

func obtenerDetalleVacaciones(idvacaciones int) []models.VacacionesDetalle {
	var d models.VacacionesDetalle
	var data []models.VacacionesDetalle
	query := `SELECT iddetallevacaciones, idvacaciones, numerodia, mes, anio FROM vacacionesdetalles where idvacaciones = ? 
	order by mes ASC, numerodia ASC
	         `
	filas, err := conexion.SessionMysql.Query(query, idvacaciones)
	if err != nil {
		panic(err)
	}
	for filas.Next() {
		errsql := filas.Scan(&d.Iddetallevacaciones, &d.Idvacaciones, &d.Numerodia, &d.Mes, &d.Anio)
		if errsql != nil {
			panic(err)
		}
		data = append(data, d)
	}
	return data
}

func verificarSiTieneAsignadaVacaciones(identificacion string, anio string) bool {
	var contador int = 0
	query := `SELECT count(*) FROM vacaciones where identificacion = ? and anio = ?`
	filas, err := conexion.SessionMysql.Query(query, identificacion, anio)
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

func createToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("fenix1920"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Validartoken(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"response": "ok"})
}

func VerificarSiEsOficialCredito(c *gin.Context) {
	codigouser := c.Param("codigouser")
	var contador int = 0

	query := `select * from Seguridades.Usuario
	inner join Seguridades.UsuarioRol on UsuarioRol.codigoUsuario = Usuario.codigo
	inner join Seguridades.Rol on Rol.codigo = UsuarioRol.codigoRol
	where Usuario.codigo = @codigouser and Rol.codigo = '005' and UsuarioRol.estaActivo = 1`

	filas, err := conexion.Session.Query(query, sql.Named("codigouser", codigouser))
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
	}

	if contador > 0 {
		c.JSON(http.StatusCreated, gin.H{"response": "ok"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"error": "No hay datos"})
	}
}

func GetAsesores(c *gin.Context) {
	var contador int = 0
	var d models.Empleados
	var datos []models.Empleados

	query := `select identificacion, nombreUnido, secuencialOficina from Seguridades.Usuario
			inner join Seguridades.UsuarioRol on UsuarioRol.codigoUsuario = Usuario.codigo
			inner join Seguridades.Rol on Rol.codigo = UsuarioRol.codigoRol
			INNER JOIN Seguridades.Usuario_Complemento ON Seguridades.Usuario_Complemento.codigoUsuario = Seguridades.Usuario.codigo
			INNER JOIN Personas.Persona ON Persona.secuencial = Usuario_Complemento.secuencialPersona
			where Rol.codigo = '005' and Rol.estaActivo = 1 and Usuario.estaActivo = 1 and UsuarioRol.estaActivo = 1`

	filas, err := conexion.Session.Query(query)
	if err != nil {
		panic(err)
	}

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Identificacion, &d.NombreUnido, &d.SecuencialOficina)
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
