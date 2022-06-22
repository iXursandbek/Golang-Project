package main
 
import "fmt"
 
func main () {
    var a, b, max, max1 int
    fmt.Scan(&a, &b)
    if a <= b {
        for i := a; i <= b; i++ {
            if i%7 == 0 {
                max = i
            }
        }
        if max == 0 {
            fmt.Println("NO")
        } else {
            fmt.Println(max)
        }
    } else if a >= b {
        for i := a; i >= b; i-- {
            if i%7 == 0 {
                max1 = i
                if max1 == 0 {
                    break
                }
            }
        }
        if max1 == 0 {
            fmt.Println(max1)
 
        } else {
            fmt.Println("NO")
        }
 
    }
}



  
// package main

// import "fmt"

// func main() {
// 	var a, b, res, count int
// 	fmt.Scan(&a, &b)

// 	for i := b; i >= a; i-- {
// 		if i%7 == 0 {
// 			res = i
// 			fmt.Println(res)
// 			count++
// 			break
// 		} else {
// 			continue
// 		}
// 	}
// 	if count == 0 {
// 		fmt.Println("NO")
// 	}
// }