package twitter

import (
	"fmt"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/eamirgh/go-twitter-cleaner/config"
)

// Twitter holds configs and clinet
type Twitter struct {
	Client *twitter.Client
	Config *config.Config
}

// New returns a new Twitter struct
func New() *Twitter {
	var t Twitter
	c := config.New()
	t.createClient(c)
	return &t
}

func (t *Twitter) createClient(c *config.Config) {
	oauthConfig := oauth1.NewConfig(c.APIKey, c.APISecret)
	oauthToken := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)
	t.Client = twitter.NewClient(httpClient)
}

// DeleteTweets delets tweets :D
func (t *Twitter) DeleteTweets() {
	params := twitter.UserTimelineParams{ScreenName: t.Config.ScreenName, Count: 200, IncludeRetweets: twitter.Bool(true)}
	lastTweetID := int64(0)
	for {
		if lastTweetID != 0 {
			params.MaxID = lastTweetID
		}
		tweets, _, err := t.Client.Timelines.UserTimeline(&params)
		if err != nil {
			log.Fatalln(err)
			return
		}
		deadline := time.Now().Local().AddDate(0, 0, -7)
		for _, tweet := range tweets {
			lastTweetID = tweet.ID
			created, _ := tweet.CreatedAtTime()
			if created.Before(deadline) {
				dt, _, err := t.Client.Statuses.Destroy(tweet.ID, nil)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("delated :%v", dt.ID)
			}
		}
		time.Sleep(60 * time.Second)
	}
}
