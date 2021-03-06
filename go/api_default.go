/*
 * Books API
 *
 * This web service offers information on books
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

var books = []Book{
	Book{BookId: "1", PublisherId: "1", Title: "Libro 1",
		Copyright: "2012", Edition: "3th", Pages: "9726"},
	Book{BookId: "2", PublisherId: "1", Title: "Libro 2",
		Copyright: "2002", Edition: "4th", Pages: "1500"},
	Book{BookId: "3", PublisherId: "1", Title: "Libro 3",
		Copyright: "2019", Edition: "7th", Pages: "15030"},
	Book{BookId: "4", PublisherId: "2", Title: "Libro 4",
		Copyright: "2018", Edition: "9th", Pages: "722"},
}

var authors = []Author{
	Author{AuthorId: "1", BookId: "3", Name: "Jose", Nationality: "Costa Rica",
		Birth: "1992", Genere: "First"},
	Author{AuthorId: "2", BookId: "2", Name: "Gabriel", Nationality: "Argentina",
		Birth: "1991", Genere: "Second"},
	Author{AuthorId: "3", BookId: "2", Name: "Carolina", Nationality: "Panama",
		Birth: "1990", Genere: "Third"},
	Author{AuthorId: "4", BookId: "1", Name: "Maritza", Nationality: "Francia",
		Birth: "1988", Genere: "Fourth"},
	Author{AuthorId: "5", BookId: "4", Name: "Josue", Nationality: "EEUU",
		Birth: "1987", Genere: "Fiveth"},
}

var publishers = []Publisher{
	Publisher{PublisherId: "1", Name: "Michael", Country: "Inglaterra", Founded: "Inglaterra",
		Genere: "First"},
	Publisher{PublisherId: "2", Name: "Tatiana", Country: "Italia", Founded: "Italia",
		Genere: "Second"},
	Publisher{PublisherId: "3", Name: "Christian", Country: "Colombia", Founded: "Colombia",
		Genere: "Third"},
	Publisher{PublisherId: "4", Name: "Maria", Country: "China", Founded: "China",
		Genere: "Sixth"},
	Publisher{PublisherId: "5", Name: "Luis", Country: "Nicaragua", Founded: "Nicaragua",
		Genere: "Seventh"},
}

func findBook(x string) int {
	for _, book := range books {
		if x == book.BookId {
			s2, _ := strconv.Atoi(book.BookId)
			return s2
		}
	}
	return -1
}

func findBookPos(x string) int {
	for i, book := range books {
		if x == book.BookId {
			return i
		}
	}
	return -1
}

func findAuthorPos(x string) int {
	for i, author := range authors {
		if x == author.AuthorId {
			return i
		}
	}
	return -1
}

func findAuthor(x string) int {
	for _, author := range authors {
		if x == author.AuthorId {
			s2, _ := strconv.Atoi(author.AuthorId)
			return s2
		}
	}
	return -1
}

//   /books/1/authors/
func findAuthorIdbyBook(x string) int {
	for _, author := range authors {
		if x == author.BookId {
			s2, _ := strconv.Atoi(author.AuthorId)
			return s2
		}
	}
	return -1
}

//   /authors/1/books/
func findBookIdbyAuthor(x string) int {
	for _, author := range authors {
		if x == author.AuthorId {
			s2, _ := strconv.Atoi(author.BookId)
			return s2
		}
	}
	return -1
}

func findPublisher(x string) int {
	for _, publishers := range publishers {
		if x == publishers.PublisherId {
			s2, _ := strconv.Atoi(publishers.PublisherId)
			return s2
		}
	}
	return -1
}

func findPublisherPos(x string) int {
	for i, publishers := range publishers {
		if x == publishers.PublisherId {
			return i
		}
	}
	return -1
}

//   /books/1/publishers/
func findPublisherIdbyBook(x string) int {
	for _, book := range books {
		if x == book.BookId {
			s2, _ := strconv.Atoi(book.PublisherId)
			return s2
		}
	}
	return -1
}

//   /publishers/1/books/
func findBookIdbyPublisher(x string) int {
	for _, book := range books {
		if x == book.PublisherId {
			s2, _ := strconv.Atoi(book.BookId)
			return s2
		}
	}
	return -1
}

//  /authors/1/books/
func AuthorsAuthorIdBooksGet(w http.ResponseWriter, r *http.Request) {
	idTemp := path.Dir(r.URL.Path)
	idTemp2 := path.Dir(idTemp)
	id := path.Base(idTemp2)
	i := findAuthor(id)
	if i != -1 {
		s := strconv.Itoa(i)
		idBook := findBookIdbyAuthor(s)
		if idBook != -1 {
			p := strconv.Itoa(idBook)
			bookId := findBook(p)
			if bookId != -1 {
				p2 := strconv.Itoa(bookId)
				bookId2 := findBookPos(p2)
				dataJson, _ := json.Marshal(books[bookId2])
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.Write(dataJson)
				w.WriteHeader(http.StatusOK)
			} else {
				fmt.Println("findBook invalido")
			}
		} else {
			fmt.Println("idBook invalido")
		}
	} else {
		fmt.Println("idAuthor invalido")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findAuthorPos(id)
	if i == -1 {
		fmt.Println("Id Invalido")
	}
	authors = append(authors[:i], authors[i+1:]...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findAuthorPos(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(authors[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdPut(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	for index, item := range authors {
		if item.AuthorId == id {
			authors = append(authors[:index], authors[index+1:]...)
			var author Author
			_ = json.NewDecoder(r.Body).Decode(&author)
			author.AuthorId = id
			authors = append(authors, author)
			json.NewEncoder(w).Encode(&author)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func AuthorsPost(w http.ResponseWriter, r *http.Request) {
	var author Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authors = append(authors, author)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdAuthorsGet(w http.ResponseWriter, r *http.Request) {
	idTemp := path.Dir(r.URL.Path)
	idTemp2 := path.Dir(idTemp)
	id := path.Base(idTemp2)
	i := findBook(id)
	if i != -1 {
		s := strconv.Itoa(i)

		idBook := findAuthorIdbyBook(s)

		for i, author := range authors {
			if s == author.BookId {
				fmt.Println("Author-AuthorId: ", author.AuthorId)
				p := strconv.Itoa(idBook)
				bookId := findAuthor(p)
				if bookId != -1 {

					dataJson, _ := json.Marshal(authors[i])
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.Write(dataJson)
					w.WriteHeader(http.StatusOK)
				} else {
					fmt.Println("findBook invalido")
				}
			}
		}
	} else {
		fmt.Println("idAuthor invalido")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBookPos(id)
	if i == -1 {
		fmt.Println("Id Invalido")
	}
	books = append(books[:i], books[i+1:]...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBookPos(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

//  /books/1/publishers/
func BooksBookIdPublishersGet(w http.ResponseWriter, r *http.Request) {
	idTemp := path.Dir(r.URL.Path)
	idTemp2 := path.Dir(idTemp)
	id := path.Base(idTemp2)
	i := findBook(id)
	if i != -1 {
		s := strconv.Itoa(i)
		idBook := findPublisherIdbyBook(s)
		if idBook != -1 {
			p := strconv.Itoa(idBook)
			bookId := findBook(p)
			if bookId != -1 {
				p2 := strconv.Itoa(bookId)
				bookId2 := findBookPos(p2)
				dataJson, _ := json.Marshal(publishers[bookId2])
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.Write(dataJson)
				w.WriteHeader(http.StatusOK)
			} else {
				fmt.Println("findBook invalido")
			}
		} else {
			fmt.Println("idBook invalido")
		}
	} else {
		fmt.Println("idAuthor invalido")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdPut(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	for index, item := range books {
		if item.BookId == id {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.BookId = id
			books = append(books, book)
			json.NewEncoder(w).Encode(&book)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksPost(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PublishersPost(w http.ResponseWriter, r *http.Request) {
	var publisher Publisher
	err := json.NewDecoder(r.Body).Decode(&publisher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	publishers = append(publishers, publisher)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

//  /publishers/1/books/
func PublishersPublisherIdBooksGet(w http.ResponseWriter, r *http.Request) {
	idTemp := path.Dir(r.URL.Path)
	idTemp2 := path.Dir(idTemp)
	id := path.Base(idTemp2)
	i := findPublisher(id)
	if i != -1 {
		s := strconv.Itoa(i)
		idBook := findBookIdbyPublisher(s)
		for i, book := range books {
			if s == book.PublisherId {
				fmt.Println("Author-AuthorId: ", book.PublisherId)
				p := strconv.Itoa(idBook)
				bookId := findBook(p)
				if bookId != -1 {
					dataJson, _ := json.Marshal(books[i])
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.Write(dataJson)
					w.WriteHeader(http.StatusOK)
				} else {
					fmt.Println("findBook invalido")
				}
			}
		}
	} else {
		fmt.Println("idAuthor invalido")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findPublisherPos(id)
	if i == -1 {
		fmt.Println("Id Invalido")
	}
	publishers = append(publishers[:i], publishers[i+1:]...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findPublisherPos(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(publishers[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdPut(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	for index, item := range publishers {
		if item.PublisherId == id {
			publishers = append(publishers[:index], publishers[index+1:]...)
			var publisher Publisher
			_ = json.NewDecoder(r.Body).Decode(&publisher)
			publisher.PublisherId = id
			publishers = append(publishers, publisher)
			json.NewEncoder(w).Encode(&publisher)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
