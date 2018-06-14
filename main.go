package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"crawler/persist"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(
		engine.Request{
			Url:       "http://www.zhenai.com/zhenghun",
			ParseFunc: parser.ParseCityList,
		})
}
