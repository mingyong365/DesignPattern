package main

import "fmt"

type fuel int

const (
	GASOLINE fuel = iota
	BIO
	ELECTRIC
	JET
)

type vehicle struct {
	make  string
	model string
}

type engine struct {
	fuel   fuel
	thrust int //推力
}

func (e *engine) start() {
	fmt.Println("Engine started.")
}

type truck struct {
	vehicle
	engine
	axels  int //轴
	wheels int
	class  int
}

func (t *truck) drive() {
	fmt.Printf("Truck %s %s, on the go!\n", t.make, t.model)
}

type plane struct {
	vehicle
	engine
	engineCount int
	fixedWings  bool
	maxAttitude int
}

func (p *plane) fly() {
	fmt.Printf("Aircraft %s %s clear for takeoff!\n", p.make, p.model)
}

func newTruck(make, model string) *truck {
	return &truck{vehicle: vehicle{make, model}}
}

func newPlane(make, model string) *plane {
	p := &plane{}
	p.make = make
	p.model = model
	return p
}

//func main() {
//	t := &truck{
//		vehicle: vehicle{"Ford", "F750"},
//		engine:  engine{GASOLINE + BIO, 700},
//		axels:   2,
//		wheels:  6,
//		class:   3,
//	}
//	t.start()
//	t.drive()
//
//	p := &plane{}
//	p.make = "HodaJet"
//	p.model = "HA-420"
//	p.fuel = JET
//	p.thrust = 2050
//	p.engineCount = 2
//	p.fixedWings = true
//	p.maxAttitude = 43000
//	p.start()
//	p.fly()
//}
