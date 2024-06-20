package config

import "net/http"

var (
	ALLOWED_HEADERS = []string{
		"Accept",
		"Accept-Language",
		"Content-Type",
	}
	ALLOWED_ORIGINS = []string{
		"http://localhost",
		"http://localhost:3000",
		"http://localhost:9001",
		"http://127.0.0.1",
		"http://127.0.0.1:3000",
		"http://socio-project.ru:9001",
	}
	ALLOWED_METHODS = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	}
)
