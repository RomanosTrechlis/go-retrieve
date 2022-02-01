package registry

//go:generate stringer -type=Protocol
type Protocol int

const (
	FTP Protocol = iota
	HTTP
	LOCAL
)

//var Protocols = []string{"ftp", "http", "local"}

type Registry struct {
	Name      string               `json:"name"`
	Protocol  Protocol             `json:"protocol"`
	Templates []RegisteredTemplate `json:"templates"`
}

type RegisteredTemplate struct {
	Name         string   `json:"name"`
	Url          string   `json:"url"`
	Dependencies []string `json:"dependencies"`
	Commands     string   `json:"commands"`
	Files        []string `json:"files"`
}
