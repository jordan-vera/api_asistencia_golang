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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(Cors())

	v1 := r.Group("api/")
	{
		v1.GET("login-empleado/:codigo/:clave", controller.Login)
		v1.GET("login-servicios/:usuario/:clave", controller.LoginServicios)
	}

	v1jwt := r.Group("apijwt/")
	v1jwt.Use(middleware.AuthorizeJWT())
	{
		v1jwt.GET("validartoken", controller.Validartoken)

		v1jwt.GET("empleados", controller.GetEmpleados)
		v1jwt.GET("verificar-si-oficial-credito/:codigouser", controller.VerificarSiEsOficialCredito)
		v1jwt.GET("asesores", controller.GetAsesores)

		v1jwt.POST("marcacion", controller.AgregarMarcacion)
		v1jwt.GET("ultima-marcacion/:idasistencia", controller.Getultimamarcacion)
		v1jwt.GET("marcacioneshoy/:idasistencia", controller.GetMarcacionesHoy)

		v1jwt.GET("sucursales", controller.GetSucursal)
		v1jwt.GET("sucursal/:idsucursal", controller.Getonesucursal)

		v1jwt.POST("asistencia", controller.AgregarAsistencia)
		v1jwt.GET("verificar-asistencia/:identificacion/:fecha", controller.Verificarsiexisteasistencia)
		v1jwt.GET("asistencia-fecha/:identificacion/:fecha", controller.Getfechaasistencia)
		v1jwt.GET("asistencia-mes-anio-identificacion/:identificacion/:anio/:mes", controller.GetAsistenciaPorMesAnioEmpleado)
		v1jwt.GET("asistencia-all-empleados/:anio/:mes", controller.GetAsistenciasMarcacionesAllEmpleados)
		v1jwt.GET("asistencia-all-servicio-profecionales/:anio/:mes", controller.GetAsistenciasMarcacionesAllServiciosProfecionales)
		v1jwt.POST("justificar-asistencia", controller.JustificarAsistencia)

		v1jwt.GET("tipo-permisos", controller.GetTipoPermiso)
		v1jwt.POST("permiso", controller.AgregarPermiso)
		v1jwt.GET("permisos/:identificacion", controller.GetAllPermisos)
		v1jwt.GET("permisos-por-fecha/:identificacion/:mes/:anio", controller.GetAllPermisoFecha)
		v1jwt.GET("permiso-delete/:idpermiso", controller.EliminarPermiso)
		v1jwt.GET("permisos-count/:identificacion", controller.GetCountPermiso)
		v1jwt.GET("permisos-admin", controller.GetAllPermisosadmin)
		v1jwt.GET("autorizar-permiso/:idpermiso", controller.AutorizarPermiso)
		v1jwt.GET("permisos-estados/:estadojefe", controller.GetAllPorestadoPermiso)
		v1jwt.GET("permisos-filtros/:estadojefe/:identificacion", controller.GetPermisosFiltro)

		v1jwt.GET("rostros", controller.GetAllRostros)
		v1jwt.GET("rostro/:identificacion", controller.Getonerostros)

		v1jwt.POST("detellepermiso", controller.AgregarDetallePermiso)
		v1jwt.GET("detallepermiso/:mes/:anio", controller.GetDetallePermisos)

		v1jwt.GET("servicios-profesionales", controller.GetAllServiciosProfesionales)

		v1jwt.POST("trabajocampo", controller.AgregarTrabajoCampo)
		v1jwt.GET("trabajocampo/:identificacion/:mes/:anio", controller.GetAllTrabajoCampoPorIdentificacion)
		v1jwt.GET("trabajocampo-filtro/:identificacion/:mes/:anio", controller.GetAllTrabajoCampoFiltro)

		v1jwt.GET("skypes", controller.GetAllSkypes)

		v1jwt.POST("solicitud-anticipo", controller.AgregarAnticipo)
		v1jwt.GET("solicitudes-anticipos-pendientes", controller.GetAnticiposPendientes)
		v1jwt.GET("solicitudes-anticipos-por-estado/:estado", controller.GetAnticiposPorEstadoGerente)
		v1jwt.GET("solicitudes-anticipos-identificacion/:identificacion", controller.GetAnticiposPorIdentificacion)
		v1jwt.GET("update-anticipos-gerente/:idanticipo", controller.AutorizarAnticiposGerente)
		v1jwt.GET("solicitud-anticipo-delete/:idanticipo", controller.EliminarAnticipo)

		v1jwt.POST("bloqueo", controller.AgregarBlqueo)
		v1jwt.GET("bloqueos/:mes/:anio", controller.GetBloqueosAll)
		v1jwt.GET("bloqueos-por-estado/:mes/:anio/:estado", controller.GetBloqueosAllPorEstado)
		v1jwt.GET("verificar-si-puede-marcar/:identificacion", controller.VerificarsiPuedeMarcarAsistencia)
		v1jwt.GET("autorizar-bloqueo/:idbloqueo", controller.AutorizarBloqueos)
	}

	r.RunTLS(":8096", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/fullchain.pem", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/privkey.pem")

	//r.Run(":8096")
	//r.RunTLS(":8096", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/fullchain.pem", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/privkey.pem")
}
