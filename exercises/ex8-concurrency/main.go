package main

import (
	"time"
	"fmt"
	"github.com/chemicL/golang-exercises/present"
	"math/rand"
	"sync"
	"context"
)

const WaitToFinish = true

func main() {

	launchGoroutine(!WaitToFinish)

	//channels()

	//bufferedChannels()

	//closingChannels()

	//sequences()

	//multiplexing()

	//multiplexingWithBufferedChannels()

	//synchronizedAccess()

	//goroutinesAndContext()
}

func launchGoroutine(waitSomeTime bool) {
	present.Header("Goroutines.")

	// go keyword delegates execution to the goroutines scheduling mechanism.
	// The execution can run in a separate thread, though it's not necessary.
	// These goroutines share address space, so synchronized access might be required.
	// We will discuss different ways to share data instead of accessing same memory regions too.

	go helloWorld()

	fmt.Println("My friend hasn't woken up yet, she'll greet you in a sec.")

	if !waitSomeTime {
		fmt.Println("Unless the program finishes first. Ooops.")
		return
	}
	time.Sleep(2 * time.Second)
}

func helloWorld() {
	time.Sleep(time.Second)

	fmt.Println("Hello, World!")
}

func channels() {
	present.Header("Channels.")

	// Go allows for communicating between goroutines via messages as an alternative to sharing memory.
	// Go's slogan is: "Do not communicate by sharing memory; instead, share memory by communicating."
	// Operations on channels need no explicit synchronization mechanisms, the runtime makes sure it's safe to use channels.

	// Create a channel for sending ints using built-in make.
	ch := make(chan int)

	go produceRandomNumber(ch)

	fmt.Printf("Waiting before trying to read from channel to experience blocking.\n\n")
	time.Sleep(10 * time.Second)

	fmt.Println("Will read from channel now.")

	// Try to receive a message from the channel. Block until value is available.
	number := <-ch

	fmt.Println("Read a number from channel:", number)

	// Wait a bit to let the goroutine print last statements.
	time.Sleep(2 * time.Second)
}

// chan T signifies the type for channel of type T
func produceRandomNumber(c chan int) {
	fmt.Printf("About to send a random number to the channel.\n\n")
	// Let's send something to the channel. The arrow operator shows the direction of data flow.
	// The call blocks until the receiver part is ready in case of non-buffered channels.
	// For buffered channels, they block only when buffer is full.
	c <- rand.Int() % 100 // Side note: every time we run the program we receive the same numbers - why?
	fmt.Println("! Successfully sent the random number to the channel.")
}

func bufferedChannels() {
	present.Header("Buffered channels.")

	// Second parameter to make when creating a channel means the buffer length.
	ch := make(chan int, 10)

	go produceRandomNumber(ch)

	fmt.Printf("Waiting before trying read from channel to experience no blocking.\n\n")
	time.Sleep(10 * time.Second)

	fmt.Println("Will read from channel now.")
	number := <-ch

	fmt.Println("Read a number from channel:", number)

	// Wait a bit to let the goroutine print last statements. It's unnecessary here, because it should have already be finished.
	time.Sleep(2 * time.Second)
}

func closingChannels() {
	ch := make(chan string, 1)

	go computeAnswer("What's up? ", ch)

	a, ok := <-ch
	if !ok {
		fmt.Println("No answer given.")
		return
	}
	fmt.Println("Provided answer:", a)
}

func computeAnswer(question string, ch chan string) {
	fmt.Println("Question:", question)
	if question == "What is the answer to the Ultimate Question of Life, the Universe, and Everything ?" {
		time.Sleep(time.Second) // Should be 7.5 million years, but using 1s for testing purposes.
		ch <- "42"
	}
	// Closing channel should always be performed by the sender, not receiver.
	// It's not necessary to close a channel. It's useful when the receiver is interested in knowing about end of communication.
	// Check the sequences function to find out when it's useful.
	close(ch)

	// Sending data to a closed channel yields a panic.
	// ch <- "Ooooops."
}

func sequences() {
	present.Header("Sequences using channels.")

	ch := make(chan uint)
	go generateEvenNumbers(10, ch)

	for i := range ch {
		fmt.Println("Got number:", i)
	}
}

func generateEvenNumbers(n uint, ch chan uint) {
	defer close(ch)
	var end uint = n
	if n%2 == 1 {
		end = n - 1
	}
	for i := uint(0); i <= end; i = i + 2 {
		ch <- i
		time.Sleep(time.Second)
	}
}

func generateEvenNumbersUntilClosed(numbers chan uint, shutdown chan bool) {
	i := uint(0)
	// Usually an endless loop.
	for {
		// We use select statement to multiplex between channels.
		select { // Blocks until one of the cases can run. Chooses one at random if multiple are ready.
		case numbers <- i:
			i += 2
		case <-shutdown: // This case reacts to signals on the shutdown channel.
			fmt.Println("Shutting down generator.")
			return
			//default: // If we want to busy spin instead of blocking.
			//	fmt.Println("Not doing anything useful.")
			//	time.Sleep(50 * time.Millisecond)
		}
	}
}

