package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.3" )

func w0mA5ZuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmGYlpVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AYGklfbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQe0yv8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDrUPiUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3F05DKRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLoNzlveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUjKs8FoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIwbM4o3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0EgVuXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwcZKjGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giSxy1a9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bcs4PMa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x171G2aaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1rUL5rZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5AhyL03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qmkr00NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgclzRAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpAoIwZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSjvbP0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xNo7AgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8bsC5lbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCMDnfmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnNDEwp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpJHl57DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODHrHUF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hfpl77MgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbNHYHS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjSy2DCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppF14esOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2WfTbrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4Qd2AwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzKeLmtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1UerwLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qboSlUxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyHH0mhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCPhax4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxoqC1zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEFHAFIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgpFcevbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjt8CmdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0STjgGnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4fwKnHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ck7XhvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2mBD9oL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dKeIYL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ztw25tgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMmO9RopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zph663PWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VXUJoBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHCeI3ZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWspRHmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6eXzMWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WE204uoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uCtBSr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJz1Ylx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBlPffDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5S59temWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9DWyuAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwuv8rpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XtezQ6dVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWtfpIYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vvQZbDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O71ftW6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pufsbKgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func il3hEd45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVPEXs43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVFKRtFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBfbridaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROimO6hIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YetNlSR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7hae4iqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PU9Jt4jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJiYVl5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqRXlHl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUc7oSsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEBsjsg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtDfmuuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hxwMoM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIPmrlChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYvU3RHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdcqFiK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtxzMY6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZyRqO6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrmtDLZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xCGxCboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XB3PLIH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYDsHrF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnQ9VMH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJ8aYAMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJUoLcqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rS7TZoDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnURAtvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCelktNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6ryqktgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwNDHhccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXliCXbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGDwMSZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4oTVHntsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIf1pUJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYEFrdAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBNZmYqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2niVugx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6otaEhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhjmKpBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BOAw4b0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkJKXbL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXqqm3vUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HE4oi1zQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktTAFN4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HnCpf90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1ym3HvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDGWy2cvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zk5jGMBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPnXEPbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2y9juSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEf9MyeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a94G1g2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXlWwudPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyum8Z8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdoJhobDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qz8Vd1ZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sF8W2KRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pblWZ65xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UByEtt4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8w3JWCdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAfPyZpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glKVXBjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KegNy10TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQyx1CpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZEnqL5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwpU9p06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TsNxPN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUlDQM9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3eNJBPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjymSTbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pb9rDwduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iv4RWv06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmMIgkAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFfnbmJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRzc0pgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZE7pnWHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1RUDlZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ku2Kc2u5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8ZWZiZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpunUxwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AC1FNZBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoapTIrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wm3S5N5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSEvisRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDuyutxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnaa50lCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxi84mJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wBEByZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kW3tND8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUofJDfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzSuhPoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHbA6SqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGkDDGJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96rV9forWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZu1Uqs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdtDWpQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvvNau7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymeQo6LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3PGsvUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83IuZ7kmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCuNMJZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnReTKrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZfC66ThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WqHnOStzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWc1khPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUcFgKsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAP78g0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpyaVoG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqSpEKqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9C8W2N1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zejubHyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZFW2VNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoSp3dl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20HNZbA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3xTXsKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLtYFyNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRXQWGVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgYf66jzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGQcE7HTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YnYdCDzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zMtDLwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DqvRs9ZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyo3XHotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHUilnkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5McZKOICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkgMuraLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnRC1hibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGFt49F4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLaTukqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNmU9J1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hetm988iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHBG0R0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hajEl0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bmdWgO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCrgLgYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAuoY7wRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reemOsuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dL7CUJLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAHzmxTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4jphme9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJG5GmayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmOTl8t7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUgoxyU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1r2YinlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8sV7KgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SObEEVi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skmCQUbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5KV6JmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TV8VSz0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3pFrsUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70XzMru9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDfoEznBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgZPDZRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9FsK1bhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XuGaSwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uprrixw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ET6I4n54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsDJ6M4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztq5cJKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qzrqHDDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V49CFRfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lN3NPXMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgJb8D7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SydXneOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIWPXGddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cGiRGEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ck7cnMpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8W5QGuhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2I3PrKjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fg7SPlrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOSh0zWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T33hfD4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5Ps7Id3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3qX2Tx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csybKfAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ae12IXJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7UUjk0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRtzAFBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvcK7YBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbZEj0SiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQKl41TFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRagS1dIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PfjjvMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4yInILWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajVjMTBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2edpuN1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6cQe7uBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLgszsfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func on4gjtnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ofQL8qSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1fjDKoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C88lPXLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQ8jMwyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovYKgwruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbyyWZ2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fa6U5wFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ms1WTkV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uq2q46GHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQk7WOXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SS6J4M5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLbVlnsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vh2NXNLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVLCRmIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLE7SuuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76fRPjAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsM2rRtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sj1UepuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTHS5Ht1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nz4QcryEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gw9hzdqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1tu3dSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEAo3S4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AChOYfTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lraFAdsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEEpB9sXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZGvfLR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4HULMuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhRbidptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7v6oa3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMwFYbeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIG48EkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0QKTISTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ts2gcyxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsgb0wbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLYTh7RpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiHhHcokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzYyIYhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mhpZhBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuBWnVzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSZXgf5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KrNlLgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRp4ywFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oL6QZz71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gPfal1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lv9QTMINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWMz69iCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqUwvB17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNz6WiTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2gwjQSD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGOTD36XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZFZOMw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQrzxvn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHBSt1NEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEYpiJ7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVxBYtbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z43e8PG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykUvHrCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ut2aSZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4zUhxJ1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRMq70cNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vv2qEiLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ev4uxElbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjNf1A45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkzY4GXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2t1gtrm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EY7vQ9OZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Al5G1OgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vgaSDdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7w37JDqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESugtxrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9S7ynxC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOeMKH4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8MMINotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQVwiLgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ufq114DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1qszLLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoF0qg9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6ljAoZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyX1zc0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydVPRBQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gMxOlZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFxVApH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCQnpGSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tskt4xUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHIYHInQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K49Q6ObQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFR3bsbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iz9Vu6IZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MEG1Ob9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqTqhwrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1j8MPekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P05Cg2jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpFJTqlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwrxWk7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c50vSUNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55UBmzQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WaLVMDR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiOaEfN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WORAUmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68fQkGWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyYtei1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhJX7fNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZsHG6SeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3LU8MfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvafiGy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dj4oJQvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJzFzr8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLao0vcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBOBc9GqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9UeLCcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyDGyEfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAv1IOj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLcH5FMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igWaOoSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYl56urUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w1WQy9eJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzibFot1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rz1lOWrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvzmyqBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klqcwV5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRSodyf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUbZGqQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3nRMRGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EZotmwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZcX3wltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FCc1PqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5Zzr2kSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tThwHmAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4d7IvvAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANxsenihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAIo9LHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vv5l5RJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4vKh0ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbjJ6kCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rb7iaJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAuww0i2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfzI2OiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLyBd1kYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LmKVMEkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxpQLpbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXF5nV9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWJZYxG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qWm3oSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZKshnkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHqa1OQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umG8WyvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBe4MyhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rXSqT0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoDicm58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okhpY128Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrmycXlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIbv1DReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMVGCuYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhqsOwBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9K45SLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ln1h2d9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TrW3Ml5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1ova235Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F29E7ctNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXcyrLF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ag6M8Vk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1rBs4HzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2RdwhwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eS6i6GvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AYcajglgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYMtZLNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4aylYuiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OpWKiW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qNe8OKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbhN4kCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wq7FnxjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSsdBxTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i18tkhdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCm1XLLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxQFMxvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3rt389iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLbYH5syWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDbEbdmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhmsgyZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nulg40uXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oru9AEV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aJdR41aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1gUOErFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMnEzZM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uf1VSlkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhMislrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQTOp32UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAnrFimbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meyefMNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSNlNCqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ET1pOdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3My90JoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojXoZmsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKANFe70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXUOkBfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1abrmVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7L9aQXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CmDHi0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rh2dTbjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBGB3dMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHKno19rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZn2Yon3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LL9SAm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwqBBBffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VstMJ8jTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiVTUZuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5y1nSPIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxSA2ePGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIQg7ZnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEBVlqsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtXKSOmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yo1NcMOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbwp6xWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTUQn9eEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSr3OmdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNMJztYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0RKw7pcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hOG6N1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KV4jdyvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDHmvkjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3axxFfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usR9hpGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVBU3T8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1HHa2d5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKj6rvwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgvcUC3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjYj6JgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQMMnz0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xAqbCZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eradwj52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIJUkjraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wVtrRWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HrJTl9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xDAAQl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2VRkHzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWDm1zMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tv0IMOP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTBe7276Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UX5VWgRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DQIBaZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJblNva0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKyJComsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vGR3vU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPc025HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdyUj0KqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYRa1xAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGT8wyc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcAJvhmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4rAqNM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCNAy06rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcqjFhIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zU7GjDUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04tOGMSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuJYeWlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Je0Pb8EDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stObDExxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baj115suWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qV3LVDC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBsbLUA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMcqxw9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnU0qNXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7q8JQUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func af0oQROAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFLDXqK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pByiNQ97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSF305ERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkasxeu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9vGZh6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Yo6iaDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMgY5FaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9x8yBbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqvBwFd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wv0rpTeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeAWdoeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oruwVTUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pg0RvncHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zc6qO6aeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vbJqIviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2oDIqkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5fujMnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbTrZ9iLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dijPaY6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o96MdCkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLKY6rnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gW30d5geWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHK0jGrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4mahwE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuM8ngVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nBtbgaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFs3Pqg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lhL2yZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpHzmjbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0N5yVWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1MLXUTtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pn5qkjhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJWvX5FdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjGSlplhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ga0ZwgRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oaVHWztUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGfhxCxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jh2sElkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRISbHfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcCnNruOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVSEmCnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBKnqMo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1E6kuwBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYyyDRAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyBiYx33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQMiBDoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaKuOcLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHoSd5kgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7FPOjxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wT1wsuBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYJCBeXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bs2i3wkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAgDVQOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Hpbo3rrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYqmWKaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZEqMky8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKHW9RoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTQyllTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmWbQJJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVh4dLVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8nKiIyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssC8fNTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8oXYQdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecl65RpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGsVPLjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ty9HBY3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJCHw2e9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKUOVabPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vu08wIeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmG1NkGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ym5RZBpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3BJTgs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGcxF9ucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ezlvbhk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1cqGxa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glVBvBPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsc2Hbz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWCqN5SZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCIbMypXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8m9OogOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zdg56K3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ESD8lTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcjSMK6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2c0I7bwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NGD2isSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ew5wCCY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWthkFruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waRo1dJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDRAfBJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5hDxIphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJwkgmNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70BpciK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dUxtLJfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjGP6bp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76Am8AtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwkEgc4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHq5261IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HykphybTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJPsDiG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GJmlr6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ss0u7EF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPSuGyliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDqob4SbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9L8Zm32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsYNlUYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqGPH5orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71JyNt60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQgkCtKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtWheKDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XDPuzFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STyuCT8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hb4XJBOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNrBQtrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilKFmcqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qtl7rktgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8MuhHBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulzsrmD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MC75nPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0x55nyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJPmnjovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZeO1ehsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDhSQumDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M75xFt3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xExrU0mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ImsopPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iYR1dFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNyxkMO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxNF9Gi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SbBN1jrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9FLcLOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ITWcZsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trnIXrx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ev4yTXh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hK2wfWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkQFL7d2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UfH2IzGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmbJThk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4KByKDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuUYzFeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqRyIi5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmRJGyEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTlXtAXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvuYUNXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8HIgyYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNUzwMHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GG8Yr6JkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bC8yfwljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubjVLhMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPwBthn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQpAcuF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knKJI6V3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4QsyUFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiSSDSUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQfuVyGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45uCa8qZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1yt0iUiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuDaITYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTlsURdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JacZ5FbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPvcKYVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7IjhEwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rdSyJd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMwSrFHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgOvjTLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkhXv7YsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huDc0CYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arzREugjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnUAtYJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXPyQJ7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPXcZiI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZ6RABYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nivuOUVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1Yjo2yZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjwEIN2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWzT3EUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTXEE91LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcdmN1HhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lm2Ro2bzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YArSM67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzCLql4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFqvl25sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJoVzbAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2ImznOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIHgAO53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDxbgYDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtEDuRiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcdJRg1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBwo07A4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5WXH1YvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqW5nt1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBMNImXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6ehC5PJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdalbhrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func saN5smBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvyKH7UKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOc0fT3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JljFOgozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1x53MBUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIPFF13hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfpmVO5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWfdO3PcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func os1OtwZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IAGqlaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qz7p1OQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lI38zYfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcmi0NDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXyrNIEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxKALC3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cfn7vx6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpIrNiY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoRGdWt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2RHcJAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zGY7mnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbzpQhfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IFeEhBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLV2QqRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXkr5Z29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJ3qNvEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eN6IEHqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZLNuiVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeUegVBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIYKOUsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5qb4ZXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHNOXQIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KomqGBgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5Ywe8pZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVow8NtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mqQrtXafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqZXYuMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SHHEskwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZivsSPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xSqVE9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUMggCcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qcBy08tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyvyRlDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VzQDESYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFKzfsHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enNQm3frWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHPCTHw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6gwhrl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQOHycThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzyWm7XIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzIL2PriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsTpbYSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNmHtJxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJ87Ly2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sjo17PmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7ouplcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3yrLvwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0ENbKCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZlCmM1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UruiCkmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DU6eLVOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBT2FAtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qtkTYk1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JGffGSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYzvp9pBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nsb11R2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zvpiggo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izSu1RMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bObXYvInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQgMmV56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNfNAwTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AM2kGmiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERvulpkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qz8NrywgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SI5uJcZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkeCqBKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFrOC6uDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQxKtU4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJgvO8sPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCapEHj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwRQHsZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a052zJTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbUW2fOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func us3ZxvoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2400mm2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dWYb1XYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6BdVcuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVrFGia0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uj3EDas6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmSjtU5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAzi2euqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySWtW1oCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFnS7AV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LalVswaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxTosRNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nexwJWULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQHoFm82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JS9pziEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 508YXWcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nz9AkMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWxSPJ64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UchvJ2R0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kkb7ymbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6izmT27hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JO75KnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOzVgd0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEZmOZTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhmlBlXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUmKvEaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVGM3EH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHhTjENxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcGDOL8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZuF1RFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orl0ZzZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLHTJO7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xn6ndjG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2R7Pr7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k69FlbUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bkv89zX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvPsT5ESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5q1ZOl8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFx8l3dMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHyKqxuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDDEbZENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSWSPkElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeitpCarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUpbGaIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 311ztwZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0BbdOoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAfvr2HbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXJyxpo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvG7Io3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glTRErsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ee4FwUZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1s4WV31jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfeaTg2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWpasm1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wdptn6G6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mu1BT82XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGLpNFdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUbjgYcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8A8QvIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njrxm16TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLMidWMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rs3jrrXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9xbzSgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtY3loLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OBrdDP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnWbpvIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRDblQ5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAlqU1AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2WgwX7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVsg0sdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCNbs5fuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTvbe8eOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6rgRejnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iR6ZIhaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZt4biQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEXA0Z2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhQctYSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btOGWJhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIPCFupbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GuQDhKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3b5MhoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDawcukXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSGbwNFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func naR9fsZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGdvMJWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfwFe7MTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mupObLsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHfjLUFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqktYvFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eVT5xstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6986hfuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kVQzt3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOFS1uXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRcdi82aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Djj8Fr0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwkeSJ81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuqLi4NwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3EyeiHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9MwuGTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p74fgXc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyoNiHfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qmu4C5ykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWIti1wLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlnWdEnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2n3mOtLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mty060NrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZ7xVzBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7k4l7vdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCBs1AwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsVGp2N6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1aOAxBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 114UAdVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBOBgX6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2fiiOQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6fN4G6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajVoKCniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29B2UKwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RRuZ30CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWX52JqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5vVzKWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpTpgGxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoyiKYUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shAQrDf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COpOhcAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5csl0u8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpkYsVj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeVuIDUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSDjyRq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hliHJ9RQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4awXCkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5AVZ3jIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuzpcZeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7x39JPQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnPpF7zxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBqW6yMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQzpy6liWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdyKwo2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAo3XbzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9sqO3qkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9eNrC1B9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMBU7p5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cu3gTzgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHKENhSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIPwViYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVJv3pBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7vn7vU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAzRh4VeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WX1pVXa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EKTaMCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQkXoi9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGoWrpS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRqbRbnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCRY6qkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQsamLQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDBNiDyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QGSUj6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIMW5WkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcxwGFShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uqt5HatmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoXDTHugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkQz3trVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnmqNEdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkMxqqZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfSg0xNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGb939VgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ci4yCuhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JaM5C9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Twf8Tl9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ehy9PhVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMdDQP0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68P6EUX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bqDoD94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUQOcT3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blaSNx0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Kq9Q0rpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bam9WZ07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kxod09b3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOJBdv1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XMmR25cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nmcwwxj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7iN8O3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4xBdWJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYeuhr1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func em3eqDqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUnm12HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KufaxoKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oP6DSncMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVcblIVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZxCWj0PAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DveKJTDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7115NVsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3plvOVupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZiXgC3BcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xN8quNeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqWtVpqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqRiFKq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rawoN0mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9YdFF58aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1ZxubvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8kPH2hIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TClDSnhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNhI58tTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNvWKgUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOBiz3ULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZebvXdGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HOY8ZQU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEaW0GzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9gbxg4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWz2O4rjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEpINMdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rydi05vOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pd7TQ9VxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mV79NongWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxqvTZkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8H7KBxWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loEXO2DqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFxMrBkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMLwq9MdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wumJwtBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l19AictVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZ7lQ3ePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DY8POTbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MIv2CY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irQAtLW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFgxMhZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbH8UJc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q0uR4JAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwDwSV8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWNWExWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBbndtcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFjMr8xaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFJAVDCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJm3Sv7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLq6obKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfGnd2ZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UN0jrdrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQ7Jf2GHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jF50dJBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdUUjgLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lZNyNZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ksf4iFLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PO555rFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6vnBzCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gV3MNJkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKauqv82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ha3CUINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZ8vBuV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0i6o61iaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdRdZVqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGg0tZzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnso0OlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VS4KW6qaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LntlrbCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LpebmBlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWRbfrb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIwyrOflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQyS6vhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fn5RMLZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TvckE32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMjOLMAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFeAxtmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXZiVCjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CintZBWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNPzgpkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvDkHFCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBbYiRn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wB0chRbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apsozexAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEOi0cMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nPwZkNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwiRoiceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p50DqThJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E38V4Q8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXcllQ9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SsTBpnMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6k4ZcQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8JV30rYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMS5GuBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g350WiRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

