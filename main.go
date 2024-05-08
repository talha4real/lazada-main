package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/asmcos/requests"
	"html"
	"net/http"
	"os"
	"regexp"
	"strconv"
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

func visitProductPage(session *requests.Request, csrfToken, productURL string) (string, string, string, string, string, string) {
	// Update session headers

	h := requests.Header{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"Referer":      productURL,
		"X-CSRF-TOKEN": csrfToken,
	}
	// Define the regular expression pattern
	pattern := `-s(\d+)\.`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find submatch in the URL
	matches := re.FindStringSubmatch(productURL)

	// Check if there is a match

	resp, err := session.Get(productURL, h)
	if err != nil {
		return "", "", "", "", "", ""
		//println(resp.Text())
	}

	// Parse HTML response using goquery
	document, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Text()))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return "", "", "", "", "", ""
	}

	// Extract product details
	productName := document.Find("h1.pdp-mod-product-badge-title").Text()
	if productName == "" {
		productName = "No product name found"
	}

	imageURL, _ := document.Find("img.pdp-mod-common-image.gallery-preview-panel__image").Attr("src")
	if imageURL == "" {
		imageURL = "No image URL found"
	}

	price := document.Find("span.pdp-price.pdp-price_type_normal.pdp-price_color_orange.pdp-price_size_xl").Text()
	if price == "" {
		priceScript := document.Find("script").FilterFunction(func(i int, s *goquery.Selection) bool {
			return regexp.MustCompile(`"pdt_price":"฿\d+\.\d+"`).MatchString(s.Text())
		}).Text()
		priceMatch := regexp.MustCompile(`"pdt_price":"(฿\d+\.\d+)"`).FindStringSubmatch(priceScript)
		if len(priceMatch) > 1 {
			price = priceMatch[1]
		} else {
			price = "Price not found"
		}
	}

	var itemID, skuID, sellerID string
	document.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptText := s.Text()
		if regexp.MustCompile(`skuId|itemId|sellerId`).MatchString(scriptText) {
			itemIDMatch := regexp.MustCompile(`"itemId":"(\d+)"`).FindStringSubmatch(scriptText)

			skuIDMatch := regexp.MustCompile(`"skuId":"(\d+)"`).FindStringSubmatch(scriptText)
			sellerIDMatch := regexp.MustCompile(`"sellerId":"(\d+)"`).FindStringSubmatch(scriptText)

			if len(itemIDMatch) > 1 {
				itemID = itemIDMatch[1]
			}
			if len(skuIDMatch) > 1 {
				if len(matches) > 1 {
					// Print the value between "-s" and "."
					skuID = matches[1]
					//fmt.Println("Value between -s and .:", matches[1])
				} else {
					skuID = skuIDMatch[0]

					//fmt.Println("No match found")
				}

			}
			if len(sellerIDMatch) > 1 {
				sellerID = sellerIDMatch[1]
			}
		}
	})

	return productName, imageURL, price, itemID, skuID, sellerID

}

func addToCart(session *requests.Request, item_id string, sku_id string, seller_id string, csrf_token string) bool {
	// API endpoint URL
	url := "https://cart.lazada.co.th/cart/api/add"
	// Define the JSON string for the payload

	payload := `{
    	"addItems": [
        		{
            	"itemId": "` + item_id + `",
            	"skuId": "` + sku_id + `",
            	"quantity": 1,
            	"extendInfo": {},
            	"attributes": {
                	"sellerId": "` + seller_id + `"
            	}
        	}
    	]
	}`

	h := requests.Header{
		"Content-Type": "application/json",
		"x-csrf-token": csrf_token,
		"Referer":      fmt.Sprintf("https://www.lazada.co.th/products/i%s-s%s.html", item_id, sku_id),
	}

	resp, err := session.PostJson(url, h, payload)
	if err == nil {
		var json map[string]interface{}
		resp.Json(&json)
		module, ok := json["module"].(map[string]interface{})
		if !ok {
			return false
		}

		status, ok := module["success"].(bool)
		if !ok {
			return false
		}
		if status {
			return true
		} else {
			return false
		}
	}
	return false
}

