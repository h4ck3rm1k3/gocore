
clean :
	find -name \*.o -exec rm {} \;
	find -name \*.a -exec rm {} \;

test123:
	gccgo test123.go

test1234:
	gccgo test1234.go

#export GCCGO := gccgo 
export LD_LIBRARY_PATH := $(HOME)/install/lib64/
#/gcc-1/build/powerpc64le-unknown-linux-gnu/libgo/.libs/:$(HOME)/gcc-1/build/gcc/:/home/h4ck3rm1k3/install/libexec/gcc/powerpc64le-unknown-linux-gnu/5.0.0/:/home/h4ck3rm1k3/install/libexec/gcc/powerpc64le-unknown-linux-gnu/5.0.0/
export PATH := $(HOME)/install/bin:$(PATH)
export GCCGO := $(HOME)/install/bin/gccgo
#export GOTOOLDIR := $(HOME)/gcc-1/build/powerpc64le-unknown-linux-gnu
#export GOBIN := $(HOME)/gcc-1/build/powerpc64le-unknown-linux-gnu
#export GCC_EXEC_PREFIX := $(HOME)/gcc-1/build/powerpc64le-unknown-linux-gnu
#export GOROOT := $(HOME)/go
export GOPATH := $(HOME)/testgo2
export GOROOT := $(HOME)/go
#export WORK := $(GOPATH)/src

#-L/home/h4ck3rm1k3/gcc-1/build/powerpc64le-unknown-linux-gnu/libgo/

image.sh :
#strace -o strace.txt -f 
	~/go/src/gocc build -n -v -a  -work  -gccgoflags '-c --verbose -save-temps -O0  ' -compiler gccgo github.com/h4ck3rm1k3/gocore/image/ > image.sh 2>&1

unsafe : unsafe.sh
	bash -x unsafe.sh

testcompile : image.sh
	bash -x image.sh

#image :
#	~/go/src/gocc build -v -a  -work  -gccgoflags '-c --verbose -save-temps -O0 -L/home/h4ck3rm1k3/gcc-1/build/powerpc64le-unknown-linux-gnu/libgo/ ' -compiler gccgo github.com/h4ck3rm1k3/gocore/image/

compile:
#strace -f -s 99 -o strace.txt ~/go/src/gocc build -v -x -compiler gccgo -work  -a github.com/h4ck3rm1k3/gocore/...

	~/go/src/gocc build -v -a  -work  -gccgoflags '--verbose -save-temps' -compiler gccgo github.com/h4ck3rm1k3/gocore/...

install:
#strace -f -s 99 -o strace.txt ~/go/src/gocc build -v -x -compiler gccgo -work  -a github.com/h4ck3rm1k3/gocore/...

	~/go/src/gocc install -v -a  -work  -gccgoflags '--verbose -save-temps' -compiler gccgo github.com/h4ck3rm1k3/gocore/...

setup:
	rm -rf /home/h4ck3rm1k3/testgo/src/github.com/h4ck3rm1k3/gocore/*
	cp -r ~/go/src/* /home/h4ck3rm1k3/testgo/src/github.com/h4ck3rm1k3/gocore/

#install:
# # why do I need to do this?
# 	ln -s /home/h4ck3rm1k3/gcc-1/build/gcc/liblto_plugin.*                      /home/h4ck3rm1k3/gcc-1/build/libexec/gcc/powerpc64le-unknown-linux-gnu/5.0.0/
# 	ln -s /home/h4ck3rm1k3/gcc-1/build/powerpc64le-unknown-linux-gnu/libgcc/*   /home/h4ck3rm1k3/gcc-1/build/lib/gcc/powerpc64le-unknown-linux-gnu/5.0.0/
# 	ln -s /home/h4ck3rm1k3/gcc-1/build/powerpc64le-unknown-linux-gnu/libgo/*    /home/h4ck3rm1k3/gcc-1/build/lib/gcc/
# 	ln -s /home/h4ck3rm1k3/gcc-1/build/powerpc64le-unknown-linux-gnu/libgo/*    /home/h4ck3rm1k3/gcc-1/build/lib/


src/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.o:
	/home/h4ck3rm1k3/install/bin/gccgo -I $(WORK) -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/unsafe -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe -o $(WORK)/github.com/h4ck3rm1k3/gocore/unsafe/_go_.o -c  -O0 /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.go
	/home/h4ck3rm1k3/install/bin/gccgo -c -I $(WORK)/github.com/h4ck3rm1k3/gocore/unsafe/ -o $(WORK)/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.o -D GOOS_linux -D OARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_unsafe\"" /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.s
	ar cru $(WORK)/github.com/h4ck3rm1k3/gocore/libunsafe.a $(WORK)/github.com/h4ck3rm1k3/gocore/unsafe/_go_.o $(WORK)/github.com/h4ck3rm1k3/gocore/unsafe/unsafe.o

symbols:
	find -name \*.o -print -exec nm  {} \; > symbols.txt
