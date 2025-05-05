package constant

type ServiceID string

const (
	User ServiceID = "user"
)

func (s ServiceID) String() string {
	return string(s)
}
