package models

type TagEvent struct {
	ObjectKind string `json:"object_kind"`
	UserName   string `json:"user_name"`
	Ref        string `json:"ref"`
	Project    struct {
		Name              string `json:"name"`
		PathWithNamespace string `json:"path_with_namespace"`
	} `json:"project"`
}
