package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func KeyInSlice(key int, list [][]string) bool {
	for k, _ := range list {
		if k == key {
			return true
		}
	}
	return false
}

func RandStringRunes(n int) string {

	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetRemoteFile(url string) (fileName string,downloadSize int64 ) {

	subStringsSlice := strings.Split(url, "/")
	fileName = subStringsSlice[len(subStringsSlice)-1]

	resp, err := http.Head(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Is our request ok?
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		os.Exit(1)
		// exit if not ok
	}

	// the Header "Content-Length" will let us know
	// the total file size to download
	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize = int64(size)

	return
}

