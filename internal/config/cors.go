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
		"http://89.111.173.104:3000",
		"http://89.111.173.104",
		"https://89.111.173.104",
		"http://north-tribune.ru:3000",
		"http://north-tribune.ru",
		"https://north-tribune.ru",
	}
	ALLOWED_METHODS = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	}
)
