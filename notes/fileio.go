package notes

import "os"
import "os/exec"
import "errors"
import "fmt"
import "log"

var DEFAULT_DNCONFIG_PATH = ".config/dn.conf"

func loadNote(index int16) (string, error) {
	dnConfig := os.Getenv("DN_CONFIG")
	if dnConfig == "" {
		dnConf := fmt.Sprintf("%s/%s", os.Getenv("HOME"), DEFAULT_DNCONFIG_PATH)
		log.Printf("DN_CONFIG not set opening at default %s", dnConf)
		fd, err := os.Open(dnConf)
		if err != nil {
			log.Printf("Unable to open file %s", dnConf)
			os.Exit(1)
		}

	}
	return "nil", nil
}
func AddNote(index int16, note, input string) error {

	editor := os.Getenv("EDITOR")
	if input == "-" {
		//read from stdin
		exec.Command("/bin/sh", "-c ", editor, " - ")
	}

	return nil
}

func EditNote(index int16, note, input string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("EDITOR env variable not set")
	}

	return nil
}
