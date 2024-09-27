package web

import (
  "fmt"
	"time"
  "strconv"

  jwt "github.com/golang-jwt/jwt/v5"
  db "github.com/romakot321/filestorage/db/sqlc"
  "github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key")

type AuthService struct {
  db *db.Queries
}

func (s AuthService) GetCurrentUser(c *gin.Context) db.User {
	tokenString, ok := c.Request.Header["Token"]
  if ok == false {
		fmt.Println("Token missing in header")
		c.AbortWithStatus(401)
		return db.User{}
	}
  token, _ := VerifyToken(tokenString[0])
  sub, _ := token.Claims.(jwt.MapClaims).GetSubject()
  userID, _ := strconv.Atoi(sub)
  user, _ := s.db.GetUserById(c, int32(userID))
  return user
}

func NewAuthService(db *db.Queries) *AuthService {
  return &AuthService{db: db}
}

func CreateToken(userID int32) (string, error) {
  claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(userID)),                    // Subject (user identifier)
		"iss": "files",                  // Issuer
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
	})

  tokenString, err := claims.SignedString(secretKey)
  if err != nil {
      return "", err
  }

	fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}

func AuthenticateMiddleware(c *gin.Context) {
	tokenString, ok := c.Request.Header["Token"]
	if ok == false {
		fmt.Println("Token missing in header")
		c.AbortWithStatus(401)
		return
	}

	token, err := VerifyToken(tokenString[0])
	if err != nil {
		fmt.Printf("Token verification failed: %v\n", err)
		c.AbortWithStatus(401)
		return
	}

	fmt.Printf("Token verified successfully. Claims: %+v\n", token.Claims)
  subj, err := token.Claims.GetSubject()
  fmt.Printf("Token subject: %+v, err: %+v\n", subj, err)
	c.Next()
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
