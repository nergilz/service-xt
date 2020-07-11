package main

import (
    "fmt"
    "html/template"
    "net/http"
    "encoding/xml"
    "encoding/json"
    "io/ioutil"
    "log"
    b64 "encoding/base64"
)

var EncryptText map[string]*Text // сохраняем в памяти

func main() {

    EncryptText = make(map[string]*Text, 0)
    fmt.Println("---listening on port :8000")

    http.HandleFunc("/", indexHendler)
    http.HandleFunc("/encrypt", encryptHendler)
    http.HandleFunc("/decrypt", decryptHendler)
    http.HandleFunc("/courses", coursesHendler)
    
    http.ListenAndServe(":8000", nil)
}

func indexHendler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/index.html")
    if err != nil{
        fmt.Fprintf(w, err.Error()) // возврат ошибки в браузер
    }

    fmt.Println(EncryptText) // вывод в консоль при входе в index

    t.ExecuteTemplate(w, "index", EncryptText) // вывод на страницу
}


func encryptHendler(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    secretkey := r.FormValue("secretkey")
    content := r.FormValue("content")

    key := []byte(secretkey)
    plaintext := []byte(content)

    ciphertextenc, err := Encryptor(key, plaintext)
    if err != nil {
        log.Fatal(err)
    }
    //s := string(ciphertextenc)
    //fmt.Println(s)
    sEnc := b64.StdEncoding.EncodeToString(ciphertextenc)
    fmt.Println(sEnc)
    post := NewText(id, secretkey, content, sEnc )
    EncryptText[post.Id] = post // записываем post в map
    http.Redirect(w, r, "/", 302)
}

func decryptHendler(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    secretkey := r.FormValue("secretkey")
    content := r.FormValue("content")
    fmt.Println(content)

    sDec, _ := b64.StdEncoding.DecodeString(content)
    text := []byte(sDec)
    key := []byte(secretkey)

    result, err := Decryptor(key, text)
    if err != nil {
        log.Fatal(err)
    }
    s := string(result)
    fmt.Println(s)

    post := NewText(id, secretkey, content, s )
    EncryptText[post.Id] = post
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