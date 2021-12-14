package main

import (
	"fmt"
)

type ObserverPattern struct {
}

func (p ObserverPattern) Show() {
	notifier := &Notifier{clients: make([]*Client, 0, 10)}
	notifier.Start()
}

func (p ObserverPattern) Name() string {
	return "Observer"
}

type Notifier struct {
	clients []*Client
}

func (n *Notifier) Start() {
	updates := []string{
		"Update1",
		"Update2",
		"Update3",
		"Update4",
	}

	for i, update := range updates {
		fmt.Println("executing update", i)
		for _, client := range n.clients {
			client.Handle(update)
		}

		n.Register(&Client{id: i + 1})

		fmt.Println()
	}
}

func (n *Notifier) Register(c *Client) {
	n.clients = append(n.clients, c)
}

type Client struct {
	id int
}

func (c Client) Handle(update string) {
	fmt.Printf("Handling update '%s' on client %d\n", update, c.id)
}
