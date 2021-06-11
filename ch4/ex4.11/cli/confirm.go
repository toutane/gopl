package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskForConfirm(s string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("%s ?  [y/n]: ", s)

		resp, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		resp = strings.ToLower(strings.TrimSpace(resp))

		if resp == "y" || resp == "yes" {
			return true, nil
		} else if resp == "n" || resp == "no" {
			return false, nil
		}
	}

}
