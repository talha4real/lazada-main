package address

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/asmcos/requests"
	"os"
	"strings"
	"sync"
	"time"
)

type Configuration struct {
	Email           string
	Password        string
	MonitorInterval int
	Proxy           string
	ProductURL      string
	Webhook         string
}

func getCSRFToken(responseText string) string {
	// Extracts CSRF token from HTML response.
	document, err := goquery.NewDocumentFromReader(strings.NewReader(responseText))
	if err != nil {
		panic(err)
	}
	token, exists := document.Find(`meta[name="csrf-token"], meta[name="X-CSRF-TOKEN"], input[name="csrf_token"]`).Attr("content")
	if exists {
		return token
	}
	return ""
}

func initializeSession(client *requests.Request) string {
	// Starts a session and fetches CSRF token and cookies.
	url := "https://member.lazada.co.th/user/login"
	h := requests.Header{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Language": "en-US,en;q=0.5",
		"Accept-Encoding": "gzip, deflate, br",
	}
	resp, err := client.Get(url, h)
	if err == nil {
		csrfToken := getCSRFToken(resp.Text())
		return csrfToken
		//println(resp.Text())
	}

	return ""
}

func login(client *requests.Request, csrfToken string, email string, password string) bool {
	// Logs in to the site with the provided credentials.
	url := "https://member.lazada.co.th/user/api/login"
	data := requests.Datas{
		"loginName": email,
		"password":  password,
	}
	h := requests.Header{
		"Content-Type": "application/json",
		"Referer":      url,
		"X-CSRF-TOKEN": csrfToken,
	}
	resp, err := client.PostJson(url, h, data)

	if err == nil {
		var json map[string]interface{}
		resp.Json(&json)

		// Check if "success" key exists in the map
		if _, ok := json["success"].(bool); ok {
			return true
		} else {
			return false
		}
	}
	return false
}

func getUserData(session *requests.Request, csrfToken string) (map[string]interface{}, error) {
	// URL for getting user data
	url := "https://member.lazada.co.th/user/api/getUser"
	h := requests.Header{
		"x-csrf-token": csrfToken,
		"Referer":      "https://www.lazada.co.th/cart",
	}

	resp, err := session.Get(url, h)
	if err != nil {
		return nil, err
	}
	var json map[string]interface{}
	resp.Json(&json)
	return json, nil
}

func addAddress(tname string, session *requests.Request, csrfToken string, xdata map[string]interface{}, postcode string, detail string, detail2 string, subDetail string, uname string, phone string, detailAddress string) bool {
	// Fetch User-Agent value from a URL
	uaURL := "https://lazada-checkout-bol-production.up.railway.app/"
	uaResponse, err := session.Get(uaURL, nil)
	if err != nil {
		return false
	}
	var xjson map[string]interface{}
	uaResponse.Json(&xjson)
	extString := xdata["module"].(map[string]interface{})["ext"].(string)
	//// Define a variable to hold the unmarshalled data
	var data map[string]interface{}
	// Unmarshal the JSON string into the map variable
	xerr := json.Unmarshal([]byte(string(extString)), &data)
	if xerr != nil {
		return false
	}
	// Access the value associated with the key "x-umidtoken"
	umid, ok := data["x-umidtoken"].(string)
	if !ok {
		return false
	}
	uaVal := xjson["uab_value"].(string)
	_ = uaVal
	_ = umid

	// Extract package and item values from params

	//URL for placing order
	checkoutURL := "https://member.lazada.co.th/address/api/createAddress"

	////// Set request headers
	h := requests.Header{
		"x-csrf-token":     csrfToken,
		"Referer":          "https://www.lazada.co.th/cart",
		"Content-Type":     "application/json",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
		"x-ua":             uaVal,
		"x-umidtoken":      umid,
	}

	// Convert submitParams to string

	// Construct payload
	payload := fmt.Sprintf(`{
    "name": "%s",
    "phone": "%s",
    "detailAddress": "%s",
    "locationTreeAddressArray": "[{\"id\":\"R1908769\",\"name\":\"%s\"},{\"id\":\"R435\",\"name\":\"%s\"},{\"id\":\"R1420\",\"name\":\"%s\"},{\"id\":\"RTH86\",\"name\":\"%s\"}]",
    "locationTreeAddressId": "R1908769-R435-R1420-RTH86",
    "locationTreeAddressName": "%s,%s,%s,%s"
}`, uname, phone, detailAddress, detail, detail2, postcode, subDetail, detail, detail2, postcode, subDetail)

	//_ = payload

	//// Send POST request to place order

	response, err := session.PostJson(checkoutURL, h, payload)
	fmt.Println(tname + ": " + response.Text())

	return true

}

func processRecord(record []string) {
	// Process the record here

	name := record[0]
	fmt.Println(name)
	startTime := time.Now()
	fmt.Println(name + ": " + "Starting...")
	fmt.Println(name + ": " + "Config Loaded.")
	email := record[1]
	password := record[2]

	uname := record[4]
	phone := record[5]

	_ = uname
	_ = phone
	_ = email
	_ = startTime
	_ = password

	detailAddress := record[6]
	detail := record[7]
	detail2 := record[8]
	postcode := record[9]
	subDetail := record[10]

	//interval := record[3]
	proxyURL := record[3]
	fmt.Println(name + ": " + "Initiating Session.")
	//num, _ := strconv.Atoi(interval)
	session := requests.Requests()
	session.Proxy(proxyURL)
	csrfToken := initializeSession(session)

	loggedIn := login(session, csrfToken, email, password)
	if !loggedIn {
		fmt.Println(name + ": " + "Failed to login. Maybe check your email for verification")
		return
	}

	fmt.Println(name + ": " + "Logged in successfully.")
	fmt.Println(name + ": " + "Adding address.")

	userdata, _ := getUserData(session, csrfToken)

	id := addAddress(name, session, csrfToken, userdata, postcode, detail, detail2, subDetail, uname, phone, detailAddress)
	if !id {
		fmt.Println(name + ": " + "Failed to add address.")
		return
	}
	fmt.Println(name + ": " + "Address request sent.,")

	// Call the function to visit the product page and extract details

	elapsed := time.Since(startTime)
	fmt.Println(name+": "+"Execution time: %s\n", elapsed)
}
func main() {
	file, err := os.Open("./address.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Read the remaining records and process them concurrently
	for {
		// Read a single record from the CSV file
		record, err := reader.Read()
		if err != nil {
			// Check if it's the end of the file
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error:", err)
			continue // Continue to the next iteration
		}

		// Process the record concurrently
		wg.Add(1)
		go func(record []string) {
			defer wg.Done()
			processRecord(record)
		}(append([]string(nil), record...)) // Pass a copy of the record slice
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
