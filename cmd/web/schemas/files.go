package web

import (
  "mime/multipart"
)

type FileSchema struct {
  Filename string `json:"filename" binding:"required"`
  OwnerID  int32    `json:"owner_id" binding:"required"`
}

type CreateFileSchema struct {
  File     *multipart.FileHeader `form:"file" binding:"required"`
}

type UpdateFileSchema struct {
  OwnerID int32 `json:"owner_id"`
}

