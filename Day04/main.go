/* Author: Mecu Sorin       Phone: 0040747020102 */

package day04

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"runtime"
	"strconv"
	"sync"
)

func findNumberThatGenerateAMD5HashWithPatternSynchronous(secret, regexPattern string, startingFrom int) (int, error) {
	regex, err := regexp.Compile(regexPattern)
	if nil != err {
		return 0, err
	}
	i := startingFrom
	for {
		hash := getHexadecimalMD5Hash(secret + strconv.Itoa(i))
		if regex.MatchString(hash) {
			return i, nil
		}
		i++
	}
}

// don't know why is slower than synchronous variant
func findNumberThatGenerateAMD5HashWithPatternAsynchronous(secret, regexPattern string, startingFrom int) (int, error) {
	regex, err := regexp.Compile(regexPattern)
	if nil != err {
		return 0, err
	}
	workers := maxParallelism()

	jobs := make(chan int)
	done := make(chan bool)
	results := make(chan int, workers)

	go func() {
		defer close(jobs)
		i := startingFrom
		for {
			select {
			case jobs <- i:
			case <-done:
				return
			}
			i++
		}
	}()

	var wg sync.WaitGroup
	for i := workers; i > 0; i-- {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				hash := getHexadecimalMD5Hash(secret + strconv.Itoa(job))
				if regex.MatchString(hash) {
					results <- job
					done <- true
				}
			}
		}()
	}
	wg.Wait()
	found := false
	result := -1
	for {
		select {
		case r := <-results:
			if result > r || !found {
				result = r
			}
		default:
			return result, nil
		}
	}
}

func getHexadecimalMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU && numCPU > 1 {
		return maxProcs
	}
	return numCPU
}
