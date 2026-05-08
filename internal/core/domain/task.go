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

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (p *TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("title can't patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf(
			"completed can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate task patch: %w", err)
	}

	tempT := *t

	if patch.Title.Set {
		tempT.Title = *patch.Title.Value
	}
	if patch.Description.Set {
		tempT.Description = patch.Description.Value
	}
	if patch.Completed.Set {
		tempT.Completed = *patch.Completed.Value
		if tempT.Completed {
			now := time.Now().UTC()
			tempT.CompletedAt = &now
		} else {
			tempT.CompletedAt = nil
		}
	}

	if err := tempT.Validate(); err != nil {
		return fmt.Errorf("valivate patched task: %w", err)
	}
	*t = tempT
	return nil
}
