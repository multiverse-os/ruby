package ruby

import (
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (self Version) String() string {
	return fmt.Sprintf("%v.%v.%v", self.Major, self.Minor, self.Patch)
}
