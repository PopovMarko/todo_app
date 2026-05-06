package domain

import (
	"time"
)

type Task struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	createdat time.Time,
	completedat *time.Time,
	authoruserid int,
) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		CreatedAt:    createdat,
		CompletedAt:  completedat,
		AuthorUserID: authoruserid,
	}
}
func NewUninitializeTask(
	title string,
	description *string,
	authorUserID int,
) Task {
	return Task{
		ID:           UninitializedID,
		Version:      UninitializedVersion,
		Title:        title,
		Description:  description,
		Completed:    false,
		CreatedAt:    time.Now().UTC(),
		CompletedAt:  nil,
		AuthorUserID: authorUserID,
	}
}

func (t *Task) Validate() error {
	return nil
}
