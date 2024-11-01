package channel

import (
	"fmt"
	"runtime"
	"sync"
)

// printFunctionName prints the function name of the caller
func printFunctionName() {
	counter, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(counter).Name()
	fmt.Println("[", funcName, "]")
}

func createArrayByGoRoutine(n int) []int {
	printFunctionName()

	arr := []int{}
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			arr = append(arr, i)
		}(i)
	}

	wg.Wait()
	return arr
}

func createArrayWithMutex(n int) []int {
	printFunctionName()

	arr := []int{}
	wg := sync.WaitGroup{}
	wg.Add(n)

	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			arr = append(arr, i)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	return arr
}

func calcSum(s []int, c chan<- int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func sumByChannel() (int, int) {
	printFunctionName()

	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	fmt.Println("c = ", c, "len = ", len(c), "cap = ", cap(c))

	go calcSum(s[len(s)/2:], c)
	go calcSum(s[:len(s)/2], c)
	first_half, second_half := <-c, <-c

	return first_half, second_half
}

func createArrayByChannel(n int) []int {
	printFunctionName()

	arr := []int{}
	wg := sync.WaitGroup{}
	wg.Add(n)

	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}

	wg.Wait()
	close(ch)

	for i := range ch {
		arr = append(arr, i)
	}

	return arr
}

func Main() {
	arr := createArrayByGoRoutine(10)
	fmt.Println("arr = ", arr)
	fmt.Println("len = ", len(arr))

	arr = createArrayWithMutex(10)
	fmt.Println("arr = ", arr)
	fmt.Println("len = ", len(arr))

	first_half, second_half := sumByChannel()
	fmt.Println("first_half =", first_half, "second_half =", second_half)

	arr = createArrayByChannel(10)
	fmt.Println("arr = ", arr)
	fmt.Println("len = ", len(arr))
}
