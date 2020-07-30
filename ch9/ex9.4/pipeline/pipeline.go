package pipeline

func Create(n int) (in, out chan int) {
	out = make(chan int)
	if n < 1 {
		in = out
		return
	}
	_, in = Create(n-1)
	go func() {
		out<- <-in
	}()
	return 
}
