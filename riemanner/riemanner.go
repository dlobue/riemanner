package riemanner

import (
	"bufio"
	"encoding/json"
	"fmt"
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

// Run Takes json input and sends it to riemann
func (r *Riemanner) Run() error {

	scanner := bufio.NewScanner(r.input)

	for scanner.Scan() {

		event := &raidman.Event{}
		json_err := json.Unmarshal(scanner.Bytes(), &event)
		if json_err == nil {
			check(r.rc.Send(event))
		} else {
			fmt.Println("Skipping invalid JSON line:", jsonErr)
		}

	}
	check(scanner.Err())

	return nil
}

// Helpers

func check(e error) {
	if e != nil {
		panic(e)
	}
}
