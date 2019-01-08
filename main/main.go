package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	facebookGraphAPIMeEndpoint = "https://graph.facebook.com/v3.2/me"
	field                      = "likes"
)

//Like ...
type Like struct {
	Name string `json:"name"`
}

//Likes ...
type Likes struct {
	DataList []Like `json:"data"`
}

//LikesResponse ...
type LikesResponse struct {
	LikesCollection Likes `json:"likes"`
}

func main() {
	accessTokenIn := flag.String("accessToken", "", "https://developers.facebook.com/docs/graph-api/faq#faq_129986974356991")
	flag.Parse()

	myLikes := get(*accessTokenIn)
	fmt.Printf("Posts I liked %s \n", parseLikes(myLikes))

}

func get(accessTokenIn string) []byte {
	facebookUserLikesURL := facebookGraphAPIMeEndpoint + "?fields=" + field + "&access_token=" + accessTokenIn
	resp, err := http.Get(facebookUserLikesURL)
	if err != nil {
		log.Fatal(err)
	}

	body, error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if error != nil {
		log.Fatal(error)
	}

	return body
}

func parseLikes(body []byte) *LikesResponse {
	var s = new(LikesResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}
	return s
}