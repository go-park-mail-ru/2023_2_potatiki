package models

import (
	"log"
	"testing"
)

func TestAddress_MarshalJSON(t *testing.T) {
	u := Address{}
	got, _ := u.MarshalJSON()
	log.Println(string(got))
}
