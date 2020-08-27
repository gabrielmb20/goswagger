/*
 * Books API
 *
 * This web service offers information on books
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen 
 */
package swagger

import (
	"encoding/json"
	"net/http"
	"path"
	"fmt"
	"log"
)

var books = []Book{
	Book{BookId: "Book1", Title: "Book1", Edition: "9th",
        Copyright: "2012", Pages: "976",
        AuthorId: "1", PublisherId: "1"},

	Book{BookId: "Book2", Title: "Book2", Edition: "9th",
        Copyright: "2012", Pages: "976",
        AuthorId: "2", PublisherId: "2"},

}

var authors = []Author{
	Author{AuthorId: "1", BookId: "1", Name: "Gabriel Martínez", Nationality: "Costa Rica",
		Birth: "1990", Genere: "Male"},
	Author{AuthorId: "2", BookId: "2", Name: "José Barboza", Nationality: "Costa Rica",
		Birth: "1990", Genere: "Male"},
}

var publishers = []Publisher{
	Publisher{PublisherId: "1", Name: "Yensie", Country: "Inglaterra", Founded: "Costa Rica",
		Genere: "First"},
	Publisher{PublisherId: "2", Name: "Tatiana", Country: "Italia", Founded: "Costa Rica",
		Genere: "Second"},
}

func findBook(x string) int {
	for i, book := range books {
		if x == book.BookId {
			return i
		}
	}
	return -1
}

func findAuthor(x string) int {
	for i, author := range authors {
		if x == author.AuthorId {
			return i
		}
	}
	return -1
}

func findPublisher(x string) int {
	for i, publisher := range publishers {
		if x == publisher.PublisherId {
			return i
		}
	}
	return -1
}

func AuthorsAuthorIdBooksGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findAuthor(id)
	if i == -1 {
		//return
		fmt.Println("Id Invalido")
	}
	authors = append(authors[:i], authors[i+1:]...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Obtiene el Author por ID
func AuthorsAuthorIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findAuthor(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(authors[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdPut(w http.ResponseWriter, r *http.Request) {
	log.Printf("AUTHORS: ", authors)
	id := path.Base(r.URL.Path)
        i := findAuthor(id)
        if i == -1 {
                //return
                fmt.Println("Id Invalido")
        }
        authors = append(authors[:i], authors[i+1:]...)

	oldAuthor := bytes.IndexAny(authors, i)
	log.Printf("TEST",oldAuthor)

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	updateAuthor := Author{}
	json.Unmarshal(body, &updateAuthor)


	for index, item := range books {
		if item.AuthorId == id {
			authors = append(authors[:index], authors[index+1:]...)
			if updateAuthor.BookId != "" {
				oldAuthor.BookId = updateAuthor.BookId
			}
			if updateAuthor.Name != "" {
				oldAuthor.Name = updateAuthor.Name
			}
			if updateAuthor.Nationality != "" {
				oldAuthor.Nationality = updateAuthor.Nationality
			}
			if updateAuthor.Birth != "" {
				oldAuthor.Birth = updateAuthor.Birth
			}
			if updateAuthor.Genere != "" {
				oldAuthor.Genere = updateAuthor.Genere
			}
			//log.Println(oldAuthor)
			authors = append(authors, oldAuthor)
			json.NewEncoder(w).Encode(authors)
		}
	}
	return
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}


func BooksBookIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBook(id)
	if i == -1 {
		//return
		fmt.Println("Id Invalido")
	}
	books = append(books[:i], books[i+1:]...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBook(id)
	if i == -1 {
		return
	}
	dataJson, _ := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdPublishersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdPut(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
        i := findBook(id)
        if i == -1 {
                //return
                fmt.Println("Id Invalido")
        }
        books = append(books[:i], books[i+1:]...)

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	updateBook := Book{}
	json.Unmarshal(body, &updateBook)

	updateBook.BookId = id
	books = append(books, updateBook)
	json.NewEncoder(w).Encode(books)

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

func PublishersPublisherIdBooksGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findPublisher(id)
	if i == -1 {
		//return
		fmt.Println("Id Invalido")
	}
	publishers = append(publishers[:i], publishers[i+1:]...)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Obtiene el Publisher por ID
func PublishersPublisherIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findPublisher(id)
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
        i := findPublisher(id)
        if i == -1 {
                //return
                fmt.Println("Id Invalido")
        }
        publishers = append(publishers[:i], publishers[i+1:]...)

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	updatePublisher := Publisher{}
	json.Unmarshal(body, &updatePublisher)

	updatePublisher.PublisherId = id
	publishers = append(publishers, updatePublisher)
	json.NewEncoder(w).Encode(publishers)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
