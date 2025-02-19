package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"strconv"
)

type exchangeRate struct {
	Base       string             `json:"base"`
	Timestamp  int                `json:"timestamp"`
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Rates      map[string]float64 `json:"rates"`
}

func initialBaseModel() CurrencyModel {
	return CurrencyModel{
		choices:  []string{"USD", "EUR", "GBP", "TRY"},
		cursor:   0,
		selected: "TRY",
	}
}

func initialTargetModel() CurrencyModel {
	return CurrencyModel{
		choices:  []string{"USD", "EUR", "GBP", "TRY"},
		cursor:   3,
		selected: "USD",
	}
}

func getConvertedValue(value float64, base, target string) (float64, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&base=%s", appId, "USD"), nil)
	if err != nil {
		fmt.Println("Error creating request")
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request")
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return 0, err
	}

	data := exchangeRate{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, nil
	}

	return (value / data.Rates[base]) * data.Rates[target], nil
}

var (
	appId       string
	defaultFlag = flag.Bool("d", true, "Default currency")
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	flag.Parse()
	if envId := os.Getenv("appId"); envId != "" {
		appId = envId
	}

	if *defaultFlag {
		val, err := getConvertedValue(1.0, "USD", "TRY")
		if err != nil {
			fmt.Println("Error getting converted value")
			return
		}
		fmt.Printf("%.2f USD = %.2f TRY\n", 1.0, val)
		return
	}

	p := tea.NewProgram(initialBaseModel())
	if model, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	} else {
		base := model.(CurrencyModel).selected
		p = tea.NewProgram(initialTargetModel())
		if model, err := p.Run(); err != nil {
			fmt.Println("Error running program:", err)
		} else {
			target := model.(CurrencyModel).selected

			fmt.Println("Enter value: ")
			reader := bufio.NewReader(os.Stdin)

			valueStr, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println("Error reading value")
				return
			}

			value, err := strconv.ParseFloat(string(valueStr), 64)
			if err != nil {
				fmt.Println("Error parsing value, using default (1.0)")
				value = 1.0
			}

			convertedValue, err := getConvertedValue(value, base, target)
			if err != nil {
				fmt.Println("Error getting converted value")
				return
			}
			fmt.Printf("%.2f %s = %.2f %s\n", value, base, convertedValue, target)
		}
	}
}
