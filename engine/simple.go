package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		//将解析的结果url放回队列
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)

	//获取网页源码
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch: error fetching url %s :  %v", r.Url, err)
		return ParseResult{}, err
	}
	//利用对应得到处理函数进行解析
	return r.ParseFunc(body), nil

}
