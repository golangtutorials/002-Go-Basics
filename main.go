package main

import "fmt"

const con string = "Paul's tutorials ROCK!"

func main() {


	values()
	variables()
	constants()
	ifthen()
	cases()
	forloop()
	arrays()

}

func values() {
	fmt.Println("Hello" + " " + "Paul\n")
	fmt.Println("My Age: ", 16*10, "\n")
	fmt.Println("My float: ", 10/5)
	fmt.Println("Am I a male? ", (!true))


}

func variables() {
	var s1, s2 string = "Hi", "Paul"

	fmt.Println(s1, s2)

	var i int = 1
	fmt.Println(i)

	s3 := "How are you?"

	fmt.Println(s3)

	i2 := 12345

	i2 = i2 *3

	fmt.Println(i2)
}

func constants() {
	fmt.Println(con)

	const con2 int = 50

	fmt.Println(con2)
}

func ifthen() {
	i := 30
	x := 300

	if i == x {
		fmt.Println("The numbers are perfect")
	} else if i < x {
		fmt.Println(i, "is less than", x)
	} else if i > x {
		fmt.Println(x, "is less than", i)
	} else {
		fmt.Println("Not so perect, try again")
	}


	if(x%i == 0){
		fmt.Println(x, "is divisiable by", i)
	}
}

func cases() {
	i := 1

	switch i {
	case 1:
		fmt.Println("This is 1")
	case 2:
		fmt.Println("This is 2")
	default:
		fmt.Println("none of the above")
	}


	whatTypeIsUsed :=  func(iface interface{}) {
		switch x := iface.(type) {
		case int:
			fmt.Println("This is an int")
		case bool:
			fmt.Println("this is a bool")
		default:
			fmt.Printf("Please tell me what %X is", x)
		}
	}

	whatTypeIsUsed("Hey Paul")
}

func forloop() {
	for i := 0; i <=10; i++ {
		fmt.Println(i)
	}

	for i := 0; i <=10000; i++ {
		if i == 13 {
			break
		}
		fmt.Println(i)
	}
}

func arrays() {
	var arr [10]string

	arr[1] = "Paul"
	arr[2] = "Wilson"
	arr[3] = "June"
	arr[4] = "16"
	arr[5] = "Howdy"

	st := [2]string{"hi", "there"}

	fmt.Println(st)

	fmt.Println(arr)

	i := [5]int{1, 2, 3, 4, 5}

	fmt.Println(i)

}
