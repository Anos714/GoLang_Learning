package main
import "fmt"

func main() {
	todos:=[]string{"learn Go","learn Hono","learn communication","be productive","be consistent","be collaborative","be confident"}

	// for range loop
	for i,val:=range todos{
		fmt.Println("\nindex: ",i,"\nvalue: ",val)
	}

	views:=[]int{10,200,450,16000,2200};

	sum:=0
	for _,val:=range views{
		sum+=val
	}
	fmt.Println("Total Combined Views Sum: ",sum)
}
