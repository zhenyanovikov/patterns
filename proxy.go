package main

import "fmt"

type ProxyPattern struct {
}

func (p ProxyPattern) Show() {
	worker := Worker{}
	clearWork := worker.Work("aaa")
	fmt.Println("Clear work:\n" + clearWork)

	fmt.Println()

	detailedWorker := DetailedWorker{Worker: worker}
	proxiedWork := detailedWorker.Work("aaa")
	fmt.Println("Work with proxy:\n" + proxiedWork)
}

func (p ProxyPattern) Name() string {
	return "Proxy"
}

type Worker struct {
}

func (w Worker) Work(s string) string {
	return s + ":" + s + ":" + s
}

type DetailedWorker struct {
	Worker Worker
}

func (w DetailedWorker) Work(s string) string {
	result := "[x] Something before work\n"
	result += w.Worker.Work(s)
	result += "\n[x] Something after work"
	return result
}
