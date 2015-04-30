package main

import (  
    "strconv"  
)

func main(){
    data := []byte{'a', 'b', 'c'}
    
    var post string  
    post += "POST /postpage HTTP/1.1\r\n"
    post += "Content-Type: application/x-www-form-urlencoded\r\n"  
    post += "Content-Length: "+strconv.Itoa(len(data))+"\r\n"
    post += "Connection: keep-alive\r\n"  
    post += "Accept-Language: zh-CN,zh;q=0.8,en;q=0.6\r\n"  
    post += "\r\n"  
    post += "mydata="+string(data)+"\r\n" 
    
    println("post=",post)
}
