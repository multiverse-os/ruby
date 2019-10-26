[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)


## Multiverse OS: `ruby` Scripting Library
**URL** [multiverse-os.org](https://multiverse-os.org)

The `ruby` scripting library is currently a simple working example but will be further refined to provide scripting solutions for a varitey of Multiverse OS applications. The binary included is intended as an example, as this library is designed to function as a library for other applications to use to incorporate Ruby scripting. 

#### Initialization
The Multiverse OS `ruby` scripting library utilizes `ruby` MRI, not MRuby, and leverages entirely in memory fileless execution of an embedded Ruby binary, then leverages IPC to pass the resulting data back into Go for processing. In memory fileless execution is handled by the Multiverse OS [`memexec` library](https://github.com/multiverse-os/memexec).

The library provides a very simple solution, with no dependencies; begins with local initialization of the binary by downloading the package from a reliable third-party source and creating the embedded binary locally for greater security. 


#### Using the Library
The API is like any other encoder and just require passing a pointer to a struct to the `ruby.Marshal()` function.


==========================
# Binary Data Store Research & Brainstorming

**Use a bit-checking format like Reed Solomon** for the binary being written to memory or being written to the binary. 

Ideally don't do it in order. And save chucks twice for important data like keys.

Reed Solomon, Raptor, Fountain, Luby Trnsform (LT) RpatorQ (Forward error correction)


**GitFS based**
[gitfs](https://github.com/wade-welles/gitfs) This looks outstand and looks to be accomplishing what we want; but we really still want the complete control that comes with parsing the ELF file, segregating out the different bits. And determining exactly where our data will be insereted. And even provide us API functionality to clean up any data that doesn't need to be in the final binary.

An example can be found in C here but its not ideal; there are complete C programs that do exactly what we want. But we can also get this functionality from things like llgo or the linker that comes in Go's runtime package.

