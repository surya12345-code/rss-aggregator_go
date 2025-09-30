package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/surya123/RSSaggregator/internal/database"
)

type user struct {
	ID         uuid.UUID `json:"id"`
	Create_at  time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	Api_key    string    `json:"apikey"`
}
type feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"apdated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseUserToUser(dbuser database.User) user {
	return user{
		ID:         dbuser.ID,
		Create_at:  dbuser.CreatedAt,
		Updated_at: dbuser.UpdatedAt,
		Name:       dbuser.Name,
		Api_key:    dbuser.ApiKey,
	}

}
func databasefeedtofeed(dbfeed database.Feed) feed {
	return feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}
func databasefeedstofeeds(dbfeeds []database.Feed) []feed {
	feeds := []feed{}
	for _, dbfeed := range dbfeeds {
		feeds = append(feeds, databasefeedtofeed(dbfeed))
	}
	return feeds

}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"apdated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedfollowstofeedfollows(dbfeedfollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbfeedfollow.ID,
		CreatedAt: dbfeedfollow.CreatedAt,
		UpdatedAt: dbfeedfollow.UpdatedAt,
		UserID:    dbfeedfollow.UserID,
		FeedID:    dbfeedfollow.FeedID,
	}
}

func databaseGetuserstousers(dbusers []database.User) []user {
	Allusers := []user{}
	for _, user := range dbusers {
		Allusers = append(Allusers, databaseUserToUser(user))

	}
	return Allusers
}
func databaseGetuserfeedstofeeds(userfeeds []database.FeedFollow) []FeedFollow {
	Allfeeds := []FeedFollow{}
	for _, Feed := range userfeeds {
		Allfeeds = append(Allfeeds, databaseFeedfollowstofeedfollows(Feed))
	}
	return Allfeeds
}
