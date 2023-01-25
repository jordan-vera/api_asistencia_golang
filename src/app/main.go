package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/controller"
	"github.com/jordan-vera/api_asistencia_golang/src/middleware"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Accept, Origin, On-behalf-of, x-sg-elas-acl, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(Cors())

	v1 := r.Group("api/")
	{
		v1.GET("login-empleado/:codigo/:clave", controller.Login)
	}

	v1jwt := r.Group("apijwt/")
	v1jwt.Use(middleware.AuthorizeJWT())
	{
		v1jwt.GET("validartoken", controller.Validartoken)

		v1jwt.GET("empleados", controller.GetEmpleados)

		v1jwt.POST("marcacion", controller.AgregarMarcacion)
		v1jwt.GET("ultima-marcacion/:idasistencia", controller.Getultimamarcacion)
		v1jwt.GET("marcacioneshoy/:idasistencia", controller.GetMarcacionesHoy)

		v1jwt.GET("sucursales", controller.GetSucursal)
		v1jwt.GET("sucursal/:idsucursal", controller.Getonesucursal)

		v1jwt.POST("asistencia", controller.AgregarAsistencia)
		v1jwt.GET("verificar-asistencia/:identificacion/:fecha", controller.Verificarsiexisteasistencia)
		v1jwt.GET("asistencia-fecha/:identificacion/:fecha", controller.Getfechaasistencia)

		v1jwt.GET("tipo-permisos", controller.GetTipoPermiso)
		v1jwt.POST("permiso", controller.AgregarPermiso)
		v1jwt.GET("permisos/:identificacion", controller.GetAllPermisos)
		v1jwt.GET("permiso-delete/:idpermiso", controller.EliminarPermiso)
		v1jwt.GET("permisos-count/:identificacion", controller.GetCountPermiso)
		v1jwt.GET("permisos-admin", controller.GetAllPermisosadmin)
		v1jwt.GET("autorizar-permiso/:idpermiso", controller.AutorizarPermiso)
		v1jwt.GET("permisos-estados/:estadojefe", controller.GetAllPorestadoPermiso)
		v1jwt.GET("permisos-filtros/:estadojefe/:identificacion", controller.GetPermisosFiltro)

		v1jwt.GET("rostros", controller.GetAllRostros)
		v1jwt.GET("rostro/:identificacion", controller.Getonerostros)
	}

	r.Run(":8096")
	//r.RunTLS(":8096", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/fullchain.pem", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/privkey.pem")
}
