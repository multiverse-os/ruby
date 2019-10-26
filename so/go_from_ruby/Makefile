# Go+Ruby Makefile
###############################################################################
#
TARGET        = go
RUBY_VERSION  = 2.5 #Would be better to just check the standard folders and get the version
RVM           = 0 
DLIB          = $(TARGET).so
CLEANOBJS     = *.o *.bak *.so *.h go.h

all: build

build: clean
	go build -o $(TARGET).so -buildmode=c-shared *.go
	chmod 700 $(TARGET).so

setup:
	cd /usr/lib/x86_64-linux-gnu/ && sudo ln -s libruby.so.$(RUBY_VERSION) libruby.so

clean:

