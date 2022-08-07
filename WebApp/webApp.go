package main

import(
	"net/http"
	"io"
)

type HandleFnc func(http.ResponseWriter, *http.Request)

const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form></html></body>`

func SimpleServer(w http.ResponseWriter, request *http.Request){
	io.WriteString(w, "<h1>Hello, World</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type", "text/html")

	switch request.Method{
	case "GET":
		io.WriteString(w, form);

	case "POST":
		io.WriteString(w, request.FormValue("in"))
	}
}

func main(){
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		panic(err)
	}
}

func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request){
		defer func(){
			if x:= recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		} ()
		function(writer, request)
	}
}