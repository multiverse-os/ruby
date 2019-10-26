
There is a lingering issue with libruby that has not been fixed in Debian  that
requires the following copy command to be run: 

```
cp /usr/include/x86_64-linux-gnu/ruby-2.5.0/ruby/config.h /usr/include/ruby-2.5.0/ruby/
```

We should just specify the file, then we don't need to symbolic link when
install ruby with apt-get
//#cgo LDFLAGS: ${SRCDIR}/../target/release/libruby.so
