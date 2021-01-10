// Package cmd is where the dadjoke magic happens
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gochuckn",
	Version: "0.0.1",
	Short:   "A random roundhouse kick to your terminal",
	Long:    `A simple go based CLI to grab a random Chuck Norris 'fact' from api.chucknorris.io `,
	Run: func(cmd *cobra.Command, args []string) {
		getChuckFact()
	},
}

// ChuckFact struct for json data returned from the API
type ChuckFact struct {
	IconURL   string `json:"icon_url"`
	ID        string `json:"id"`
	URL       string `json:"url"`
	ChuckFact string `json:"value"`
}

func getChuckFact() {
	url := "https://api.chucknorris.io/jokes/random"
	responseBytes := getChuckFactData(url)
	fact := ChuckFact{}

	if err := json.Unmarshal(responseBytes, &fact); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	fmt.Println(string(fact.ChuckFact))
}

func getChuckFactData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not request a Chuck fact. %v", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
