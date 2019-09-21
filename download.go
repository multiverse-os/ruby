package ruby

import (
	"fmt"
)

type Feed struct {
}

const rubyReleasesRSS = "http://www.ruby-lang.org/en/feeds/news.rss"

type ReleaseChecksums struct {
	SHA1   string
	SHA256 string
	SHA512 string
}

// TODO: Or maybe CommitsURL since its not []string of commits
type Release struct {
	URL         string
	Title       string
	Description string
	Commits     string
	Comment     string
	Size        int
	Checksum    ReleaseChecksums
}

func LoadReleasesFeed() (feed *Feed) {
	if feed, err := rss.Fetch("http://example.com/rss"); err != nil {
		panic(err)
	}
	return feed
}

func (self *Feed) Newest() *Release {
	if err = feed.Update(); err != nil {
		panic(err)
	}
	self.Update()

}
