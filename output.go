package mdv

import (
  "regexp"
  "bufio"
  "fmt"
  "io"
  "os"
  "os/exec"
  "bytes"

  "github.com/mgutz/ansi"
)

func Output(scanner *bufio.Scanner) error {
  buf := bytes.NewBufferString("")

  for scanner.Scan() {
    re, err := regexp.Compile("(#{1,6}) (.*)")
    if err != nil {
      return err
    }

    level := re.ReplaceAllString(scanner.Text(), "$1")
    text  := re.ReplaceAllString(scanner.Text(), "$2")
    if scanner.Text() != text {
      var format string = ""
      switch len(level) {
        case 1: format = "red+b"
        case 2: format = "green+b"
        case 3: format = "yellow+b"
        case 4: format = "blue+b"
        case 5: format = "magenta+b"
        case 6: format = "cyan+b"
      }
      buf.WriteString(ansi.Color(level + " " + text, format))
    } else {
      buf.WriteString(scanner.Text())
    }
    buf.WriteString("\n")
  }
  if err := scanner.Err(); err != nil {
    return err
  }

  pager(buf.String())

  return nil
}

func pager(text string) {
  cmd := exec.Command(os.Getenv("PAGER"))
  r, stdin := io.Pipe()
  cmd.Stdin = r
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  c := make(chan struct{})
  go func() {
    defer close(c)
    cmd.Run()
  }()

  fmt.Fprintf(stdin, text)

  stdin.Close()

  <-c
}
