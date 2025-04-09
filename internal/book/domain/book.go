// api_hexagonal_go/internal/book/domain/book.go
package domain

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   int    `json:"year"`
}