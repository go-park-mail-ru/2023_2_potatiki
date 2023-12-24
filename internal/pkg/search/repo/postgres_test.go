package repo

//func TestSearchRepo_ReadProductsByName(t *testing.T) {
//	tests := []struct {
//		name       string
//		mockRepoFn func(pool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows)
//		err        error
//		columns    []string
//	}{
//		{
//			name: "SuccessfullReadProductsByName",
//			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
//				mockPool.EXPECT().Query(
//					gomock.Any(),
//					getProductsByName,
//					gomock.Any(),
//				).Return(pgxRows, nil)
//			},
//			columns: []string{
//				"id",
//				"name",
//				"description",
//				"price",
//				"imgsrc",
//				"rating",
//				"category_id",
//				"category_name",
//				"count",
//			},
//			err: nil,
//		},
//		{
//			name: "UnSuccessfullReadProductsByName",
//			mockRepoFn: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows) {
//				mockPool.EXPECT().Query(
//					gomock.Any(),
//					getProductsByName,
//					gomock.Any(),
//				).Return(pgxRows, pgx.ErrNoRows)
//			},
//			columns: []string{
//				"id",
//				"name",
//				"description",
//				"price",
//				"imgsrc",
//				"rating",
//				"category_id",
//				"category_name",
//				"count",
//			},
//			err: ErrProductNotFound,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ctr := gomock.NewController(t)
//			mockPool := pgxpoolmock.NewMockPgxPool(ctr)
//			defer ctr.Finish()
//
//			pgxRows := pgxpoolmock.NewRows(tt.columns).AddRow(
//				uuid.UUID{},
//				"",
//				"",
//				int64(0),
//				"",
//				float32(0),
//				int64(0),
//				"",
//				int64(0),
//			).ToPgxRows()
//			tt.mockRepoFn(mockPool, pgxRows)
//
//			repo := NewSearchRepo(mockPool)
//			_, err := repo.ReadProductsByName(context.Background(), "")
//
//			assert.Equal(t, tt.err, err)
//		})
//	}
//}
