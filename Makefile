# Go+Ruby Makefile
###############################################################################
#DLIB          = $(TARGET).so
RUBY_VERSION   = 2.5
CLEANOBJS      = *.o *.bak *.so *.h go.h
RUBY_MODE      = embed
#RUBY_MODE      = vm
ARCH           = x86_64
RUBY_BIN       = /usr/bin/ruby

#all: setup ruby irb rheap
all: ruby

setup: download vendor
	cd /usr/lib/$(ARCH)-linux-gnu/ && sudo ln -s libruby.$(RUBY_VERSION).so libruby.so 

download: 
	sudo apt-get install ruby libruby ruby-dev

vendor:
	cp /usr/lib/$(ARCH)-linux-gnu/libruby-$(RUBY_VERSION).so ./include/libruby.so
	cp $(RUBY_BIN) ./include/ruby
	chmod 770 ./include/ruby
	chmod 770 ./include/libruby.so

ruby: clean
	go build -o bin/ cmd/$(RUBY_MODE)/ruby/main.go

irb: clean
	go build -o bin/ cmd/$(RUBY_MODE)/irb/main.go

rheap: clean
	go build -o bin/ cmd/rheap/main.go 

# TODO: Write tools to help generate and install go files in a ruby scripts in binary folder  
# and ensure the library supports calling those scripts easily. 
#
# WRite tools to copy downloaded ruby binary and embed it next time it needs to be updates so future updates iwll be trivial

###############################################################################

clean:
	-@rm -rf bin
	-@rm -rf include
