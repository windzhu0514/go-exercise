package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/ChimeraCoder/gojson"
)

func TestSameName(t *testing.T) {
	// 编码后如果出现了相同的json key，这个key不会编码到json字符串中。
	// 结果 {}
	type Price struct {
		TicketPrice  float64 `json:"ticketPrice"`
		TicketPrice2 float32 `json:"ticketPrice"`
	}

	var p = Price{10.2, 25.6}
	d, err := json.Marshal(p)
	fmt.Println(string(d), err)

	// Output:
	// {} <nil>
}

func TestEmptyArraySlice(t *testing.T) {
	arr := [3]int{}
	s := []int{}

	d, err := json.Marshal(arr)
	fmt.Println(string(d), err)
	d, err = json.Marshal(s)
	fmt.Println(string(d), err)

	// Output:
	// [0,0,0] <nil>
	// [] <nil>
}

func TestAnonymous(t *testing.T) {
	type Cat struct {
	}

	type Animal struct {
		Cat `json:"cat"`
	}

	type Animal2 struct {
		Cat interface{}
	}

	var a Animal
	fmt.Println(a)
}

func TestHtml(t *testing.T) {
	jsonStr := `{"abx":"\u0027business\u0027:null,\u0027channelSourceType\u0027:11"}`

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		t.Error(err)
		return
	}

	byteData, err := json.Marshal(jsonData)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(byteData))

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(jsonData); err != nil {
		t.Error(err)
		return
	}
	str := buf.String()
	ioutil.WriteFile("111.txt", buf.Bytes(), os.ModePerm)
	fmt.Printf("%s\n", str)
}

func TestUnmarshal(t *testing.T) {
	var result interface{}
	jsonStr := `{"id":7344683,"node_id":"MDEwOlJlcG9zaXRvcnk3MzQ0Njgz","name":"gojson","full_name":"ChimeraCoder/gojson","private":false,"owner":{"login":"ChimeraCoder","id":376414,"node_id":"MDQ6VXNlcjM3NjQxNA==","avatar_url":"https://avatars0.githubusercontent.com/u/376414?v=4","gravatar_id":"","url":"https://api.github.com/users/ChimeraCoder","html_url":"https://github.com/ChimeraCoder","followers_url":"https://api.github.com/users/ChimeraCoder/followers","following_url":"https://api.github.com/users/ChimeraCoder/following{/other_user}","gists_url":"https://api.github.com/users/ChimeraCoder/gists{/gist_id}","starred_url":"https://api.github.com/users/ChimeraCoder/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/ChimeraCoder/subscriptions","organizations_url":"https://api.github.com/users/ChimeraCoder/orgs","repos_url":"https://api.github.com/users/ChimeraCoder/repos","events_url":"https://api.github.com/users/ChimeraCoder/events{/privacy}","received_events_url":"https://api.github.com/users/ChimeraCoder/received_events","type":"User","site_admin":false},"html_url":"https://github.com/ChimeraCoder/gojson","description":"Automatically generate Go (golang) struct definitions from example JSON","fork":false,"url":"https://api.github.com/repos/ChimeraCoder/gojson","forks_url":"https://api.github.com/repos/ChimeraCoder/gojson/forks","keys_url":"https://api.github.com/repos/ChimeraCoder/gojson/keys{/key_id}","collaborators_url":"https://api.github.com/repos/ChimeraCoder/gojson/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/ChimeraCoder/gojson/teams","hooks_url":"https://api.github.com/repos/ChimeraCoder/gojson/hooks","issue_events_url":"https://api.github.com/repos/ChimeraCoder/gojson/issues/events{/number}","events_url":"https://api.github.com/repos/ChimeraCoder/gojson/events","assignees_url":"https://api.github.com/repos/ChimeraCoder/gojson/assignees{/user}","branches_url":"https://api.github.com/repos/ChimeraCoder/gojson/branches{/branch}","tags_url":"https://api.github.com/repos/ChimeraCoder/gojson/tags","blobs_url":"https://api.github.com/repos/ChimeraCoder/gojson/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/ChimeraCoder/gojson/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/ChimeraCoder/gojson/git/refs{/sha}","trees_url":"https://api.github.com/repos/ChimeraCoder/gojson/git/trees{/sha}","statuses_url":"https://api.github.com/repos/ChimeraCoder/gojson/statuses/{sha}","languages_url":"https://api.github.com/repos/ChimeraCoder/gojson/languages","stargazers_url":"https://api.github.com/repos/ChimeraCoder/gojson/stargazers","contributors_url":"https://api.github.com/repos/ChimeraCoder/gojson/contributors","subscribers_url":"https://api.github.com/repos/ChimeraCoder/gojson/subscribers","subscription_url":"https://api.github.com/repos/ChimeraCoder/gojson/subscription","commits_url":"https://api.github.com/repos/ChimeraCoder/gojson/commits{/sha}","git_commits_url":"https://api.github.com/repos/ChimeraCoder/gojson/git/commits{/sha}","comments_url":"https://api.github.com/repos/ChimeraCoder/gojson/comments{/number}","issue_comment_url":"https://api.github.com/repos/ChimeraCoder/gojson/issues/comments{/number}","contents_url":"https://api.github.com/repos/ChimeraCoder/gojson/contents/{+path}","compare_url":"https://api.github.com/repos/ChimeraCoder/gojson/compare/{base}...{head}","merges_url":"https://api.github.com/repos/ChimeraCoder/gojson/merges","archive_url":"https://api.github.com/repos/ChimeraCoder/gojson/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/ChimeraCoder/gojson/downloads","issues_url":"https://api.github.com/repos/ChimeraCoder/gojson/issues{/number}","pulls_url":"https://api.github.com/repos/ChimeraCoder/gojson/pulls{/number}","milestones_url":"https://api.github.com/repos/ChimeraCoder/gojson/milestones{/number}","notifications_url":"https://api.github.com/repos/ChimeraCoder/gojson/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/ChimeraCoder/gojson/labels{/name}","releases_url":"https://api.github.com/repos/ChimeraCoder/gojson/releases{/id}","deployments_url":"https://api.github.com/repos/ChimeraCoder/gojson/deployments","created_at":"2012-12-27T19:10:50Z","updated_at":"2020-06-18T05:19:18Z","pushed_at":"2020-06-11T11:36:00Z","git_url":"git://github.com/ChimeraCoder/gojson.git","ssh_url":"git@github.com:ChimeraCoder/gojson.git","clone_url":"https://github.com/ChimeraCoder/gojson.git","svn_url":"https://github.com/ChimeraCoder/gojson","homepage":"","size":210,"stargazers_count":2213,"watchers_count":2213,"language":"Go","has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":true,"has_pages":false,"forks_count":152,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":39,"license":{"key":"other","name":"Other","spdx_id":"NOASSERTION","url":null,"node_id":"MDc6TGljZW5zZTA="},"forks":152,"open_issues":39,"watchers":2213,"default_branch":"master","temp_clone_token":null,"network_count":152,"subscribers_count":44}`
	json.Unmarshal([]byte(jsonStr), &result)
	fmt.Println(result)
	fmt.Printf("%T\n", result)
	s, err := gojson.Generate(strings.NewReader(jsonStr), gojson.ParseJson, "name", "main", []string{"json"}, true, true)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(s))
}

func TestEmptyStruct(t *testing.T) {
	a := struct {
		err struct{}
	}{}

	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
