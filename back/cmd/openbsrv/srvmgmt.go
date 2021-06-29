package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/OpenEugene/openboard/back/internal/log"
)

type server interface {
	Serve() error
	Stop() error
}

type serverMgmt struct {
	ss  []server
	log *log.Log
}

func newServerMgmt(log *log.Log, ss ...server) *serverMgmt {
	return &serverMgmt{
		ss:  ss,
		log: log,
	}
}

func (m *serverMgmt) serve() error {
	var wg sync.WaitGroup
	wg.Add(len(m.ss))

	for _, s := range m.ss {
		go func(s server) {
			defer wg.Done()

			// TODO: gather returned errors
			if err := s.Serve(); err != nil {
				fmt.Fprintln(os.Stderr, "server error:", err)
			}
		}(s)

		time.Sleep(time.Millisecond * 200)
	}

	wg.Wait()

	return nil
}

func (m *serverMgmt) stop() error {
	for _, s := range m.ss {
		go func(s server) {
			// TODO: gather returned errors
			if err := s.Stop(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}(s)
	}

	return nil
}
