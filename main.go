package main

import (
	"fmt"
	"physics/vector3"
)

func main() {
	v1, err1 := vector3.NewVector3()
	v2, err2 := vector3.NewVector3(1, 4, 5)
	v3, err3 := vector3.NewVector3(3, 4, 5, 2)

	fmt.Println(v1, err1)
	fmt.Println(v2, err2)
	fmt.Println(v3, err3)

	v1.Invert()
	v2.Invert()
	v2.Invert()
}
