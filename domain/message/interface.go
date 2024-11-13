package message

type Service interface {
	Create(msg Message) error
	Delete(id int) error
	Get(id int) (Message, error)
	GetAll() ([]Message, error)
}
