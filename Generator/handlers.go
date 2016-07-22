package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
)

// FormatRules Defines the rules that are used when generating a string:
type FormatRules struct {
	// Format This provides the template string that is used to generate the
	//        data value, the following special characters are observed:
	//   # - Replaced with a random digit from those provided in 'Digits'
	//   @ - Replaced with a random character from those provided in 'Characters'
	Format string `json:"format"`

	// Digits Defines the digits that can be used to replace a random digit.
	Digits string `json:"digits"`

	// Characters Defines the characters that can be used to replace a random
	//            character.
	Characters string `json:"characters"`
}

// Generate Generates a random data string following the rules given in 'FormatRules'
func Generate(w http.ResponseWriter, r *http.Request) {
	format := FormatRules{
		Digits:     "0123456789",
		Characters: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &format); err != nil {
		w.WriteHeader(422) // Unprocessable Entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	b := make([]byte, len(format.Format))
	for i := range b {
		if format.Format[i] == '#' {
			b[i] = format.Digits[rand.Intn(len(format.Digits))]
		} else if format.Format[i] == '@' {
			b[i] = format.Characters[rand.Intn(len(format.Characters))]
		} else {
			b[i] = format.Format[i]
		}
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(string(b)); err != nil {
		panic(err)
	}
}
