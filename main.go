package main

import (
	"fmt"

	myotherpackage "github.com/michal-polaczyk/myOtherPackage"
)

func main() {
	fmt.Println("something")
	doSometing("something else")
	myotherpackage.DoSomethingOther()
}
