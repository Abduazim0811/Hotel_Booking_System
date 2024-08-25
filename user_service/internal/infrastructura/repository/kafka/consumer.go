package kafka

import (
	"context"
	"encoding/json"
	"log"
	userservice "user_service/service/user_service"
	"user_service/userproto"

	"github.com/twmb/franz-go/pkg/kgo"

)

type ConsumerUser struct {
	C   *userservice.Service
	Ctx context.Context
}

func (u *ConsumerUser) Consumer() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("hotelusers"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	for {
		fetches := client.PollFetches(ctx)
		if err := fetches.Errors(); len(err) > 0 {
			log.Fatal(err)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				if err := u.Adjust(record); err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func (u *ConsumerUser) Adjust(record *kgo.Record) error {
	switch string(record.Key) {
	case "create":
		if err := u.Create(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "update":
		if err := u.Update(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "delete":
		if err := u.Delete(record.Value); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (u *ConsumerUser) Create(req []byte) error {
	var req1 userproto.UserRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = userproto.UserRequest{
		Username: req1.Username,
		Age:      req1.Age,
		Email:    req1.Email,
		Password: req1.Password,
	}
	_, err := u.C.CreateUser(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *ConsumerUser) Update(req []byte) error {
	var req1 userproto.UpdateUserReq

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = userproto.UpdateUserReq{
		UserId:   req1.UserId,
		Username: req1.Username,
		Age:      req1.Age,
		Email:    req1.Email,
	}
	_, err := u.C.UpdateUser(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *ConsumerUser) Delete(req []byte) error {
	var req1 userproto.GetUserRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = userproto.GetUserRequest{
		Id: req1.Id,
	}

	_, err := u.C.DeleteUser(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
