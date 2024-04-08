package gox

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Email    string
	Password MaskedSting
}

type MaskedSting string

// 序列化时加密字段
func (m MaskedSting) MarshalJSON() ([]byte, error) {
	//return []byte("{}"), nil
	return []byte(`"****"`), nil
}

func TestMasked(t *testing.T) {
	u := User{
		Email:    "envkt@example.com",
		Password: "123456",
	}
	val, err := json.Marshal(u)
	assert.NoError(t, err)
	t.Log(string(val))
	// {"Email":"envkt@example.com","Password":{}}
	// {"Email":"envkt@example.com","Password":"****"}
}
