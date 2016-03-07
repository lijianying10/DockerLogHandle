package watch

import "fmt"

func FileModuleLoad(watchingContainer, outputfile string, c chan bool) {
	fmt.Println("FileMod loded")

	c <- true

}
