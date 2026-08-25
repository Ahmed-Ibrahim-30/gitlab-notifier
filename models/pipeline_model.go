package models

type PipelineEvent struct {
	ObjectKind string `json:"object_kind"`
	User       struct {
		Name string `json:"name"`
	} `json:"user"`
	ObjectAttributes struct {
		Ref    string `json:"ref"`
		Status string `json:"status"`
	} `json:"object_attributes"`
	Project struct {
		Name              string `json:"name"`
		PathWithNamespace string `json:"path_with_namespace"`
	} `json:"project"`

	Commit struct {
		Message string `json:"message"`
		Title   string `json:"title"` // ✅ First line only
		Author  struct {
			Name string `json:"name"`
		} `json:"author"`
		URL string `json:"url"`
	} `json:"commit"`
}
