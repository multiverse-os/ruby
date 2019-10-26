require 'ffi'
#require 'benchmark'

module Go
  extend FFI::Library
  ffi_lib './go.so'

  #module Functions
  #  extend FFI::Library
  #  ffi_lib './go.so'
  #end

  class Slice < FFI::Struct
    layout :data,  :pointer,
           :len,   :long_long,
           :cap,   :long_long 

    def self.value 
      return self.val
    end

    #write_array_of_int
    #write_array_of_uint
    #write_array_of_int8
    #write_array_of_long
    #write_array_of_short
    #write_array_of_int64
    #write_array_of_int32
    #write_array_of_int16
    #write_array_of_uint8
    #write_array_of_pointer
    #write_array_of_ulong
    #write_array_of_ushort
    #write_array_of_uint64
    #write_array_of_uint32
    #write_array_of_uint16
    #write_array_of_char
    #write_array_of_type
    #write_array_of_float
    #write_array_of_uchar
    #write_array_of_double

    def initialize(slice)
      p "slice[0].class: #{slice[0].class}"
      if slice[0].is_a?(Integer)
        self[:data] = FFI::MemoryPointer.new(:long_long, slice.size)
        self[:data].write_array_of_long_long(slice)
      elsif slice[0].class.to_s == "String" # No idea why this works but without the to_s it does not
        p "String"
        # Not working
        #self[:data] = FFI::MemoryPointer.new(:pointer, slice.size)
        #strings = Array.new 
        #slice.size.times do |index|
        #  strings << FFI::MemoryPointer.from_string(slice[index])
        ##  strings << Go::String.new(slice[index]) 
        #end
        #self[:data].write_array_of_pointer(strings)
      else 
        p "else"
      end
      self[:len] = slice.size
      self[:cap] = slice.size 
      return self
    end
  end

  # define class GoString to map:
  # C type struct { const char *p; GoInt n; }
  class String < FFI::Struct
    layout :p,     :pointer,
           :len,   :long_long

    def self.value 
      return self.val 
    end

    def initialize(str) 
      self[:p] = FFI::MemoryPointer.from_string(str) 
      self[:len] = str.size
      return self
    end
  end

  attach_function :greeter, [String.value], :void
  attach_function :int_slice_size, [Slice.value], :int
  attach_function :str_slice_size, [Slice.value], :int
end

Go.greeter(Go::String.new("test"))

p "size: #{Go.int_slice_size(Go::Slice.new([5, 3, 2]))}"
p "size: #{Go.str_slice_size(Go::Slice.new(["one", "two"]))}"
