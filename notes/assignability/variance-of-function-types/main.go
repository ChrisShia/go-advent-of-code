package main

func main() {
	//Covariance (not allowed)
	//Note: func() Dog not assignable to func() Barker
	//dog.UseCallback(dog.ReturnDog) WRONG!

	//function types on the other hand are contravariant in their parameter types
	//dog.UseCallBack(dog.AcceptBarker) WRONG!
}