func multiplexing() {
	shutdown := make(chan bool)
	numbers := make(chan uint)

	go generateEvenNumbersUntilClosed(numbers, shutdown)

	fmt.Println(<-numbers)
	fmt.Println(<-numbers)

	shutdown <- true

	time.Sleep(2 * time.Second)
}

func multiplexingWithBufferedChannels() {
	shutdown := make(chan bool)
	numbers := make(chan uint, 100) // Allow to generate 100 numbers ready to be processed.

	// Uncomment the default case in select to observe difference between buffered and unbuffered channels.
	go generateEvenNumbersUntilClosed(numbers, shutdown)

	// The numbers channel is buffered, so the generator doesn't need to wait for the receiver to read the next number.
	// It will fill the buffer and only then jump to the default case in select statement.
	// When we start consuming, it can produce the next numbers.
	fmt.Println(<-numbers)
	fmt.Println(<-numbers)

	shutdown <- true

	time.Sleep(2 * time.Second)
}

func synchronizedAccess() {
	present.Header("Synchronized access.")

	count(&UnsafeCounter{})

	count(&SafeCounter{})
}

func count(counter Counter) {
	for i := 0; i < 1000; i++ {
		// Launch a goroutine for each increment.
		go counter.Inc()
	}

	// Wait one second for the result.
	time.Sleep(time.Second)

	// Check expected value.
	if value := counter.Get(); value != 1000 {
		fmt.Println("")
		fmt.Println(fmt.Sprintf("Oh my gosh. Got %d from %T", value, counter))
		return
	} else {
		fmt.Println(fmt.Sprintf("Ok. Works. Got %d from %T", value, counter))
	}
}

type Counter interface {
	Inc()
	Get() int
}

type UnsafeCounter struct {
	Counter int
}

func (c *UnsafeCounter) Inc() {
	c.Counter++
}

func (c *UnsafeCounter) Get() int {
	return c.Counter
}

type SafeCounter struct {
	// Mutex synchronizes both reads and writes.
	// Check more available synchronization primitives at https://godoc.org/sync.
	mux     sync.Mutex
	Counter int
}

func (c *SafeCounter) Inc() {
	// Call lock before entering critical section.
	c.mux.Lock()
	// Call unlock once we're done. Make sure to catch any errors and avoid deadlocks.
	// In our case it's just adding an int value, so we're good.
	defer c.mux.Unlock()
	c.Counter++
}

func (c *SafeCounter) Get() int {
	// Also call lock when reading to make sure no corruption happens to the structure we're investigating.
	// With ints there's not that much risk, but consider iterating over a slice or accessing a map.
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.Counter
}

func goroutinesAndContext() {
	// Go provides a useful tool for controlling goroutines' lifecycle - context.Context (https://golang.org/pkg/context/).
	// You start with context.Background() which should be the root of all created contexts.
	// It's never cancelled, doesn't have a deadline nor values.
	handleGeneratedNumbers(context.Background())
	time.Sleep(time.Second)
}

func handleGeneratedNumbers(ctx context.Context) {
	// Contexts form a tree.
	// context.WithTimeout, context.WithCancel, context.WithDeadline create a child context and return a cancel function.
	// Here we branch from the provided context and make sure the context expires in 2 seconds.
	intGenerationCtx, cancel := context.WithTimeout(ctx, 2 * time.Second)

	// Always call the cancel function after the results have been made available, to free resources associated with the context.
	defer cancel()

	const MaxInteger = 10

	integers := make(chan int)

	// The created goroutine uses the provided Context and can react to it's cancellation or deadline.
	go generateIntegers(intGenerationCtx, integers)

	var last int
	for n := range integers {
		fmt.Println("Got", n)
		last = n
		if n == MaxInteger {
			fmt.Println("Got enough integers.")
			break
		}
		time.Sleep(300 * time.Millisecond) // Try changing it to 100ms.
	}

	if last != MaxInteger {
		fmt.Println("Didn't get enough integers. Got to", last)
	}
}

// Always pass context.Context as first param. Never wrap in a struct - static analysis won't help you if you do so.
func generateIntegers(ctx context.Context, integers chan int) {
	n := 0
	for {
		select {
		case integers <- n:
			n++
		case <-ctx.Done(): // React to Context being done. ctx.Done is a channel.
			// We can investigate ctx.Err() to check what has actually happened and act accordingly.
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Context timed out.")
			} else if ctx.Err() == context.Canceled {
				fmt.Println("Context cancelled.")
			}

			// Remember to close the channel to which we will no longer produce data.
			close(integers) // Try commenting it out and see what happens.
			return
		default:
			fmt.Println("Waiting for something to happen.")
			time.Sleep(100 * time.Millisecond) // Always sleep a while in the default case, otherwise you waste CPU cycles.
		}
	}
}
