package main

import "fmt"

func SumNeighbors(arr [10]int) (newArr [10]int) {
    for i := range arr {
        switch {
        case i == 0:
            newArr[0] = arr[1]
        case i == 9:
            newArr[9] = arr[8]
        default:
            newArr[i] = arr[i-1] + arr[i+1]
        }
    }
    return
}

func main() {
    // Пример массива
    arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println("Исходный массив:")
    fmt.Println(arr)
    
    result := SumNeighbors(arr)
    
    fmt.Println("Результат (сумма соседей):")
    fmt.Println(result)
    
    // Проверим каждый элемент
    fmt.Println("\nПодробно:")
    for i, val := range result {
        switch {
        case i == 0:
            fmt.Printf("newArr[%d] = arr[1] = %d\n", i, val)
        case i == 9:
            fmt.Printf("newArr[%d] = arr[8] = %d\n", i, val)
        default:
            fmt.Printf("newArr[%d] = arr[%d] + arr[%d] = %d + %d = %d\n", 
                i, i-1, i+1, arr[i-1], arr[i+1], val)
        }
    }
}