package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "os"
        "os/exec"
       "strings"
        router2 "github.com/fasthttp/router"
        "github.com/valyala/fasthttp"
)

var (
        HTTP_PORT = 4551
        DEBUG     = 1
        CONF_FILE = "/etc/shutdown-server.conf"
        LOG_FILE = "/var/log/shutdown-server.log"
        API_KEY  = ""
        logger *log.Logger
)


func Index(ctx *fasthttp.RequestCtx) {
  ctx.SetStatusCode(fasthttp.StatusOK)
}


func CaptureCmd(cmd []string) (string, bool) {
  out, err := exec.Command(cmd[0], cmd[1:]...).Output()
  if err != nil {
    fmt.Printf("Exec failed: %s\n", err)
    return string(out), false
  }
  return string(out), true
}

func ShutdownAction() bool {
  logger.Printf("Shutdown action reached\n")
  params := []string{"shutdown", "-h"}
  out, ok := CaptureCmd(params)
  if !ok {
    fmt.Printf("Shutdown command have failed: %s\n",out)
    return false
  }
 return ok
}

func ShutdownCancelAction() bool {
  logger.Printf("Shutdown cancel action reached\n")
  params := []string{"shutdown", "-c"}
  out, ok := CaptureCmd(params)
  if !ok {
    fmt.Printf("Shutdown cancel command have failed: %s\n",out)
    return false
  }
 return ok
}

func ShutdownRequest(ctx *fasthttp.RequestCtx) {
  method := ctx.Method()
  if string(method) == "POST" {
    auth := ctx.Request.Header.Peek("Authorization")
    if CheckAuth(string(auth)) {
      logger.Printf("Authorization ok, requesting shutdown device...\n")
      stat := ShutdownAction()
      if stat {
        ctx.SetStatusCode(fasthttp.StatusOK)
      } else {
        ctx.SetStatusCode(401)
      }
    } else {
      logger.Printf("Bad authorization\n")
      ctx.SetStatusCode(403)
    }
  }
}


func ShutdownCancel(ctx *fasthttp.RequestCtx) {
  method := ctx.Method()
  if string(method) == "POST" {
    auth := ctx.Request.Header.Peek("Authorization")
    if CheckAuth(string(auth)) {
      logger.Printf("Authorization ok, requesting to cancel shutdown device...\n")
      stat := ShutdownCancelAction()
      if stat {
        ctx.SetStatusCode(fasthttp.StatusOK)
      } else {
        ctx.SetStatusCode(401)
      }
    } else {
      logger.Printf("Bad authorization\n")
      ctx.SetStatusCode(403)
    }
  }
}

func CheckAuth(key string) bool {
  if API_KEY == "" {
    return true
  }
  if API_KEY == key {
    return true
  }
  return false
}

func ReadConf() {
  if _, err := os.Stat(CONF_FILE); err == nil {
    tmp_conf, err := ioutil.ReadFile(CONF_FILE)
      if err != nil {
        return
      }
      API_KEY = TruncateMaterial(string(tmp_conf))
  }
}

func TruncateMaterial(str string) string {
  str = strings.Replace(str, "\t", "", -1)
  str = strings.Replace(str, "\n", "", -1)
  return str
}

func HttpPool() {
        router := router2.New()
        router.GET("/", Index)
        router.POST("/shutdown",ShutdownRequest)
        router.POST("/cancel",ShutdownCancel)
        logger.Printf("Binding web service to 0.0.0.0:%d...\n", HTTP_PORT)
        s := &fasthttp.Server{
                Handler:            router.Handler,
                Name:               "Damn it's just http server", // we will use some fake identity
                MaxRequestBodySize: 32 << 20,
        }
        s.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", HTTP_PORT))
}


func Logging() {
   f, err := os.OpenFile(LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("debug log file not created", err.Error())
    }
    logger = log.New(f, "[LOG]", log.Ldate|log.Ltime|log.Lmicroseconds)
    logger.Println("log started")
/*
  logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    log.Panic(err)
  }
  defer logFile.Close()
  logger.SetOutput(logFile)
  logger.SetFlags(log.Lshortfile | log.LstdFlags)
  logger.Println("Logging started")
*/
}

func main() {
        Logging()
        ReadConf()
        HttpPool()
}
