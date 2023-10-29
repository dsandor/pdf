package main

import (
	"fmt"
	"strings"

	"github.com/dsandor/pdf"
)

func main() {

	pdf.DebugOn = true
	content, err := readPdf("ds-resume.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	b, err := r.GetPlainTextSlice()
	if err != nil {
		return "", err
	}

	return strings.Join(b, "\n"), nil
}
