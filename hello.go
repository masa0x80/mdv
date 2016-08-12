package main

import (
  "regexp"
  "bufio"
  "fmt"
  "os"
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

  scanner := bufio.NewScanner(fp)
  for scanner.Scan() {
    re, err := regexp.Compile("(#{1,6}) (.*)")
    if err != nil {
      panic(err)
    }

    level    := re.ReplaceAllString(scanner.Text(), "$1")
    headline := re.ReplaceAllString(scanner.Text(), "$2")
    if scanner.Text() != headline {
      var template string = ""
      switch len(level) {
        case 1: template = "\x1b[31m"
        case 2: template = "\x1b[32m"
        case 3: template = "\x1b[33m"
        case 4: template = "\x1b[34m"
        case 5: template = "\x1b[35m"
        case 6: template = "\x1b[36m"
      }
      template += "%s %s\x1b[0m\n"
      fmt.Printf(template, level, headline)
    } else {
      fmt.Printf("%s\n", scanner.Text())
    }
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
