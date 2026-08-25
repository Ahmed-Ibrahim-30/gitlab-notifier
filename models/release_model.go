package models

type ReleaseEvent struct {
	ObjectKind string `json:"object_kind"`
	Action     string `json:"action"` // create, update, delete
	Tag        string `json:"tag"`
	CreatedAt  string `json:"created_at"`
	Project    struct {
		Name              string `json:"name"`
		PathWithNamespace string `json:"path_with_namespace"`
	} `json:"project"`
	Commit struct {
		Title  string `json:"title"`
		Author struct {
			Name string `json:"name"`
		} `json:"author"`
	} `json:"commit"`
}
