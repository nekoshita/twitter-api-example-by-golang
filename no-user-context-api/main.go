package main

import (
	"flag"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
}

func main() {
	// read credentials from environment variables if available
	config := &Config{
		TwitterConsumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
	}
	// allow consumer credential flags to override config fields
	consumerKey := flag.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flag.String("consumer-secret", "", "Twitter Consumer Secret")
	flag.Parse()
	if *consumerKey != "" {
		config.TwitterConsumerKey = *consumerKey
	}
	if *consumerSecret != "" {
		config.TwitterConsumerSecret = *consumerSecret
	}
	if config.TwitterConsumerKey == "" {
		log.Fatal("Missing Twitter Consumer Key")
	}
	if config.TwitterConsumerSecret == "" {
		log.Fatal("Missing Twitter Consumer Secret")
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	clientConfig := &clientcredentials.Config{
		ClientID:     config.TwitterConsumerKey,
		ClientSecret: config.TwitterConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := clientConfig.Client(oauth2.NoContext)

	// Twitter client
	twitterClient := twitter.NewClient(httpClient)

	tweets, _, err := twitterClient.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: "nekoshita_yuki",
		Count:      5,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(tweets)
}
