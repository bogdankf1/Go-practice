package main

import "fmt"
import myPackage"golang-book/chapter2/packages"

func main() {
    xs := []float64{1,2,3,4}
    avg := myPackage.Average(xs)
    fmt.Println(avg)
}
