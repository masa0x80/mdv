package mdv

import (
  "bufio"
  "os"
)

func Init() {
  var fp *os.File
  var err error

	fp, err = os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()

  Output(bufio.NewScanner(fp))
}