func proceedToCheckout(session *requests.Request, csrfToken string) (map[string]interface{}, error) {
	// URL for proceeding to the checkout page
	url := "https://checkout.lazada.co.th/shipping?spm=a2o42.cart.proceed_to_checkout.proceed_to_checkout"

	h := requests.Header{
		"Referer":      "https://www.lazada.co.th/cart",
		"X-CSRF-TOKEN": csrfToken,
	}

	resp, err := session.Get(url, h)
	if err != nil {
		return nil, err
	}
	//fmt.Println(resp.Text())
	// Extract initialization data using regex pattern
	pattern := regexp.MustCompile(`window\.__initData__\s*=\s*({.*?});`)
	match := pattern.FindStringSubmatch(resp.Text())
	// If match found, parse initialization data as JSON
	if len(match) > 1 {
		initdataJSON := match[1]
		var initData map[string]interface{}
		err := json.Unmarshal([]byte(initdataJSON), &initData)
		if err != nil {
			return nil, err
		}
		return initData, nil
	}

	return nil, fmt.Errorf("Initialization data not found")
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

func placeOrder(session *requests.Request, csrfToken, itemID, skuID string, params, xdata map[string]interface{}) string {
	// Fetch User-Agent value from a URL
	uaURL := "https://lazada-checkout-bol-production.up.railway.app/"
	uaResponse, err := session.Get(uaURL, nil)
	if err != nil {
		return ""
	}
	var xjson map[string]interface{}
	uaResponse.Json(&xjson)
	extString := xdata["module"].(map[string]interface{})["ext"].(string)
	//// Define a variable to hold the unmarshalled data
	var data map[string]interface{}
	// Unmarshal the JSON string into the map variable
	xerr := json.Unmarshal([]byte(string(extString)), &data)
	if xerr != nil {
		return ""
	}
	// Access the value associated with the key "x-umidtoken"
	umid, ok := data["x-umidtoken"].(string)
	if !ok {
		return ""
	}
	uaVal := xjson["uab_value"].(string)
	_ = uaVal
	_ = umid
	// Extract package and item values from params
	packageValues := make([]interface{}, 0)
	itemValues := make([]interface{}, 0)
	dataMap := params["module"].(map[string]interface{})["data"].(map[string]interface{})
	for key, value := range dataMap {
		if strings.HasPrefix(key, "package_") {
			packageValues = append(packageValues, value)
		}
		if strings.HasPrefix(key, "item_") {
			itemValues = append(itemValues, value)
		}
	}

	// Get the value of "id" from itemValues[0]
	packageitemID := ""
	if itemMap, ok := itemValues[0].(map[string]interface{}); ok {
		packageitemID = itemMap["id"].(string)
	}
	// Get the value of "id" from packageValues
	packageID := ""
	if len(packageValues) > 0 {
		if packageMap, ok := packageValues[0].(map[string]interface{}); ok {
			packageID = packageMap["id"].(string)
		}
	}

	//URL for placing order
	checkoutURL := "https://checkout.lazada.co.th/placeOrder"

	//// Set request headers
	h := requests.Header{
		"x-csrf-token":     csrfToken,
		"Referer":          "https://www.lazada.co.th/cart",
		"Content-Type":     "application/json",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
		"x-ua":             uaVal,
		"x-umidtoken":      umid,
	}

	// Convert signature to string
	signature, ok := params["module"].(map[string]interface{})["linkage"].(map[string]interface{})["signature"].(string)
	if !ok {
		// Handle the case where signature is not a string or doesn't exist
		return "" // or provide a default value or handle the error in another way
	}

	// Convert submitParams to string
	submitParams, ok := params["module"].(map[string]interface{})["linkage"].(map[string]interface{})["common"].(map[string]interface{})["submitParams"].(string)
	if !ok {
		// Handle the case where submitParams is not a string or doesn't exist
		return "" // or provide a default value or handle the error in another way
	}

	// Construct payload
	payload := fmt.Sprintf(`
{
    "hierarchy": {
        "structure": {
            "addressV2_DELIVERY#2": [],
            "container_10008": [
                "leftContainer_10009",
                "rightContainer_10010"
            ],
            "leftContainer_10009": [
                "addressV2_DELIVERY#2",
                "package_%s"
            ],
            "orderSummary_10005": [],
            "orderTotal_1": [],
            "package_%s": [
                "delivery_%s_delivery",
                "item_%s"
            ],
            "rightContainer_10010": [
                "paymentCard_11021",
                "voucherInput_1",
                "additionalInfo_11022",
                "additionalDetail_11023",
                "orderSummary_10005",
                "orderTotal_1"
            ],
            "root_10000": [
                "container_10008"
            ]
        }
    },
    "data": {},
    "linkage": {
        "common": {
            "compress": true,
            "submitParams": "%s"
        },
        "signature": "%s"
    }
}`, packageID, packageID, packageID, packageitemID, submitParams, signature)

	_ = payload

	// Send POST request to place order
	response, err := session.PostJson(checkoutURL, h, payload)
	if err == nil {
		var yjson map[string]interface{}
		response.Json(&yjson)

		decodedInput := html.UnescapeString(response.Text())

		var zdata map[string]interface{}
		zerr := json.Unmarshal([]byte(decodedInput), &zdata)
		if zerr != nil {
			fmt.Println("Error:", zerr)
			return ""
		}

		//Order Placed Successfully
		nextURL, ok := zdata["nextUrl"].(string)
		if ok {
			// The key "nextUrl" exists and its value is a string
			return nextURL
		} else {
			// The key "nextUrl" does not exist or its value is not a string
			// Handle the case where "nextUrl" does not exist or its value is not a string
			return "" // or any default value
		}
	}
	return ""
}

func sendDiscordNotification(webhookURL, checkoutOrderID, productName, imageURL, price, email, password, proxy string) {
	paymentURL := checkoutOrderID
	expirationTime := time.Now().Add(4320 * time.Minute).Unix()

	// Create the embed data
	embedData := map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       "Order Placed 🚀",
				"description": "Please complete your payment.",
				"Payment Url": paymentURL,
				"color":       15526829, // RGB value (239, 176, 69) converted to decimal
				"fields": []map[string]interface{}{
					{"name": "Store", "value": "Lazada - TH", "inline": true},
					{"name": "Product", "value": productName, "inline": true},
					{"name": "Product Price", "value": price, "inline": true},
					{"name": "Account", "value": email, "inline": true},
					{"name": "Password", "value": password, "inline": true},
					{"name": "Quantity", "value": "1", "inline": true},
					{"name": "Payment Link Expires", "value": fmt.Sprintf("<t:%d:R>", expirationTime), "inline": true},
					{"name": "Proxy Used", "value": proxy, "inline": false},
					{"name": "Complete Payment", "value": fmt.Sprintf("[Click to Pay 💵](%s)", paymentURL), "inline": false},
				},
				"thumbnail": map[string]interface{}{"url": imageURL},
				"footer":    map[string]interface{}{"text": "LazadaBot"},
			},
		},
	}

	// Marshal the embed data to JSON
	jsonData, err := json.Marshal(embedData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Send POST request to Discord webhook
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("Notification sent to Discord successfully.")
	} else {
		fmt.Printf("Failed to send notification to Discord: %d\n", resp.StatusCode)
	}
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
	productURL := record[5]

	interval := record[3]
	webhook := record[6]
	proxyURL := record[4]
	fmt.Println(name + ": " + "Initiating Session.")
	num, _ := strconv.Atoi(interval)
	session := requests.Requests()
	session.Proxy(proxyURL)
	csrfToken := initializeSession(session)

	loggedIn := login(session, csrfToken, email, password)
	if !loggedIn {
		fmt.Println(name + ": " + "Failed to login. Maybe check your email for verification")
		return
	}

	fmt.Println(name + ": " + "Logged in successfully.")
	fmt.Println(name + ": " + "Visiting Product URL.")

	// Call the function to visit the product page and extract details
	productName, imageURL, price, itemID, skuID, sellerID := visitProductPage(session, csrfToken, productURL)

	fmt.Println(productName, imageURL, price, itemID, skuID, sellerID)
	fmt.Println(name + ": " + "Monitoring Product.")

	for {
		if addToCart(session, itemID, skuID, sellerID, csrfToken) {
			fmt.Println(name + ": " + "Product successfully added to cart.")
			break // Exit the loop when addToCart returns true
		} else {
			fmt.Println(name + ": " + "Product out of stock. Monitoring..")
		}
		time.Sleep(time.Duration(num) * time.Second) // Sleep for 10 seconds before the next iteration
	}

	initData, err := proceedToCheckout(session, csrfToken)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	userdata, _ := getUserData(session, csrfToken)
	id := placeOrder(session, csrfToken, itemID, skuID, initData, userdata)
	if len(id) > 0 {
		fmt.Println(name + ": " + "Order successfully placed.")
		sendDiscordNotification(webhook, id, productName, imageURL, price, email, password, proxyURL)
	} else {
		fmt.Println(name + ": " + "Failed to place order.")
	}
	elapsed := time.Since(startTime)
	fmt.Println(name+": "+"Execution time: %s\n", elapsed)
}
func main() {
	file, err := os.Open("./tasks.csv")
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
