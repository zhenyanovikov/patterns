package main

import "fmt"

type AdapterPattern struct {
}

func (p AdapterPattern) Show() {
	InsertTypeC(TypeC{})
	fmt.Println()
	fmt.Println("Let's try lightning...")
	fmt.Println()
	InsertTypeC(Lighting2TypeCAdapter{Lightning{}})
}

func (p AdapterPattern) Name() string {
	return "Adapter"
}

type TypeCInterface interface {
	TypeCData() string
}

func InsertTypeC(c TypeCInterface) {
	fmt.Println(c.TypeCData())
}

type Lightning struct {
}

func (l Lightning) Data() string {
	return "Some lighting data"
}

type TypeC struct {
	data string
}

func (t TypeC) TypeCData() string {
	return "Some Type-C data"
}

type Lighting2TypeCAdapter struct {
	Lightning Lightning
}

func (l Lighting2TypeCAdapter) TypeCData() string {
	fmt.Println("Lightning connected via Type-C!")
	return l.Lightning.Data()
}
