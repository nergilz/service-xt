package main

// модель
type Text struct {
	Id string
	Secretkey string
	Content string
	X []byte
}

/*type JsonData struct {
	
}*/

// конструктор для кодируемого текста
func NewText(
	id string, 
	secretkey string, 
	content string, 
	x []byte) *Text { // возвращаем указатель на Text
	return &Text{id, secretkey, content, x}
}