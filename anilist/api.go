package anilist

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

var url = "https://graphql.anilist.co"

type reqbody = map[string]string

// Request makes a GraphQL request to the AniList endpoint
func Request(queryTemplateString string, ReqBody interface{}) []byte {
	query := getReqBodyString(queryTemplateString, ReqBody)
	req, _ := json.Marshal(reqbody{"query": query})
	res, _ := http.Post(url, "application/json", bytes.NewBuffer(req))
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func getReqBodyString(queryTemplateString string, ReqBody interface{}) string {
	var query bytes.Buffer
	queryTemplate, _ := template.New("query").Parse(queryTemplateString)
	queryTemplate.Execute(&query, ReqBody)
	return query.String()
}
