package user_filters

type UserFilter struct {
	ID       uint64   `json:"id"`
	Email    string   `json:"email"`
	InIDs    []uint64 `json:"in_ids"`
	NotInIDs []uint64 `json:"not_in_ids"`
}
