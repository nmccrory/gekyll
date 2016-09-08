package main 

import(
	"log"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/blogs/", blog)
	http.HandleFunc("/", index)

	log.Println("Minimal Blog server running on port 7899...")
	http.ListenAndServe(":7899", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//get all blogs and order by date
	data := readAll()
	sortByDate(data)
	t, _ := template.ParseFiles("./views/templates/blogs.html")
	t.Execute(w, data)
}

func blog(w http.ResponseWriter, r *http.Request) {
	data := toUnsafe(read(r.URL.Path[6:]))
	t, _ := template.ParseFiles("./views/templates/article.html")
	t.Execute(w, data)
}
