[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse OS: `ruby` Scripting Library
**URL** [multiverse-os.org](https://multiverse-os.org)

The `ruby` scripting library is currently a simple working example but will be
further refined to provide scripting solutions for a varitey of Multiverse OS
applications. 

#### Initialization
The Multiverse OS `ruby` scripting library utilizes `ruby` MRI, not MRuby, and
leverages entirely in memory fileless execution of an embedded Ruby binary, then
leverages IPC to pass the resulting data back into Go for processing. In memory
fileless execution is handled by the Multiverse OS (`memexec`
library)[https://github.com/multiverse-os/memexec).

The library provides a very simple solution, with no dependencies; begins with
local initialization of the binary by downloading the package from a reliable
third-party source and creating the embedded binary locally for greater
security. 


#### Using the Library
The API is like any other encoder and just require passing a pointer to a struct
to the `ruby.Marshal()` function.

