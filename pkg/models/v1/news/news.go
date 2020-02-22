package news

import (
	"encoding/json"
	"time"

	"github.com/bungysheep/news-api/pkg/models/v1/basemodel"
)

// News type
type News struct {
	basemodel.BaseModel
	ID      int64     `json:"id"`
	Author  string    `json:"author" mandatory:"true" max_length:"64"`
	Body    string    `json:"body" mandatory:"true"`
	Created time.Time `json:"created"`
}

// NewNews - Creates news
func NewNews() *News {
	return &News{}
}

// GetID - Returns news id
func (news *News) GetID() int64 {
	return news.ID
}

// GetAuthor - Returns author
func (news *News) GetAuthor() string {
	return news.Author
}

// GetBody - Returns body
func (news *News) GetBody() string {
	return news.Body
}

// GetCreated - Returns created
func (news *News) GetCreated() time.Time {
	return news.Created
}

// MarshalBinary - Marshal binary
func (news *News) MarshalBinary() ([]byte, error) {
	return json.Marshal(news)
}

// UnmarshalBinary - Unmarshal binary
func (news *News) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &news); err != nil {
		return err
	}

	return nil
}

// DoValidate - Validates news
func (news *News) DoValidate() (bool, string) {
	return news.DoValidateBase(*news)
}
