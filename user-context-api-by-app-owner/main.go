package main

import (
	"flag"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Config struct {
	TwitterConsumerKey      string
	TwitterConsumerSecret   string
	TwitterUserAccessToken  string
	TwitterUserAccessSecret string
}

const followTargetTwitterUserScreenName = "nekoshita_yuki"

func main() {
	// read credentials from environment variables if available
	config := &Config{
		TwitterConsumerKey:      os.Getenv("TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret:   os.Getenv("TWITTER_CONSUMER_SECRET"),
		TwitterUserAccessToken:  os.Getenv("TWITTER_USER_ACCESS_TOKEN"),
		TwitterUserAccessSecret: os.Getenv("TWITTER_USER_ACCESS_SECRET"),
	}
	// allow consumer credential flags to override config fields
	consumerKey := flag.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flag.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flag.String("access-token", "", "Twitter User Access Token")
	accessSecret := flag.String("access-secret", "", "Twitter User Access Secret")
	flag.Parse()
	if *consumerKey != "" {
		config.TwitterConsumerKey = *consumerKey
	}
	if *consumerSecret != "" {
		config.TwitterConsumerSecret = *consumerSecret
	}
	if *accessToken != "" {
		config.TwitterUserAccessToken = *accessToken
	}
	if *accessSecret != "" {
		config.TwitterUserAccessSecret = *accessSecret
	}
	if config.TwitterConsumerKey == "" {
		log.Fatal("Missing Twitter Consumer Key")
	}
	if config.TwitterConsumerSecret == "" {
		log.Fatal("Missing Twitter Consumer Secret")
	}
	if config.TwitterUserAccessToken == "" {
		log.Fatal("Missing Twitter User Access Token")
	}
	if config.TwitterUserAccessSecret == "" {
		log.Fatal("Missing Twitter User Access Secret")
	}

	oauthConfig := oauth1.NewConfig(config.TwitterConsumerKey, config.TwitterConsumerSecret)
	token := oauth1.NewToken(config.TwitterUserAccessToken, config.TwitterUserAccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	_, _, err := twitterClient.Friendships.Create(&twitter.FriendshipCreateParams{
		ScreenName: followTargetTwitterUserScreenName,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully followed @%s", followTargetTwitterUserScreenName)
}
