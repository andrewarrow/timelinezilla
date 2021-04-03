package network

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justincampbell/timeago"
)

type TimelineProfileWrapper struct {
	Profile []Timeline `json:"profile"`
}
type TimelineInboxWrapper struct {
	Inbox []Timeline `json:"inbox"`
}
type TimelineRecentWrapper struct {
	Recent []Timeline `json:"recent"`
}

func ShowRecent(c *gin.Context) {
	i := TlzIndex(c)
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	c.JSON(200, gin.H{"recent": universes[uids[i]].Recent})
}
func ShowInbox(c *gin.Context) {
	i := TlzIndex(c)
	from := c.Request.Header["Username"]
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	fromIndex := universes[uids[i]].UsernameToIndex(from[0]) - 1
	if fromIndex < 255 {
		c.JSON(200, gin.H{"inbox": universes[uids[i]].Inboxes[fromIndex]})
		return
	}
	c.JSON(200, gin.H{"inbox": "well..."})
}

func DisplayRecentTimelines(s string) {
	var tw TimelineRecentWrapper
	json.Unmarshal([]byte(s), &tw)
	fmt.Println("Recent")
	for i, t := range tw.Recent {
		fmt.Printf("%02d. %20s %20s %s\n", i+1, t.From,
			timeago.FromDuration(time.Since(t.AsTime())), t.Text)
		if i > 20 {
			break
		}
	}
}
func DisplayInboxTimelines(s string) {
	var tw TimelineInboxWrapper
	json.Unmarshal([]byte(s), &tw)
	fmt.Println("Inbox")
	for i, t := range tw.Inbox {
		fmt.Printf("%02d. %20s %20s %s\n", i+1, t.From,
			timeago.FromDuration(time.Since(t.AsTime())), t.Text)
		if i > 20 {
			break
		}
	}
}
func DisplayProfileTimelines(s string) {
	var tw TimelineProfileWrapper
	json.Unmarshal([]byte(s), &tw)
	fmt.Println("Profile")
	for i, t := range tw.Profile {
		fmt.Printf("%02d. %20s %20s %s\n", i+1, t.From,
			timeago.FromDuration(time.Since(t.AsTime())), t.Text)
	}
}