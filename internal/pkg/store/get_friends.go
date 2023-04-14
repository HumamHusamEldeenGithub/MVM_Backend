package store

func (repository *MVMRepository) GetFriends(userID string) ([]string, error) {
	// filter := bson.D{{Key: "id", Value: userID}}

	// var user model.User

	// if err := userDB.FindOne(repository.ctx, filter, opt).Decode(&user); err != nil {
	// 	return nil, err
	// }
	return []string{}, nil
}
