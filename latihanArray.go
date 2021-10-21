package main

//reflect : ngecek tipe data dari variabel
import (
	"fmt"
	"reflect"
)

func main() {
	var primes [5]int
	var countries [5]string

	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(reflect.ValueOf(countries).Kind())
	//
	x := [5]int{1, 3, 5, 7, 9}
	var y [5]int = [5]int{0, 2, 4}

	fmt.Println(x)
	fmt.Println(y)
	//
	prime := [5]int{2, 3, 5}
	for index := 0; index < len(prime); index++ {
		fmt.Println(prime[index])
	}
	//
	for index, element := range prime {
		fmt.Println(index, "=>", element)
	}
	//
	index := 0
	for range prime {
		fmt.Println(prime[index])
		index++
	}
}
