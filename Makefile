
#export GCCGO := gccgo 
export LD_LIBRARY_PATH := $(HOME)/gcc-1/build/powerpc64le-unknown-linux-gnu/libgo/.libs/:$(HOME)/gcc-1/build/gcc/:/home/h4ck3rm1k3/install/libexec/gcc/powerpc64le-unknown-linux-gnu/5.0.0/:/home/h4ck3rm1k3/install/libexec/gcc/powerpc64le-unknown-linux-gnu/5.0.0/
export PATH := $(HOME)/install/bin:$(PATH)
export GCCGO := $(HOME)/gcc-1/build/gcc/gccgo
export GOTOOLDIR := $(HOME)/gcc-1/build/powerpc64le-unknown-linux-gnu
export GOBIN := $(HOME)/gcc-1/build/powerpc64le-unknown-linux-gnu
#export GCC_EXEC_PREFIX := $(HOME)/gcc-1/build/powerpc64le-unknown-linux-gnu
#export GOROOT := $(HOME)/go
export GOPATH := $(HOME)/testgo2
export GOROOT := $(HOME)/go
export WORK := /tmp/gccgo/work

module :
	strace -f -e open -o strace.txt ~/go/src/gocc build -v -a  -work  -gccgoflags '-c --verbose -save-temps -O0 -L/home/h4ck3rm1k3/gcc-1/build/powerpc64le-unknown-linux-gnu/libgo/ ' -compiler gccgo github.com/h4ck3rm1k3/gocore/image/color/palette/...

compile:
#strace -f -s 99 -o strace.txt ~/go/src/gocc build -v -x -compiler gccgo -work  -a github.com/h4ck3rm1k3/gocore/...

	~/go/src/gocc build -v -a  -work  -gccgoflags '--verbose -save-temps' -compiler gccgo github.com/h4ck3rm1k3/gocore/...

setup:
	rm -rf /home/h4ck3rm1k3/testgo/src/github.com/h4ck3rm1k3/gocore/*
	cp -r ~/go/src/* /home/h4ck3rm1k3/testgo/src/github.com/h4ck3rm1k3/gocore/
