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
		v1jwt.GET("empleados-sin-vacaciones/:anio", controller.GetEmpleadosSinVacaciones)
		v1jwt.GET("empleados-con-vacaciones/:anio", controller.GetEmpleadosConVacaciones)
		v1jwt.GET("empleados-con-vacaciones-identificacion/:anio/:identificacion", controller.GetEmpleadoConVacacionesIdentificacion)
		v1jwt.GET("verificar-si-oficial-credito/:codigouser", controller.VerificarSiEsOficialCredito)
		v1jwt.GET("asesores", controller.GetAsesores)
		v1jwt.GET("empleados-sin-horarios-almuerzo", controller.GetEmpleadosConHorarioAlmuerzo)

		v1jwt.POST("marcacion", controller.AgregarMarcacion)
		v1jwt.POST("nueva-marcacion", controller.NuevaMarcacionParaEdit)
		v1jwt.POST("marcacion-general/:identificacion/:tienepermiso", controller.RealizarMarcacionGeneral)
		v1jwt.GET("ultima-marcacion/:idasistencia", controller.Getultimamarcacion)
		v1jwt.GET("marcacioneshoy/:idasistencia", controller.GetMarcacionesHoy)
		v1jwt.GET("marcacion-delete/:idmarcacion", controller.EliminarMarcacion)

		v1jwt.GET("sucursales", controller.GetSucursal)
		v1jwt.GET("sucursal/:idsucursal", controller.Getonesucursal)

		v1jwt.POST("asistencia", controller.AgregarAsistencia)
		v1jwt.GET("verificar-asistencia/:identificacion/:fecha", controller.Verificarsiexisteasistencia)
		v1jwt.GET("asistencia-fecha/:identificacion/:fecha", controller.Getfechaasistencia)
		v1jwt.GET("asistencia-mes-anio-identificacion/:identificacion/:anio/:mes", controller.GetAsistenciaPorMesAnioEmpleado)
		v1jwt.GET("asistencia-all-empleados/:anio/:mes", controller.GetAsistenciasMarcacionesAllEmpleados)
		v1jwt.GET("asistencia-all-servicio-profecionales/:anio/:mes", controller.GetAsistenciasMarcacionesAllServiciosProfecionales)
		v1jwt.POST("justificar-asistencia", controller.JustificarAsistencia)
		v1jwt.GET("asistencias-por-fecha/:fecha", controller.GetAsistenciaAllPorFecha)

		v1jwt.GET("tipo-permisos", controller.GetTipoPermiso)
		v1jwt.POST("permiso", controller.AgregarPermiso)
		v1jwt.GET("permisos/:identificacion", controller.GetAllPermisos)
		v1jwt.GET("permisos-por-fecha/:identificacion/:mes/:anio", controller.GetAllPermisoFecha)
		v1jwt.GET("permiso-delete/:idpermiso", controller.EliminarPermiso)
		v1jwt.GET("permisos-count/:identificacion", controller.GetCountPermiso)
		v1jwt.GET("permisos-admin/:anio", controller.GetAllPermisosadmin)
		v1jwt.GET("autorizar-permiso/:idpermiso/:usuario", controller.AutorizarPermiso)
		v1jwt.GET("negar-permiso/:idpermiso/:usuario", controller.NegarPermiso)
		v1jwt.GET("permisos-estados/:estadojefe", controller.GetAllPorestadoPermiso)
		v1jwt.GET("permisos-filtros/:estadojefe/:identificacion/:anio", controller.GetPermisosFiltro)
		v1jwt.GET("permisos-para-calcular-vacaciones/:identificacion", controller.GetPermisosParaCalcularVacaciones)
		v1jwt.GET("update-escargo-vacaciones/:idpermiso/:escargovacaciones", controller.EditEsCargoVacaciones)
		v1jwt.POST("permisos-por-anio-mes-dia", controller.GetPermisoPorAnioMesDia)

		v1jwt.GET("rostros", controller.GetAllRostros)
		v1jwt.GET("rostro/:identificacion", controller.Getonerostros)

		v1jwt.POST("detellepermiso", controller.AgregarDetallePermiso)
		v1jwt.GET("detallepermiso/:mes/:anio", controller.GetDetallePermisos)
		v1jwt.GET("detallepermiso-fecha/:mes/:anio/:dia", controller.GetDetallePermisosFecha)

		v1jwt.GET("servicios-profesionales", controller.GetAllServiciosProfesionales)
		v1jwt.GET("servicios-profesionales-all", controller.GetAllServiciosProfesionalesAll)
		v1jwt.GET("servicios-sin-horarios-almuerzo", controller.GetServiciosProfesionalesConHorarioAlmuerzo)
		v1jwt.GET("servicios-profesional-inactivar/:idservicio", controller.InactivarServicioProfesional)

		v1jwt.POST("busqueda", controller.AgregarBusqueda)

		v1jwt.POST("trabajocampo", controller.AgregarTrabajoCampo)
		v1jwt.GET("trabajocampo/:identificacion/:mes/:anio", controller.GetAllTrabajoCampoPorIdentificacion)
		v1jwt.GET("trabajocampo-filtro/:identificacion/:mes/:anio", controller.GetAllTrabajoCampoFiltro)
		v1jwt.GET("trabajocampo-fecha/:mes/:anio/:dia", controller.GetAllTrabajoCampoPorFecha)

		v1jwt.GET("skypes", controller.GetAllSkypes)

		v1jwt.POST("solicitud-anticipo", controller.AgregarAnticipo)
		v1jwt.GET("solicitudes-anticipos-pendientes/:anio", controller.GetAnticiposPendientes)
		v1jwt.GET("solicitudes-anticipos-por-estado/:estado/:anio", controller.GetAnticiposPorEstadoGerente)
		v1jwt.GET("get-numerocliente/:identificacion", controller.GetNumerocliente)
		v1jwt.GET("solicitudes-anticipos-identificacion/:identificacion/:anio", controller.GetAnticiposPorIdentificacion)
		v1jwt.GET("solicitudes-anticipos-identificacion-mes-anio/:identificacion/:mes/:anio", controller.GetAnticiposPorIdentificacionMesAnio)
		v1jwt.GET("update-anticipos-gerente/:idanticipo", controller.AutorizarAnticiposGerente)
		v1jwt.GET("negar-anticipo-gerente/:idanticipo", controller.NegarAnticiposGerente)
		v1jwt.GET("solicitud-anticipo-delete/:idanticipo", controller.EliminarAnticipo)

		v1jwt.POST("bloqueo", controller.AgregarBlqueo)
		v1jwt.GET("bloqueos/:mes/:anio", controller.GetBloqueosAll)
		v1jwt.GET("bloqueos-por-estado/:mes/:anio/:estado", controller.GetBloqueosAllPorEstado)
		v1jwt.GET("bloqueos-por-fecha/:mes/:anio/:dia", controller.GetBloqueosAllPorFecha)
		v1jwt.GET("bloqueos-por-identificacion-mes-anio/:mes/:anio/:identificacion", controller.GetBloqueosIdentificacionMesAnio)
		v1jwt.GET("verificar-si-puede-marcar/:identificacion", controller.VerificarsiPuedeMarcarAsistencia)
		v1jwt.GET("autorizar-bloqueo/:idbloqueo/:usuario", controller.AutorizarBloqueos)

		v1jwt.POST("vacaciones", controller.AgregarVacaciones)
		v1jwt.PUT("vacaciones", controller.Updatevacaciones)
		v1jwt.GET("vacaciones/:anio", controller.GetVacacionesAll)
		v1jwt.GET("vacaciones-por-empleado/:identificacion", controller.GetVacacionesPorIdentificacion)

		v1jwt.POST("vacacionesdetalle", controller.AgregarVacacionesDetalle)
		v1jwt.GET("vacacionesdetalle/:anio/:mes", controller.GetVacacionesDetalleAll)
		v1jwt.GET("vacacionesdetalle-identificacion/:anio/:mes/:identificacion", controller.GetVacacionesDetalleAllIdentificacion)
		v1jwt.GET("vacacionesdetalle-delete/:idvacaciones", controller.EliminarDetallesVacaciones)
		v1jwt.GET("vacacionesdetalle-delete-one/:iddetallevacaciones", controller.EliminarDetallesVacacionesOne)

		v1jwt.POST("horarioalmuerzo", controller.AgregarHorarioAlmuerzo)
		v1jwt.GET("horarioalmuerzosall", controller.GetHorasAlmuerzo)
		v1jwt.GET("horarioalmuerzosidentificacion/:identificacion", controller.GetHorasAlmuerzoPorIdentificacion)
		v1jwt.GET("horarioalmuerzosentrada/:entrada", controller.GetHorasAlmuerzoPorEntrada)
		v1jwt.PUT("horarioalmuerzoedit", controller.ActualizarHorario)

		v1jwt.POST("bloqueofuerarango", controller.AgregarBlqueoFueraRango)
		v1jwt.GET("bloqueosfuerarango/:mes/:anio", controller.GetBloqueosFueraRangoAll)
		v1jwt.GET("bloqueos-por-estado-fuerarango/:mes/:anio/:estado", controller.GetBloqueosFueraRanfoAllPorEstado)
		v1jwt.GET("bloqueos-por-fecha-fuerarango/:mes/:anio/:dia", controller.GetBloqueosFueraRangoAllPorFecha)
		v1jwt.GET("bloqueos-por-identificacion-mes-anio-fuerarango/:mes/:anio/:identificacion", controller.GetBloqueosFueraRangoIdentificacionMesAnio)

	}

	r.RunTLS(":8096", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/fullchain.pem", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/privkey.pem")

	//r.Run(":8096")
}
