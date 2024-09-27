package web

import (
  "strings"
  "net/http"
  files "github.com/romakot321/filestorage/internal/files"
  schemas "github.com/romakot321/filestorage/cmd/web/schemas"
  services "github.com/romakot321/filestorage/cmd/web/services"
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
  auth *services.AuthService
}

func (h fileHandler) Register(router *gin.RouterGroup) {
  router.GET("/", h.handleGetList)
  router.GET("/:filename", h.handleGet)
  router.POST("/", h.handleCreate)
}

// GetFilesList godoc
// @Summary      Show a list of accessable files
// @Tags         files
// @Accept       json
// @Produce      json
// @Success      200  {object}  []db.File
// @Router       /files [get]
// @Security ApiKeyAuth
func (h fileHandler) handleGetList(c *gin.Context) {
  args := &db.ListFilesParams{Limit: 10000000, Offset: 0}

  fileList, err := h.db.ListFiles(c, *args)
  if err != nil {
    c.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving file", "error": err.Error()})
    return
  }
  currentUser := h.auth.GetCurrentUser(c)

  filteredFileList := fileList[:0]
  for _, file := range fileList {
    if file.OwnerID == currentUser.ID {
      filteredFileList = append(filteredFileList, file)
    }
  }

  c.JSON(200, filteredFileList)
}

// GetFile godoc
// @Summary      Stream of file content
// @Tags         files
// @Accept       mpfd
// @Produce      multipart/form-data
// @Success      200  {binary}  string
// @Param        filename path string true "Filename UUID"
// @Router       /files/{filename} [get]
// @Security ApiKeyAuth
func (h fileHandler) handleGet(c *gin.Context) {
  filename := c.Param("filename")
  if strings.TrimPrefix(filename, "/") == "" {
    c.AbortWithStatus(404)
    return
  }
  filenameParsed, _ := uuid.Parse(filename)

  currentUser := h.auth.GetCurrentUser(c)
  fileModel, err := h.db.GetFileById(c, filenameParsed)
  if fileModel.OwnerID != currentUser.ID || err != nil{
    c.AbortWithStatus(401)
    return
  }

  filenameVerified, _ := uuid.Parse(filename)
  file := files.File{Filename: filenameVerified}
  c.File(file.Path())
}

// CreateFile godoc
// @Summary      Create a file
// @Tags         files
// @Accept       mpfd
// @Produce      json
// @Param        file formData file true "File"
// @Success      200  {object}  db.File
// @Router       /files [post]
// @Security ApiKeyAuth
func (h fileHandler) handleCreate(c *gin.Context) {
  var payload schemas.CreateFileSchema
  if err := c.ShouldBind(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
    return
  }
  currentUser := h.auth.GetCurrentUser(c)

  fileObject := files.File{}
  fileObject.GenerateFilename()
  args := &db.CreateFileParams{Filename: fileObject.Filename, OwnerID: currentUser.ID}
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

func NewFileHandler(db *db.Queries, auth *services.AuthService) FileHandler {
  h := fileHandler{db: db, auth: auth}
  return h
}
