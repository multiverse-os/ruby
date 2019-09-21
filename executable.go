package ruby

import (
	"crypto/sha256"
	"io"
	"os"

	memfd "github.com/multiverse-os/ruby/memfd"
)

type Executable struct {
	*os.File
}

// https://github.com/mithro/go-archive-ar/blob/master/archive/ar/reader.go
func (fi *arFileInfoData) Name() string       { return fi.name }
func (fi *arFileInfoData) Size() int64        { return fi.size }
func (fi *arFileInfoData) Mode() os.FileMode  { return os.FileMode(fi.mode) }
func (fi *arFileInfoData) ModTime() time.Time { return time.Unix(int64(fi.modtime), 0) }
func (fi *arFileInfoData) IsDir() bool        { return fi.Mode().IsDir() }
func (fi *arFileInfoData) Sys() interface{}   { return fi }

// Extra
func (fi *arFileInfoData) UserID() int  { return fi.uid }
func (fi *arFileInfoData) GroupID() int { return fi.gid }

type ExecutableType int

const (
	cExecutable ExecutableType = iota
	rubyExecutable
	goExecutable
)

func LoadExecutable(name string, bytes []byte) *Executable {
	mem := memfd.New(name)
	bytesWritten, err := mem.Write(bytes)
	if err != nil {
		panic("[error] failed to write data to fd:" + err.Error())
	}
	return &Executable{mem.File}
}

func (self *Executable) Execute(arguments ...string) int { return 0 }

func (self *Executable) Checksum() string {
	hash := sha256.New()
	if _, err := io.Copy(hash, self.Bytes()); err != nil {
		panic(err)
	}
	return hash.Sum(nil)
}
