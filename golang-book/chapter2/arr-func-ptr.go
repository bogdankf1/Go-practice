package main

import "fmt"

//Function for calculationg average value in array
func average(arr []float64) float64{
  var avg float64 = 0
  for _, value := range arr {
    avg += value
  }
  return avg / float64(len(arr))
}

//Function for returning some values
func returnSomeValues() (int, int){
  return 7283749823, 2312312
}

//Function with variative number of arguments
func add(args ...int) int {
  total := 0
  for _, v := range args {
      total += v
  }
  return total
}

func makeEvenGenerator() func() int {
  i := 0
  return func () (result int) {
    result = i
    i += 2
    return
  }
}

//Example of classic recursion (factorial)
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }

    return x * factorial(x-1)
}

//Functions for demonstration of defer call
func first() {
    fmt.Println("1st")
}
func second() {
    fmt.Println("2nd")
}

func half(digit int) (int, bool) {
  if (digit/2)%2 == 0 {
    return 1, true
  } else {
    return 0, false
  }
}

func findMax(args ...int) int {
  max := args[0]
  for _, value := range args {
    if value > max {
      max = value
    }
  }
  return max
}

//Good example how ptr works
func zero(xPtr *int) {
    *xPtr = 0
}

func swap(x *int, y *int) {
  *x, *y = *y, *x
}

func main() {
  //****************FUNCTIONS CALLS EXAMPLES*******************
  //Array example
  arr := []float64{22,33,44,55,66}
  fmt.Println("Average value is", average(arr))
  funcValue1, funcValue2 := returnSomeValues()
  fmt.Println(funcValue1, funcValue2)
  fmt.Println("Sum:",add(1,2,3,4,5,6,7,8,9))

  //Closures
  add := func(x,y int) int {
    return x + y
  }
  fmt.Println("Result of add():",add(2,2))

  nextEven := makeEvenGenerator()
  for i := 0; i < 5; i++ {
    fmt.Println("Result of nextEven():",nextEven())
  }

  fmt.Println("Result of findMax():",findMax(1,5,6,-4,56,23,434,-34,0,-100,1000, 3232))

  defer second()
  first()

  //Call function half()
  half1, half2 := half(11)
  fmt.Println("Result of half():",half1,half2)
  half1, half2 = half(20)
  fmt.Println("Result of half():",half1,half2)

  //panic and recover
  defer func() {
        str := recover()
        fmt.Println(str)
    }()
  // panic("PANIC")

  //************END OF FUNCTIONS BLOCK*****************

  //************ARRAYS*********************
  //Slices(срезы)
  // срез, связан с массивом такого типа, срез длиной 5, указывает на массив длиной 10
  // slice1 := make([]float64, 5, 10)

  //Brief form of array declaration
  x := [5] float64{ 98, 93, 77, 82, 83 }
  fmt.Println(x)

  slice2 := x[0:3]//98 93 77
  fmt.Println(slice2)

  //arr[low:high]
  //arr[0:] == arr[:len(arr)] == arr[:]

  //Two functions with slices:
  //APPEND
  slice3 := append(slice2, 100, 1000)
  fmt.Println(slice3) //98 93 77 100 1000

  //COPY
  slice4 := make([]float64, 4)
  copy(slice4, slice3)
  fmt.Println(slice4)//98 93 77 100

  //Maps(карты)
  xMap := make(map[string]int)
  xMap["first"] = 1
  fmt.Println(xMap["first"])
  if name, ok := xMap["first"]; ok {
    fmt.Println(name, ok)//1, true
  }

  someArr := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }
  var min = someArr[0]
  for _, value := range someArr {
    if value < min {
      min = value
    }
  }
  fmt.Println("Min:", min)
  //*************END OF ARRAYS BLOCK****************

  //**********************POINTERS*****************
  xPtr := 5
  zero(&xPtr)
  fmt.Println("Xptr = ", xPtr) // x is 0
  //& - для получения адреса переменной
  //* - для получения значения указателя

  xValue, yValue := 10, 20
  fmt.Println("X and Y before swap:", xValue, yValue)
  swap(&xValue, &yValue)
fmt.Println("X and Y after swap:", xValue, yValue)
}
