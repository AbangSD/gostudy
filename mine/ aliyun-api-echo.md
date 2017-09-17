```go
package handler

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
  
  	"github.com/labstack/echo"
)

type Phone struct {
	Phone string `json:"phone"`
}

const (
	host    string = "http://jshmgsdmfb.market.alicloudapi.com"
	path    string = "/shouji/query"
	querys  string = "shouji="
	method  string = "GET"
	appcode string = "APPCODE ae401443f92c4cee9d06051dd6304881"
)

func QueryPhoneAttribution(c echo.Context) error {
	var (
		phone  Phone
		url    string
		err    error
		data   map[string]interface{}
		client = &http.Client{}
	)

	if err := c.Bind(&phone); err != nil {
		fmt.Println(err)

		return err
	}

	url = host + path + "?" + querys + phone.Phone

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
      	fmt.Println(err)
      
		return err
	}

	req.Header.Add("Authorization", appcode)

	resp, err := client.Do(req)
	if err != nil {
      	fmt.Println(err)
      
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)

	return c.JSON(0, &data)
}
```

