package main

import "fmt"

func main() {

  //How to declare few variables
  var (
    x= "1"
    z = "333"
  )
  y:= x + "2a"

  //Print it on screen
  fmt.Println(len(y + "qqqq" + z))
  fmt.Println("Enter some number:")

  //Example of Scanf
  var input float64
  fmt.Scanf("%f", &input)
  fmt.Println("Square of input value: ", input * input)

  //FizzBuzz usage
  for i := 0; i <= 50; i++ {
    fizzBuzz(i)
  }
}

//FizzBuzz function
func fizzBuzz(digit int) {
  if digit % 3 == 0 && digit % 5 == 0 {
    fmt.Println(digit, "FizzBuzz")
  } else if digit % 5 == 0 {
    fmt.Println(digit, "Buzz")
  } else if digit % 3 == 0  {
    fmt.Println(digit, "Fizz")
  }
}
