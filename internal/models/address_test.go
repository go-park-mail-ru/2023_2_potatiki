package models

import (
	"fmt"
	"github.com/mailru/easyjson"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestAddress_Sanitize(t *testing.T) {
	a := Address{
		Id:        uuid.NewV4(),
		ProfileId: uuid.NewV4(),
		City:      uuid.NewV4().String(),
		Street:    uuid.NewV4().String(),
		House:     uuid.NewV4().String(),
		Flat:      uuid.NewV4().String(),
		IsCurrent: false,
	}
	b, err := easyjson.Marshal(&a)
	fmt.Println(string(b), err)

	a2 := AddressSlice{{
		Id:        uuid.NewV4(),
		ProfileId: uuid.NewV4(),
		City:      uuid.NewV4().String(),
		Street:    uuid.NewV4().String(),
		House:     uuid.NewV4().String(),
		Flat:      uuid.NewV4().String(),
		IsCurrent: false,
	}}
	b, err = easyjson.Marshal(&a2)
	fmt.Println(string(b), err)

	a3 := Product{
		Id:            uuid.NewV4(),
		Name:          uuid.NewV4().String(),
		Description:   uuid.NewV4().String(),
		Price:         0,
		ImgSrc:        uuid.NewV4().String(),
		Rating:        0,
		CountComments: 0,
		Category: Category{
			Id:     0,
			Name:   uuid.NewV4().String(),
			Parent: 0,
		},
	}
	b, err = easyjson.Marshal(&a3)
	fmt.Println(string(b), err)
}
