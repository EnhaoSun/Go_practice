// Netcat1 is a read-only TCP client.
package main

import (
    "io"
    "log"
    "net"
    "os"
    //"fmt"
)

func main() {
    ports := []string{ "8010", "8020", "8030"}
    for i, p := range ports {
        conn, err := net.Dial("tcp", "localhost:" + p)
        if err != nil {
            log.Fatal(err)
        }
        defer conn.Close()
        if i == len(ports) - 1 {
            mustCopy(os.Stdout, conn)
        } else {
            go mustCopy(os.Stdout, conn)
        }
    }
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}

