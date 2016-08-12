package main

import (
  "flag"
  "fmt"
  "io"
  "os"

  "github.com/masa0x80/mdv"
)

const (
  ExitCodeOK int = 0
  ExitCodeFlagPraseError
  ExitCodeBadArgs
)

type CLI struct {
  outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
  var version, edit bool

  flags := flag.NewFlagSet("mdv", flag.ContinueOnError)
  flags.SetOutput(os.Stderr)
  flags.Usage = func() {
    fmt.Fprint(os.Stderr, helpText)
  }

  flags.BoolVar(&version, "version", false, "")
  flags.BoolVar(&edit, "edit", false, "")
  flags.BoolVar(&edit, "e",    false, "")

  if err := flags.Parse(os.Args[1:]); err != nil {
    return ExitCodeFlagPraseError
  }
  if version {
    fmt.Fprintf(cli.errStream, "%s v%s\n", Name, Version)
    return ExitCodeOK
  }

  if flags.NArg() < 1 {
    fmt.Fprintln(cli.errStream, "mdv: ", fmt.Errorf("too few arguments"))
    return ExitCodeBadArgs
  }

  filePath := flags.Args()[0]

  if edit {
    mdv.Edit(filePath)
    return ExitCodeOK
  }

  if err := mdv.Init(filePath); err != nil {
    panic(err)
  }

  return ExitCodeOK
}

var helpText = `Usage: mdv [options] [path]
  mdv is a simple markdown viewer

Options:
  --edit, -e         Edit the file

  --version          Print the version of this application
`
