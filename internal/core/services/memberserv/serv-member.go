package memberserv

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/oasis-prime/oas-platform-firebase-core/domain/firebasedm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain/googledm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	memberRepo         ports.MemberRepository
	pub                *pubsub.Client
	firebaseMemberRepo ports.FirebaseMemberRepository
}

func NewService(repo ports.MemberRepository, pub *pubsub.Client, firebaseMemberRepo ports.FirebaseMemberRepository) *service {
	return &service{
		memberRepo:         repo,
		pub:                pub,
		firebaseMemberRepo: firebaseMemberRepo,
	}
}

func (serv *service) MemberRegister(condition firebasedm.MemberRegister) (err error) {
	_, err = serv.firebaseMemberRepo.MemberRegister(condition)

	if err != nil {
		return err
	}

	return nil
}

func (serv *service) PublisherVerifyEmail(condition domain.PublisherVerifyEmail) (err error) {
	topic := serv.pub.Topic("oas-trigger-event")

	b := googledm.VerifyEmailPublish{
		Email: condition.Email,
		Token: condition.Token,
	}

	byteData, _ := json.Marshal(b)

	result := topic.Publish(context.Background(), &pubsub.Message{
		Data: []byte(byteData),
		Attributes: map[string]string{
			"email": condition.Email,
			"token": condition.Token,
		},
	})

	id, err := result.Get(context.Background())
	if err != nil {
		return err
	}

	log.Printf("Public message to pub/sub complate : %s", id)

	return nil
}
