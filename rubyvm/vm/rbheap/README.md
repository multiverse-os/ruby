# rbheap [![Build Status](https://travis-ci.org/jimeh/rbheap.png)](https://travis-ci.org/jimeh/rbheap)

### `leak`

The leak command is intended to help track down memory leaks. By requiring three
heap dumps as input, it attempts to find memory that showed up in dump #2, and
is still there in #3.

The idea is to take a heap dump shortly after the application starts and before
it's had much of a chance to leak memory. Then take another heap dump after it's
been running for a while and leaked memory. And finally take a third heap dump
after it's been running for a while longer and leaked even more.

But comparing these three dumps and extracting only the objects which are held
in memory during heap dumps #2 and #3, we should mostly be left with objects
which are leaked memory.

```
Usage:
  rbheap leak [flags] <dump-A> <dump-B> <dump-C>

Flags:
  -f, --format string   output format: "hex" / "json" (default "hex")
  -h, --help            help for leak
  -v, --verbose         print verbose information
```

