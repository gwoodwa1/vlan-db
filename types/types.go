package types


type Vlan struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ReservedVlans struct {
	IDs []int `json:"ids"`
}

var Reserved = ReservedVlans{
	IDs: []int{1, 2, 3, 4094},
}
