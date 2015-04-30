package main
 
import (
    "fmt"
    "net"
    "io/ioutil"
    "encoding/json"
    "./process"
)

func main() {
    mode :=1
    
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
    addr_listen :=conf["listen"].(string)
    addr_connect :=conf["connect"].(string)
    
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
        
        go process.Server(mode, conn_listen, addr_connect, "0", "", "", "", "")
    }
}

