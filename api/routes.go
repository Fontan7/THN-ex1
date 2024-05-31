package api

import (
	"THN-ex1/middleware"

	_ "THN-ex1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initializes and returns a new Gin engine with configured routes.
func InitRouter(app App) *gin.Engine {
	gin.SetMode(app.GinMode())
	router := gin.Default()
	router.Use(middleware.ErrorManager())
	router.Use(middleware.CorsConfig())
	defineRoutes(router, app)
	return router
}

// defineRoutes sets up routes and their corresponding groups.
func defineRoutes(router *gin.Engine, app App) {
	registerHealthAndSwaggerRoutes(router)
	registerVersionedAPIRoutes(router, app)
}

// registerHealthAndSwaggerRoutes registers routes for health checks and Swagger UI.
func registerHealthAndSwaggerRoutes(router *gin.Engine) {
	router.GET("/health", handleHealthCheck)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// registerAPIRoutes sets up API versioning and their specific routes.
func registerVersionedAPIRoutes(router *gin.Engine, app App) {
	v1 := router.Group("/v1")
	routesWithLogsV1(v1, app)
	routesWithNoLogsV1(v1, app)
}

// definePublicRoutes registers public routes. /v1/public/...
func routesWithLogsV1(rg *gin.RouterGroup, app App) {
	rg.GET("/feature", func(c *gin.Context) { handleGetFeature(c, app) })
}

// definePrivateRoutes registers private routes requiring authentication.
func routesWithNoLogsV1(rg *gin.RouterGroup, app App) {
	rg.Use(middleware.CheckAPIKey(app.ClientKey())) // API key check middleware
	rg.GET("/metrics/:ip", func(c *gin.Context) { handleGetMetrics(c, app) })
}
