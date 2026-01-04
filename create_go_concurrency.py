import os

BASE_DIR = "go concurrency"

topics = {
    "130_concurrency_vs_parallelism.go": """package main

/*
Concurrency vs Parallelism

Concurrency:
- Managing multiple tasks at once
- Tasks may not run at the same time

Parallelism:
- Executing multiple tasks simultaneously
- Requires multiple CPU cores

Go supports concurrency via goroutines and channels
*/

func main() {
    // Example
    // go task1()
    // go task2()
}
""",

    "131_race_conditions.go": """package main

/*
Race Conditions

Occurs when:
- Multiple goroutines access shared data
- At least one write happens
- No synchronization is used
*/

import "sync"

func main() {
    var count int
    var mu sync.Mutex

    mu.Lock()
    count++
    mu.Unlock()
}
""",

    "132_deadlocks.go": """package main

/*
Deadlocks

A deadlock happens when goroutines wait forever
for each other and no progress can be made
*/

func main() {
    ch := make(chan int)
    // ch <- 1 // Uncommenting causes deadlock
}
""",

    "133_rwmutex.go": """package main

/*
RWMutex

- Allows multiple readers
- Only one writer at a time
*/

import "sync"

func main() {
    var mu sync.RWMutex

    mu.RLock()
    // read operation
    mu.RUnlock()

    mu.Lock()
    // write operation
    mu.Unlock()
}
""",

    "134_sync_newcond.go": """package main

/*
sync.NewCond

Used for goroutine signaling and coordination
*/

import "sync"

func main() {
    cond := sync.NewCond(&sync.Mutex{})

    cond.L.Lock()
    cond.Wait()
    cond.L.Unlock()
}
""",

    "135_sync_once.go": """package main

/*
sync.Once

Ensures a piece of code runs only once
*/

import "sync"

var once sync.Once

func initResource() {
    // initialization logic
}

func main() {
    once.Do(initResource)
}
""",

    "136_sync_pool.go": """package main

/*
sync.Pool

Used to reuse objects and reduce GC pressure
*/

import "sync"

func main() {
    pool := sync.Pool{
        New: func() interface{} {
            return make([]byte, 1024)
        },
    }

    buf := pool.Get().([]byte)
    pool.Put(buf)
}
""",

    "137_for_select_statement.go": """package main

/*
for-select statement

Used to listen on multiple channels
*/

import "fmt"

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    for {
        select {
        case msg := <-ch1:
            fmt.Println(msg)
        case msg := <-ch2:
            fmt.Println(msg)
        default:
            // non-blocking
        }
    }
}
"""
}

def main():
    os.makedirs(BASE_DIR, exist_ok=True)

    for filename, content in topics.items():
        file_path = os.path.join(BASE_DIR, filename)
        with open(file_path, "w") as f:
            f.write(content)

    print("✅ 'go concurrency' folder and Go files created successfully")

if __name__ == "__main__":
    main()
