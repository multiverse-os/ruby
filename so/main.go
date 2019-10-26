package main

// #cgo CFLAGS: -O2 -fpic -I/usr/include/ruby-2.5.0/
// #cgo LDFLAGS:  -Wl,-rpath,/usr/lib/x86_64-linux-gnu/ -L/usr/lib/x86_64-linux-gnu/ -lruby -lm -lc -rdynamic -lm -ldl -lcrypt -lpthread -lobjc
// #include <stdlib.h>
// #include <ruby.h>
import "C"
import (
//"fmt"
//"unsafe"
)

//const ModuleRoot = uintptr(0)
//
//var gcmap = map[interface{}]bool{}
//
//func init() {
//	call := C.CString("call")
//	defer C.free(unsafe.Pointer(call))
//	enumFor := C.CString("enum_for")
//	defer C.free(unsafe.Pointer(enumFor))
//}

//func DefineModule(name string) uintptr {
//	cname := C.CString(name)
//	defer C.free(unsafe.Pointer(cname))
//	return uintptr(C.rb_define_module(cname))
//}

//export Init_trans
func Init_goruby() {
	//C.ruby_init()
	C.ruby_init()
	//C.ruby_init_loadpath()
	//ok := C.ruby_setup()
	_ = C.rb_define_module(C.CString("GoRuby"))
	//mod := DefineModule("Trans")

	//C.rb_define_module_function(mod, "shutdown",
}

func main() {}
