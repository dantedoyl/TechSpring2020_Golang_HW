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

type multiHashType struct {
	th   int
	hash string
}

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

		crc32Hash := make(chan string)

		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}(str, crc32Hash)

		md5Hash := make(chan string)

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

		innerWg := &sync.WaitGroup{}
		hash := make(chan multiHashType)

		for th := 0; th < 6; th++ {
			innerWg.Add(1)
			go func(data string, th int, out chan<- multiHashType) {
				out <- multiHashType{
					th:   th,
					hash: DataSignerCrc32(data),
				}
				innerWg.Done()
			}(strconv.Itoa(th)+str, th, hash)
		}

		go func(innerWg *sync.WaitGroup, ch chan multiHashType) {
			innerWg.Wait()
			close(hash)
		}(innerWg, hash)

		go func() {
			defer wg.Done()
			var arr = make([]string, 6)
			for h := range hash {
				arr[h.th] = h.hash
			}
			out <- strings.Join(arr, "")
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
