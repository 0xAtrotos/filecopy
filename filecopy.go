package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "os"
    "io"
)

func filecopy(w http.ResponseWriter, r *http.Request){
  r.ParseForm()

  if r.Method == "GET" {
      //fmt.Fprintf(w, "Error Method")
      file := strings.TrimSpace(r.FormValue("file"))
      giveSource := strings.TrimSpace(r.FormValue("source"))
      giveDest := strings.TrimSpace(r.FormValue("dest"))
      fmt.Println("file:", file)

      source := giveSource + ":/" + file
      dest := "C:/" + giveDest + "/" + file
      from, err := os.Open(source)
      check(err)
      defer from.Close()

      to, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0666)
      check(err)
      defer to.Close()

      _, err = io.Copy(to, from)
      check(err)

  } else {

  }
}

func handleRequests() {
    http.HandleFunc("/", filecopy)
    log.Fatal(http.ListenAndServe(":8181", nil))
}

func main() {
    handleRequests()
}

func check(err error) {
    if err != nil {
        fmt.Println("Error : %s", err.Error())
        //os.Exit(1)
    }
}
