package main
import(
    "log"
    "os"
    "runtime"
    "runtime/pprof"
    "time"
)

var ch4cpu chan uint64
var chTimer chan struct{}
var memMap map[int]interface{}

func init(){
    ch4cpu = make(chan uint64,10000)
    chTimer = make(chan struct{},20)
    memMap = make(map[int]interface{})
}

func main(){
    c,err:= os.Create("/")
    if err != nil{
        log.Fatal(err)
    }
    defer c.Close()
    m1,err:= os.Create("/")

    if err != nil{
        log.Fatal(err)
    }
    m2,err:= os.Create("/")

    if err != nil{
        log.Fatal(err)
    }
    m3,err:= os.Create("/")

    if err != nil{
        log.Fatal(err)
    }
    m4,err:= os.Create("/")

    if err != nil{
        log.Fatal(err)
    }

    defer m1.Close()
    defer m2.Close()
    defer m3.Close()
    defer m4.Close()

    pprof.StartCPUProfile(c)
    defer pprof.StopCPUProfile()

    memMap[1] = runMEMTest()
    runtime.GC()
    pprof.Lookup("heap").WriteTo(m1,0)
    

}