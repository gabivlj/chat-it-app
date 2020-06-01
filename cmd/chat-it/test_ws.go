package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/gabivlj/chat-it/internals/domain"
)

func testWs(h http.Handler) {
	go func() {
		time.Sleep(time.Second * 2)

		c := client.New(h)

		sub := c.WebsocketWithPayload(`subscription @user(id:"5ed1b3ab0f94ac59a503574a") { newMessage(postId:"5ed1b3ab0f94ac59a503574a") { id text user { username } createdAt } }`, map[string]interface{}{"Authorization": "test"})
		defer sub.Close()

		go func() {
			var resp interface{}
			time.Sleep(10 * time.Millisecond)
			options := func(bd *client.Request) {
				bd.HTTP.Header["Authorization"] = []string{"test"}
			}
			c := client.New(h, options)
			err := c.Post(`mutation { 
					a:sendMessage(text:"Hello!", postId:"5ed1b3ab0f94ac59a503574a", userId: "1") { id user { username } } 
					b:sendMessage(text:"Hello Vektah!", postId:"5ed1b3ab0f94ac59a503574a", userId:"1") { id user { username }} 
					c:sendMessage(text:"Whats up?", postId:"5ed1b3ab0f94ac59a503574a", userId:"1") { id user { username }} 
				}`, &resp, options)
			fmt.Println(err, resp)
		}()

		type msg struct {
			resp struct {
				NewMsg *domain.Message `json:"newMessage"`
			}
			err error
		}
		newMessage := msg{}
		newMessage.err = sub.Next(&newMessage.resp)
		fmt.Println(newMessage.err, *newMessage.resp.NewMsg)
		newMessage.err = sub.Next(&newMessage.resp)
		fmt.Println(newMessage.err, *newMessage.resp.NewMsg)
		newMessage.err = sub.Next(&newMessage.resp)
		fmt.Println(newMessage.err, *newMessage.resp.NewMsg)
	}()
}
