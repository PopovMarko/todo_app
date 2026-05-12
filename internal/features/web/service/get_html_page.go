package web_service

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *WebService) GetHTMLPage() ([]byte, error) {
	htmlFilePath := filepath.Join(os.Getenv("PROJECT_ROOT"), "/web/index.html")
	htmlFile, err := s.webRepository.GetHTMLPage(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("get html file %s from repository: %w", htmlFilePath, err)
	}

	return htmlFile, nil

}
