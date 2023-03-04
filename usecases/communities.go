package usecases

import "mckp/vesta/database"

type NewCommunity struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func SaveNewCommunity(newCommunity NewCommunity) (database.Community, error) {
	sqlQuery := `
		INSERT INTO
			communities (name, slug)
		VALUES
			($1, $2)
		RETURNING *
	`

	defaultValue := database.Community{}

	result, err := database.GenericDBCall[[]database.Community](
		sqlQuery,
		[]database.Community{},
		&database.Community{},
		newCommunity.Name,
		newCommunity.Slug,
	)

	if err != nil {
		return defaultValue, err
	}

	return result[0], nil
}

func AddUserToCommunity(userId string, communityId string) error {
	sqlQuery := `
		INSERT INTO
			user_communities (user_id, community_id)
		VALUES
			($1, $2)
	`

	_, err := database.GenericDBCall[[]database.UserCommunity](
		sqlQuery,
		[]database.UserCommunity{},
		&database.UserCommunity{},
		userId,
		communityId,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetUsersInCommunity(communityId string) ([]database.User, error) {
	sqlQuery := `
		SELECT
			users.*
		FROM
			user_communities
		JOIN users ON users._id = user_communities.user_id
		WHERE
			user_communities.community_id = $1
	`

	defaultValue := []database.User{}

	result, err := database.GenericDBCall[[]database.User](
		sqlQuery,
		[]database.User{},
		&database.DatabaseUser{},
		communityId,
	)

	if err != nil {
		return defaultValue, err
	}

	return result, nil

}

func DeleteUserFromCommunity(communityId string, userId string) error {
	sqlQuery := `
		DELETE FROM
			user_communities
		WHERE
			community_id = $1
			AND user_id = $2
	`

	_, err := database.GenericDBCall[[]database.Community](
		sqlQuery,
		[]database.Community{},
		&database.Community{},
		communityId,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetCommunities(limit int64, offset int64) ([]database.Community, error) {
	sqlQuery := `
		SELECT
			*
		FROM
			communities
		ORDER BY created_at DESC
		LIMIT $1
		OFFSET $2
	`

	defaultValue := []database.Community{}

	result, err := database.GenericDBCall[[]database.Community](
		sqlQuery,
		[]database.Community{},
		&database.Community{},
		limit,
		offset,
	)

	if err != nil {
		return defaultValue, err
	}

	return result, nil
}

func GetCommunityById(communityId string) (database.Community, error) {
	sqlQuery := `
		SELECT
			*
		FROM
			communities
		WHERE
			_id = $1
	`
	defaultValue := database.Community{}

	result, err := database.GenericDBCall[[]database.Community](
		sqlQuery,
		[]database.Community{},
		&database.Community{},
		communityId,
	)

	if err != nil {
		return defaultValue, err
	}

	return result[0], nil
}
