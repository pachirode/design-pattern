package memento

import "fmt"

type History struct {
	body string
}

func NewHistory(body string) *History {
	return &History{body: body}
}

func (h *History) SetBody(body string) {
	h.body = body
}

func (h *History) GetBody() string {
	return h.body
}

type Doc struct {
	title string
	body  string
}

func NewDoc(title string, body string) *Doc {
	return &Doc{title: title, body: body}
}

func (d *Doc) SetTitle(title string) {
	d.title = title
}

func (d *Doc) GetTitle() string {
	return d.title
}

func (d *Doc) SetBody(body string) {
	d.body = body
}

func (d *Doc) GetBody() string {
	return d.body
}

func (d *Doc) CreateHistory() *History {
	return NewHistory(d.body)
}

func (d *Doc) RestoreHistory(history *History) {
	d.body = history.GetBody()
}

type Editor struct {
	doc       *Doc
	histories []*History
	position  int
}

func NewEditor(doc *Doc) *Editor {
	return &Editor{
		doc:       doc,
		histories: make([]*History, 0),
		position:  -1,
	}
}

func (e *Editor) Append(text string) {
	e.backup()
	e.doc.SetBody(e.doc.body + text + "\n")

	fmt.Println("===> 插入，文件如下：")
	e.Show()
}

func (e *Editor) Save() {
	fmt.Println("===> 保存，文件如下：")
	e.Show()
}

func (e *Editor) Delete() {
	e.backup()
	fmt.Println("===> 删除操作，文档内容如下：")
	e.doc.SetBody("")
}

func (e *Editor) Show() {
	fmt.Println(e.doc.GetBody())
}

func (e *Editor) backup() {
	e.histories = append(e.histories, e.doc.CreateHistory())
	e.position++
}

func (e *Editor) Undo() {
	if e.position > 0 {
		e.doc.RestoreHistory(e.histories[e.position])
		fmt.Println("===> 撤销操作，文档内容如下：")
		e.Show()
		e.position--
	}
}
