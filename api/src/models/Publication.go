package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("The title is mandatory and can't be blank")
	}

	if publication.Content == "" {
		return errors.New("The content is mandatory and can't be blank")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
