package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

type mapResponse struct {
	Count		int		`json:"count"`
	Next		string	`json:"next"`
	Previous	string	`json:"previous"`
	Results		[]struct {
		Name	string	`json:"name"`
		URL		string	`json:"url"`
	}	`json:"results"`
}

func commandMapNext(cfg *config, args []string) error {
	body, ok := cfg.Cache.Get(cfg.Next)

	if !ok {
		res, err := http.Get(cfg.Next)
		if err != nil {
			return err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("API call was not successful, response Code: %v", res.StatusCode)
		}

		if err != nil {
			return err
		}

		cfg.Cache.Add(cfg.Next, body)
	}

	mres := mapResponse{}
	err := json.Unmarshal(body, &mres)

	if err != nil {
		return err
	}

	cfg.Next = mres.Next
	cfg.Previous = mres.Previous

	for _, result := range mres.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}

func commandMapPrevious(cfg *config, args []string) error {
	if cfg.Previous == "" {
		fmt.Println("Cannot go back any further!")
		return nil
	}

	res, err := http.Get(cfg.Previous)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("API call was not successful, response Code: %v", res.StatusCode)
	}

	if err != nil {
		return err
	}

	mres := mapResponse{}
	err = json.Unmarshal(body, &mres)

	if err != nil {
		return err
	}

	cfg.Next = mres.Next
	cfg.Previous = mres.Previous

	for _, result := range mres.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}
