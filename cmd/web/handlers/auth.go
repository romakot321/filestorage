package web

import (
  "net/http"
  "fmt"

  db "github.com/romakot321/filestorage/db/sqlc"
  services "github.com/romakot321/filestorage/cmd/web/services"
  schemas "github.com/romakot321/filestorage/cmd/web/schemas"
  "github.com/gin-gonic/gin"
)

type AuthHandler interface {
  Register(r *gin.RouterGroup)
}

type authHandler struct {
  db *db.Queries
}

func (h authHandler) Register(r *gin.RouterGroup) {
  r.POST("/login", h.handleLogin)
  r.POST("/register", h.handleRegister)
}

// Login godoc
// @Summary      Do Log in
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        schema body schemas.LoginSchema true "Login schema"
// @Router       /auth/login/ [post]
func (h authHandler) handleLogin(c *gin.Context) {
  var schema schemas.LoginSchema
  c.ShouldBind(&schema)

  if user, err := h.db.GetUserByName(c, schema.Username); err != nil {
    c.String(http.StatusUnauthorized, "Invalid credentials")
  } else {
    tokenString, err := services.CreateToken(user.ID)
    if err != nil {
      c.String(http.StatusInternalServerError, "Error creating token")
      return
    }

    resp := schemas.TokenSchema{Token: tokenString}
    c.JSON(200, resp)
  } 
}

// Register godoc
// @Summary      Register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        schema body schemas.RegisterSchema true "Register schema"
// @Router       /auth/register/ [post]
func (h authHandler) handleRegister(c *gin.Context) {
  var schema schemas.RegisterSchema
  c.ShouldBind(&schema)

  args := &db.CreateUserParams{Name: schema.Username, PasswordHash: schema.Password}

  if user, err := h.db.CreateUser(c, *args); err != nil {
    fmt.Println(err)
    c.String(400, "Invalid request")
  } else {
    c.JSON(201, user)
  } 
}

func NewAuthHandler(dbCon *db.Queries) AuthHandler {
  return authHandler{db: dbCon}
}
