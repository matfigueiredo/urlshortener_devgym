package domain

type URL struct {
	Original string
	Code     string
}

type URLRepository interface {
	Save(url URL) error // Aqui, removemos o retorno (domain.URL, error) para apenas error
	FindByCode(code string) (URL, error)
}
