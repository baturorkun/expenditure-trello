package lib

import (
	"errors"
	"expenditure/utils"
	"io/ioutil"
	"net/http"
	"os"
)

const maxUploadSize = 10 * 1024 * 1024 // 10 mb
const uploadPath = "/tmp"

func UploadFile(r *http.Request) (filename string, err error) {

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return "", errors.New("FILE_TOO_BIG")
	}

	file, _, err := r.FormFile("attachment")

	if err != nil {
		return "", errors.New("READ_ERROR")
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/zip", "application/x-zip-compressed":
	case "application/x-rar-compressed":
	case "application/pdf":
		break
	default:
		return "", errors.New("INVALID_FILE_TYPE")
	}

	filename = utils.RandStringRunes(10)

	filepath := uploadPath + "/exp_" + filename

	newFile, err := os.Create(filepath)

	if err != nil {
		return "", errors.New("CANT_WRITE_FILE")

	}
	defer newFile.Close() // idempotent, okay to call twice

	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return "", errors.New("CANT_WRITE_FILE")
	}

	return
}
