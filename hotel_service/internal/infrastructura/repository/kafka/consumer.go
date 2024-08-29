package kafka

import (
	"context"
	"encoding/json"
	"hotel_service/hotelproto"
	hotelservice "hotel_service/service/hotel_service"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

type ConsumerUser struct {
	C   *hotelservice.Service
	Ctx context.Context
}

func (u *ConsumerUser) Consumer() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("broker:9092"),
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
	var req1 hotelproto.HotelRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = hotelproto.HotelRequest{
		Address: req1.Address,
		Location: req1.Location,
		Name: req1.Name,
		Rating: req1.Rating,
	}
	_, err := u.C.CreateHotel(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *ConsumerUser) Update(req []byte) error {
	var req1 hotelproto.Hotel

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = hotelproto.Hotel{
		HotelId: req1.HotelId,
		Name: req1.Name,
		Address: req1.Address,
		Location: req1.Location,
		Rating: req1.Rating,
	}
	_, err := u.C.UpdateHotel(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *ConsumerUser) Delete(req []byte) error {
	var req1 hotelproto.HotelResponse
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = hotelproto.HotelResponse{
		HotelId: req1.HotelId,
	}

	_, err := u.C.DeleteHotel(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
