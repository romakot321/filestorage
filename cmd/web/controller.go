package web

import (
  handlers "github.com/romakot321/filestorage/cmd/web/handlers"
  dbCon "github.com/romakot321/filestorage/db/sqlc"

  "github.com/gin-gonic/gin"
)

func Init(db *dbCon.Queries) {
  filesHandler := handlers.NewFileHandler(db)

  router := gin.New()
  filesGroup := router.Group("/files")
  filesHandler.Register(filesGroup)

  router.Run()
}
