package main

import (
    "encoding/xml"
)

type Text struct {
    Id string
    Secretkey string
    Content string
    X string
}

type Rates struct {
    XMLName xml.Name `xml:"rates"`
    Items [] Item `xml:"item"`
}
type Item struct {
        From      string `xml:"from"`
        To        string `xml:"to"`
        In        string `xml:"in"`
        Out       string `xml:"out"`
        Amount    string `xml:"amount"`
        Minamount string `xml:"minamount"`
        Maxamount string `xml:"maxamount"`
        Param     string `xml:"param"`
        City      string `xml:"city"`
}


// конструктор для кодируемого текста
func NewText(id, secretkey, content, x string) *Text { // возвращаем указатель на Text
    return &Text{id, secretkey, content, x}
}