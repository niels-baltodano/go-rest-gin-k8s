package bulletin

import "time"

type Getter interface {
	GetAll() []Bulletin
}

type Adder interface {
	Add(b Bulletin)
}
type Bulletin struct {
	Author    string    `json:"author" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
type Repo struct {
	Bulletins []Bulletin
}

func New() *Repo{
	return &Repo{
		Bulletins: []Bulletin{},
	}
}
func(r *Repo) Add(b Bulletin){
	r.Bulletins = append(r.Bulletins, b)
}

func (r *Repo) GetAll() []Bulletin {
	return r.Bulletins
}