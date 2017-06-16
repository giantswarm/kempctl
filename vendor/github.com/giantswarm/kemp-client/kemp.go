package kempclient

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/juju/errgo"
)

type Config struct {
	User     string
	Password string
	Endpoint string
	Debug    bool
}

type Client struct {
	user     string
	password string
	endpoint string
	debug    bool
}

type ParameterResponse struct {
	Debug   string        `xml:",innerxml"`
	XMLName xml.Name      `xml:"Response"`
	Data    ParameterList `xml:"Success>Data"`
}

type ParameterList struct {
	Parameters []Parameter `xml:",any"`
}

type Parameter struct {
	XMLName xml.Name `xml:""`
	Value   string   `xml:",chardata"`
}

func NewClient(config Config) *Client {
	c := &Client{
		user:     config.User,
		password: config.Password,
		endpoint: config.Endpoint,
		debug:    config.Debug,
	}

	return c
}

func (c *Client) Get(param string) (string, error) {
	parameters := make(map[string]string)
	parameters["param"] = param

	data := ParameterResponse{}
	err := c.Request("get", parameters, &data)
	if err != nil {
		return "", errgo.NoteMask(err, fmt.Sprintf("kemp get '%s' failed", param), errgo.Any)
	}

	if c.debug {
		fmt.Println("DEBUG:", data.Debug)
	}

	result := make(map[string]string)
	for _, param := range data.Data.Parameters {
		result[param.XMLName.Local] = param.Value
	}

	return result[param], nil
}

func (c *Client) Set(param, value string) (string, error) {
	data, err := c.Get(param)
	if err != nil {
		return "", errgo.Mask(err)
	}

	parameters := make(map[string]string)
	parameters["param"] = param
	parameters["value"] = value
	err = c.Request("set", parameters, &ParameterResponse{})
	if err != nil {
		return "", errgo.NoteMask(err, fmt.Sprintf("kemp set '%s %s' failed", param, value), errgo.Any)
	}

	return data, nil
}

func (c *Client) Request(cmd string, parameters map[string]string, data interface{}) error {
	params := url.Values{}
	for key, val := range parameters {
		params.Set(key, val)
	}

	requestURL := fmt.Sprintf("%s%s?%s", c.endpoint, cmd, params.Encode())
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return errgo.NoteMask(err, fmt.Sprintf("kemp request to '%s' failed", requestURL), errgo.Any)
	}

	req.SetBasicAuth(c.user, c.password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Do(req)
	if err != nil {
		return errgo.NoteMask(err, fmt.Sprintf("kemp request to '%s' failed", requestURL), errgo.Any)
	}

	if res.StatusCode >= 400 {
		return c.parseError(res.StatusCode, res.Body)
	}

	return c.parseSuccess(res.Body, data)
}
