package stdlib

//Ruby Standard Library¶ ↑
//
//The Ruby Standard Library is a vast collection of classes and modules that you can require in your code for additional features.
//
//Below is an overview of libraries and extensions followed by a brief description.
//Libraries¶ ↑
//
//Abbrev
//
//    Calculates a set of unique abbreviations for a given set of strings
//Base64
//
//    Support for encoding and decoding binary data using a Base64 representation
//Benchmark
//
//    Provides methods to measure and report the time used to execute code
//CGI
//
//    Support for the Common Gateway Interface protocol
//DEBUGGER__
//
//    Debugging functionality for Ruby
//Delegator
//
//    Provides three abilities to delegate method calls to an object
//DRb
//
//    Distributed object system for Ruby
//English.rb
//
//    Require 'English.rb' to reference global variables with less cryptic names
//ERB
//
//    An easy to use but powerful templating system for Ruby
//Find
//
//    This module supports top-down traversal of a set of file paths
//GetoptLong
//
//    Parse command line options similar to the GNU C getopt_long()
//MakeMakefile
//
//    Module used to generate a Makefile for C extensions
//Monitor
//
//    Provides an object or module to use safely by more than one thread
//Net::FTP
//
//    Support for the File Transfer Protocol
//Net::HTTP
//
//    HTTP client api for Ruby
//Net::IMAP
//
//    Ruby client api for Internet Message Access Protocol
//Net::POP3
//
//    Ruby client library for POP3
//Net::SMTP
//
//    Simple Mail Transfer Protocol client library for Ruby
//Observable
//
//    Provides a mechanism for publish/subscribe pattern in Ruby
//OpenURI
//
//    An easy-to-use wrapper for Net::HTTP, Net::HTTPS and Net::FTP
//Open3
//
//    Provides access to stdin, stdout and stderr when running other programs
//OptionParser
//
//    Ruby-oriented class for command-line option analysis
//PP
//
//    Provides a PrettyPrinter for Ruby objects
//PrettyPrinter
//
//    Implements a pretty printing algorithm for readable structure
//profile.rb
//
//    Runs the Ruby Profiler__
//Profiler__
//
//    Provides a way to profile your Ruby application
//PStore
//
//    Implements a file based persistence mechanism based on a Hash
//Racc
//
//    A LALR(1) parser generator written in Ruby.
//RbConfig
//
//    Information of your configure and build of Ruby
//resolv-replace.rb
//
//    Replace Socket DNS with Resolv
//Resolv
//
//    Thread-aware DNS resolver library in Ruby
//Rinda
//
//    The Linda distributed computing paradigm in Ruby
//Gem
//
//    Package management framework for Ruby
//SecureRandom
//
//    Interface for secure random number generator
//Set
//
//    Provides a class to deal with collections of unordered, unique values
//Shellwords
//
//    Manipulates strings with word parsing rules of UNIX Bourne shell
//Singleton
//
//    Implementation of the Singleton pattern for Ruby
//Tempfile
//
//    A utility class for managing temporary files
//Time
//
//    Extends the Time class with methods for parsing and conversion
//Timeout
//
//    Auto-terminate potentially long-running operations in Ruby
//tmpdir.rb
//
//    Extends the Dir class to manage the OS temporary file path
//TSort
//
//    Topological sorting using Tarjan's algorithm
//un.rb
//
//    Utilities to replace common UNIX commands
//URI
//
//    A Ruby module providing support for Uniform Resource Identifiers
//WeakRef
//
//    Allows a referenced object to be garbage-collected
//YAML
//
//    Ruby client library for the Psych YAML implementation
//
//Extensions¶ ↑
//
//Coverage
//
//    Provides coverage measurement for Ruby
//Digest
//
//    Provides a framework for message digest libraries
//IO
//
//    Extensions for Ruby IO class, including wait and ::console
//NKF
//
//    Ruby extension for Network Kanji Filter
//objspace
//
//    Extends ObjectSpace module to add methods for internal statistics
//Pathname
//
//    Representation of the name of a file or directory on the filesystem
//PTY
//
//    Creates and manages pseudo terminals
//Readline
//
//    Provides an interface for GNU Readline and Edit Line (libedit)
//Ripper
//
//    Provides an interface for parsing Ruby programs into S-expressions
//Socket
//
//    Access underlying OS socket implementations
//Syslog
//
//    Ruby interface for the POSIX system logging facility
//WIN32OLE
//
//    Provides an interface for OLE Automation in Ruby
//
//Default gems¶ ↑
//Libraries¶ ↑
//
//Bundler
//
//    Manage your Ruby application's gem dependencies
//CMath
//
//    Provides Trigonometric and Transcendental functions for complex numbers
//CSV
//
//    Provides an interface to read and write CSV files and data
//E2MM
//
//    Module for defining custom exceptions with specific messages
//FileUtils
//
//    Several file utility methods for copying, moving, removing, etc
//Forwardable
//
//    Provides delegation of specified methods to a designated object
//IPAddr
//
//    Provides methods to manipulate IPv4 and IPv6 IP addresses
//IRB
//
//    Interactive Ruby command-line tool for REPL (Read Eval Print Loop)
//Logger
//
//    Provides a simple logging utility for outputting messages
//Matrix
//
//    Represents a mathematical matrix.
//Mutex_m
//
//    Mixin to extend objects to be handled like a Mutex
//OpenStruct
//
//    Class to build custom data structures, similar to a Hash
//Prime
//
//    Prime numbers and factorization library
//RDoc
//
//    Produces HTML and command-line documentation for Ruby
//REXML
//
//    An XML toolkit for Ruby
//RSS
//
//    Family of libraries that support various formats of XML “feeds”
//Scanf
//
//    A Ruby implementation of the C function scanf(3)
//Shell
//
//    An idiomatic Ruby interface for common UNIX shell commands
//Synchronizer
//
//    A module that provides a two-phase lock with a counter
//ThreadsWait
//
//    Watches for termination of multiple threads
//Tracer
//
//    Outputs a source level execution trace of a Ruby program
//WEBrick
//
//    An HTTP server toolkit for Ruby
//
//Extensions¶ ↑
//
//BigDecimal
//
//    Provides arbitrary-precision floating point decimal arithmetic
//Date
//
//    A subclass of Object includes Comparable module for handling dates
//DateTime
//
//    Subclass of Date to handling dates, hours, minutes, seconds, offsets
//DBM
//
//    Provides a wrapper for the UNIX-style Database Manager Library
//Etc
//
//    Provides access to information typically stored in UNIX /etc directory
//Fcntl
//
//    Loads constants defined in the OS fcntl.h C header file
//Fiddle
//
//    A libffi wrapper for Ruby
//GDBM
//
//    Ruby extension for the GNU dbm (gdbm) library
//IO::console
//
//    Console interface
//JSON
//
//    Implements Javascript Object Notation for Ruby
//OpenSSL
//
//    Provides SSL, TLS and general purpose cryptography for Ruby
//Psych
//
//    A YAML parser and emitter for Ruby
//SDBM
//
//    Provides a simple file-based key-value store with String keys and values
//StringIO
//
//    Pseudo I/O on String objects
//StringScanner
//
//    Provides lexical scanning operations on a String
//Zlib
//
//    Ruby interface for the zlib compression/decompression library
//
//Bundled gems¶ ↑
//Libraries¶ ↑
//
//DidYouMean
//
//    “Did you mean?” experience in Ruby
//MiniTest
//
//    A test suite with TDD, BDD, mocking and benchmarking
//Net::Telnet
//
//    Telnet client library for Ruby
//PowerAssert
//
//    Power Assert for Ruby.
//Rake
//
//    Ruby build program with capabilities similar to make
//Test::Unit
//
//    A compatibility layer for MiniTest
//XMLRPC
//
//    Remote Procedure Call over HTTP support for Ruby
//
//
//
