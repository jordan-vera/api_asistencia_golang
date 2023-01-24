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
		c.JSON(http.StatusAccepted, gin.H{"persona": u, "secuencialpersona": secuencialPersona, "token": token})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"error": "Contraseña Incorrecta!!"})
	}
}

func createToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("fenix1920"))
	if err != nil {
		return "", err
	}
	return token, nil
}
