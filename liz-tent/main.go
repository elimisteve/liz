package main

import (
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tent/hawk-go"
	"github.com/tent/tent-client-go"
)

// This package based off of:
// https://github.com/tent/tent-examples-go
// Commit:
// f2f75c056333f6c6015403ff9910371aad4d4381

func main() {

	datafile := "/home/traveller/.liz/data"
	data := make(map[string]string)
	tent.HTTP.Transport = &roundTripRecorder{roundTripper: tent.HTTP.Transport}

	_, err := os.Stat(datafile)
	if err != nil {

		discover()

		creds := createApp()
		data["ID"] = creds.ID
		data["Key"] = creds.Key
		data["App"] = creds.App
		data["Delegate"] = creds.Delegate

		rawdata, _ := json.Marshal(data)
		ioutil.WriteFile(datafile, rawdata, 0644)

	} else {
		meta, _ = tent.Discover("https://ian.cupcake.is")
		client = &tent.Client{Servers: meta.Servers}

		rawdata, _ := ioutil.ReadFile(datafile)
		json.Unmarshal(rawdata, &data)

		creds := &hawk.Credentials{
			ID:       data["ID"],
			Key:      data["Key"],
			Hash:     sha256.New,
			App:      data["App"],
			Delegate: data["Delegate"]}
		client.Credentials = creds

		text := os.Args[1]
		// Actual newlines (0x0A) don't work here. Backslash N (0x5C 0x6E) does.
		escaped_text := strings.Replace(text, "\x0a", "\x5c\x6e", -1)
		newPost(escaped_text)
	}
}
