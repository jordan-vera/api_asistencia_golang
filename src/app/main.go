package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jordan-vera/api_asistencia_golang/src/controller"
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
		v1.GET("login/:codigo/:clave", controller.Login)
	}

	r.Run(":8095")
	//r.RunTLS(":8088", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/fullchain.pem", "/etc/letsencrypt/live/sistemflm.futurolamanense.fin.ec/privkey.pem")
}
