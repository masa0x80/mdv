package mdv

import (
  "bufio"
  "os"
)

func Init(filePath string) error {
  var fp *os.File
  var err error

  fp, err = os.Open(filePath)
  if err != nil {
    return err
  }
  defer fp.Close()

  return Output(bufio.NewScanner(fp))
}
