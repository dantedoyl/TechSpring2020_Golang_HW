package main

// сюда писать код

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mtx sync.Mutex

func toString(data interface{}) (string, error) {
	if str, ok := data.(string); ok {
		return str, nil
	}

	if intStr, ok := data.(int); ok {
		return strconv.Itoa(intStr), nil
	}

	return "", errors.New("Can't convert expression to string.")
}

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for val := range in {
		wg.Add(1)
		str, _ := toString(val)

		crc32Hash := make(chan string, 1)

		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}(str, crc32Hash)

		md5Hash := make(chan string, 1)

		go func(data string, out chan<- string) {
			mtx.Lock()
			md5Data := DataSignerMd5(data)
			mtx.Unlock()
			out <- DataSignerCrc32(md5Data)
		}(str, md5Hash)

		go func() {
			defer wg.Done()
			out <- (<-crc32Hash) + "~" + (<-md5Hash)
		}()
	}
	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for val := range in {
		wg.Add(1)
		str, _ := toString(val)

		hash0 := make(chan string, 1)
		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}("0"+str, hash0)

		hash1 := make(chan string, 1)
		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}("1"+str, hash1)

		hash2 := make(chan string, 1)
		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}("2"+str, hash2)

		hash3 := make(chan string, 1)
		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}("3"+str, hash3)

		hash4 := make(chan string, 1)
		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}("4"+str, hash4)

		hash5 := make(chan string, 1)
		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}("5"+str, hash5)

		go func() {
			defer wg.Done()
			resStr := ""
			resStr = <-hash0 + <-hash1 + <-hash2 + <-hash3 + <-hash4 + <-hash5
			out <- resStr
		}()
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var strTempSlice []string
	for val := range in {
		str, _ := toString(val)
		strTempSlice = append(strTempSlice, str)
	}

	sort.Strings(strTempSlice)

	out <- strings.Join(strTempSlice, "_")
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{}, MaxInputDataLen)

	for _, currJob := range jobs {
		wg.Add(1)
		out := make(chan interface{}, MaxInputDataLen)
		go func(job job, in, out chan interface{}) {
			defer wg.Done()
			job(in, out)
			close(out)
		}(currJob, in, out)
		in = out
	}
	wg.Wait()
}
