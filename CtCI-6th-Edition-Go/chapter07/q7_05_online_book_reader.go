package chapter07

import (
	"fmt"
	"sort"
)

type Page struct {
	Number  int
	Content string
}

func NewPage(number int, content string) *Page {
	if number < 1 {
		panic("page number must be positive")
	}
	return &Page{Number: number, Content: content}
}

type Book struct {
	Title string
	Pages []*Page
}

func NewBook(title string, pages []*Page) *Book {
	if title == "" {
		panic("title required")
	}
	copied := append([]*Page{}, pages...)
	sort.Slice(copied, func(i, j int) bool { return copied[i].Number < copied[j].Number })
	if len(copied) == 0 {
		panic("book must have pages")
	}
	return &Book{Title: title, Pages: copied}
}

func (b *Book) PageCount() int { return len(b.Pages) }

type BookReader struct {
	book  *Book
	index int
}

func NewBookReader() *BookReader { return &BookReader{} }

func (r *BookReader) Open(book *Book) {
	if book == nil {
		panic("book is nil")
	}
	r.book = book
	r.index = 0
}

func (r *BookReader) CurrentPage() *Page {
	if r.book == nil {
		panic("no book open")
	}
	return r.book.Pages[r.index]
}

func (r *BookReader) NextPage() *Page {
	if r.book == nil {
		panic("no book open")
	}
	if r.index+1 >= len(r.book.Pages) {
		return nil
	}
	r.index++
	return r.book.Pages[r.index]
}

func (r *BookReader) PrevPage() *Page {
	if r.book == nil {
		panic("no book open")
	}
	if r.index == 0 {
		return nil
	}
	r.index--
	return r.book.Pages[r.index]
}

func (r *BookReader) Goto(pageNumber int) *Page {
	if r.book == nil {
		panic("no book open")
	}
	for i, page := range r.book.Pages {
		if page.Number == pageNumber {
			r.index = i
			return page
		}
	}
	panic("page not found")
}

func RunQ705() {
	book := NewBook("Demo", []*Page{NewPage(1, "a"), NewPage(2, "b"), NewPage(3, "c")})
	reader := NewBookReader()
	reader.Open(book)
	fmt.Println(reader.CurrentPage().Content)
	reader.Goto(3)
	fmt.Println(reader.CurrentPage().Content)
}
