package main

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
