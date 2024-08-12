package dog

import "fmt"

type Dog string

type Barker interface {
	Bark()
}

func (d Dog) Bark() {
	fmt.Println("Woof!")
}

func Speak(barker Barker) {
	barker.Bark()
}

func SpeakAll(barkers []Barker) {
	for _, barker := range barkers {
		barker.Bark()
	}
}

func ReturnDog() Dog {
	return "Rex"
}

func UseCallback(callback func() Barker) {
	barker := callback()
	barker.Bark()
}

func AcceptBarker(barker Barker) {
	barker.Bark()
}

func UseCallBack(callback func(Dog)) {
	callback("Laika")
}
