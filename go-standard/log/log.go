package log

import "fmt"

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	Lmsgprefix
	LstdFlags = Ldate | Ltime
)

func PrintLogConst() {
	fmt.Println("Ldate = ", Ldate)
	fmt.Println("Ltime = ", Ltime)
	fmt.Println("Lmicroseconds = ", Lmicroseconds)
	fmt.Println("Llongfile = ", Llongfile)
	fmt.Println("Lshortfile = ", Lshortfile)
	fmt.Println("LUTC = ", LUTC)
	fmt.Println("Lmsgprefix = ", Lmsgprefix)
	fmt.Println("LstdFlags = ", LstdFlags)
}

type Logger struct {

}