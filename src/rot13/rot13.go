package main

import (
//    "fmt"
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
    a := make([]byte, len(p))
    //fmt.Printf("made array of len %v\n", len(a))
    
    n, err = r.r.Read(a)
    //fmt.Printf("read returned %v, %v\n", n, err)
    
    if err != nil {
        return n, err
    }
    //fmt.Printf("Contents of array: %v\n", a)

    for i := 0; i < n; i++ {
        v := a[i]
        //fmt.Printf("v=%v, ", v)
        if v >= 65 && v <= 90 {
            v = v + 13 
            if v > 90 {
                v -= 26
            }
        } else if v >= 97 && v <= 122 {
            v = v + 13
            if v > 122 {
                v -= 26
            }
        }
        p[i] = v
        //fmt.Printf("added a %v\n", v)
    }
    return n, err
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
