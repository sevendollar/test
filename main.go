package test

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func Shuffle(i interface{}) (err error) {
    switch v := i.(type) {
    case *[]string:
        if len(*v) == 0 {
            return errors.New("can not be empty slice")
        }
        rand.Seed(time.Now().UnixNano())

        f := *v
        for i := 0; i < len(f); i++ {
            j := rand.Intn(i + 1)
            f[i], f[j] = f[j], f[i]
        }
    case *[]int:
        if len(*v) == 0 {
            return errors.New("can not be empty slice")
        }
        rand.Seed(time.Now().UnixNano())

        f := *v
        for i := 0; i < len(f); i++ {
            j := rand.Intn(i + 1)
            f[i], f[j] = f[j], f[i]
        }
    default:
        return errors.New("must be either pointer of int or string")
    }
    return
}
