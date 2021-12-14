package main

import "fmt"

type StrategyPattern struct {
}

func (p StrategyPattern) Show() {
	sendTelegramMessage(TelegramMTProto{})
	sendTelegramMessage(TelegramBotAPI{})
}

func (p StrategyPattern) Name() string {
	return "Strategy"
}

type TelegramProvider interface {
	Send(string)
}

func sendTelegramMessage(tp TelegramProvider) {
	tp.Send("Message to send!")
}

type TelegramMTProto struct {
}

func (t TelegramMTProto) Send(s string) {
	fmt.Println("[MTPTORO]", s)
}

type TelegramBotAPI struct {
}

func (t TelegramBotAPI) Send(s string) {
	fmt.Println("[BOT_API]", s)
}
