package pkg

type Meta struct {
	Total       int64  `json:"total"`
	TotalPages  int    `json:"total_pages"`
	PerPage     int    `json:"per_page,omitempty;query:per_page"`
	CurrentPage int    `json:"current_page,omitempty;query:page"`
	LastPage    int    `json:"last_page,omitempty"`
	FistPage    int    `json:"first_page,omitempty"`
	Search      string `json:"search,omitempty;query:search"`
	Sort        string `json:"sort,omitempty;query:sort"`
	Order       string `json:"order,omitempty;query:order"`
}

type Pagination struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func (m *Meta) GetOffset() int {
	return (m.GetCurrentPage() - 1) * m.GetPerPage()
}

func (m *Meta) GetPerPage() int {
	if m.PerPage == 0 || m.PerPage < 0 {
		return 10
	} else if m.PerPage > 100 {
		return 100
	}

	return m.PerPage
}

func (m *Meta) GetCurrentPage() int {
	if m.CurrentPage == 0 || m.CurrentPage < 0 {
		return 1
	}
	return m.CurrentPage
}

func (m *Meta) GetSort() string {
	if m.Sort == "" {
		return "id" + m.GetOrder()
	}
	return m.Sort
}

func (m *Meta) GetOrder() string {
	if m.Order == "" {
		return "ASC"
	}
	return m.Order
}

func (m *Meta) GetSearch() string {
	return m.Search
}

func (p *Pagination) SetMeta(meta Meta) {
	p.Meta = meta
}

func (p *Pagination) SetData(data interface{}) {
	p.Data = data
}
