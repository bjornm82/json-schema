package drill

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	query = "query.json"
)

type Client struct {
	url    string
	client *http.Client
}

func NewClient(host string, port int, useSSL bool) *Client {
	client := &Client{}

	var scheme = ""
	if useSSL {
		scheme = "https"
	} else {
		scheme = "http"
	}

	client.url = fmt.Sprintf("%s://%s:%d/%s", scheme, host, port, query)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client.client = &http.Client{Transport: tr}

	return client
}

type RequestBody struct {
	QueryType string `json:"queryType"`
	Query     string `json:"query"`
}

type ResponseBody struct {
	QueryID string   `json:"queryId"`
	Columns []string `json:"columns"`
	Rows    []struct {
		Summary string `json:"summary"`
		Ok      string `json:"ok"`
	} `json:"rows"`
	Metadata           []string `json:"metadata"`
	QueryState         string   `json:"queryState"`
	AttemptedAutoLimit int      `json:"attemptedAutoLimit"`
}

func (c *Client) ValidateSQL() {

}

func (c *Client) UpsertView(name string, sql string) error {
	base := fmt.Sprintf("CREATE OR REPLACE VIEW `s3`.`tmp`.`%s` AS ", name)

	u := RequestBody{
		QueryType: "SQL",
		Query:     base + sql,
	}

	b, err := json.Marshal(u)
	if err != nil {
		return errors.Wrap(err, "unable to marshal upsert object")
	}

	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "unable to make request")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("something went wrong, expected status code 200 but got %d", resp.StatusCode))
	}

	return nil
}

func (c *Client) DeleteView(name string) error {
	query := fmt.Sprintf("DROP VIEW `s3`.`tmp`.`%s`", name)
	u := RequestBody{
		QueryType: "SQL",
		Query:     query,
	}

	b, err := json.Marshal(u)
	if err != nil {
		return errors.Wrap(err, "unable to marshal upsert object")
	}

	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "unable to make request")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("something went wrong, expected status code 200 but got %d", resp.StatusCode))
	}

	return nil
}
