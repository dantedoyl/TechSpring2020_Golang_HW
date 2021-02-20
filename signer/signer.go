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
	if str, ok := data.(string); ok{
		return str, nil
	}

	if intStr, ok := data.(int); ok{
		return strconv.Itoa(intStr), nil
	}

	return "", errors.New("Can't convert expression to string.")
}

func SingleHash(in, out chan interface{}){
	wg := &sync.WaitGroup{}
	for val := range in {
		wg.Add(1)
		str, _ := toString(val)

		hashes := make([]chan string, 2)
		for i := range hashes {
			hashes[i] = make(chan string, 1)
		}

		// crc32(data)
		go func(data string, out chan<- string) {
			out <- DataSignerCrc32(data)
		}(str, hashes[0])

		// crc32(md5(data))
		go func(data string, out chan<- string) {
			mtx.Lock()
			md5Data := DataSignerMd5(data)
			mtx.Unlock()
			out <- DataSignerCrc32(md5Data)
		}(str, hashes[1])

		// result hash
		go func() {
			defer wg.Done()
			out <- (<-hashes[0])+"~"+(<-hashes[1])
		}()
	}
	wg.Wait()
}

func MultiHash (in, out chan interface{}){
	wg := &sync.WaitGroup{}
	for val := range in {
		wg.Add(1)
		str, _ := toString(val)

		hashes :=make([]chan string, 6)
		for th := range hashes {
			hashes[th] = make(chan string, 1)
			//crc32(th+data)
			go func(data string, out chan<- string) {
				out <- DataSignerCrc32(data)
			}(strconv.Itoa(th) + str, hashes[th])
		}

		go func() {
			defer wg.Done()
			resStr := ""
			for i := range hashes{
				resStr += <-hashes[i]
			}
			out <- resStr
		}()
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}){
	var strTempSlice []string
	for val := range in {
		str, _ := toString(val)
		strTempSlice = append(strTempSlice, str)
	}

	//sorting
	sort.Strings(strTempSlice)

	out <- strings.Join(strTempSlice, "_")
}

func ExecutePipeline(jobs ...job){
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