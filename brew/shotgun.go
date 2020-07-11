package brew

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)


const shotgunURL = "https://brazenanimation.shotgunstudio.com/api/v1"

type ShotgunAuth struct {
	ExpiresIn		int		`json:"expires_in"`
	RefreshToken	string	`json:"refresh_token"`
	AccessToken		string	`json:"access_token"`
	TokenType		string	`json:"token_type"`
}

type Project struct {
	ID			int		`json:"id"`
	Code 		string 	`json:"code"`
	Name		string 	`json:"name"`
	Description string 	`json:"sg_description"`
}

type GetAllEntityRequest struct {
	Filters 	[][]string 				`json:"filters"`
	Fields 		[]string 				`json:"fields"`
	Page 		*PaginationParameter 	`json:"page,omitempty"`
	SortKeys 	[]string 				`json:"sort,omitempty"`
}

type GetProjectsResponse struct {
	Data []ShotgunProjectData `json:"data"`
}

type ShotgunProjectData struct {
	ID 			int 						`json:"id"`
	Attributes 	ShotgunProjectAttributes 	`json:"attributes"`
}

type ShotgunProjectAttributes struct {
	Name 		string `json:"name"`
	Description string `json:"sg_description"`
	Code 		string `json:"code"`
}

type PaginationParameter struct {
	Size 	int `json:"size,omitempty"`
	Number 	int `json:"number,omitempty"`
}

type BrewConfiguration struct {
	RefreshToken string `json:"token"`
}

func GetConfig() *BrewConfiguration {
	usr, _ := user.Current()
	configPath := filepath.Join(usr.HomeDir, ".brazen", "brew-config.json")

	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			logrus.WithError(err).Error("failed to read Brew config")
			return nil
		}
		config := BrewConfiguration{}
		err = json.Unmarshal(data, &config)
		if err != nil {
			logrus.WithError(err).Error("failed to parse Brew config JSON data")
			return nil
		}

		return &config
	}

	return nil
}


func StoreRefreshToken(token string) {
	config := GetConfig()
	if config == nil {
		config = &BrewConfiguration{
			RefreshToken: token,
		}
	}

	usr, _ := user.Current()
	configDir := filepath.Join(usr.HomeDir, ".brazen")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.Mkdir(configDir, 0755)

	}

	configPath := filepath.Join(configDir, "brew-config.json")
	f, err := json.Marshal(*config)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal BrewConfiguration")
		return
	}

	err = ioutil.WriteFile(configPath, f, 0755)
	if err != nil {
		logrus.WithError(err).Error("failed to save BrewConfiguration to file")
		return
	}
}

//func ReAuthenticateUser() *ShotgunAuth {
//	authURL := shotgun_url + "/auth/access_token"
//
//	data := url.Values{}
//	data.Set("refresh_token", token)
//	data.Set("grant_type", "refresh_token")
//	requestBody := strings.NewReader(data.Encode())
//
//	req, _ := http.NewRequest("POST", authURL, requestBody)
//	req.Header.Add("Accept", "application/json")
//	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//
//	if err != nil {
//		logrus.WithError(err).Error("failed to authenticate user, unhandled exception")
//		return nil
//	}
//
//	if resp.StatusCode == 400 {
//		logrus.Error("failed to authenticate user, Bad Request")
//		return nil
//	}
//
//	responseBody, _ := ioutil.ReadAll(resp.Body)
//	a := ShotgunAuth{}
//	err = json.Unmarshal(responseBody, &a)
//	if err != nil {
//		logrus.WithError(err).Error("failed to unmarshal auth response body")
//		return nil
//	}
//
//	StoreRefreshToken(a.RefreshToken)
//
//	return &a
//}

func AuthenticateAsUser(username, password string) *ShotgunAuth {
	authURL := shotgunURL + "/auth/access_token"

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("grant_type", "password")
	requestBody := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", authURL, requestBody)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		logrus.WithError(err).Error("failed to authenticate user, unhandled exception")
		return nil
	}

	if resp.StatusCode == 400 {
		logrus.Error("failed to authenticate user, Bad Request")
		return nil
	}

	responseBody, _ := ioutil.ReadAll(resp.Body)
	a := ShotgunAuth{}
	err = json.Unmarshal(responseBody, &a)
	if err != nil {
		logrus.WithError(err).Error("failed to unmarshal auth response body")
		return nil
	}

	StoreRefreshToken(a.RefreshToken)

	return &a
}

func GetAllProjects(authToken, username string) []Project{
	searchURL := shotgunURL + "/entity/project/_search"
	body := GetAllEntityRequest{
		Filters: [][]string{
			{
				"users.HumanUser.login", "is", username,
			},
		},
		Fields: []string{
			"id", "name", "code", "sg_description",
		},
		Page: &PaginationParameter{
			Size:25,
		},
		SortKeys: []string{
			"-updated_at",
		},
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		logrus.Error("failed to create request body")
		return nil
	}
	data := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", searchURL, data)
	if err != nil {
		logrus.Error("failed to create request to get all projects")
		return nil
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/vnd+shotgun.api3_array+json")
	req.Header.Add("Authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("failed to request all projects")
		return nil
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	projectsResp := GetProjectsResponse{}
	err = json.Unmarshal(responseBody, &projectsResp)
	if err != nil {
		logrus.WithError(err).Error("failed to unmarshal shotgun response body")
	}

	var projects []Project
	for _, p := range projectsResp.Data {
		projects = append(projects, Project{
			ID: p.ID,
			Name: p.Attributes.Name,
			Code: p.Attributes.Code,
			Description: p.Attributes.Description,
		})
	}
	return projects
}