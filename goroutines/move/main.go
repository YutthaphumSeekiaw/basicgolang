// //=========================================================================== Example Goroutines
// // package main

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // func main() {
// // 	// สร้าง Goroutines ใหม่ 2 ตัว
// // 	go doSomething1()
// // 	go doSomething2()

// // 	// รอให้ Goroutines ทั้งสองตัวทำงานเสร็จสิ้น
// // 	time.Sleep(time.Second)

// // 	// พิมพ์ข้อความ
// // 	fmt.Println("Done")
// // }

// // func doSomething1() {
// // 	fmt.Println("Doing something 1")
// // }

// // func doSomething2() {
// // 	fmt.Println("Doing something 2")
// // }

// // =========================================================================== Example Channel
// //========การใช้งานเเบบ  sync
// package main

// import (
//   "fmt"
// )

// func main() {
//   // สร้าง Channel ใหม่
//   ch := make(chan int, 1)

//   // ส่งข้อมูลไปยัง Channel
//   ch <- 1

//   // รับข้อมูลจาก Channel
//   v := <-ch

//   // พิมพ์ข้อมูล
//   fmt.Println(v)
// }

// //========การใช้งานเเบบ  async
// package main

// import (
//   "fmt"
// )

// func main() {
//   // สร้าง Channel ใหม่
//   ch := make(chan int)

//   // สร้าง Goroutine ใหม่
//   go func() {
//     // ส่งข้อมูลไปยัง Channel
//     ch <- 1
//   }()

//   // รับข้อมูลจาก Channel
//   v := <-ch

//   // พิมพ์ข้อมูล
//   fmt.Println(v)
// }

// //=========   buffer channel

// package main

// import (
//   "fmt"
// )

// func main() {
//   // สร้าง Buffer channel ขนาด 10 ตัว
//   ch := make(chan int, 10)

//   // ส่งข้อมูลไปยัง Buffer channel
//   for i := 0; i < 10; i++ {
//     ch <- i
//   }

//   // รับข้อมูลจาก Buffer channel
//   for i := 0; i < 10; i++ {
//     fmt.Println(<-ch)
//   }

// }

// // ผลลัพธ์
// // 0
// // 1
// // 2
// // 3
// // 4
// // 5
// // 6
// // 7
// // 8
// // 9

// //=================  select statement
// package main

// import (
//   "fmt"
//   "time"
// )

// func main() {
//   ch := make(chan string)

//   go func() {
//     ch <- "Hello, world!"
//   }()

//   time.Sleep(1 * time.Second)

//   select {
//   case msg := <-ch:
//     fmt.Println(msg)
//   default:
//     fmt.Println("No message received")
//   }
// }

// //=================== Wait group
// package main

// import (
//   "fmt"
//   "math/rand"
//   "sync"
//   "time"
// )

// func worker(id int, wg *sync.WaitGroup) {
//   defer wg.Done() // Decrement the counter when the goroutine completes

//   fmt.Printf("Worker %d starting\n", id)

//   // Simulate some work by sleeping
//   sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
//   time.Sleep(sleepDuration)

//   fmt.Printf("Worker %d done\n", id)
// }

// func main() {
//   var wg sync.WaitGroup

//   // Launch several goroutines and increment the WaitGroup counter for each
//   for i := 1; i <= 5; i++ {
//     wg.Add(1)
//     go worker(i, &wg)
//   }

//   wg.Wait() // Block until the WaitGroup counter goes back to 0; all workers are done

//   fmt.Println("All workers completed")
// }

// //========================= mutex
// package main

// import (
//   "fmt"
//   "sync"
//   "time"
// )

// var m sync.Mutex

// var n = 10

// func p() {
//   m.Lock()
//   fmt.Println("LOCK")
//   fmt.Println(n)
//   time.Sleep(1 * time.Second)
//   m.Unlock()
//   fmt.Println("UNLOCK")
// }

// func main() {
//   fmt.Println("FIRST")
//   go p()
//   fmt.Println("SECOND")
//   p()
//   fmt.Println("THIRD")
//   time.Sleep(3 * time.Second)
//   fmt.Println("DONE")
// }

// //=============  mutex ทำ  counter

// // ตัวอย่างการใช้งาน

// // การป้องกันการใช้ Resource ร่วมกัน เช่น ไฟล์ ฐานข้อมูล หรือโครงสร้างข้อมูล
// // การประสานงานระหว่าง goroutines เช่น การเข้าแถวรอทรัพยากรหรือการสื่อสารระหว่าง goroutines
// package main

// import (
//   "fmt"
//   "sync"
// )

// // Counter struct holds a value and a mutex
// type Counter struct {
//   value int
//   mu    sync.Mutex
// }

// // Increment method increments the counter's value safely using the mutex
// func (c *Counter) Increment() {
//   c.mu.Lock()   // Lock the mutex before accessing the value
//   c.value++     // Increment the value
//   c.mu.Unlock() // Unlock the mutex after accessing the value
// }

// // Value method returns the current value of the counter
// func (c *Counter) Value() int {
//   return c.value
// }

// func main() {
//   var wg sync.WaitGroup
//   counter := Counter{}

//   // Start 10 goroutines
//   for i := 0; i < 10; i++ {
//     wg.Add(1)
//     go func() {
//       defer wg.Done()
//       for j := 0; j < 100; j++ {
//         counter.Increment()
//       }
//     }()
//   }

//   wg.Wait() // Wait for all goroutines to finish
//   fmt.Println("Final counter value:", counter.Value())
// }

// //========================================================================= once
// // ตัวอย่างการใช้งาน

// // การเปิดการเชื่อมต่อฐานข้อมูล
// // การเริ่มต้นบริการ
// package main

// import (
//   "fmt"
//   "sync"
// )

// func main() {
//   var once sync.Once
//   var wg sync.WaitGroup

//   initialize := func() {
//     fmt.Println("Initializing only once")
//   }

//   doWork := func(workerId int) {
//     defer wg.Done()
//     fmt.Printf("Worker %d started\n", workerId)
//     once.Do(initialize) // This will only be executed once
//     fmt.Printf("Worker %d done\n", workerId)
//   }

//   numWorkers := 5
//   wg.Add(numWorkers)

//   // Launch several goroutines
//   for i := 0; i < numWorkers; i++ {
//     go doWork(i)
//   }

//   // Wait for all goroutines to complete
//   wg.Wait()
//   fmt.Println("All workers completed")
// }

//=================================================== cond
// ตัวอย่างการใช้งาน

// การรอเหตุการณ์บางอย่าง เช่น ข้อความจากเครือข่ายหรือความพร้อมใช้งานของทรัพยากร
// การประสานงานระหว่าง goroutines เช่น การรอให้ goroutine อื่นทำงานเสร็จสิ้น
package move

import (
	"fmt"
	"sync"
	"time"
)

func move() {
	// Create a new condition variable
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	// A shared resource
	ready := false

	// A goroutine that waits for a condition
	go func() {
		fmt.Println("Goroutine: Waiting for the condition...")

		mutex.Lock()
		for !ready {
			cond.Wait() // Wait for the condition
		}
		fmt.Println("Goroutine: Condition met, proceeding...")
		mutex.Unlock()
	}()

	// Simulate some work (e.g., loading resources)
	time.Sleep(2 * time.Second)

	// Signal the condition
	mutex.Lock()
	ready = true
	cond.Signal() // Signal one waiting goroutine
	mutex.Unlock()
	fmt.Println("Push signal !")

	// Give some time for the goroutine to complete
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Work is done.")
}
