package main
import "fmt"

func shiftMatrix(numArray [][]int, n int) {

    rows := len(numArray)
    cols := len(numArray[0])
    
    for i := 0; i < n; i++ {
        temp := make([]int, n)
        for j := cols - n; j < cols; j++ {
            temp[j-cols+n] = numArray[i][j]
        }
        for j := cols - 1; j >= n; j-- {
            numArray[i][j] = numArray[i][j-n]
        }
        for j := 0; j < n; j++ {
            numArray[i][j] = temp[j]
        }
        
    }
}
 

func main() {
    var rows,cols int
    rows = 5 //n
    cols = 4 //m
    
    var numArray[5][4]
    
    // Заполняем массив значениями, введенными пользователем
    fmt.Println("Введите элементы матрицы:")
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Scan(&numArray[i][j])
        }
    }
    
    // Выводим массив на экран
    fmt.Println("Матрица:")
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Print(numArray[i][j], "\t")
        }
        fmt.Println()
    }
    
    var k,n int
    fmt.Print("Введите количество шагов: ")
    fmt.Scanln(&n)
    //k - номер слоя, n - количество шагов
    
    shiftMatrix(numArray[rows][cols],n)
    
    // Выводим массив на экран
    fmt.Println("Матрица:")
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Print(numArray[i][j], "\t")
        }
        fmt.Println()
    }
    
}
