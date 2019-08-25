package main;
 
import (
    "github.com/fsnotify/fsnotify"
    "log"
    "fmt"
    "os"
    "time"
    "bufio"
    "io"
   

)


func ReadFile (filname string) {
     
     f,err := os.Open(filname)
     if err == nil {
        buf := bufio.NewReader(f)
            for {
                line ,err := buf.ReadString('\n')
                 if err != nil || io.EOF == nil {
                 time.Sleep(1e9)
                 }
                fmt.Println(line)
               }
     }
}

 
func ReadFileOnce (filname string) {

     f,err := os.Open(filname)
     if err == nil {
        buf := bufio.NewReader(f)
            for {
                line ,err := buf.ReadString('\n')
                 if err != nil || io.EOF == nil {
                 break
                 }
                fmt.Println(line)
               }
     }
}

func main() {
     ReadFileOnce(os.Args[2])
    //创建一个监控对象
    f := os.Args[1]
    watch, err := fsnotify.NewWatcher();
    if err != nil {
        log.Fatal(err);
    }
    defer watch.Close();
    //添加要监控的对象，文件或文件夹
    err = watch.Add(f);
    if err != nil {
        log.Fatal(err);
    }
    //我们另启一个goroutine来处理监控对象的事件
    go func() {
        for {
            select {
            case ev := <-watch.Events:
                {
                    //判断事件发生的类型，如下5种
                    // Create 创建
                    // Write 写入
                    // Remove 删除
                    // Rename 重命名
                    // Chmod 修改权限
                    if ev.Op&fsnotify.Create == fsnotify.Create {
                        log.Println("创建文件 : ", ev.Name);
                        s :=  fmt.Sprintf("%s",ev.Name)
             go           ReadFile(s)
                    }
                }
            case err := <-watch.Errors:
                {
                    log.Println("error : ", err);
                    return;
                }
            }
        }
    }();
 
    //循环
    select {};
}
