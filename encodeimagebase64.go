package tommy

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
)

// EncodeImageToBase64 trasfmorma una immagine nella sua versione base64.
func EncodeImageToBase64(imagefile string) (string, error) {
	// Open file on disk.
	f, err := os.Open(imagefile)
	if err != nil {
		return "", err
	}

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded, nil
}
