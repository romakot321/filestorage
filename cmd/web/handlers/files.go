package web

import (
  "strings"
  "net/http"
  files "github.com/romakot321/filestorage/internal/files"
  schemas "github.com/romakot321/filestorage/cmd/web/schemas"
  db "github.com/romakot321/filestorage/db/sqlc"
  
  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
)

type FileHandler interface {
  handleGetList(c *gin.Context)
  handleGet(c *gin.Context)
  handleCreate(c *gin.Context)
  Register(router *gin.RouterGroup)
}

type fileHandler struct {
  db *db.Queries
}

func (h fileHandler) Register(router *gin.RouterGroup) {
  router.GET("/", h.handleGetList)
  router.GET("/:filename", h.handleGet)
  router.POST("/", h.handleCreate)
}

func (h fileHandler) handleGetList(c *gin.Context) {
  args := &db.ListFilesParams{Limit: 10000000, Offset: 0}

  fileList, err := h.db.ListFiles(c, *args)
  if err != nil {
    c.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving file", "error": err.Error()})
    return
  }

  c.JSON(200, fileList)
}

func (h fileHandler) handleGet(c *gin.Context) {
  filename := c.Param("filename")
  if strings.TrimPrefix(filename, "/") == "" {
    c.AbortWithStatus(404)
    return
  }

  filenameVerified, _ := uuid.Parse(filename)
  file := files.File{Filename: filenameVerified}
  c.File(file.Path())
}

func (h fileHandler) handleCreate(c *gin.Context) {
  var payload schemas.CreateFileSchema
  if err := c.ShouldBind(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
    return
  }

  fileObject := files.File{}
  fileObject.GenerateFilename()
  args := &db.CreateFileParams{Filename: fileObject.Filename, OwnerID: payload.OwnerID}
  fileModel, err := h.db.CreateFile(c, *args)
  if err != nil {
    c.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving file", "error": err.Error()})
    return
  }

  content, _ := payload.File.Open()
  defer content.Close()
  fileObject.Create(content)

  c.JSON(201, fileModel)
}

func NewFileHandler(db *db.Queries) FileHandler {
  h := fileHandler{db: db}
  return h
}
