package web

type GetUserSchema struct {
  ID int32 `json:"id" binding: "required"`
}

type CreateUserSchema struct {
  Name string `json:"name" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type UpdateUserSchema struct {
  Name string `json:"name"`
}

