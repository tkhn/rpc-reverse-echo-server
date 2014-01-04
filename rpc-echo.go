package main

import (
    "fmt"
    "net"
    "net/rpc"
    "log"
)

type Server struct {}
func (this *Server) Echo(i string, reply *string)  error {
    *reply = Reverse (i)
    return nil
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func server() {
    rpc.Register(new(Server))
    ln, err := net.Listen("tcp", "127.0.0.1:28099")
    if err != nil {
	    log.Fatal("server:listen error:", err)
        return
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Printf("server: error ecountered while accepting a connectionr..")      
            continue

        }
        go rpc.ServeConn(conn)
    }
}
func client(s string){
    c, err := rpc.Dial("tcp", "127.0.0.1:28099")
    if err != nil {
        log.Fatal("listen error:", err)
        return
    }
    var result string
    err = c.Call("Server.Echo", s, &result)
    if err != nil {
        fmt.Println(err)
    } else {
	    fmt.Println(result)
    }
}

func main() {

    var input string
    fmt.Println("Enter whatever you want:")
    fmt.Scanln(&input)

    go server()
    go client(input)
 
    fmt.Scanln(&input)

}
