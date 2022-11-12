package main

import (
        "flag"
        "fmt"
        "strings"
        "net/http"
        "time"
)

var (
  DEFAULT_PORT = 4551
  GLOBAL_TIMEOUT = 10
  api_key string
  action string
)

type HostList []string

func (hl *HostList) String() string {
  return fmt.Sprintln(*hl)
}

func (ml *HostList) Set(s string) error {
    *ml = strings.Split(s, ",")
    return nil
}

func PostAction(host string) {
  fmt.Printf("Shutting down host: %s\n",host)
  url := fmt.Sprintf("http://%s:%d/%s",host,DEFAULT_PORT,action)
  ApiClient := http.Client{
    Timeout: time.Second * time.Duration(GLOBAL_TIMEOUT),
  }
  req, err := http.NewRequest(http.MethodPost, url, nil)
  if err != nil {
    fmt.Printf("Unable to make a request to shutdown-server api %s\n",err)
    return
  }
  req.Header.Set("Authorization", api_key)
  res, getErr := ApiClient.Do(req)
  if getErr != nil {
    fmt.Printf("some other error then posting to shutdown-server api %s\n",getErr)
    return
  }
  if res.StatusCode == 200 {
    fmt.Printf("Shutdown confirmed for host: %s\n",host)
    return
  } else if res.StatusCode == 403 {
    fmt.Printf("Unauthorized, bad api key\n")
    return
 } else if res.StatusCode == 401 {
    fmt.Printf("Shutdown command failed on remote host!\n")
 } else {
    fmt.Printf("Unknown state of error, status code: %s\n",res.StatusCode)
    return
  }
}

func ShutdownHosts(hl HostList) {
  for _, host := range hl {
    PostAction(host)
  }
}

func main() {
  var hlist HostList
  flag.StringVar(&action,"action","","What action to take shutdown or cancel ?")
  flag.StringVar(&api_key, "api_key", "", "Shutdown api key")
  flag.Var(&hlist,"hosts","List of hosts to shutdown")
  flag.Parse()
  if action == "shutdown" {
    ShutdownHosts(hlist)
  } else if action == "cancel" {
    ShutdownHosts(hlist)
  } else {
    fmt.Printf("Unknown action provided, please use shutdown or cancel\n")
  } 
}
