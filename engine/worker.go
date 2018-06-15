package engine

import (
	"log"
	"crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s ", r.Url)

	//获取网页源码
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch: error fetching url %s :  %v", r.Url, err)
		return ParseResult{}, err
	}
	//利用对应得到处理函数进行解析
	return r.ParseFunc(body), nil
}
