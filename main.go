package main
import "fmt"


const  (
first = "some"
second ="second value"

)

func main() {

//println(second)

	/*println("Hello...Go")

    var myint int
    myint= 50

    println(myint)

    myString := "this is a string"

    println(myString); println(myint)

    myComplex := complex(5,10)

    println(myComplex) ; println(real(myComplex))*/

    myArray := [...] int {5,10,15,20}

    mySlice := myArray[:1]

fmt.Println(myArray)
fmt.Println(mySlice)

}
