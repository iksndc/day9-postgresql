package main

import (
	"fmt"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	connection.DataBaseConnection()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/contactme", contact).Methods("GET")
	route.HandleFunc("/myproject", myproject).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")

	port := "5000"

	fmt.Print("Server sedang berjalan diport " + port)
	http.ListenAndServe("localhost:"+port, route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset utf-8")
	tmpt, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	data := map[string]interface{}{
		"Title":   "DumbWays Web App",
		"Content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus id pharetra quam. Integer sagittis quis dolor sit amet faucibus. Nullam id ante finibus arcu porta iaculis. Praesent felis augue, pellentesque sit amet cursus id, gravida maximus ante. Quisque et ex sollicitudin, consequat lorem quis, auctor ante. In hac habitasse platea dictumst. In rhoncus pellentesque porta. Integer nibh justo, scelerisque at blandit vitae, posuere et neque. Suspendisse sapien lacus, ornare ut lectus nec, pulvinar pellentesque massa. Proin sagittis luctus lacus, a porta mauris porta non.<br> Duis interdum diam augue, sed sagittis dui tristique a. Curabitur vehicula massa orci, sit amet faucibus tortor lobortis a. Fusce pellentesque mollis leo, et vulputate risus gravida et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Sed lobortis, purus nec tincidunt vulputate, mi eros sollicitudin neque, eget eleifend eros tellus sit amet eros. Nam sit amet elit molestie, imperdiet elit semper, ultrices magna. Vivamus faucibus scelerisque faucibus. Duis scelerisque lorem eget placerat lacinia. Nullam suscipit quam in felis mattis luctus.<br> Phasellus lacus turpis, sollicitudin nec sem ac, ornare posuere mi. Morbi accumsan elit ut justo viverra, eget elementum dolor dapibus. Fusce venenatis, neque vel volutpat volutpat, diam purus vehicula libero, eget faucibus dui risus sit amet dolor. Donec vehicula sed lectus et bibendum. Nullam malesuada, neque sed tempus faucibus, erat quam tristique felis, pharetra euismod est urna quis est. Etiam nec leo nunc. Proin interdum nec erat vel auctor. In gravida lacinia magna, a pretium nisl tempus ut. Fusce id odio sapien. Cras leo nulla, condimentum at auctor vitae, scelerisque sit amet ex. Donec vitae diam a massa venenatis rutrum elementum at enim. Curabitur nec dui magna.<br> Donec tempus dignissim nibh, a cursus massa convallis ut. Sed neque risus, fermentum a purus sit amet, rhoncus malesuada velit. Sed condimentum diam id enim iaculis, eu accumsan nulla finibus. Morbi eget facilisis erat, eu dignissim ante. Morbi cursus efficitur mi at semper. Sed ligula risus, accumsan eu porta quis, porta at leo. Maecenas porta, orci id eleifend pulvinar, nunc lacus facilisis leo, at pharetra urna sapien eu ante. Quisque malesuada massa non tortor ultrices condimentum.",
		"Id":      id,
	}

	tmpt.Execute(w, data)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func myproject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset utf-8")
	tmpt, err := template.ParseFiles("views/myproject.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tittle : " + r.PostForm.Get("title"))

}
