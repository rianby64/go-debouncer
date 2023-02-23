package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_debounce_case0(t *testing.T) {
	t.Parallel()

	noise := make(chan struct{})
	debounced := make(chan struct{})

	now := time.Now()
	noiseFinished := make(chan struct{})

	go Debounce(noise, debounced, time.Second)

	go func() {
		noise <- struct{}{}
		noiseFinished <- struct{}{}
	}()

	select {
	case <-debounced:
		elapsed := time.Since(now)

		select {
		case <-noiseFinished:
		case <-time.After(time.Second * 2):
			t.Error("finished noise not caught by timeout")
		}

		t.Log(elapsed, "elapsed")
		assert.True(t, elapsed < time.Second*2)
		assert.True(t, time.Second < elapsed)

	case <-time.After(time.Second * 2):
		t.Error("finished by timeout")
	}
}

func Test_debounce_case1(t *testing.T) {
	t.Parallel()

	noise := make(chan struct{})
	debounced := make(chan struct{})

	now := time.Now()
	noiseFinished := make(chan struct{})

	go Debounce(noise, debounced, time.Second)

	go func() {
		noise <- struct{}{}
		noise <- struct{}{}
		noise <- struct{}{}
		noise <- struct{}{}
		noise <- struct{}{}
		noise <- struct{}{}
		noise <- struct{}{}
		noise <- struct{}{}
		noise <- struct{}{}
		noiseFinished <- struct{}{}
	}()

	select {
	case <-debounced:
		elapsed := time.Since(now)

		select {
		case <-noiseFinished:
		case <-time.After(time.Second * 2):
			t.Error("finished noise not caught by timeout")
		}

		t.Log(elapsed, "elapsed")
		assert.True(t, elapsed < time.Second*2)
		assert.True(t, time.Second < elapsed)

	case <-time.After(time.Second * 2):
		t.Error("finished by timeout")
	}
}

func Test_debounce_case3(t *testing.T) {
	t.Parallel()

	noise := make(chan struct{})
	debounced := make(chan struct{})

	now := time.Now()
	noiseFinished := make(chan struct{})

	go Debounce(noise, debounced, time.Second)

	go func() {
		noise <- struct{}{}

		time.Sleep(time.Second / 2)

		noise <- struct{}{}

		time.Sleep(time.Second / 2)

		noise <- struct{}{}

		noiseFinished <- struct{}{}
	}()

	select {
	case <-debounced:
		elapsed := time.Since(now)

		select {
		case <-noiseFinished:
		case <-time.After(time.Second * 4):
			t.Error("finished noise not caught by timeout")
		}

		t.Log(elapsed, "elapsed")
		assert.True(t, elapsed < time.Second*3)
		assert.True(t, time.Second*2 <= elapsed)

	case <-time.After(time.Second * 4):
		t.Error("finished by timeout")
	}
}
