package events

import (
	"fmt"

	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type AtndClient struct {
}

func newAtndClient() *AtndClient {
	client := AtndClient{}
	return &client
}

func (c *AtndClient) queryEvent(id string, key string) (*Event, error) {
	doc, err := goquery.NewDocument(fmt.Sprintf("https://atnd.org/events/%s?k=%s", id, key))
	if err != nil {
		return nil, err
	}
	event := &Event{}
	event.Title = doc.Find("#events h1>a").First().Text()
	event.Limit, _ = strconv.Atoi(doc.Find("#members-info>h3>span").First().Text())
	members := []Member{}
	doc.Find("#go_members_ol>li").Each(func(_ int, s *goquery.Selection) {
		member := Member{
			Name: s.Find("a").Text(),
		}
		members = append(members, member)
	})
	event.Members = members
	return event, nil
}
