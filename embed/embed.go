package ruby

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func DownloadRuby() {}

func EmbedRuby(workingDirectory string) {
	var downloadFilename string

	outputFile, err := os.Create(filepath.Join(workingDirectory, "bin", "ruby"))
	if err != nil {
		panic(fmt.Errorf("[fatal] failed open output file for writing:", err))
	}
	defer outputFile.Close()

	basename := strings.Split(filepath.Base(downloadFilename), ".")[0]
	if len(basename) == 0 {
		fmt.Println("[bin2go] binary name must not be empty `"+basename+"`:", err)
		os.Exit(1)
	}

	outputFile.Write([]byte("package " + strings.ToLower(basename) + "\n\nvar (\n\tBinary = []byte{"))
	binaryFile, err := ioutil.ReadFile(downloadFilename)
	if err != nil {
		fmt.Println("[bin2go] failed to open binary file:", err)
		os.Exit(1)
	}

	var data []string
	for _, binaryData := range binaryFile {
		byteString := fmt.Sprintf("%v", binaryData)
		data = append(data, byteString)
	}

	outputFile.WriteString(strings.Join(data, ", "))
	outputFile.Write([]byte("}\n)"))
	fmt.Println("[bin2go] Binary to byte slice conversion completed!")
}
