package main

/**
 * Author:Lichen
 * Description: interface实战
 * Date: 2022-07-18 22:20
 */
type action interface {
	getName() string
}

type peopleA struct {
	Name string
	Age  int
}

func (a *peopleA) getName() string {
	return a.Name
}

type peopleB struct {
	Name string
	Sex  int
}

func (b peopleB) getName() string {
	return b.Name
}

func main() {
	actions := []action{}
	a := new(peopleA)
	a.Name = "a"
	a.Age = 20
	actions = append(actions, a)
	b := peopleB{
		Name: "b",
		Sex:  1,
	}
	actions = append(actions, b) //无论是指针还是值只要实现了方法就可以算的上实现了interface

}
