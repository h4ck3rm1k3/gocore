# gocore
testing the boostrapping of go with gccgc without gc

it is checked out to ~/testgo2 and compiled with "make compile" the full error is in error.txt

You will notice changes against go core. 
I have stared to move the includes over to my gitrepo name to get around that go build problem.

The go build is from here :
git@github.com:h4ck3rm1k3/go.git branch Patch-10264
https://github.com/h4ck3rm1k3/go/tree/Patch-10264

The gcc build is here :
https://github.com/h4ck3rm1k3/gcc-1

Built like this :
../configure --enable-languages=go

all of this is checked out and build here :
/home/h4ck3rm1k3/
on the server gcc2-power8.osuosl.org
gcc10.fsffrance.org
https://gcc.gnu.org/wiki/CompileFarm
