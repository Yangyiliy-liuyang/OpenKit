package gox

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Email    string
	Password MaskedSting
	Phone    MaskedPhone
}

type MaskedPhone string

// 序列化时加密字段
func (m MaskedPhone) MarshalJSON() ([]byte, error) {
	str := string(m)
	if len(str) == 11 {
		str = str[0:3] + "****" + str[7:]
	}
	return []byte(`"` + str + `"`), nil
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
		Phone:    "13812345678",
	}
	val, err := json.Marshal(u)
	assert.NoError(t, err)
	t.Log(string(val))
	// {"Email":"envkt@example.com","Password":{}}
	// {"Email":"envkt@example.com","Password":"****"}
	// {"Email":"envkt@example.com","Password":"****","Phone":"138****5678"}
}
