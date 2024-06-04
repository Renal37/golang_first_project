package main

import (
	"encoding/base64"
	"os"
	"io"
	"strings"
)

func main() {
	input := "Zm9vAGJhcg=="
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))
        io.Copy(os.Stdout, r)
}