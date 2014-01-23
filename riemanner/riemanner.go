package riemanner

import (
	"github.com/amir/raidman"
	"io"
)

// RaidmanClient is the interface exposed by raidman.
type RaidmanClient interface {
	Send(event *raidman.Event) error
}

// Riemanner is an io.Writer that you can send events to as lines of json objects.
type Riemanner struct {
	rc    RaidmanClient
	input io.Reader
}

// NewRiemanner initializes a Riemanner.
func NewRiemanner(rc RaidmanClient, input io.Reader) Riemanner {
	return Riemanner{rc: rc, input: input}
}

func (r *Riemanner) Run() error {
	// TODO: read input using bufio.NewScanner, unmarshal line using json.Unmarshal, send event to riemann via raidman client

	return nil
}
