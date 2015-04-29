package iocopy
 
import (
    "net"
    "fmt"
)

func IoCopy(conn_src net.Conn, conn_dst net.Conn, ch_err chan string){
    println("IoCopy. RemoteAddr:",conn_src.RemoteAddr().String())
    defer println("IoCopy leave...")
    buf := make([]byte, 4096)
    for {
        n,err := conn_src.Read(buf)
        println("IoCopy,Read. RemoteAddr:",conn_src.RemoteAddr().String())
        if err != nil {
            println("Error reading :", err.Error()) 
            select {
                case ch_err <- err.Error():
                default:
            }
            return
        }
        
        n, err = conn_dst.Write(buf[0:n]) 
        fmt.Printf("IoCopy,Write. RemoteAddr:%s, n:%d \n",conn_dst.RemoteAddr().String(),n)
        if err != nil { 
            println("Error send :", err.Error()) 
            select {
                case ch_err <- err.Error():
                default:
            }
            return
        } 
    }
}
