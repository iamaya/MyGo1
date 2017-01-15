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

    mySlice := myArray[:]

    mySlice[1] = 55

fmt.Println(mySlice)
    mySlice = append(mySlice, 100)


fmt.Println(myArray)
fmt.Println(mySlice)


newSlice := []int{3,4,5,6}

fmt.Println(newSlice)
fmt.Println(len(newSlice))


fmt.Println("********** MAP  ***********")

myMap := make(map[int]string)

myMap[1] = "First"
myMap[2] = "Second"

fmt.Println(myMap)

}
