# Go+Ruby Makefile
###############################################################################
#DLIB          = $(TARGET).so
RUBY_VERSION   = 2.5
CLEANOBJS      = *.o *.bak *.so *.h go.h

#all: setup ruby irb rheap
all: ruby

setup:
	sudo apt-get install ruby libruby ruby-dev
	cd /usr/lib/x86_64-linux-gnu/ && sudo ln -s libruby.$(RUBY_VERSION).so libruby.so 
	cp /usr/lib/x86_64-linux-gnu/libruby-$(RUBY_VERSION).so

ruby: clean
	go build -o bin/ cmd/ruby/main.go

irb: clean
	go build -o bin/ cmd/irb/main.go

rheap: clean
	go build -o bin/ cmd/rheap/main.go

# TODO: Write tools to help generate and install go files in a ruby scripts in binary folder  
# and ensure the library supports calling those scripts easily. 
#
# WRite tools to copy downloaded ruby binary and embed it next time it needs to be updates so future updates iwll be trivial

###############################################################################

clean:
	-@rm -rf bin/ruby 
	-@rm -rf bin/irb
	-@rm -rf bin/rheap
