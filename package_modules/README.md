# Go Packages
- A ``package`` is a project directory containing ``.go`` files with the ``same package statement`` at
the beginning. A package contains many source files each ending in .go extension and
belonging to a single directory.
There are 2 types of packages:
- **executable packages** that generate executable files which can be run. The name of an
executable package is predefined and is called **main.**
- **non-executable packages** (libraries or dependencies) that are used by other packages
and can have any name. They can not be executed, only imported.


```bash
go env
```

## Output

```bash
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\PATH\USER\AppData\Local\go-build
set GOENV=C:\PATH\USER\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\PATH\USER\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\PATH\USER\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=c:\go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=c:\go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 
```