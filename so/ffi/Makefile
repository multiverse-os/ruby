# translang Makefile
###############################################################################
#
TARGET     = go
DLIB       = $(TARGET).so
CLEANOBJS  = *.o *.bak *.so *.h

all: build

build: clean
	go build -o $(TARGET).so -buildmode=c-shared *.go
	chmod 700 $(TARGET).so

clean:
	-@rm -rf $(CLEAN_OBJS) $(DLIB) $(TARGET).h $(TARGET).o $(TARGET)
