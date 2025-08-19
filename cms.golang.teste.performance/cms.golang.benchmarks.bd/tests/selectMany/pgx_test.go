package tests

// 	for i := 0; i < b.N; i++ {
// 		model := make([]models.PersonModel, len(rows))
// 		rows, err := db.Query(ctx, utils.DBSelectAll)
// 		for j := 0; rows.Next() && j < len(model); j++ {
// 			err = rows.Scan(&model[j].ID, &model[j].Name, &model[j].CreatedAt)
// 			rows.Close()
// 		}
// 	}
