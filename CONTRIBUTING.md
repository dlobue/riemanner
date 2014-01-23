## Code Layout

The files in the top level directory define the `riemanner` executable, usable on the command line.
This is a consumer of the `riemanner` module in the riemanner directory.

## Developing

Prerequisites:

* Install Go 1.2. Ensure this is the version of Go you're running with `go version`.
* Make sure your GOPATH is set, e.g. `export GOPATH=~/go`.
* Make sure `$GOPATH/bin` is in your PATH.
* Install `golint` with `go get github.com/golang/lint/golint`. Linting is done automatically as part of running tests.
* Clone the repository to a location outside your GOPATH, and symlink it to `$GOPATH/src/github.com/Clever/riemanner`.
If you have [gvm](https://github.com/moovweb/gvm) installed, you can make this symlink easily by running the following from the root of where you have cloned the repository: `gvm linkthis github.com/Clever/riemanner`.
The reason for making this symlink is that the Makefile runs tests, builds, and installs by specifying the full package or sub-package name, e.g. `go test github.com/Clever/riemanner`. Running `go ...` commands specifying the full package name eliminates a lot of the complexity of running these tools specifying a relative path.

If you have done all of the above, then you should be able to run

```
make
```

This will test, build, and install riemanner and all sub-packages.

To see code coverage across all packages you can run

```
COVERAGE=1 make
```
