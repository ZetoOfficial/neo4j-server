package models

type Node struct {
	ID         int64   `json:"id"`
	Label      string  `json:"label"`
	Name       *string `json:"name,omitempty"`
	ScreenName *string `json:"screen_name,omitempty"`
	Sex        *int    `json:"sex,omitempty"`
	City       *string `json:"city,omitempty"`
}

type Relationship struct {
	Type      string `json:"type"`
	EndNodeID int64  `json:"end_node_id"`
}
