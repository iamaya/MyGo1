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

println("*********** Example for make() function ***********")

slice2 := make([]int, 15)
slice2[0] = 1
slice2[1] = 2

fmt.Println(slice2)


println("********* Example map() make() *****")

string2 := make(map[int] string, 2)

string2[0] = "1222"
string2[1] = "2222"
string2[2] = "3222"
string2[3] = "4222"


fmt.Println(string2)

fmt.Println("********** MAP  ***********")

myMap := make(map[int]string)

myMap[1] = "First"
myMap[2] = "Second"

fmt.Println(myMap)

fmt.Println(myMap[100])

fmt.Println("*************** Branching *************")

i := 5
if i == 0 {
    fmt.Println("i is 0")
} else  {
    fmt.Println("i is not 0")
}

if j :=5; j <=6{
    fmt.Println("j is less than 6")
}

println("********** For Loop **********")

for t:=1; t<=5; t++ {
    println(t)
}

k:=10
for{
    k++
    println(k)

if k>15 {
    break
}
}

}

