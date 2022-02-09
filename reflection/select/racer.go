package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func WebsiteRacer(a, b string) (string, error) {
	da, err := measureResponseTime(a)
	if err != nil {
		return "", err
	}

	db, err := measureResponseTime(b)
	if err != nil {
		return "", err
	}

	if da < db {
		return a, nil
	}
	return b, nil
}

func measureResponseTime(url string) (time.Duration, error) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return time.Since(start), nil
}
