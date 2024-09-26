package files

import (
  "os"
  "io"
  "log"
  "bufio"
  "path/filepath"
  "path"
  "github.com/google/uuid"
)

var StorageFolder string = "storage/"

func (f *File) GenerateFilename() {
  f.Filename = uuid.New()
}

func (f *File) Create(content io.Reader) {
  file, err := os.Create(StorageFolder + f.Filename.String())
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  io.Copy(w, content)
  w.Flush()
}

func (f *File) Path() string {
  return filepath.Join(StorageFolder, filepath.FromSlash(path.Clean("/" + f.Filename.String())))
}

func (f *File) Read() *bufio.Reader {
  file, err := os.Open(f.Path())

  if err != nil {
    log.Fatal(err)
  }
  r := bufio.NewReader(file)
  return r
}

func GetFiles() []File {
  var files []File = make([]File, 0, 5)

  filenames, _ := os.ReadDir(StorageFolder)
  for _, fn := range filenames {
    if fn.Name() == "" { continue }
    filename, _ := uuid.Parse(fn.Name())
    files = append(files, File{Filename: filename})
  }
  return files
}
