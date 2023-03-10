package controller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-vera/api_asistencia_golang/src/conexion"
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
		INNER JOIN Seguridades.UsuarioRol ON Seguridades.UsuarioRol.codigoUsuario = Seguridades.Usuario.codigo
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
		token, err := createToken(2)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusAccepted, gin.H{"empleado": u, "secuencialpersona": secuencialPersona, "token": token})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"error": "Contraseña Incorrecta!!"})
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
	codigouser := c.Param("codigouser")
	var contador int = 0
	var d models.Empleados
	var datos []models.Empleados

	query := `select identificacion, nombreUnido, secuencialOficina from Seguridades.Usuario
	inner join Seguridades.UsuarioRol on UsuarioRol.codigoUsuario = Usuario.codigo
	inner join Seguridades.Rol on Rol.codigo = UsuarioRol.codigoRol
	INNER JOIN Seguridades.Usuario_Complemento ON Seguridades.Usuario_Complemento.codigoUsuario = Seguridades.Usuario.codigo
			INNER JOIN Personas.Persona ON Persona.secuencial = Usuario_Complemento.secuencialPersona
	where Rol.codigo = '005' and Rol.estaActivo = 1 and Usuario.estaActivo = 1 and UsuarioRol.estaActivo = 1`

	filas, err := conexion.Session.Query(query, sql.Named("codigouser", codigouser))
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
