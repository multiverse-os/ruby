package ruby

import (
	"github.com/multiverse-os/libexec"
)

// Alias for a more intuitive API
func New(rubyBytes []byte) *Binaries {
	binaries := &Binaries{
		Ruby: LoadExecutable("ruby", rubyBytes),
		Ruby: LoadExecutable("irb", libexec.IRB),
	}
	Binary := binaries.Ruby
	return binaries
}

func Load() *Binaries {
	// TODO: Load from the already embedded binary data
	// ...
	return nil
}

// NOTE: The binaries below are actually Ruby not C, unlike `ruby`. However the
// will remain under the Binaries struct and object, to keep things inline with
// `ruby` naming. Because by default they are in the `/bin` folder in`ruby`
// releases.

// TODO: When PortalGun is finished, we will launch a potal and compile locally

// TODO: Need to also have control over:
// [bundler, gem, irb, rdoc, erb, ri]
type Binaries struct {
	Version Version
	Ruby    *Executable
	Bundler *Executable
	Gem     *Executable
	RDoc    *Executable
	ERB     *Executable
	RI      *Executable
	Gems    []string
}

// Alias for the Ruby Binary for a more please intutive API
var Binary *Executable

// TODO: **Ideally we would statically build the Gems into this binary/package
// too, to make it work more as a Go program would be expected to. And that
// means we can carry them with us, if we say move to another system. But that
// can be carried along in a PortalGun VM

// TODO: Prefix tree for autocomplete
