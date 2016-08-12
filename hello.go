package main

import (
  "regexp"
  "bufio"
  "fmt"
  "os"
)

const (
  Black   = "\x1b[30m"
  Red     = "\x1b[31m"
  Green   = "\x1b[32m"
  Yellow  = "\x1b[33m"
  Blue    = "\x1b[34m"
  Magenta = "\x1b[35m"
  Cyan    = "\x1b[36m"
  White   = "\x1b[37m"

  Underline = "\x1b[4m"
  Bold      = "\x1b[1m"
  Reverse   = "\x1b[7m"

  Reset = "\x1b[0m"
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
      var template string = ""
      switch len(level) {
        case 1: template = Red
        case 2: template = Green
        case 3: template = Yellow
        case 4: template = Blue
        case 5: template = Magenta
        case 6: template = Cyan
      }
      template += Bold + "%s %s" + Reset + "\n"
      fmt.Printf(template, level, text)
    } else {
      fmt.Printf("%s\n", scanner.Text())
    }
  }

  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
