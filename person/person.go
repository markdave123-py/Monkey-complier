package person

import "fmt"

type Person interface {
	Shout();
}


type Animall interface {
	Breathe()
}



type Student struct {

}

type Teacher struct {

}


func (s Student) Shout() {
	fmt.Println("Student is shouting")
}

func (s Student) Breathe() {
	fmt.Println("Student is breathing")
}

func (t Teacher) Breathe() {
	fmt.Println("Teacher breathing.")
}







// func New(name string, age int) Person {
// 	return Person{name, age}
// }

// func (p Person) GetName() string {
// 	return p.name
// }

// func (p *Person) SetName(name string) {
// 	p.name = name
// }