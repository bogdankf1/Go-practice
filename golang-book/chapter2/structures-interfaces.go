package main

import ("fmt"; "math")

//Struct for circle
type Circle struct {
  x, y, r float64
}
//Метод для вычисления площади круга - особый тип функции
func (c *Circle) calculateArea() float64 {
  return math.Pi * math.Pow(c.r, 2)
}
func (c *Circle) calculatePerimeter() float64 {
  return math.Pi * 2 * c.r
}

//Function for calculating circle area
//Принимаем аргумент с указателем (*)!!
//Аргументы без звездочки всегда копируются
// func calculateCircleArea(c *Circle) float64 {
//   return math.Pi * math.Pow(c.r, 2)
// }

//Struct for rectangle
type Rectangle struct {
  x1, y1, x2, y2 float64
}
//Function for calculating vector distance
func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}
//Method for calculating rectangle area
func (r *Rectangle) calculateArea() float64 {
  length := distance(r.x1, r.y1, r.x1, r.y2)
  width := distance(r.x1, r.y1, r.x2, r.y1)
  return length * width
}
func (r *Rectangle) calculatePerimeter() float64 {
  length := distance(r.x1, r.y1, r.x1, r.y2)
  width := distance(r.x1, r.y1, r.x2, r.y1)
  return 2*(length + width)
}

//Демонстрация встраиваемых типов
type Person struct {
  name string
}
func (p *Person) talk() {
  fmt.Println("My name is", p.name)
}
type Android struct {
  Person
  model string
}

//INTERFACES
type Shape interface {
  calculateArea() float64
  calculatePerimeter() float64
}
func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.calculateArea()
  }
  return area
}

type MultiShape struct {
  shapes []Shape
}
//Calculate total area of all shapes
func (multiShape *MultiShape) area() float64 {
  var area float64
    for _, s := range multiShape.shapes {
        area += s.calculateArea()
    }
    return area
}
//Calculate total perimter of all shapes
func (multiShape *MultiShape) perimeter() float64 {
  var perimeter float64
    for _, s := range multiShape.shapes {
        perimeter += s.calculatePerimeter()
    }
    return perimeter
}

func main()  {
  //***************STRUCTURES (USAGE) *********************
  // var c Circle
  //Или другой вариант:
  // c := new(Circle)

  //Явное присваивание значений:
  c := Circle{x: 0, y: 0, r: 5}
  //Если порядок полей известен:
  // c := Circle{0, 0, 5}
  fmt.Println(c.x, c.y, c.r)// 0 0 5
  //Передаем с & чтобы изменить экземпляр типа Circle в main()
  // fmt.Println("Circle area is:", calculateCircleArea(&c))

  //При использовании методов вместо функций не нужно использовать амперсанд
  //Go автоматически предоставит доступ к указателю на объект
  fmt.Println("Circle area is:", c.calculateArea())

  r := Rectangle{0, 0, 10, 10}
  fmt.Println("Rectnagle area is", r.calculateArea())

  //Встраиваемые типы:
  a := Person{"Bogdan"}
  android := Android{a, "model3"}
  a.talk()
  android.talk()
  //*******************END************************

  //******************INTERFACES**********************
  fmt.Println("Total area of all shapes is", totalArea(&r, &c))
  var m MultiShape
  m.shapes = []Shape{&c, &r}
  fmt.Println("Total area using multishape interface:",m.area())
  fmt.Println("Total perimeter using multishape interface", m.perimeter())
}
