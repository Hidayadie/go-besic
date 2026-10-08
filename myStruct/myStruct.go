package mystruct

import "fmt"

type TheStruct struct {
	id int
	desc string
	isDone bool 
}



func MyStruct() {

	MyJobs := TheStruct{
		id : 1,
		desc : "Breathe...",
		isDone: false,
	}

	fmt.Println(MyJobs)
}

