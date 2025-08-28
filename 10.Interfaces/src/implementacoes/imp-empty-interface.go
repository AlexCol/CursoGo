package implementacoes

import (
	"fmt"
	"reflect"
	"time"
)

func EmptyInterface() {
	var myDynamic interface{} //! pode ser usado 'any' no lugar de interface{}

	myDynamic = 5 //inteiro
	fmt.Println(myDynamic)

	myDynamic = "Um texto" //texto
	fmt.Println(myDynamic)

	myDynamic = []int{1} //slice
	fmt.Println(myDynamic)

	myDinamicSlice := []interface{}{"Campo1", 2, true, []int{1, 2}}
	fmt.Println(myDinamicSlice)
	myDinamicSlice = append(myDinamicSlice, time.Now()) //.Format("02/01/2006 15:04:05")
	//fmt.Println(myDinamicSlice)
	fmt.Println("------------------------------------imprimindo lista dynamica-----------------------------------------")
	imprimeLista(myDinamicSlice)
}

func imprimeLista(lista []interface{}) {
	for _, item := range lista {
		fmt.Printf("tipo=%s, valor=%v\n", reflect.TypeOf(item).Kind(), item)
	}
}
