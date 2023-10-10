// for.go : A simple for loop program written in golang 
package main

// get fmt package for Printf()
import "fmt"

// Our main()
func main() {
        // set a for loop
        for i := 1;  i<=5; i++ {
                fmt.Printf("Welcome %d times\n",i)
        }
}
