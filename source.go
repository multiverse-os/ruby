package ruby

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Feed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Language    string `xml:"language"`
	TTL         string `xml:"ttl"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PublishDate string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Description string `xml:"description"`
}

type Release struct {
	Version  Version
	URL      string
	Checksum ReleaseChecksums
	Files    map[string][]byte
}

type ReleaseChecksums struct {
	SHA1   string
	SHA256 string
	SHA512 string
}

func LoadReleasesFeed() (Feed, error) {
	response, err := http.Get("https://www.ruby-lang.org/en/feeds/news.rss")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	feed := Feed{}
	if err := xml.Unmarshal(data, &feed); err != nil {
		panic(err)
	}
	return feed, nil
}

func (self Feed) Newest() Release {
	release := Release{
		Files: make(map[string][]byte),
	}
	item := self.Channel.Items[0]

	versionRegex := regexp.MustCompile("\\d{0,2}\\.\\d{0,2}\\.\\d{0,2}")
	versionSection := versionRegex.FindStringSubmatch(item.Title)
	version := versionSection[0]
	versionParts := strings.Split(version, ".")
	major, _ := strconv.Atoi(versionParts[0])
	minor, _ := strconv.Atoi(versionParts[1])
	patch, _ := strconv.Atoi(versionParts[2])

	release.Version = Version{Major: major, Minor: minor, Patch: patch}

	ulRegex := regexp.MustCompile("\\<ul\\>(.|\\s)*?\\</ul\\>")
	liRegex := regexp.MustCompile("\\<li\\>(.|\\s)*?\\</li\\>")
	codeRegex := regexp.MustCompile("\\<code\\>(.|\\s)*?\\</code\\>")
	aRegex := regexp.MustCompile("\\<a href=\"(.)*?\"\\>")

	ulSections := ulRegex.FindAllStringSubmatch(item.Description, -1)
	lastUL := ulSections[(len(ulSections) - 1)][0]

	liSections := liRegex.FindAllStringSubmatch(lastUL, -1)
	latestRelease := liSections[1][0]
	codeSection := codeRegex.FindStringSubmatch(latestRelease)

	aHrefSection := aRegex.FindStringSubmatch(latestRelease)
	aHrefSection = strings.Split(aHrefSection[0], "\"")
	downloadURL := aHrefSection[1]
	release.URL = downloadURL

	codeBlock := codeSection[0]
	codeBlockLines := strings.Split(codeBlock, "\n")
	checksums := ReleaseChecksums{}
	for _, line := range codeBlockLines {
		lineSections := strings.Split(line, ":")
		if len(lineSections) == 2 {
			checksum := strings.TrimSpace(lineSections[1])
			switch lineSections[0] {
			case "SHA1":
				checksums.SHA1 = checksum
			case "SHA256":
				checksums.SHA256 = checksum
			case "SHA512":
				checksums.SHA512 = checksum
			}
		}
	}
	release.Checksum = checksums
	return release
}

func (self Release) Download() Release {
	response, err := http.Get(self.URL)
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
		self.Files[file.Name] = fileBuffer.Bytes()

	}

	return self
}
