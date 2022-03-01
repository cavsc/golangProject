// package main

// import (
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// )

// const (
// 	UPLOAD_DIR = "./uploads"
// )

// func viewHandler(w http.ResponseWriter,r *http.Request){
// 	imageId := r.FormValue("id")
// 	imagePath := UPLOAD_DIR + "/" + imageId
// 	if exists := isExists(imagePath); !exists{
// 		http.NotFound(w,r)
// 		return
// 	}

// 	w.Header().Set("Content-Type","image")
// 	http.ServeFile(w,r,imagePath)
// }

// func isExists(path string) bool{
// 	_,err := os.Stat(path)
// 	if err == nil{
// 		return true
// 	}
// 	return os.IsExist(err)
// }

// func uploadHandler(w http.ResponseWriter,r *http.Request){
// 	if r.Method == "GET"{
// 		io.WriteString(w,"<form method=\"POST\" action=\"/upload\""+" enctype=\"multipart/form-dat4a\">"+"Choose an image to upload: <input name=\"image\" type=\"file\" />"+"<input type=\"submit\" value=\"Upload\" />"+"</form>")
// 		return
// 	}

// 	if r.Method == "POST"{
// 		f,h,err := r.FormFile("image")
// 		if err != nil{
// 			http.Error(w,err.Error(),http.StatusInternalServerError)
// 			return
// 		}

// 		filename:= h.Filename
// 		defer f.Close()
// 		t,err := os.Create(UPLOAD_DIR+"/"+filename)
// 		if err!=nil{
// 			http.Error(w,err.Error(),http.StatusInternalServerError)
// 			return
// 		}

// 		defer t.Close()
// 		if _,err := io.Copy(t,f);err!=nil{
// 			http.Error(w,err.Error(),http.StatusInternalServerError)
// 			return
// 		}

// 		http.Redirect(w,r,"/view?id="+filename,http.StatusFound)
// 	}
// }

// func main(){
// 	http.HandleFunc("/upload",uploadHandler)
// 	http.HandleFunc("/view",viewHandler)
// 	err := http.ListenAndServe(":8080",nil)
// 	if err!=nil{
// 		log.Fatal("ListenAndServe:",err.Error())
// 	}
// }

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"
)

// func main() {
// 	a := make([]int, 0, 5)
// 	// addElem(a, 5)
// 	fmt.Println(a)
// }

// func addElem(a []int, i int) {
// 	a = append(a, 5)
// }

type HelloService struct{}
func (p *HelloService) Hello(request string,reply *string) error{
	*reply = "hello:" +request
	return nil
}

func main(){
	rpc.RegisterName("HelloService",new(HelloService))
	listener,err := net.Listen("tcp",":1234")
	if err!=nil{
		log.Fatal("LIstenTCP error:",err)
	}
	conn,err := listener.Accept()
	if err!=nil{
		log.Fatal("Accept error:",err)
	}

	rpc.ServeConn(conn)
}