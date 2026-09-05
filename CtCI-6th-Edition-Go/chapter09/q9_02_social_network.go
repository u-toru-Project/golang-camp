package chapter09

import (
	"fmt"
	"sort"
)

type SocialNetwork struct {
	following map[string]map[string]struct{}
}

func NewSocialNetwork() *SocialNetwork {
	return &SocialNetwork{following: map[string]map[string]struct{}{}}
}

func (n *SocialNetwork) AddUser(userId string) {
	if userId == "" {
		panic("user_id required")
	}
	if _, ok := n.following[userId]; !ok {
		n.following[userId] = map[string]struct{}{}
	}
}

func (n *SocialNetwork) Follow(follower, followee string) {
	if follower == followee {
		panic("cannot follow self")
	}
	n.AddUser(follower)
	n.AddUser(followee)
	n.following[follower][followee] = struct{}{}
}

func (n *SocialNetwork) FriendsOfFriends(userId string, limit ...int) []string {
	if _, ok := n.following[userId]; !ok {
		panic("unknown user")
	}
	lim := 10
	if len(limit) > 0 {
		lim = limit[0]
	}
	direct := n.following[userId]
	seen := map[string]struct{}{userId: {}}
	for friend := range direct {
		seen[friend] = struct{}{}
	}
	scores := map[string]int{}
	type item struct {
		node  string
		depth int
	}
	queue := make([]item, 0)
	for friend := range direct {
		queue = append(queue, item{friend, 1})
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.depth > 2 {
			continue
		}
		for next := range n.following[cur.node] {
			if _, ok := seen[next]; ok {
				continue
			}
			scores[next]++
			if cur.depth < 2 {
				queue = append(queue, item{next, cur.depth + 1})
			}
		}
	}
	type pair struct {
		key   string
		value int
	}
	pairs := make([]pair, 0, len(scores))
	for k, v := range scores {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].value != pairs[j].value {
			return pairs[i].value > pairs[j].value
		}
		return pairs[i].key < pairs[j].key
	})
	if lim > len(pairs) {
		lim = len(pairs)
	}
	out := make([]string, lim)
	for i := 0; i < lim; i++ {
		out[i] = pairs[i].key
	}
	return out
}

func RunQ902() {
	network := NewSocialNetwork()
	for _, user := range []string{"a", "b", "c", "d"} {
		network.AddUser(user)
	}
	network.Follow("a", "b")
	network.Follow("b", "c")
	network.Follow("a", "d")
	fmt.Printf("Suggestions for a: %v\n", network.FriendsOfFriends("a"))
}
