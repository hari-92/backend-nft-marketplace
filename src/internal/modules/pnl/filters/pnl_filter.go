package pnl_filters

type PnlFilter struct {
	UserId       string   `json:"user_id"`
	PairId       string   `json:"pair_id"`
	InUserIds    []string `json:"in_user_ids"`
	InPairIds    []string `json:"in_pair_ids"`
	NotInUserIds []string `json:"not_in_user_ids"`
	NotInPairIds []string `json:"not_in_pair_ids"`
}
