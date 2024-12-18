// hello.go
package myhelloworld

import "fmt"

// Hello returns a greeting message.
func Hello() string {
    return "Hello, World!"
}

// PrintHello prints the greeting message directly to the console.
func PrintHello() {
    fmt.Println(Hello())
}
