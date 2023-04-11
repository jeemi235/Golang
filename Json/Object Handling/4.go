package main
import "fmt"
func main(){
    types := []interface{} {"a",6,6.0,true}
    for _,v := range types{
        fmt.Printf("%T\n",v)
        switch v.(type) {
        case int:
           fmt.Printf("Twice %v is %v\n", v, v.(int) * 2)
        case string:
           fmt.Printf("%q is %v bytes long\n", v, len(v.(string)))
       default:
          fmt.Printf("I don't know about type %T!\n", v)
      }
    }
}