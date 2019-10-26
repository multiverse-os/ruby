package leak

import (
	"fmt"
	"time"

	"github.com/jimeh/rbheap/obj"
)

// NewFinder returns a new *Finder instance, populated with the three given file
// paths.
func NewFinder(filePath1, filePath2, filePath3 string) *Finder {
	return &Finder{
		FilePaths: [3]string{filePath1, filePath2, filePath3},
	}
}

// Finder helps with finding a memory leak across three different memory dumps
// from a Ruby process.
type Finder struct {
	FilePaths [3]string
	Dumps     [3]*obj.Dump
	Leaks     []*string
	Verbose   bool
}

// Process will will load and process each of the dump files.
func (s *Finder) Process() error {
	for i, filePath := range s.FilePaths {
		start := time.Now()
		s.log(fmt.Sprintf("Parsing %s", filePath))
		dump := obj.NewDump(filePath)

		err := dump.Process()
		if err != nil {
			return err
		}

		s.Dumps[i] = dump
		elapsed := time.Now().Sub(start)
		s.log(fmt.Sprintf(
			"Parsed %d objects in %.6f seconds",
			len(dump.Index),
			elapsed.Seconds(),
		))
	}

	return nil
}

// PrintLeakedAddresses prints the memory addresses in hex (0x...) format for
// all objects which are likely to be leaked memory.
func (s *Finder) PrintLeakedAddresses() {
	s.log("\nLeaked Addresses:")
	s.Dumps[1].PrintEntryAddress(s.FindLeaks())
}

// PrintLeakedObjects prints the full JSON blobs for all objects which are
// likely to be memory leaks.
func (s *Finder) PrintLeakedObjects() error {
	s.log("\nLeaked Objects:")
	return s.Dumps[1].PrintEntryJSON(s.FindLeaks())
}

// FindLeaks finds potential memory leaks by removing all objects in heap dump
// #1 from heap dump #2, and then also removing all entries from heap dump #2
// which are not present in heap dump #3.
func (s *Finder) FindLeaks() []*string {
	if s.Leaks != nil {
		return s.Leaks
	}

	mapA := map[string]bool{}
	mapC := map[string]bool{}

	for _, x := range s.Dumps[0].Index {
		mapA[*x] = true
	}

	for _, x := range s.Dumps[2].Index {
		mapC[*x] = true
	}

	for _, x := range s.Dumps[1].Index {
		_, okA := mapA[*x]
		_, okC := mapC[*x]

		if !okA && okC {
			s.Leaks = append(s.Leaks, x)
		}
	}

	return s.Leaks
}

func (s *Finder) log(msg string) {
	if s.Verbose {
		fmt.Println(msg)
	}
}
