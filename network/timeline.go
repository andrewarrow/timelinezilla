package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowTimelines(c *gin.Context) {

	ByFromLock.Lock()
	c.JSON(200, gin.H{"from": ByFrom})
	ByFromLock.Unlock()
}

func mapIt(tokens []string) (key, val string) {
	if len(tokens) < 2 {
		return "", ""
	}
	return tokens[0], tokens[1]
}
func mapBody(c *gin.Context) map[string]string {
	defer c.Request.Body.Close()
	body, _ := ioutil.ReadAll(c.Request.Body)
	m := map[string]string{}
	for _, line := range strings.Split(string(body), "\n") {
		k, v := mapIt(strings.Split(line, "="))
		if k != "" {
			m[k] = v
		}
	}
	return m
}

type TimelineWrapper struct {
	From map[string][]Timeline `json:"from"`
}

type Timeline struct {
	Text     string `json:"text"`
	From     string `json:"from"`
	PostedAt int64  `json:"posted_at"`
}

var ByFromLock sync.Mutex
var ByFrom map[string][]Timeline = map[string][]Timeline{}

func CreateTimeline(c *gin.Context) {
	m := mapBody(c)
	t := Timeline{}
	t.Text = m["text"]
	t.From = m["username"]
	t.PostedAt = time.Now().Unix()

	ByFromLock.Lock()
	ByFrom[t.From] = append([]Timeline{t}, ByFrom[t.From]...)
	ByFromLock.Unlock()
}

func DisplayTimelines(s string) {
	var tw TimelineWrapper
	json.Unmarshal([]byte(s), &tw)
	for k, v := range tw.From {
		fmt.Println(k)
		for i, t := range v {
			fmt.Printf("%02d. %s\n", i+1, t.Text)
		}
	}
}
