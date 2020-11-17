package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Websites struct which contains
type Websites struct {
	Websites []websites `json:"Websites"`
}

// websites struct
type websites struct {
	WebsiteName string `json:"websitename"`
	OutcomeURL  string `json:"outcomeurl"`
	IP          string `json:"ip"`
	Flavor      string `json:"flavor"`
	OpenPORTS   string `json:"openports"`
	OS          string `json:"os"`
}

func main() {

	filePath := os.Args[1]
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened ex.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Websites array
	var Websites Websites

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'Websites' which we defined above
	json.Unmarshal(byteValue, &Websites)

	// we iterate through every websites within our Websites array and

	for i := 0; i < len(Websites.Websites); i++ {
		fmt.Println("website name: " + Websites.Websites[i].WebsiteName)
		fmt.Println("outcome url: " + Websites.Websites[i].OutcomeURL)
		fmt.Println("IP: " + Websites.Websites[i].IP)
		fmt.Println("Flavor: " + Websites.Websites[i].Flavor)
		fmt.Println("open ports: " + Websites.Websites[i].OpenPORTS)
		fmt.Println("OS: " + Websites.Websites[i].OS)
	}

}
