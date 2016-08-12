package mdv

import (
  "os"
  "os/exec"
  "syscall"
)

func Edit(filePath string) {
  binary, lookErr := exec.LookPath(os.Getenv("EDITOR"))
  if lookErr != nil {
    panic(lookErr)
  }

  args := []string{os.Getenv("EDITOR"), filePath}
  env  := os.Environ()

  execErr := syscall.Exec(binary, args, env)
  if execErr != nil {
    panic(execErr)
  }
}
