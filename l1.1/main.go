package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Printf("Hi, my name is %s, Im %d years.\n", h.Name, h.Age)
}

type Action struct {
	Human
	Skill string
}

func (a Action) DoSomething() {
	fmt.Printf("%s is doing by its profession: %s\n", a.Name, a.Skill)
}

func main() {
	person := Action{
		Human: Human{
			Name: "Nikita",
			Age:  22,
		},
		Skill: "Programming",
	}

	person.Speak()
	person.DoSomething()
}
