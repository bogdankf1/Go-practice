package pckg

//Find average value in float numbers array
func Average(digits []float64) float64 {
    total := float64(0)
    for _, number := range digits {
        total += number
    }
    return total / float64(len(digits))
}

//Find min value in float numbers array
func Min(digits []float64) float64 {
  min := digits[0]
  for _, number := range digits {
    if number < min {
      min = number
    }
  }
  return min
}

//Find max value in float numbers array
func Max(digits []float64) float64 {
  max := digits[0]
  for _, number := range digits {
    if number > max {
      max = number
    }
  }
  return max
}
