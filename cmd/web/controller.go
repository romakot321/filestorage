package web

import (
	"github.com/romakot321/filestorage/docs" // docs is generated by Swag CLI, you have to import it.
  handlers "github.com/romakot321/filestorage/cmd/web/handlers"
  services "github.com/romakot321/filestorage/cmd/web/services"
  dbCon "github.com/romakot321/filestorage/db/sqlc"
  swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
  "github.com/gin-gonic/gin"
)

func Init(db *dbCon.Queries) {
  authService := services.NewAuthService(db)
  authHandler := handlers.NewAuthHandler(db)
  filesHandler := handlers.NewFileHandler(db, authService)
  
	docs.SwaggerInfo.Title = "Filestorage"
	docs.SwaggerInfo.Description = "Cloud storage in golang"
	docs.SwaggerInfo.Version = "1.0"
  docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

  router := gin.New()
  authGroup := router.Group("/auth")
  authHandler.Register(authGroup)
  filesGroup := router.Group("/files", services.AuthenticateMiddleware)
  filesHandler.Register(filesGroup)

  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  router.Run()
}
