package main 

import (
    // "reflect"
    "os/exec"
    "strconv"
    "bytes"
    "time"
    "fmt"
)

var (
    auth = "G0d1Ike&me"
    port = "7005"
    queue = "ping_hang_check_queue"
    queue_limit int32 = 100
    check_interval = "10m"
)

func main() {
    cmd := exec.Command("/usr/local/bin/redis-cli", "-p", port, "-n", "10", "-a", auth, "llen", queue)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Printf("cmd.Run() failed with %s\n", err)
    }

    // remove line terminator for int parsing later
    queue_len_str := string(stdout.Bytes())
    queue_len_str = queue_len_str[:len(queue_len_str)-1]      // queue_len_str = strings.TrimSuffix(queue_len_str, "\n")
    queue_len, _ := strconv.ParseInt(queue_len_str, 10, 32)  // still int64 to be converted
                                                            // fmt.Println(reflect.TypeOf(queue_len))

    if int32(queue_len) > queue_limit {
        fmt.Printf("WANRING: %s unhealthy in the last %s. %d\n", time.Now(), check_interval, queue_len)
        clear()
    }else{
        fmt.Printf("INFO: %s healthy in the last %s. %d\n", time.Now(), check_interval, queue_len)
    }
}

func clear() {
    cmd := exec.Command("redis-cli", "-p", port, "-n", "10", "-a", auth, "del", queue)
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Queue cleared! %s", time.Now())
    }
}