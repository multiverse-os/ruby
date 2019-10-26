package main

// #cgo CFLAGS: -O2 -fpic -I/usr/include/ruby-2.5.0/
// #cgo LDFLAGS:  -Wl,-rpath,/usr/lib/x86_64-linux-gnu/ -L/usr/lib/x86_64-linux-gnu/libruby.so.2.5 -lruby -lm -lc -rdynamic -lm -ldl -lcrypt -lpthread -lobjc
// #include <stdlib.h>
// #include <ruby.h>
//
//static char *rstring_ptr(VALUE rstring) {
//  return RSTRING_PTR(rstring);
//}
//
//static int rstring_len(VALUE rstring) {
//  return RSTRING_LEN(rstring);
//}
import "C"

import (
	"bufio"
	"fmt"
	"os"
)

func rstringToGostring(rstring C.VALUE) string {
	return C.GoStringN(C.rstring_ptr(rstring), C.rstring_len(rstring))
}

func evalString(code string) C.VALUE {
	val := C.gorby_eval(C.CString(code))
	return C.rb_funcall(val, C.rb_intern(C.CString("inspect")), 0)
}

func main() {
	C.boot_vm()
	r := bufio.NewReader(os.Stdout)
	for {
		fmt.Printf("ruby> ")
		line, _, _ := r.ReadLine()
		result := evalString(string(line))
		ln := rstringToGostring(result)
		fmt.Printf("=> %s\n", ln)
	}
}
