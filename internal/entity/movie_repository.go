package entity

type MovieRepository interface {
	Save(movie *Movie) (string, error)
	Update(movie *Movie) error
	FindById(id string) (*Movie, error)
	FindAll() ([]Movie, error)
	Remove(id string) error
}
