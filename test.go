package main

import "fmt"
import "os"
import "net/http"

func main() {
	fmt.Println("hola")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><head></head><body>")
		dir, err := os.Open("pages")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		names, err := dir.Readdirnames(-1)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		for i := 0; i < len(names); i++ {
			fmt.Fprintln(w, "<a href='pages/"+names[i]+"'>"+names[i]+"</a><br>")
		}
		fmt.Println(names)
		fmt.Fprintln(w, "</body></html>")
		dir.Close()
	})
	http.ListenAndServe(":8071", nil)
}
