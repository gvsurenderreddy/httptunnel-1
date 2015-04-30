package iocopy
 
import (
    "net"
    "fmt"
)


func IoCopy(process func([]byte) []byte, conn_src net.Conn, conn_dst net.Conn, ch_err chan string){
    println("IoCopy. RemoteAddr:",conn_src.RemoteAddr().String())
    defer println("IoCopy leave...")
    
    buf := make([]byte, 65536)
    for {
        nr,errr := conn_src.Read(buf)
        println("IoCopy,Read. RemoteAddr:",conn_src.RemoteAddr().String())
        if errr != nil {
            println("Error reading :", errr.Error()) 
            select {
                case ch_err <- errr.Error():
                default:
            }
            return
        }
        
        dataw :=process(buf[0:nr])
        nw, errw := conn_dst.Write(dataw)
        fmt.Printf("IoCopy,Write. RemoteAddr:%s, nw:%d \n",conn_dst.RemoteAddr().String(),nw)
        if errw != nil { 
            println("Error send :", errw.Error()) 
            select {
                case ch_err <- errw.Error():
                default:
            }
            return
        } 
    }
}
