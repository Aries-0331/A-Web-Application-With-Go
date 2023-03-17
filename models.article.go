package main

import (
	"errors"
	"testing"
)

type article struct {
	ID      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1's content"},
	article{ID: 2, Title: "Article 2", Content: "Article 2's content"},
}

func getAllArticles() []article {
	return articleList
}

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
