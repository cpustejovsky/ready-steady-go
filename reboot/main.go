package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name       string
	Alive      bool
	Productive bool
	Thoughts   []string
}

func NewPerson(name string) *Person {
	return &Person{
		Name:       name,
		Alive:      true,
		Productive: true,
		Thoughts:   []string{},
	}
}

func (p *Person) action(task string) {
	fmt.Println(task)
}

func (p *Person) Work() error {
	if len(p.Thoughts) > 10 {
		p.Productive = false
		return errors.New("too many thoughts; failing to do the work")
	}
	p.action("set forest app")
	p.action("clean desk")
	p.action("pray")
	p.action("do work")
	return nil
}

func (p *Person) Reboot() {
	for _, thought := range p.Thoughts {
		p.action(fmt.Sprintf("processing %s", thought))
	}
	p.Thoughts = []string{}
	p.Productive = true
}

func main() {
	cpustejovsky := NewPerson("Charles")
	for cpustejovsky.Alive {
		err := cpustejovsky.Work()
		if err != nil && !cpustejovsky.Productive {
			cpustejovsky.Reboot()
		}
	}
}
