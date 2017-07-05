package roper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// Unmarshal ...
func Unmarshal(input string, result interface{}) error {

	if input == "" {
		return errors.New("Missing input")
	}

	bytes, err := getBytesFromInput(input)

	if err != nil {
		return err
	}

	if bytes == nil {
		return errors.New("File contained no data")
	}

	if input == "-" {
		if yaml.Unmarshal(bytes, &result) != nil {
			if json.Unmarshal(bytes, &result) != nil {
				return errors.New("failed to decode from stdin")
			}
		}
		return nil
	}

	ext := path.Ext(input)

	switch strings.ToLower(ext) {

	case ".yaml", ".yml":
		err = yaml.Unmarshal(bytes, &result)

	case ".json":
		err = json.Unmarshal(bytes, &result)

	default:
		err = fmt.Errorf("Unsupported file type: %s", ext)
	}

	return err
}

// getBytesFromInput read info from STDIN, URL or from a file.
func getBytesFromInput(inputFile string) ([]byte, error) {

	switch {

	case inputFile == "-":
		fi, err := os.Stdin.Stat()
		if err != nil {
			return nil, err
		}

		if fi.Mode()&os.ModeNamedPipe == 0 {
			return nil, errors.New("Not a stdin pipe")
		}

		return ioutil.ReadAll(os.Stdin)

	case strings.Index(inputFile, "http://") == 0 || strings.Index(inputFile, "https://") == 0:
		_, err := url.Parse(inputFile)
		if err != nil {
			return nil, err
		}

		resp, err := http.Get(inputFile)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		return ioutil.ReadAll(resp.Body)

	default:
		bytes, err := ioutil.ReadFile(inputFile)
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, errors.New("File empty")
		}

		return bytes, err
	}

}
