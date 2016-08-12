package main

import (
  "regexp"
  "bufio"
  "fmt"
  "os"

  "github.com/mgutz/ansi"
)

func main() {
  var fp *os.File
  var err error

  if len(os.Args) < 2 {
    fp = os.Stdin
  } else {
    fp, err = os.Open(os.Args[1])
    if err != nil {
      panic(err)
    }
    defer fp.Close()
  }

  Output(bufio.NewScanner(fp))
}

func Output(scanner *bufio.Scanner) {
  for scanner.Scan() {
    re, err := regexp.Compile("(#{1,6}) (.*)")
    if err != nil {
      panic(err)
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
      msg := ansi.Color(level + " " + text, format)
      fmt.Println(msg)
    } else {
      fmt.Printf("%s\n", scanner.Text())
    }
  }

  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
