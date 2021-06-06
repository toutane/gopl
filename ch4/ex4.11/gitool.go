// Gitool let the user create, read, update and delete an issue on GitHub.
package main

import (
	"fmt"

	"github.com/toutane/gopl/ch4/ex4.11/cli"
)

func main() {
	data, err := cli.GetInputFromEditor()
	if err != nil {
		fmt.Printf("Fail at getting input from editor: %s\n", err)
	}
	fmt.Printf("%s\n", data)
}
