package externalservice

import (
	"errors"
)

// Post is the data structure representing the data sent and received from the
// external service
type Post struct {
	ID int `json:"id"` // the primary key

	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// Client represents the client interface to the external service
type Client interface {
	GET(id int) (*Post, error)
	POST(id int, post *Post) (*Post, error)
}

type ClientService struct {
	POSTCallCounter int
	GETCallCounter  int
}

func (c ClientService) GET(id int) (*Post, error) {
	c.GETCallCounter++
	return nil, errors.New("Bad Request")
}

func (c ClientService) POST(id int, post *Post) (*Post, error) {
	c.POSTCallCounter++
	return &Post{
		ID:          id,
		Title:       "Hello World!",
		Description: "Lorem Ipsum Dolor Sit Amen.",
	}, nil
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// func (e Error) Error() string {
// 	return fmt.Sprintf(`{"code": %d, "message": "%s"}`, e.Code, e.Message)
// }

// func (e Error) ToJson() ([]byte, error) {
// 	errJson, err := json.Marshal(e)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return errJson, nil
// }
