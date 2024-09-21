package common

import "time"

func CreateThrottledFunc[T func()](f T, timeout time.Duration) func(bool) {
	lastTime := time.Now().Add(timeout)

	return func(force bool) {
		if force || time.Since(lastTime) >= timeout {
			f()
			lastTime = time.Now()
		}
	}
}
