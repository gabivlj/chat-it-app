package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/repository"
)

func BenchmarkRegister(b *testing.B) {
	b.StopTimer()
	userRep := repository.NewRepository()
	defer userRep.Disconnect(context.TODO())
	b.StartTimer()
	user, err := userRep.SaveUser(context.TODO(), &domain.User{Username: "gabivlj02", Password: "123456"})
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	b.Log(user)
}

func BenchmarkUserLog(b *testing.B) {
	b.StopTimer()
	userRep := repository.NewRepository()
	defer userRep.Disconnect(context.TODO())
	b.StartTimer()
	user, err := userRep.FindByID(context.TODO(), "5ecff485476916b6cc2d180f")
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	if user.ID != "5ecff485476916b6cc2d180f" {
		b.Errorf("ERROR, ids don't match")
		b.FailNow()
	}

	b.Log(user)
}

func BenchmarkUserLogIn(b *testing.B) {
	timeNow := time.Now().Unix()
	userRep := repository.NewRepository()
	defer userRep.Disconnect(context.TODO())
	log.Println(time.Now().Unix() - timeNow)
	timeNow = time.Now().Unix()
	user, session, err := userRep.LogUser(context.TODO(), &domain.User{Password: "123456", Username: "gabivlj02"})
	log.Println(time.Now().Unix() - timeNow)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	us, err := userRep.VerifySession(session)
	log.Println(user, session, us, err)
}
