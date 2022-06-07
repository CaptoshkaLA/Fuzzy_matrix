package main

import (
	"fmt"
	"main/matrix"
	"os"
)

func main() {
    fmt.Print("Размерность: ")
    var n int
    fmt.Fscan(os.Stdin, &n)

    fmt.Println("Как будет заполнена матрица? 0 - рандомные значения в диапазоне [0;100] / 1 - заполнение вручную")
    fl := 1

    matrixA := make([][]float64, n) 
    for i := 0; i < n; i++ {
        matrixA[i] = make([]float64, n) 
    }

    matrixB := make([][]float64, n) 
    for i := 0; i < n; i++ {
        matrixB[i] = make([]float64, n) 
    }
    

    if(fl == 0) {
       matrixA = matrix.Generate(matrixA)
       matrixB = matrix.Generate(matrixB)
    } else {
       fmt.Println("Введите значения матрицы A:")
       matrixA = matrix.EnterTheMatrix(matrixA)
       fmt.Println("Введите значения матрицы B:")
       matrixB = matrix.EnterTheMatrix(matrixB)
    }
  
    fmt.Println("Matrix A")
    matrix.Mprint(matrixA)
    fmt.Println("Matrix B")
    matrix.Mprint(matrixB)

    matrixCompoz := make([][]float64, n) 
    for i := 0; i < n; i++ {
        matrixCompoz[i] = make([]float64, n) 
    }
  
    for i := 0; i < n; i++ {
      var str = make([]float64, n)
      str = matrix.Stroka(matrixA, n, i)
      for j := 0; j < n; j++ {
        var stb = make([]float64, n)
        stb = matrix.Stolbez(matrixB, n,j)
        matrixCompoz[i][j] = matrix.Compozichion(str,stb,n)
      }
    }

    fmt.Println("AxB(composition)")

    for i := 0; i < n; i++ {
      for j := 0; j < n; j++ {
        fmt.Print(matrixCompoz[i][j], " ")
      }
      fmt.Println(" ")
    }
    fmt.Println(" ")

    var alpha float64
    fmt.Println("Введите значение alpha")
    fmt.Fscan(os.Stdin, &alpha)
    fmt.Println("Альфа срез композиции матрицы А = {")
    fmt.Print("{")
    for i := 0; i < n; i++ {
      for j := 0; j < n; j++ {
        if matrixA[i][j] >= alpha {
          fmt.Print("(", i+1, ",", j+1, ")")
        }
      }
      fmt.Println(" ")
    }
    fmt.Println("}")

    fmt.Println("Альфа срез композиции матрицы B = {")
    fmt.Print("{")
    for i := 0; i < n; i++ {
      for j := 0; j < n; j++ {
        if matrixB[i][j] >= alpha {
          fmt.Print("(", i+1, ",", j+1, ")")
        }
      }
      fmt.Println(" ")
    }
    fmt.Println("}")
    
    fmt.Println("Альфа срез композиции матриц A & B = {")
    fmt.Print("{")
    for i := 0; i < n; i++ {
      for j := 0; j < n; j++ {
        if matrixCompoz[i][j] >= alpha {
          fmt.Print("(", i+1, ",", j+1, ")")
        }
      }
      fmt.Println(" ")
    }
   fmt.Println("}")
  
}



