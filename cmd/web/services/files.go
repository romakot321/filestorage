package web

import (
  "io"
  "github.com/romakot321/filestorage/internal/files"
)

type FileService interface {
  Create(content io.Reader) string
  GetList() []files.File
}

type fileService struct {
  files []files.File
}

func New() FileService {
  return fileService{}
}

func (s *fileService) Create(content io.Reader) string {
  file := files.File{}
  file.GenerateFilename()
  file.Create(content)
  s.files = append(s.files, file)
  return file.Filename
}

func (s *fileService) GetList() []FileSchema {
  var schemas []FileSchema = make([]FileSchema, len(s.files))
  for i, f := range s.files {
    schemas[i] = FileSchema{Filename: f.Filename}
  }
  return schemas
}
