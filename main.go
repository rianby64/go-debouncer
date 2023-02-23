package main

import "time"

const (
	defaultDebounceTimeout = time.Second * time.Duration(5)
)

func Debounce(noise <-chan struct{}, ready chan<- struct{}, timeout time.Duration) {
	if timeout == 0 {
		timeout = defaultDebounceTimeout
	}

	timer := time.NewTimer(timeout)
	timer.Stop()

	go func() {
		for {
			select {
			case <-noise:
				timer.Reset(timeout)
			case <-timer.C:
				ready <- struct{}{}

				timer.Stop()
			}
		}
	}()
}
