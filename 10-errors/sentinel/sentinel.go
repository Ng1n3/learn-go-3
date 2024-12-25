package sentinel

import (
	"archive/zip"
	"bytes"
	"fmt"
)
	//[Lecture] -> Sentinel errors are errors that are predefined and can be compared to see if they are the same. They are used to indicate that you cannot continue or start processing.

// ? Example of a sentinel error
func sentinel() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("This is not a zip file")
	}
}