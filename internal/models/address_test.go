package models

//func TestAddress_MarshalJSON(t *testing.T) {
//	type fields struct {
//		Id        uuid.UUID
//		ProfileId uuid.UUID
//		City      string
//		Street    string
//		House     string
//		Flat      string
//		IsCurrent bool
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    string
//		wantErr bool
//	}{
//		{"First Test", fields{House: "SSSSS"}, "SSSS", false},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := Address{
//				Id:        tt.fields.Id,
//				ProfileId: tt.fields.ProfileId,
//				City:      tt.fields.City,
//				Street:    tt.fields.Street,
//				House:     tt.fields.House,
//				Flat:      tt.fields.Flat,
//				IsCurrent: tt.fields.IsCurrent,
//			}
//			got, err := u.MarshalJSON()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			assert.Equal(t, tt.want, string(got))
//		})
//	}
//}
