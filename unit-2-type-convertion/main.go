package main
import (
	"fmt" 
	"strconv"
)

func main(){

	// ========  Basic numeric type conversions =====

	// integers conversion
	var i int = 32
	var i8 int8 = int8(i)
	var i32 int32 = int32(i)
	var i64 int64 = int64(i)
	fmt.Println(i,i8,i32,i64)
	fmt.Printf("Type of i: %T , i8: %T, i32: %T, i64: %T\n",i,i8,i32,i64)

	// converting int to string
	strInt := strconv.Itoa(i)
	fmt.Printf("The value is %s and type is %T\n", strInt,strInt) 

	// converting int to float 
	var f32 float32 = float32(i)
	var f64 float64 = float64(i)
	fmt.Printf("f32 value: %f, f64 vlue: %g\n",f32,f64)
	fmt.Printf("Type f32:%T , f64:%T\n",f32,f64)

	// convering float to str
	floatToString := fmt.Sprintf("%f",f64)
	fmt.Printf("Type floatToString : %T, Value: %s\n",floatToString,floatToString)


	// converting sign int to unsigned integers
	var signed int = -42
	var unsigned uint = uint(signed) // unable to convert negative -42 it will produced a very large number
	fmt.Println("Type unsigned: ", unsigned) //18446744073709551574
	fmt.Printf("Type unsigned:%T\n", unsigned)



	// converting string to integers
	str := "32"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error: %v\n", err) // %v means print a value regardless of the type
	}else {
		fmt.Printf("String type:%T %s converted to type:%T ineger: %d\n",str, str,num,num)
	}

	// coverting string to float
	str2 := "3.14159"
	newf64, err := strconv.ParseFloat(str2,64)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}else{
		fmt.Printf("newf64 value is %f with type %T\n",newf64,newf64)
	}


}