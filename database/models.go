package database

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Community struct {
	ID        uuid.UUID       `json:"_id"`
	Name      string          `json:"name"`
	Slug      string          `json:"slug"`
	Profile   json.RawMessage `json:"profile"`
	CreatedAt time.Time       `json:"created_at"`
}

func (dbU *Community) Scan(rows *sql.Rows) ([]Community, error) {
	result := []Community{}

	for rows.Next() {
		holder := Community{}
		err := rows.Scan(&holder.ID, &holder.Name, &holder.Slug, &holder.Profile, &holder.CreatedAt)

		if err != nil {
			return []Community{}, err
		}

		result = append(result, holder)
	}

	return result, nil
}

type DatabaseUser struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Profile   json.RawMessage
	CreatedAt time.Time
}

func (*DatabaseUser) Scan(rows *sql.Rows) ([]User, error) {
	result := []User{}

	for rows.Next() {
		holder := DatabaseUser{}
		err := rows.Scan(&holder.ID, &holder.Email, &holder.Password, &holder.Profile, &holder.CreatedAt)

		if err != nil {
			return []User{}, err
		}

		result = append(result, User{
			ID:        holder.ID,
			Email:     holder.Email,
			CreatedAt: holder.CreatedAt,
			Profile:   holder.Profile,
		})
	}

	return result, nil
}

type User struct {
	ID        uuid.UUID       `json:"_id"`
	Email     string          `json:"email"`
	Profile   json.RawMessage `json:"profile"`
	CreatedAt time.Time       `json:"created_at"`
}

type UserCommunity struct {
	UserID      uuid.UUID
	CommunityID uuid.UUID
}

func (*UserCommunity) Scan(rows *sql.Rows) ([]UserCommunity, error) {
	result := []UserCommunity{}

	for rows.Next() {
		holder := UserCommunity{}
		err := rows.Scan(&holder.UserID, &holder.CommunityID)

		if err != nil {
			return []UserCommunity{}, err
		}

		result = append(result, holder)
	}

	return result, nil
}
