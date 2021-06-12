// Cli provides the implementation for invoking user's preferred text editor in
// the command line.
package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const DefaultEditor = "vim"

var editor = os.Getenv("EDITOR")

func CreateBody() (string, error) {
	if editor == "" {
		editor = DefaultEditor
	}

	fmt.Printf("\n\n? Body (press [e] to launch %s, other to skip)", editor)

	key, err := listenToKey()
	if err != nil {
		return "", err
	}

	if key == "e" {
		bytes, err := GetInputFromEditor("")
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	return "", nil
}

func UpdateContent(content, name string) (string, error) {
	if editor == "" {
		editor = DefaultEditor
	}

	fmt.Printf("\n\n? Update %s (press [e] to launch %s, other to skip)", name, editor)

	key, err := listenToKey()
	if err != nil {
		return "", err
	}

	if key == "e" {
		bytes, err := GetInputFromEditor(content)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	return content, nil
}

func listenToKey() (string, error) {
	cmd1 := exec.Command("stty", "cbreak", "min", "1")
	cmd1.Stdin = os.Stdin
	err := cmd1.Run()
	if err != nil {
		return "", err
	}

	cmd2 := exec.Command("stty", "-echo")
	cmd2.Stdin = os.Stdin
	err = cmd2.Run()
	if err != nil {
		return "", err
	}

	result := make(chan string)

	go startStdioLoop(result)

	value := <-result

	close(result)

	return string(value), nil
}

func startStdioLoop(result chan string) {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		result <- string(b)
		cmd := exec.Command("stty", "-echo")
		cmd.Stdin = os.Stdin
		cmd.Run()
		return
	}
}

// OpenFileInEditor opens a file in a text editor.
func OpenFileInEditor(filename string) error {
	if editor == "" {
		editor = DefaultEditor
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func GetInputFromEditor(content string) ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")

	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()
	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return []byte{}, err
	}

	// Defer removal of the temporary file in case any of the next steps
	// fail.
	defer os.Remove(filename)

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = OpenFileInEditor(filename); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
