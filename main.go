package main

import "fmt"

type PatternDemo interface {
	Show()
	Name() string
}

func main() {
	patterns := []PatternDemo{
		// creational patterns
		SimpleFabricPattern{},
		BuilderPattern{},
		// structural pattern
		AdapterPattern{},
		ProxyPattern{},
		// behavioral patterns
		ObserverPattern{},
		StrategyPattern{},
	}

	for _, pattern := range patterns {
		fmt.Printf("--- Demo for %s pattern:\n", pattern.Name())
		pattern.Show()
		fmt.Print("---\n\n")
	}
}
