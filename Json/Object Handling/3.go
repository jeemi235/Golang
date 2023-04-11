package main

import (
  "fmt"
  "reflect"
)

type testObject struct {
  Name   string
  Age    int
  Height float64
}

func main() {
   tstObj := testObject{Name: "yog prakash", Age: 24, Height: 5.6}
   val := reflect.ValueOf(&tstObj).Elem()
   typeOfTstObj := val.Type()
   for i := 0; i < val.NumField(); i++ {
       fieldType := val.Field(i)
       fmt.Printf("object field %d key=%s value=%v type=%s \n",
          i, typeOfTstObj.Field(i).Name, fieldType.Interface(),
          fieldType.Type())
   }
}