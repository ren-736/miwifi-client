package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// MiwifiClient ...
type MiwifiClient struct {
	ip    string
	token string
}

// MustDial tries to log in and then returns a MiwifiClient
// that can be used to modify configurations.
func MustDial(ip, username, password string) *MiwifiClient {

	client := &MiwifiClient{
		ip:    ip,
		token: "",
	}

	var loginInfo = &LoginInfo{}
	client.mustGetJSON(loginInfo, fmt.Sprintf("http://%s/cgi-bin/luci/api/xqsystem/login?username=%s&password=%s", ip, username, password))
	if loginInfo.TOKEN == "" {
		log.Fatal("login failed.")
	}
	client.token = loginInfo.TOKEN

	return client
}

func (c *MiwifiClient) settingURL(name string) string {
	return fmt.Sprintf("http://%s/cgi-bin/luci/admin/settings/%s", c.ip, name)
}

func (c *MiwifiClient) apiURL(name string) string {
	return fmt.Sprintf("http://%s//cgi-bin/luci/;stok=%s/api/%s", c.ip, c.token, name)
}

func (c *MiwifiClient) mustGet(u string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Fatalf("cannot get url: %s: %v\n", u, err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("cannot get url: %s: %v\n", u, err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("cannot get url: http status != 200: %v\n", resp.Status)
	}
	return resp
}

func (c *MiwifiClient) mustPostForm(u string, data map[string]interface{}) *http.Response {
	values := url.Values{}
	for key, value := range data {
		values.Set(key, fmt.Sprint(value))
	}
	values.Set("token", c.token)
	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatalf("cannot post url: %s: %v\n", u, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("cannot post url: %s: %v\n", u, err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("cannot post url: http status != 200: %v\n", resp.Status)
	}
	return resp
}

func (c *MiwifiClient) mustGetJSON(out interface{}, url string) {
	resp := c.mustGet(url)
	defer resp.Body.Close()
	if out != nil {
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(out); err != nil {
			log.Fatalf("cannot unmarshal json: %v\n", err)
		}
	}
}

func (c *MiwifiClient) mustPostFormGetJSON(out interface{}, url string, data map[string]interface{}) {
	resp := c.mustPostForm(url, data)
	defer resp.Body.Close()
	if out != nil {
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(out); err != nil {
			log.Fatalf("cannot unmarshal json: %v\n", err)
		}
	}
}

// ListPortMappings list port mappings.
func (c *MiwifiClient) ListPortMappings() []*PortMappingRule {
	var portForward PortForward
	c.mustGetJSON(&portForward, c.apiURL("xqnetwork/portforward?ftype=1"))
	rules := make([]*PortMappingRule, 0, len(portForward.List))
	for i := range portForward.List {
		rules = append(rules, &portForward.List[i])
	}
	return rules
}

// CreatePortMapping creates a port mapping.
//
// protocol: TCP, UDP, BOTH
func (c *MiwifiClient) CreatePortMapping(name string, protocol string, outerPort int, innerIP string, innerPort int) {
	var respCode RespCode
	c.mustPostFormGetJSON(&respCode, c.apiURL("xqnetwork/add_redirect"), map[string]interface{}{
		"name":  name,
		"proto": protocol,
		"sport": outerPort,
		"ip":    innerIP,
		"dport": innerPort,
	})
	if respCode.code != 0 {
		log.Fatalf("cannot create port mapping: %v\n", respCode.msg)
	}
}

// DeletePortMapping ...
func (c *MiwifiClient) DeletePortMapping(out_port string) {
	var respCode RespCode
	c.mustPostFormGetJSON(&respCode, c.apiURL("xqnetwork/delete_redirect"), map[string]interface{}{
		"port": out_port,
	})
	if respCode.code != 0 {
		log.Fatalf("cannot delete port mapping: %v\n", respCode.msg)
	}
}

// GetInformation ...
func (c *MiwifiClient) GetInformation() Information {
	var info _Information
	c.mustGetJSON(&info, c.apiURL(`xqsystem/information`))
	return info.ToInformation()
}

// ListDevices ...
func (c *MiwifiClient) ListDevices() (devices []*Device) {
	var deviceListInfo DeviceListInfo
	c.mustGetJSON(&deviceListInfo, c.apiURL(`misystem/devicelist`))
	for index := range deviceListInfo.Devices {
		devices = append(devices, &deviceListInfo.Devices[index])
	}
	return devices
}
