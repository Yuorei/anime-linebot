package get_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type MovieInfomation struct {
	Results []Result `json:"results"`
}
type Result struct {
	Originalname string `json:"original_name"`
	Posterpath   string `json:"poster_path"`
}

func SearchTvGET(query string) MovieInfomation {
	url := "https://api.themoviedb.org/3/search/tv?api_key=" + os.Getenv("TMDBAPI") + "&page=1&query=" + query
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//data := make([]Result, 0)
	var data MovieInfomation
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("err", err)
		log.Fatal(err)
	}

	return data
}
