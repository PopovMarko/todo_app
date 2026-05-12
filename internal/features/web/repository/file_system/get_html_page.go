package web_repository

import (
	"fmt"
	"os"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

func (r *WebRepository) GetHTMLPage(htmlFilePath string) ([]byte, error) {
	file, err := os.ReadFile(htmlFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(
				"File %s not exist: %v: %w",
				htmlFilePath, err, core_errors.ErrNotFound,
			)
		}
		return nil, fmt.Errorf("File %s: %w", htmlFilePath, err)
	}

	return file, nil
}
