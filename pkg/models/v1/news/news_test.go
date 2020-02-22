package news

import (
	"testing"
	"time"
)

func TestCreateNews(t *testing.T) {
	news := NewNews()

	if news == nil {
		t.Fatalf("Expect news is not nil")
	}

	timeNow := time.Now()

	bodyNews := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent sapien mauris, viverra vel egestas sit amet, mattis sed libero. Etiam sed diam et felis venenatis porta. Aliquam semper sem eget lectus tristique vulputate. Aliquam euismod nisi at justo congue tempus. Sed faucibus non sapien sit amet condimentum. Sed rutrum ligula odio, sit amet bibendum diam sagittis a. Phasellus sit amet risus tellus. Ut elementum venenatis arcu vitae vulputate. Nulla venenatis, magna et luctus gravida, mi lorem molestie ipsum, sed malesuada erat justo id nunc. Integer sodales sem ac ipsum dapibus lobortis. Vivamus auctor felis non magna ultricies, laoreet posuere tellus posuere."

	news.ID = 1
	news.Author = "Author A"
	news.Body = bodyNews
	news.Created = timeNow

	ok, msg := news.DoValidate()
	if !ok {
		t.Fatalf("Expect news is valid, but got %s", msg)
	}

	if news.GetID() != 1 {
		t.Errorf("Expect news id %v, but got %v", 1, news.GetID())
	}

	if news.GetAuthor() != "Author A" {
		t.Errorf("Expect author %v, but got %v", "Author A", news.GetAuthor())
	}

	if news.GetBody() != bodyNews {
		t.Errorf("Expect body %v, but got %v", bodyNews, news.GetBody())
	}

	if news.GetCreated() != timeNow {
		t.Errorf("Expect craeted %v, but got %v", timeNow, news.GetCreated())
	}
}
