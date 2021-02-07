# Go Modules
- Modules are supported starting with Go v1.11 but the implementation was finalized in Go
v1.13.
- But what is in fact a module?
○ A module is a collection of related Go packages stored in a directory tree with a
go.mod file at its root.
- A module contains one or more packages like a package contains one or more .go files.
- A file called go.mod defines the module’s path, which is also the import path, and its
dependency requirements.
- The go command has direct support to work with modules, including recording and
resolving dependencies on other modules.


## Create Module

```bash
go mod init hello
go build
```


```sh
$ cat go.mod
module hello

go 1.15

require github.com/ddadumitrescu/hellomod v1.0.0
```

