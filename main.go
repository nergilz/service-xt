package main

import (
    "fmt"
    "html/template"
    "net/http"
    "encoding/xml"
    "encoding/json"
    "io/ioutil"
)

// сохраняем в памяти
var TextForEncrypt map[string]*Text


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

func coursesHendler(w http.ResponseWriter, r *http.Request) {

    xmlResp, err := http.Get("https://test.cryptohonest.ru/request-exportxml.xml")
    if err != nil {
        fmt.Println("---get xml error : ", err)
        return
    }
    xmlData, _ := ioutil.ReadAll(xmlResp.Body)

    var rates Rates
    xml.Unmarshal(xmlData, &rates)

    // Convert to JSON
    var item Item
    var items []Item

    for _, value := range rates.Items {
        item.From = value.From
        item.To = value.To
        item.In = value.In
        item.Out = value.Out
        item.Amount = value.Amount
        item.Minamount = value.Minamount
        item.Maxamount = value.Maxamount
        item.Param = value.Param
        item.City = value.City

        items = append(items, item)
    }

    jsonData, err := json.MarshalIndent(items, "", " ")
    if err != nil {
        fmt.Println(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
    fmt.Println("---return json: /courses---")
}


func main() {
	TextForEncrypt = make(map[string]*Text, 0)
	fmt.Println("---listening on port :8000")

	http.HandleFunc("/", indexHendler)
	http.HandleFunc("/encrypt", encryptHendler)
	http.HandleFunc("/courses", coursesHendler)
	
	http.ListenAndServe(":8000", nil)
}