package web

import (
  "net/http"
  schemas "github.com/romakot321/filestorage/cmd/web/schemas"
  db "github.com/romakot321/filestorage/db/sqlc"
  
  "github.com/gin-gonic/gin"
)

type UserHandler interface {
  handleGetList(c *gin.Context)
  handleGet(c *gin.Context)
  handleCreate(c *gin.Context)
  Register(router *gin.RouterGroup)
}

type userHandler struct {
  db *db.Queries
}

func (h userHandler) Register(router *gin.RouterGroup) {
  router.GET("/", h.handleGetList)
  router.GET("/:id", h.handleGet)
  router.POST("/", h.handleCreate)
}

func (h userHandler) handleGetList(c *gin.Context) {
  args := &db.ListUsersParams{Limit: 10000000, Offset: 0}

  userList, err := h.db.ListUsers(c, *args)
  if err != nil {
    c.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
    return
  }

  c.JSON(200, userList)
}

func (h userHandler) handleGet(c *gin.Context) {
  var payload schemas.GetUserSchema
  if err := c.ShouldBind(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
    return
  }

  userModel, err := h.db.GetUserById(c, payload.ID)
  if err != nil {
    c.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
    return
  }

  c.JSON(200, userModel)
}

func (h userHandler) handleCreate(c *gin.Context) {
  var payload schemas.CreateUserSchema
  if err := c.ShouldBind(&payload); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
    return
  }

  args := &db.CreateUserParams{Name: payload.Name, PasswordHash: payload.Password}
  userModel, err := h.db.CreateUser(c, *args)
  if err != nil {
    c.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
    return
  }

  c.JSON(201, userModel)
}

func NewUserHandler(db *db.Queries) UserHandler {
  h := userHandler{db: db}
  return h
}
