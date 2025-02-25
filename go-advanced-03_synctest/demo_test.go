package main

import (
	"context"
	"testing"
	"time"

	"github.com/mkch/sync2"

	"testing/synctest"
)

func TestSleep(t *testing.T) {
	const D = time.Second * 2

	start := time.Now()

	sync2.Sleep(context.Background(), D)

	elapse := time.Since(start)

	if elapse < D || elapse-D > D/100 {
		t.Fatal(elapse - D)
	}
}

func TestSleep2(t *testing.T) {
	synctest.Run(func() { TestSleep(t) })
}
