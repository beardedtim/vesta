package usecases

import (
	"mckp/vesta/database"

	"golang.org/x/crypto/bcrypt"
)

const (
	saltCost = 10
)

/*
Hashes a plaintext password for storage. DO NOT LOG
*/
func hashUserPassword(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), saltCost)

	return string(bytes), err
}

/*
Given a plaintext password and the hashed password, returns if they match. DO NOT LOG
*/
func ValidatePasswordAgainstHash(plaintext, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))

	return err == nil
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SaveNewUser(newUser NewUser) (database.User, error) {
	sqlQuery := `
		INSERT INTO
			users (email, password)
		VALUES
			($1, $2)
		RETURNING *
	`
	hashedPassed, err := hashUserPassword(newUser.Password)
	defaultValue := database.User{}

	if err != nil {
		return defaultValue, err
	}

	result, err := database.GenericDBCall[[]database.User](
		sqlQuery,
		[]database.User{},
		&database.DatabaseUser{},
		newUser.Email,
		hashedPassed,
	)

	if err != nil {
		return defaultValue, err
	}

	return result[0], nil
}

func QueryUsers() ([]database.User, error) {
	sqlQuery := `
		SELECT
			_id, email, profile, created_at
		FROM
			users
	`

	defaultValue := []database.User{}

	result, err := database.GenericDBCall[[]database.User](
		sqlQuery,
		[]database.User{},
		&database.DatabaseUser{},
	)

	if err != nil {
		return defaultValue, err
	}

	return result, nil
}

func GetCommunitiesOfUser(userId string) ([]database.Community, error) {
	sqlQuery := `
		SELECT
			communities.*
		FROM
			user_communities
		JOIN communities ON communities._id = user_communities.community_id
		WHERE
			user_communities.user_id = $1
	`
	defaultValue := []database.Community{}

	result, err := database.GenericDBCall[[]database.Community](
		sqlQuery,
		[]database.Community{},
		&database.Community{},
		userId,
	)

	if err != nil {
		return defaultValue, err
	}

	return result, nil

}

func GetUserById(userId string) (database.User, error) {
	sqlQuery := `
		SELECT
			*
		FROM
			users
		WHERE
			_id = $1
	`
	defaultValue := database.User{}

	result, err := database.GenericDBCall[[]database.User](
		sqlQuery,
		[]database.User{},
		&database.DatabaseUser{},
		userId,
	)

	if err != nil {
		return defaultValue, err
	}

	return result[0], nil
}
