package main

import "fmt"

type P interface {
	info()
}

type Pa struct {
}

func (p Pa) info() {
	fmt.Println("我是产品A")
}

type Pb struct {
}

func (p Pb) info() {
	fmt.Println("我是产品B")
}

type Fa struct {
}

func (f Fa) CreateP() P {
	return Pa{}
}

type Fb struct {
}

func (f Fb) CreateP() P {
	return Pb{}
}

func main() {
	var f1 Fa
	p1 := f1.CreateP()
	p1.info()

	var f2 Fb
	p2 := f2.CreateP()
	p2.info()
}
