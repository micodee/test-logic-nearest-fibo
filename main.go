package main

import "fmt"

func main() {
    arr := []int{33, 23, 19}
    nearestFibo(arr)
}

func nearestFibo(arr []int) {
    sum := 0
    for _, value := range arr {
        sum += value
    }
    
    fmt.Println("Hasil:", findFibo(sum))
}

func findFibo(num int) int {
    first, second, third := 0, 1, 1
    
    for third <= num {
        first = second
        second = third
        third = first + second
        // fmt.Printf("%d + %d = %d\n", first, second, third)
								fmt.Println(third)
    }
    
    var ans int
    if abs(third-num) >= abs(second-num) {
        ans = num - second
    } else {
        ans = third - num
    }
    
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}