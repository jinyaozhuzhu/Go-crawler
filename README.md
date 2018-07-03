# Go-crawler
the crawl of using golang 

1. model/profile.go 征婚人数据模型
2. fetcher/fetcher.go  用于网页数据的抓取，并解决网页编码问题
3. persist 使用goroutine并发保存数据到elasticsearch
4. zhenai/parse 用于具体页面数据的解析
5. engine/types.go 页面元素的表示，包括页面的url，解析此页面的函数
   解析结果，包括此页面解析出数据和页面解析出的url，Item结构用于存储于elasticsearch
6. engine/worker.go 调用页面获取模块和数据解析模块
7. engine/simple.go 简单版爬虫引擎  engine/concurrent.go 并发版爬虫引擎
8. main.go 整个爬虫的入口
