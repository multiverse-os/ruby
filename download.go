package ruby

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

type DebianPackage struct {
	Files map[string][]byte
}

func DownloadDebianPackage() DebianPackage {

	deb := DebianPackage{
		Files: make(map[string][]byte),
	}

	response, err := http.Get("https://packages.debian.org/bullseye/amd64/ruby/download")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//io.Copy(&download, response.Body)
	//downloadedBytes := download.Bytes()
	//fmt.Println("size of download:", len(downloadedBytes))

	// Decompress reads in, decompresses it, and writes it to out.

	// TODO: handle both the un gzip and untar to get the files within
	//pipeRead, pipeWrite := io.Pipe()
	compressedArchive, err := gzip.NewReader(response.Body)
	if err != nil {
		panic(err)
	}

	archive := tar.NewReader(compressedArchive)
	fmt.Println("archive:", archive)
	for {
		file, err := archive.Next()
		if err == io.EOF {
			break
		}
		var fileBuffer bytes.Buffer

		fmt.Println("unarchived file:", file.Name)
		io.Copy(&fileBuffer, archive)
		fmt.Println("unarchived size:", len(fileBuffer.Bytes()))
		deb.Files[file.Name] = fileBuffer.Bytes()

	}

	return deb
}
