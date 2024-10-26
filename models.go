package main

import (
	"github.com/google/uuid"
	"github.com/mrminko/rssagg/internal/database"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}
