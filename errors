mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/errors/
mkdir -p $WORK/github.com/h4ck3rm1k3/gocore/
#COMPILER /home/h4ck3rm1k3/install/bin/gccgo
cd /home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/errors
/home/h4ck3rm1k3/install/bin/gccgo -I $WORK -c -g -fgo-pkgpath=github.com/h4ck3rm1k3/gocore/errors -fgo-relative-import-path=_/home/h4ck3rm1k3/testgo2/src/github.com/h4ck3rm1k3/gocore/errors -o $WORK/github.com/h4ck3rm1k3/gocore/errors/_go_.o -c  -O0 ./errors.go
/home/h4ck3rm1k3/install/bin/gccgo -c -I $WORK/github.com/h4ck3rm1k3/gocore/errors/ -o $WORK/github.com/h4ck3rm1k3/gocore/errors/errors.o -D GOOS_linux -D GOARCH_ppc64le -D "GOPKGPATH=\"github_com_h4ck3rm1k3_gocore_errors\"" ./errors.s
ar cru $WORK/github.com/h4ck3rm1k3/gocore/liberrors.a $WORK/github.com/h4ck3rm1k3/gocore/errors/_go_.o $WORK/github.com/h4ck3rm1k3/gocore/errors/errors.o
