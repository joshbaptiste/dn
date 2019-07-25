package notes

import "os"
import "os/exec"
import "errors"

func EditNote(index int16, note string, input string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("EDITOR env variable not set")
	}

	if input == "-" {
		//read from stdin
		exec.Command("/bin/sh", "-c ", editor, " - ")
	}

	return nil
}
