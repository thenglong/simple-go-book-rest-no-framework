package handlers

import (
	"boot-rest-api/entities"
	"boot-rest-api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct{}

func (bh *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bh.get(w, r)
	case "POST":
		bh.post(w, r)
	case "PUT":
		bh.put(w, r)
	case "DELETE":
		bh.delete(w, r)
	default:
		bh.error(w)
	}
}

func (bh *BookHandler) get(w http.ResponseWriter, r *http.Request) {
	bm := models.Book{}
	w.Header().Add("content-type", "application/json")
	id, ok := getIdParamFromRequest(r)

	// book with id
	if ok {
		book, err := bm.GetBookById(id)
		if err != nil {
			log.Fatal(err)
		}

		bs, _ := json.Marshal(book)
		w.Write(bs)
		return
	}

	// all books
	books, err := bm.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := json.Marshal(books)
	w.Write(bs)
}

func (bh *BookHandler) post(w http.ResponseWriter, r *http.Request) {
	bm := models.Book{}
	w.Header().Add("content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	// validate request body
	if err != nil {
		msg, _ := json.Marshal(map[string]string{"error": "body data is invalid"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msg)
		return
	}

	// validate content type
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		msg, _ := json.Marshal(map[string]string{"error": "only 'application/json' content is allowed"})
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(msg)
		return
	}

	var book entities.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Fatal(err)
	}

	books, err := bm.Create(&book)
	if err != nil {
		log.Fatal(err)
	}

	bs, _ := json.Marshal(books)
	w.WriteHeader(http.StatusCreated)
	w.Write(bs)
}

func (bh *BookHandler) put(w http.ResponseWriter, r *http.Request) {
	bm := models.Book{}
	w.Header().Add("content-type", "application/json")
	id, ok := getIdParamFromRequest(r)

	// book with id
	if ok {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var book entities.Book
		json.Unmarshal(body, &book)

		b, err := bm.Update(&book, id)
		if err != nil {
			log.Fatal(err)
		}

		bs, _ := json.Marshal(b)
		w.Write(bs)
		return
	}

	// response with error because cannot get the id
	bs, _ := json.Marshal(map[string]string{"error": "id is not passed"})
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(bs)
}

func (bh *BookHandler) delete(w http.ResponseWriter, r *http.Request) {
	bm := models.Book{}
	w.Header().Add("content-type", "application/json")
	id, ok := getIdParamFromRequest(r)

	// book with id
	if ok {

		isSuccess, err := bm.Delete(id)
		if err != nil {
			log.Fatal(err)
		}


		if isSuccess {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte{})
		}

		// delete fail
		msg, _ := json.Marshal(map[string]string{"error": "only 'application/json' content is allowed"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msg)
	}

	// response with error because cannot get the id
	bs, _ := json.Marshal(map[string]string{"error": "id is not passed"})
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(bs)
}

func (bh *BookHandler) error(w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]string{"error": "Invalid Request"})
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(msg)
}

func getIdParamFromRequest(r *http.Request) (int, bool) {
	parts := strings.Split(r.URL.String(), "/")

	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, false
	}
	return id, true
}
