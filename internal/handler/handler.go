package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())         // to recover gin automatically
	router.Use(jsonLoggerMiddleware()) // we'll define it later
	//swagger
	router.StaticFile("/swagger-json", "./doc/swagger-ui/api.swagger.json")
	url := ginSwagger.URL("/swagger-json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return router
}

// func prometheusHandler() gin.HandlerFunc {
// 	h := promhttp.Handler()

// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})
			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}
