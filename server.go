package main 

import(
	"log"
	"time"
	"sort"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"github.com/gorilla/pat"
)

type safeArticle struct {
	Title string
	Author string
	Date string
	Route string
	First string
	Body string
	Image string
}

type unsafeArticle struct {
	Title string
	Author string
	Date time.Time
	Route string
	First string
	Body template.HTML
	Image string
}

func readArticle(route string) unsafeArticle {
	s := "./blogs/" + route + ".json"
	
	f, _ := ioutil.ReadFile(s)
	
	r := safeArticle{}
	json.Unmarshal(f, &r)

	t, _ := time.Parse("01/02/2006", r.Date)
	xr := unsafeArticle{Title: r.Title, Author: r.Author, Date: t, Route: r.Route, First: r.First, Body: template.HTML(r.Body), Image: r.Image}
	
	return xr
}

// AxisSorter sorts planets by axis.
type DateSorter []safeArticle

func (a DateSorter) Len() int           { return len(a) }
func (a DateSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DateSorter) Less(i, j int) bool { return a[i].Date > a[j].Date }

func readAll() []unsafeArticle {
	fs, _ := ioutil.ReadDir("./blogs")

	ua := make([]unsafeArticle, len(fs))
	
	sa := make([]safeArticle, len(fs))

	for i := range fs {
		s := "./blogs/" + fs[i].Name()
		f, _ := ioutil.ReadFile(s)

		r := safeArticle{}
		json.Unmarshal(f, &r)
		
		sa[i] = r
	}

	sort.Sort(DateSorter(sa))

	for x := range sa {
		t, _ := time.Parse("01/02/2006", sa[x].Date)
		ua[x] = unsafeArticle{Title: sa[x].Title, Author: sa[x].Author, Date: t, Route: sa[x].Route, First: sa[x].First, Body: template.HTML(sa[x].Body), Image: sa[x].Image}
	}

	
	
	return ua
}

func showBlogs(w http.ResponseWriter, r *http.Request) {
	data := readAll()
	t, _ := template.ParseFiles("./views/templates/blogs.html")
	t.Execute(w, data)
}

func showArticle(w http.ResponseWriter, r *http.Request) {
	data := readArticle(r.URL.Query().Get(":route"))
	t, _ := template.ParseFiles("./views/templates/article.html")
	t.Execute(w, data)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/templates/about.html")
	t.Execute(w, r)
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/templates/projects.html")
	t.Execute(w, r)
}
func index(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/blog", 301)
}

func main(){
	fs := http.FileServer(http.Dir("views/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	p := pat.New()		
		
	http.Handle("/", p)
	p.Get("/blogs/{route}", showArticle)
	p.Get("/blog", showBlogs)
	p.Get("/about", aboutPage)
	p.Get("/portfolio", projectsPage)
	p.Get("/", index)

	log.Println("Server running at port 8080.")
	http.ListenAndServe(":8080", nil)
}