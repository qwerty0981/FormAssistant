package form

type Form interface {
	Endpoints() map[string]string
	AddEndpoint(string, string) error
	Post(map[string]string) error
}
