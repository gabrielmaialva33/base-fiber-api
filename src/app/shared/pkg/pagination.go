package pkg

type Pagination struct {
	Total      int64       `json:"total"`
	Page       int         `json:"current_page,omitempty;query:page"`
	PerPage    int         `json:"per_page,omitempty;query:per_page"`
	Order      string      `json:"order,omitempty;query:order"`
	Search     string      `json:"search,omitempty;query:search"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPerPage()
}

func (p *Pagination) GetPerPage() int {
	if p.PerPage == 0 {
		return 10
	}
	return p.PerPage
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 || p.Page < 0 {
		return 1
	}
	return p.Page
}

func (p *Pagination) GetOrder() string {
	if p.Order == "" {
		return "id ASC"
	}
	return p.Order
}

func (p *Pagination) GetSearch() string {
	return p.Search
}
