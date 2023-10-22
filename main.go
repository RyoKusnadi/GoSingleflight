package main

import (
	"errors"
	"log"
	"sync"

	"golang.org/x/sync/singleflight"
)

const (
	concurrentLimit = 10
)

var (
	errorNotExist = errors.New("not exist")
	key           = "key"
	sfg           = singleflight.Group{}
	isRunWithSF   = true
)

func main() {
	if isRunWithSF {
		simulateConcurrentProcessesWithSf()
	} else {
		simulateConcurrentProcessesWithNoSf()
	}
}

func simulateConcurrentProcessesWithNoSf() {
	var wg sync.WaitGroup
	wg.Add(concurrentLimit)

	for i := 0; i < concurrentLimit; i++ {
		go func() {
			defer wg.Done()
			data, err := getDataWithNoSf(key)
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}

func simulateConcurrentProcessesWithSf() {
	var wg sync.WaitGroup
	wg.Add(concurrentLimit)

	for i := 0; i < concurrentLimit; i++ {
		go func() {
			defer wg.Done()
			data, err := getDataWithSf(key)
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}

func getDataWithNoSf(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		data, err = getDataFromDB(key)
		if err != nil {
			log.Println(err)
			return "", err
		}

	} else if err != nil {
		return "", err
	}
	return data, nil
}

func getDataWithSf(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errorNotExist {
		v, err, _ := sfg.Do(key, func() (interface{}, error) { // The Difference Is In Here
			return getDataFromDB(key)
		})
		if err != nil {
			log.Println(err)
			return "", err
		}

		data = v.(string)
	} else if err != nil {
		return "", err
	}
	return data, nil
}

// Simulate retrieving value from the cache, where the cache doesn't have the value
func getDataFromCache(key string) (string, error) {
	return "", errorNotExist
}

// Simulate fetching data from the database
func getDataFromDB(key string) (string, error) {
	log.Printf("get %s from the database", key)
	return "sample result", nil
}
