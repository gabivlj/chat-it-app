package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/stretchr/testify/assert"
)

func TestChatSubscriptions(t *testing.T) {
	c := client.New(handler.NewDefaultServer(NewExecutableSchema(New())))

	options := func(bd *client.Request) {
		bd.HTTP.Header.Add("Authorization", "test")
	}
	sub := c.Websocket(`subscription @user(id:"5ed0626ecbb3a60797995377") { newMessage(postId:"5ed1b3ab0f94ac59a503574a") { id text user createdAt } }`, options)
	defer sub.Close()

	go func() {
		var resp interface{}
		time.Sleep(10 * time.Millisecond)
		err := c.Post(`mutation { 
				a:sendMessage(text:"Hello!", postId:"5ed0626ecbb3a60797995377") { id } 
				b:sendMessage(text:"Hello Vektah!", postId:"5ed0626ecbb3a60797995377") { id } 
				c:sendMessage(text:"Whats up?", postId:"5ed0626ecbb3a60797995377") { id } 
			}`, &resp)
		assert.NoError(t, err)
	}()

	var msg struct {
		resp struct {
			MessageAdded struct {
				Text      string
				User      *domain.User
				ID        string
				CreatedAt int64
			}
		}
		err error
	}

	msg.err = sub.Next(&msg.resp)
	fmt.Println(msg)
}
