package domain

import (
	"fmt"
	"time"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
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
func NewUninitializedTask(
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
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"title length must be between 1 and 100 char: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if t.Description != nil {
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"description must be between 1 and 1000 char: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}
	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf(
				"completed at can't be nil if completed = true: %w",
				core_errors.ErrInvalidArgument,
			)
		}
		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf(
				"completed at can't be before created at: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	} else if t.CompletedAt != nil {
		return fmt.Errorf(
			"completed at should be nil if completed = false: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
