package pckg

import "testing"

//Structure for expected results and default values
type testpair struct {
  values []float64
  expected float64
}

//Test function for Average()
func TestAverage(t *testing.T) {
  //Initialize structure
  var averageTestValues = []testpair{
    { []float64{1,2}, 1.5 },
    { []float64{1,1,1,1,1,1}, 1 },
    { []float64{-1,1}, 0 },
  }
  //Check all items for passing
  for _, pair := range averageTestValues {
    result := Average(pair.values)
    if result != pair.expected {
      //Show error if smth isn't OK
      t.Error(
        "\nFunction TestAverage:\n",
        "For:", pair.values,
        "expected:", pair.expected,
        "got:", result,
      )
    }
  }
}

//Test function for Min()
func TestMin(t *testing.T) {
  //Initialize structure
  var minTestValues = []testpair{
    { []float64{1,2}, 1 },
    { []float64{1,5,1,4,-1,0}, -1 },
    { []float64{-100,10}, -100 },
  }
  //Check all items for passing
  for _, pair := range minTestValues {
    result := Min(pair.values)
    if result != pair.expected {
      //Show error if smth isn't OK
      t.Error(
         "\nFunction TestMin:\n",
         "For:", pair.values,
         "expected:", pair.expected,
         "got:", result,
      )
    }
  }
}

//Test function for Max()
func TestMax(t *testing.T) {
  //Initialize structure
  var maxTestValues = []testpair{
    { []float64{1,2}, 2 },
    { []float64{1,5,1,4,-1,0}, 5 },
    { []float64{-100,10}, 10 },
  }
  //Check all items for passing
  for _, pair := range maxTestValues {
    result := Max(pair.values)
    if result != pair.expected {
      //Show error if smth isn't OK
      t.Error(
         "\nFunction TestMax:\n",
         "For:", pair.values,
         "expected:", pair.expected,
         "got:", result,
      )
    }
  }
}
