package channel

type Service interface {
	Create(channel Channel) error
	Delete(id int) error
	Get(id int) (Channel, error)
	GetAll() ([]Channel, error)
}
