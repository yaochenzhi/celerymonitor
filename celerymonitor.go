package main 

import (
    "github.com/go-redis/redis"
    "time"
    "fmt"
)

var (
    addr = "localhost:port"
    password = "password"
    db = 10

    queue = "hang_check_queue"
    queue_limit int32 = 100
    check_interval = "10m"
)

func main() {

    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })


    queue_len, _ := client.LLen(queue).Result()

    if int32(queue_len) > queue_limit {
        fmt.Printf("WANRING: %s unhealthy in the last %s. %d\n", time.Now(), check_interval, queue_len)
        // clear
        client.Del(queue)
        fmt.Printf("Queue cleared! %s", time.Now())
    }else{
        fmt.Printf("INFO: %s healthy in the last %s. %d\n", time.Now(), check_interval, queue_len)
    }
}