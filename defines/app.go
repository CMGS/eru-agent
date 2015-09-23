package defines

type Meta struct {
	ID         string
	Pid        int
	Name       string
	EntryPoint string
	Ident      string
	Extend     map[string]interface{}
}
