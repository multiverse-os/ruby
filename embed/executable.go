package ruby

import (
	"os"
)

type Executable struct {
	*os.File
}

type ExecutableType int

const (
	cExecutable ExecutableType = iota
	rubyExecutable
	goExecutable
)

//func LoadExecutable(name string, bytes []byte) *Executable {
//	mem := memfd.New(name)
//	bytesWritten, err := mem.Write(bytes)
//	if err != nil {
//		panic("[error] failed to write data to fd:" + err.Error())
//	}
//	fmt.Println("bytesWritten:", bytesWritten)
//	return &Executable{mem.File}
//}
//
//func (self *Executable) Execute(arguments ...string) int { return 0 }

//func (self *Executable) Checksum() string {
//	hash := sha256.New()
//	if _, err := io.Copy(hash, self.Bytes()); err != nil {
//		panic(err)
//	}
//	return hash.Sum(nil)
//}
