package utils

import "time"

func Benchmark(fn func() error, name string) error {
	timer := time.Now()
	err := fn()
	if err != nil {
		return err
	}

	Log("%s took %dms", name, time.Since(timer).Milliseconds())
	return nil
}
