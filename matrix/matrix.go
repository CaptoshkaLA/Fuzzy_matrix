package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"os"
  "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func Generate(randMatrix [][]float64) [][] float64{
    for i, innerArray := range randMatrix {
        for j := range innerArray {
            randMatrix[i][j] = math.Round(rand.Float64()*float64(100) * 100)/100
        }
    }
    return randMatrix
}

func Mprint(randMatrix[][] float64) {
  for _, innerArray := range randMatrix {
    fmt.Println(innerArray)
  }
}

func Stroka(matrix [][] float64, n int, index int) [] float64{
  var str = make([]float64, n) 
  for i := 0; i < n; i++ {
    str[i] = matrix[index][i]
  }
  return str
}

func Stolbez(matrix [][] float64, n int, index int) [] float64{
  var str = make([]float64, n) 
  for i := 0; i < n; i++ {
    str[i] = matrix[i][index]
  }
  return str
}

func Maximum(array [] float64, n int) float64{
  var max = array[0]
  for i := 0; i<n;i++ {
    if(max < array[i]) {
      max = array[i]
    }
  }
  return max
}

func Compozichion(stroka[] float64, stolb[] float64, n int) float64{
  var minim = make([]float64, n)
  for i:=0;i<n;i++ {
    minim[i] = math.Min(stroka[i],stolb[i])
  }
  return Maximum(minim,n)
}

func EnterTheMatrix(Matrix [][]float64) [][] float64{
    for i, innerArray := range Matrix {
        for j := range innerArray {
            // randMatrix[i][j] = math.Round(rand.Float64()*float64(100) * 100)/100
          fmt.Fscan(os.Stdin, &Matrix[i][j])
        }
    }
    return Matrix
}