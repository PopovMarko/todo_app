package web_service

import ()

type WebRepository interface {
	GetHTMLPage(htmlFilePath string) ([]byte, error)
}

type WebService struct {
	webRepository WebRepository
}

func NewWebService(webRepository WebRepository) *WebService {
	return &WebService{
		webRepository: webRepository,
	}
}
