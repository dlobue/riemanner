# riemanner

riemanner is a command-line interface to submit events to [Riemann](http://riemann.io).

## Installation

The following downloads and builds riemanner.
You must have Go 1.2 installed to do this:

```bash
mkdir /tmp/gopath
export GOPATH=/tmp/gopath
go get github.com/Clever/riemanner
mv $GOPATH/bin/riemanner /usr/local/bin/riemanner
rm -r $GOPATH
```

## Usage

```bash
$ riemanner -h
Usage of riemanner:
  -help=false: Display help text and exit.
  -port=5555: Use the specified port. The default port is 5555
  -server="localhost": Send events to the specified remote Riemann server. The default is localhost.
  -udp=false: Use UDP instead of the default stream connection (TCP).
```

## Example

riemanner accepts events over stdin as json:

```
$  echo '{"Ttl":10,"Time":0,"Tags":null,"Host":"raidman","State":"success","Service":"raidman-sample","Metric":100,"Description":""}' | riemanner
```
