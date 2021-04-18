package main
import "fmt"

func main() {
	for i := 0;i < 100; i++ {
		if i % 7 == 0 || i / 10 == 7 || i % 10 == 7 {
			fmt.Println("敲桌子")
		} else {
			fmt.Println(i)
		}
	}
}


