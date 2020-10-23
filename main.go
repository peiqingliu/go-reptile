package main

import (
	"fmt"
	"go-reptile/engine"
	"go-reptile/persist"
	"go-reptile/scheduler"
	"go-reptile/targetset/parser"
)

func main()  {
	fmt.Println("你好,reptile")

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaverSql(),
	}

	e.Run(engine.Request{
		Url: "https://www.danke.com/room/sz/d.html",
		ParserFunc: parser.ParseAreaList,
	})
}
