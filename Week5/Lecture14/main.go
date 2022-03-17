package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

func main() {
	var numberOfConcurency int
	flag.IntVar(&numberOfConcurency, "c", 2, "buffer size")

	flag.Parse()

	filePaths := flag.Args()

	if len(filePaths) < 1 {
		flag.PrintDefaults()
		return
	}

	resultChan := fetchURLs(filePaths, numberOfConcurency)
	for url := range resultChan {
		pingURL(url)
	}

}

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)

	return nil
}

func fetchURLs(urls []string, conurrency int) chan string {
	processQueue := make(chan string, conurrency)
	var wg sync.WaitGroup

	go func() {
		for _, urlToProcess := range urls {
			wg.Add(1)

			go func(url string) {
				defer wg.Done()
				processQueue <- url
			}(urlToProcess)
		}
		wg.Wait()
		close(processQueue)

	}()
	return processQueue
}

// Output:
// ./multiping -c 2 http://www.yahoo.com http://www.google.com http://www.facebook.com http://www.bing.com http://www.twitter.com http://www.yahoo.com
// 2022/03/17 16:01:06 Got response for http://www.yahoo.com: 200
// 2022/03/17 16:01:06 Got response for http://www.google.com: 200
// 2022/03/17 16:01:07 Got response for http://www.facebook.com: 200
// 2022/03/17 16:01:07 Got response for http://www.bing.com: 200
// 2022/03/17 16:01:08 Got response for http://www.twitter.com: 200
// 2022/03/17 16:01:08 Got response for http://www.yahoo.com: 200

// Output:
// ./multiping -c 5 http://www.yahoo.com http://www.google.com http://www.facebook.com http://www.bing.com http://www.twitter.com http://www.yahoo.com
// 2022/03/17 16:01:40 Got response for http://www.yahoo.com: 200
// 2022/03/17 16:01:41 Got response for http://www.google.com: 200
// 2022/03/17 16:01:41 Got response for http://www.facebook.com: 200
// 2022/03/17 16:01:41 Got response for http://www.bing.com: 200
// 2022/03/17 16:01:42 Got response for http://www.twitter.com: 200
