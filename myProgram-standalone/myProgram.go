 package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"
	"math"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("myProgram")
	if status != syscall.StatusOk {
		efmt.Fprintf (syscall.Stderr,"Error opening %v: %v\n", path, status)
	syscall.Exit(syscall.StatusOk)
	}


	boxvals	:= MyType {Field1: 0.0, Field2: 0.0, Field3: 5.0, Field4: 5.0}
	fd, status :=ethos.OpenDirectoryPath(path)
	boxes := &boxvals
	xes := boxes.Field3 - boxes.Field1
	yess := boxes.Field4 - boxes.Field2
	dxes := xes * xes
	dyes := yess * yess
	sumxy := dxes + dyes
	answers := math.Sqrt( sumxy )
	boxvals.Write(fd)
	boxvals.WriteVar(path +"foobar")
	efmt.Fprintf(syscall.Stderr,"This also")
	efmt.Println("Hello Samrudhi!!!")
	efmt.Println(" Here is the diff " , xes)
	efmt.Println(" Square of x is " , dxes)
	efmt.Println(" Square of y is " , dyes)
	efmt.Println(" Distance between the points is " , answers)	
}
