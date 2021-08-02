package meta

import (
	"fmt"
	"testing"
	"time"
)

var m Meta
var Neeson struct {
	tableName string    `json:"neeson_table" `
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Age       int       `json:"age,omitempty"`
	Birthday  time.Time `json:"birthday"`
	CTime     time.Time `json:"c_time"`
}

func init() {
	m.Init(&Neeson)
}

func TestMeta_Tag(t *testing.T) {
	fmt.Println(m.TagName(&Neeson.CTime))
}
