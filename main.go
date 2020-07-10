package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// сохраняем в памяти
var TextForEncrypt map[string]*Text
//var JsonFromXml map[string]*JsonData

func indexHendler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil{
		fmt.Fprintf(w, err.Error()) // возврат ошибки в браузер
	}

	fmt.Println(TextForEncrypt) // вывод в консоль при входе в index

	t.ExecuteTemplate(w, "index", TextForEncrypt) // вывод на страницу
}


func writeHendler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html")
	if err != nil{
		fmt.Fprintf(w, err.Error()) // возврат ошибки в браузер
	}

	t.ExecuteTemplate(w, "write", nil)
}


func encryptHendler(w http.ResponseWriter, r *http.Request) {
	//id := GenerateId()
	id := r.FormValue("id")
	secretkey := r.FormValue("secretkey")
	content := r.FormValue("content")

	x := Encryptor(secretkey, content)

	post := NewText(id, secretkey, content, x)
	TextForEncrypt[post.Id] = post // записываем post в map

	http.Redirect(w, r, "/", 302)
}


/*func getjsonHendler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html")
	if err != nil{
		fmt.Fprintf(w, err.Error()) // возврат ошибки в браузер
	}

	t.ExecuteTemplate(w, "write", JsonFromXml) // вывод на страницу
}*/


func main() {
	TextForEncrypt = make(map[string]*Text, 0)
	//JsonFromXml = make(map[string]*JsonData, 0)
	
	fmt.Println("---listening on port :8000")

	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/encrypt", encryptHendler)
	//http.HandleFunc("/getjsonfromxml", getjsonHendler)
	
	http.ListenAndServe(":8000", nil)


}