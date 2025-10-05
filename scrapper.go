package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/surya123/RSSaggregator/internal/database"
)

func startscrapping(
	db *database.Queries,
	concurrency int,
	timebetweenrequests time.Duration,
) {
	log.Printf("Scrapping on %v goroutines for every %s duration", concurrency, timebetweenrequests)
	ticker := time.NewTicker(timebetweenrequests)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching the feeds", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapefeed(db, wg, feed)
		}
		wg.Wait()

	}
}
func scrapefeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking the feed as fetched:", err)
		return
	}

	rssfeed, err := urltofeed(feed.Url)
	if err != nil {
		log.Println("Error fetching the feed as fetched:", err)
		return
	}
	for _, item := range rssfeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		Pubdate, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Printf("failed to parse date %v, with err %v", item.PubDate, err)
			continue
		}
		_, errr := db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				FeedID:      feed.ID,
				PublishedAt: Pubdate,
				Url:         item.Link,
			})
		if errr != nil {
			if errr.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
				// log.Println("Post already exists, skipping:", item.Link)
				continue
			}
			log.Println("Failed to create a post", err)
			return
		}

	}
	log.Printf("feed %s collected %v posts found", feed.Name, len(rssfeed.Channel.Item))
}
