package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/repository"
)

func BenchmarkUserLog(b *testing.B) {
	b.StopTimer()
	userRep := repository.NewRepository()
	b.StartTimer()
	user, err := userRep.FindByID(context.TODO(), "5ecf19dc95da119a79c577d9")
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	b.Log(user)
}

func BenchmarkUserLogIn(b *testing.B) {
	b.StopTimer()
	userRep := repository.NewRepository()
	b.StartTimer()
	user, session, err := userRep.LogUser(context.TODO(), &domain.User{Password: "mhm?"})
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	us, err := userRep.VerifySession(session)
	fmt.Println(user, session, us, err)
}
