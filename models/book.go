package models

import (
	"boot-rest-api/config"
	"boot-rest-api/entities"
	"log"
)

type Book struct{}

func (bm *Book) GetAll() ([]entities.Book, error) {
	db, err := config.GetDb()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		return nil, err
	}

	var books []entities.Book

	for rows.Next() {
		var book entities.Book
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Language, &book.TotalPages, &book.ImagePath, &book.WikipediaLink, &book.Country)

		books = append(books, book)
	}

	return books, nil
}

func (bm *Book) Search(q string) ([]entities.Book, error) {
	db, err := config.GetDb()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT * FROM books WHERE books.title like ?`, "%"+q+"%")

	if err != nil {
		return nil, err
	}

	var books []entities.Book

	for rows.Next() {
		var book entities.Book
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Language, &book.TotalPages, &book.ImagePath, &book.WikipediaLink, &book.Country)

		books = append(books, book)
	}

	return books, nil
}

func (bm *Book) Create(b *entities.Book) (entities.Book, error) {
	db, err := config.GetDb()
	if err != nil {
		return entities.Book{}, err
	}

	result, err := db.Exec(`
								INSERT INTO 
    							books ( title, author, language, total_pages, image_path, wikipedia_link, country) 
    							VALUES (?, ?, ?, ?, ?, ?, ?)
    							`, b.Title, b.Author, b.Language, b.TotalPages, b.ImagePath, b.WikipediaLink, b.Country)

	if err != nil {
		return entities.Book{}, err
	}

	id , err := result.LastInsertId()
	if err != nil {
		return entities.Book{}, err
	}

	// retrieve the created book
	book, err := bm.GetBookById(int(id))
	if err != nil {
		log.Fatal(err)
	}
	return book, nil
}

func (*Book) Delete(id int) (bool, error) {
	db, err := config.GetDb()
	if err != nil {
		return false, err
	}

	result, err := db.Exec(`DELETE FROM books WHERE id = ?`, id)

	if err != nil {
		return false, err
	}

	ra, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return ra > 0, nil
}

func (bm *Book) Update(b *entities.Book, bId int) (entities.Book, error) {
	db, err := config.GetDb()
	if err != nil {
		return entities.Book{}, err
	}

	_, err = db.Exec(`
								UPDATE books SET
									title = ?,
									author = ?,
									language = ?,
									total_pages = ?, 
								    image_path = ?, 
									wikipedia_link = ?,
								    country = ?
								WHERE id = ?
    							`, b.Title, b.Author, b.Language, b.TotalPages, b.ImagePath, b.WikipediaLink, b.Country, bId)

	if err != nil {
		return entities.Book{}, err
	}

	// retrieve the updated book
	book, err := bm.GetBookById(bId)
	if err != nil {
		log.Fatal(err)
	}
	return book, nil
}

func (bm *Book) GetBookById(id int) (entities.Book, error) {
	db, err := config.GetDb()
	if err != nil {
		return entities.Book{}, err
	}

	rows, err := db.Query(`SELECT * FROM books WHERE id = ?`, id)

	if err != nil {
		return entities.Book{}, err
	}

	var book entities.Book
	for rows.Next() {
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Language, &book.TotalPages, &book.ImagePath, &book.WikipediaLink, &book.Country)
	}

	return book, nil
}
