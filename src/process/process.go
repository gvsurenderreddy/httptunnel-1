package process
 
import (
    "net"
    "time"
    "../iocopy"
)

    
func Tcp2http(data []byte) []byte {
     //构造post请求
    var post string  
    post += "POST /postpage HTTP/1.1\r\n"
    post += "Content-Type: application/x-www-form-urlencoded\r\n"  
    post += "Content-Length: "+strconv.Itoa(len(data))+"\r\n"
    post += "Connection: keep-alive\r\n"  
    post += "Accept-Language: zh-CN,zh;q=0.8,en;q=0.6\r\n"  
    post += "\r\n"
    post += "focustm="+string(data)+"\r\n"
    
    rsl :=make([]byte,len(data)+6)
    copy(rsl, "<body>")
    copy(rsl[len(data):], data)
    println("Tcp2http, rsl:", string(rsl))
    return rsl
}
    
func Http2tcp(data []byte) []byte {
    if len(data) <5 {
        println("Http2tcp, data:", string(data))
        return data
    }    
    return data[5:]
}

func Server(mode int,
            conn_listen net.Conn,
            addr_connect string,
            puse string,
            phost string,
            pport string,
            puser string,
            ppwd string) {
    defer println("Server leave...")
    defer conn_listen.Close()
    
    // connect to next.
    conn_connect, err := net.Dial("tcp", addr_connect)
    if err != nil {
        println("connect failed.", err.Error())
        return
    }
    defer conn_connect.Close()
    
    // if use proxy
    
    
    ch_err :=make(chan string)
    if mode==0{
        go iocopy.IoCopy(Tcp2http, conn_listen, conn_connect, ch_err)
        go iocopy.IoCopy(Http2tcp, conn_connect, conn_listen, ch_err)
    } else {
        go iocopy.IoCopy(Http2tcp, conn_listen, conn_connect, ch_err)
        go iocopy.IoCopy(Tcp2http, conn_connect, conn_listen, ch_err)
    }
    
    for{
        select {
            case <- time.After(time.Second*6000):
                println("time out")
            case <-  ch_err:
                return
        }
    }
}
 
