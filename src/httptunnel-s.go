package main
 
import (
    "fmt"
    "net"
    "time"
    "io/ioutil"
    "encoding/json"
    "./iocopy"
)

var addr_listen =""
var addr_connect =""
var puse ="0"
var phost =""
var pport =""
var puser =""
var ppwd =""
    
func Server(conn_listen net.Conn) {
    println("Server. RemoteAddr:",conn_listen.RemoteAddr().String())
    defer println("Server leave...")
    defer conn_listen.Close()
    
    // connect to next.
    conn_connect, err := net.Dial("tcp", addr_connect)
    if err != nil {
        println("connect failed.", err.Error())
        return
    }
    defer conn_connect.Close()
    
    ch_err :=make(chan string)
    go iocopy.IoCopy(conn_listen, conn_connect, ch_err)
    go iocopy.IoCopy(conn_connect, conn_listen, ch_err)
    
    for{
        select {
            case <- time.After(time.Second*6000):
                println("time out")
            case <-  ch_err:
                return
        }
    }
}
 
    
func main() {
    // 配置信息默认值
    confStr := `{
        "listen": "0.0.0.0:18080", 
        "connect": "127.0.0.1:10443"
        }`

    // 读配置文件
    fbuf, err := ioutil.ReadFile("./httptunnel-s.conf")
	if err != nil {
		println("open file failed!")
	} else {
        confStr =string(fbuf)
    }
    
    // 解析配置文件
    var conf map[string]interface{}
    err = json.Unmarshal([]byte(confStr), &conf) 
    if err != nil {  
        fmt.Println("error in translating,", err.Error())  
        return  
    }
    addr_listen =conf["listen"].(string)
    addr_connect =conf["connect"].(string)
    
    listener, err := net.Listen("tcp", addr_listen)
    if err != nil { 
        println("error listening:", err.Error()) 
        return  
    }
    defer listener.Close()

    for {
        conn_listen, err := listener.Accept()
        if err != nil {
            println("Error accept:", err.Error())
            return
        }
        go Server(conn_listen)
    }
}

