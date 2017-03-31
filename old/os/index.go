package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("./index.go")
  if err != nil {
    fmt.Println(err)
  }
  /**
  func (f *File) Stat() (FileInfo, error)

  type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
  }
  */
  inf, err := file.Stat()

  data := make([]byte, inf.Size())

  file.Read(data)
  file.Close()
  s := string(data)
  fmt.Println(s)

  /**
  func Create(name string) (*File, error)  
  */
  os.Create("./test.tmp")
}
