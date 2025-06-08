package model

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	ID   int64  `db:"groups_id" json:"groupID"`
	Name string `db:"groupsname" json:"groupName"`
}

func (g *Group) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", value)
	}
	return json.Unmarshal(b, g)
}
