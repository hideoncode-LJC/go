package main

type singal struct{}

func doTask(i int, s singal) {
	for {
		select {
		case <- s :
			fmt.Println("Task ", i, "Done !")
		default:
			fmt.Println("Task ", i, "Go on!")
		}
	}
}

func main() {
	stop := singal{}
	for i := 0 ; i < 3 ; i++ {
		go doTask(i, stop)
	}
	time.Sleep(time.)
	for {

	}
}