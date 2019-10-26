# Go Code From Ruby 
Using the `libruby.so` we are executing Go code from Ruby without the
requirement of the `ffi` gem or `fiddle` (We do have that implemented as well
with helpers to make the process easier). We are going to use this to experiment
with building tools to communicate by building a very simple function that will
be exposed to Ruby over `*.so` then embed that shared object in the Ruby code
and send data across using the exposed function. We are experimenting with
prefixing this to scripts executed using our embedded `ruby` binary to pass data
up to the Go software that encapsulates both the `ruby` binary and the ruby
script being executed. It will provide simple access to the memfd unix socket
using the code we already wrote in Go and just exposing that access via a simple
function call. 

### Other

There is a lingering issue with libruby that has not been fixed in Debian  that
requires the following copy command to be run: 

```
cp /usr/include/x86_64-linux-gnu/ruby-2.5.0/ruby/config.h /usr/include/ruby-2.5.0/ruby/
```

We should just specify the file, then we don't need to symbolic link when
install ruby with apt-get
//#cgo LDFLAGS: ${SRCDIR}/../target/release/libruby.so
