// Cli provides the implementation for invoking user's preferred text editor in
// the command line.
package cli

import (
	"io/ioutil"
	"os"
	"os/exec"
)

const DefaultEditor = "vim"

// OpenFileInEditor opens a file in a text editor.
func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
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

func GetInputFromEditor() ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")

	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

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
