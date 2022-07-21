package models

var books []*Book

type Book struct {
	Id_book    int64 `json:"id"`
	Name_book  string `json:"name_book"`
	Genre_book string `json:"genre_book"`
}


func GetAllBuku() ([]*Book, error) {

	books = append(books, &Book{
		Id_book : 1,
		Name_book: "Panji Petualang",
		Genre_book: "Adventure",
	})

	return books, nil
}
