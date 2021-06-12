package cli

import (
	"io"
	"os/exec"
)

func Clip(str string) {
	c1 := exec.Command("echo", str)
	c2 := exec.Command("/usr/bin/pbcopy")

	c1stdout, _ := c1.StdoutPipe()
	c2stdin, _ := c2.StdinPipe()

	c1.Start()
	c2.Start()

	io.Copy(c2stdin, c1stdout)

	c2stdin.Close()
	c2.Wait()
}
