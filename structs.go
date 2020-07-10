package main

// модель
type Text struct {
	Id string
	Secretkey string
	Content string
	X []byte
}

type Rates struct {
	XMLName xml.Name `xml:"rates"`
	Item    struct {
		From      string `xml:"from"`
		To        string `xml:"to"`
		In        string `xml:"in"`
		Out       string `xml:"out"`
		Amount    string `xml:"amount"`
		Minamount string `xml:"minamount"`
		Maxamount string `xml:"maxamount"`
		Param     string `xml:"param"`
		City      string `xml:"city"`
	} `xml:"item"`
}

// конструктор для кодируемого текста
func NewText(
	id string, 
	secretkey string, 
	content string, 
	x []byte) *Text { // возвращаем указатель на Text
	return &Text{id, secretkey, content, x}
}