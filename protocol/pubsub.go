package protocol

import (
	"context"
	"encoding/base64"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func PubSubInit() {
	var err error
	sDec, _ := base64.StdEncoding.DecodeString(con.Google.Pubsubkey)
	opt := option.WithCredentialsJSON(sDec)
	pub, err = pubsub.NewClient(context.Background(), con.Google.Projectid, opt)
	if err != nil {
		panic(err)
	}
	defer pub.Close()
}
