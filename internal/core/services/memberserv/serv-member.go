package memberserv

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain/googledm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
	"google.golang.org/api/option"
)

type service struct {
	repo      ports.MemberRepository
	projectId string
	pubsubkey string
}

func NewService(repo ports.MemberRepository, projectId, pubsubkey string) *service {
	return &service{
		repo:      repo,
		projectId: projectId,
		pubsubkey: pubsubkey,
	}
}

func (serv *service) PublisherVerifyEmail() (err error) {
	sDec, _ := base64.StdEncoding.DecodeString(serv.pubsubkey)
	opt := option.WithCredentialsJSON(sDec)

	// json.Unmarshal()

	// fmt.Println(string(sDec))

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, serv.projectId, opt)
	if err != nil {
		return err
	}
	defer client.Close()

	topic := client.Topic("oas-trigger-event")

	b := googledm.VerifyEmailPublish{
		Email: "sittha.raksasawat@gmail.com",
		Token: "tokenXXXYYY",
	}

	byteData, _ := json.Marshal(b)

	// msg := "Hello World!!!!"

	fmt.Println(byteData)

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(byteData),
		Attributes: map[string]string{
			"email": "reworldt@gmail.com",
			"token": "tokenXXXYYY",
		},
	})

	id, err := result.Get(ctx)
	if err != nil {
		return err
	}

	log.Printf("Public message to pub/sub complate : %s", id)

	return nil
}
