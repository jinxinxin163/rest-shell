package ssov3

import (
	"rest-shell/module/httputil"
	"rest-shell/pkg/utils/syslog"
	"errors"
	"github.com/json-iterator/go"
	"os"
)

const (
	TestUser = "xxx"
	TestPwd = "xxx"
)
//for test
//var SSoUrl string = "http://api-sso-master.es.wise-paas.cn/v4.0/"
// for production
var SSoUrl string = os.Getenv("sso_url")
func init(){
	LOG.Info("ssourl:", SSoUrl)
}

//from ssov3/pkg/models/user.go
const (
	DATACENTERADMIN  = "datacenterAdmin"
	CLUSTERADMIN = "clusterAdmin"
	CLUSTEROWNER = "clusterOwner"
	WORKSPACEOWNER = "workspaceOwner"
	DEVELOPER = "namespaceDeveloper"
)
type TokenPackage struct {
	TokenType string `json:"tokenType,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	ExpiresIn int32 `json:"expiresIn,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type UserRole struct {
	RoleName         string `json:"roleName"`
	Workspace        string `json:"workspace"`
	Namespace        string `json:"namespace"`
	Cluster          string `json:"cluster"`
	Datacenter       string `json:"datacenter"`
	SubscriptionName string `json:"SubscriptionName"`
	SubscriptionId   string `json:"SubscriptionId"`
}

type UserSubscriptionDto struct {
	SubscriptionId    string `json:"subscriptionId"`
	SubscriptionName       string `json:"subscriptionName"  xorm:"name"`
	Subscription_role string `json:"subscriptionRole"`
	Company    string `json:"company"`
}


type SignedInUser struct {
	Id                          string          `json:"id"`
	Address                     string          `json:"address"`
	Alternative_email           string          `json:"alternativeEmail"`
	Avatar                      string          `json:"avatar"`
	City                        string          `json:"city"`
	Company                     string          `json:"company"`
	Contact_phone               string          `json:"contactPhone"`
	Country                     string          `json:"country"`
	Creation_time               int64           `json:"creationTime"`
	Creator                     string          `json:"creator"`
	Expiration_time             int64           `json:"expirationTime"`
	First_name                  string          `json:"firstName"`
	Groups                      []string     `json:"groups"`
	Industry                    string          `json:"industry"`
	Last_city                   string          `json:"lastCity"`
	Last_ip                     string          `json:"lastIp"`
	Last_lat                    float64         `json:"lastLat"`
	Last_long                   float64         `json:"lastLong"`
	Last_modified_pwd_time      int64           `json:"lastModifiedPwdTime"`
	Last_modified_time          int64           `json:"lastModifiedTime"`
	Last_name                   string          `json:"lastName"`
	Last_signed_in_time         int64           `json:"lastSignedInTime"`
	Mobile_phone                string          `json:"mobilePhone"`
	Office_phone                string          `json:"officePhone"`
	Origin                      string          `json:"origin"`
	Postal_code                 string          `json:"postalCode"`
	Role                        string          `json:"role"`
	Signed_in_frequency_counter int64           `json:"signedInFrequencyCounter"`
	Username                    string          `json:"username"`
	Total_signed_in_times       int64           `json:"totalSignedInTimes"`
	Salt                        string          `json:"salt"`
	Roles                       []*UserRole     `json:"roles"`
	UserSubscriptions           []*UserSubscriptionDto `json:"userSubscriptions"`
	Scopes                      []string        `json:"scopes"`
	Status                      string          `json:"status"`
}


type CreateSrpForm struct {
	Name        string   `json:"name" binding:"Required"`
	WorkspaceId string   `json:"workspaceId" binding:"Required"`
	NamespaceId string   `json:"namespaceId" binding:"Required"`
	Scopes      []string `json:"scopes"`
	SrpAppId string   `json:"srpAppId"  binding:"Required"`
	SrpAppName string `json:"srpAppName" binding:"Required"`
	ClusterId   string   `json:"clusterId"  binding:"Required"`
}

type SsoSrp struct {
	SrpId            string      `json:"srpId"`
	CreationTime     int64       `json:"creationTime"`
	LastModifiedTime int64       `json:"lastModifiedTime"`
	Name             string      `json:"name"`
	Workspace        string      `json:"workspace"`
	WorkspaceId      string      `json:"workspaceId"`
	Cluster          string      `json:"cluster"`
	Scope            []string `json:"scope"`
	ClusterId        string      `json:"clusterId"`
	Namespace        string      `json:"namespace"`
	NamespaceId      string      `json:"namespaceId"`
	SrpSecret        string      `json:"srpSecret"`
	SrpAppId         string      `json:"srpAppId"`
	SrpAppName       string      `json:"srpAppName"`
}

type SubscriptionUserDto struct {
	SubscriptionId    string `json:"subscriptionId"`
	Name       string `json:"subscriptionName"`
	UserId     string `json:"userId"`
	Username   string `json:"userName"`
	Subscription_role string `json:"subscriptionRole"`
	Company    string `json:"company"`
	Status     string `json:"status"`
	LastSignedInTime int64 `json:"lastSignedInTime"`
}

type SubscriptionUsers struct {
	Resources     []*SubscriptionUserDto `json: "resources"`
	Total		int `json: "total"`
}



func GetEIToken(username string, passwd string) string {
	if username == "" || passwd == "" {
		LOG.Error("username or passwd is null")
		return ""
	}
	var url = SSoUrl+"auth/native"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["accept"] = "application/json"
	userPayload := make(map[string]string)
	userPayload["username"] = username
	userPayload["password"] = passwd
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload, err := json_iterator.MarshalToString(userPayload)
	if err != nil {
		LOG.Error("json marshal fail", err)
		return ""
	}
	//LOG.Info("payload: ", payload)
	response,status := httputil.HttpPost(url, header, payload)
	if status == 200 {
		//LOG.Info("resp: ", response)
		var tokenPackage = new(TokenPackage)
		bResp := []byte(response)
		error := json_iterator.Unmarshal(bResp, tokenPackage)
		//fmt.Printf("value:%+v\n", tokenPackage)
		if error != nil {
			return ""
		} else {
			return tokenPackage.AccessToken
		}

	} else {
		return ""
	}
}

func GetUsersMe(token string) (*SignedInUser, error) {
	var url = SSoUrl+"users/me"
	var header map[string]string
	header = make(map[string]string)
	authString := "Bearer " + token
	header["Content-type"] = "application/json"
	header["Authorization"] = authString
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("get user info fail")
		return nil, error
	}
	var userInfo = new(SignedInUser)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	ret := json_iterator.Unmarshal(bResp, userInfo)
	//fmt.Printf("userInfo:%+v\n", userInfo)
	if ret != nil {
		error := errors.New("json unmarshal fail")
		return nil, error
	} else {
		return userInfo, nil
	}
}

func GetUsersByName(token string, userName string) (*SignedInUser, error) {
	var url = SSoUrl+"users/"+userName
	var header map[string]string
	header = make(map[string]string)
	authString := "Bearer " + token
	header["Content-type"] = "application/json"
	header["Authorization"] = authString
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("get user info fail")
		return nil, error
	}
	var userInfo = new(SignedInUser)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	ret := json_iterator.Unmarshal(bResp, userInfo)
	//fmt.Printf("userInfo:%+v\n", userInfo)
	if ret != nil {
		error := errors.New("json unmarshal fail")
		return nil, error
	} else {
		return userInfo, nil
	}
}

// get users ids from subid
func GetUsersIdBySubid(token string, subscription string) []string {
	var url = SSoUrl+"subscriptions/"+subscription+"/users"
	var header map[string]string
	header = make(map[string]string)
	authString := "Bearer " + token
	header["Content-type"] = "application/json"
	header["Authorization"] = authString
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		return nil
	}
	var users = new(SubscriptionUsers)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	ret := json_iterator.Unmarshal(bResp, users)
	//fmt.Printf("userInfo:%+v\n", userInfo)
	if ret != nil {
		return nil
	} else {
		if users.Total <= 0 {
			return nil
		} else {
			var strUsers []string
			for _, user := range users.Resources {
				strUsers = append(strUsers, user.UserId)
			}
			return strUsers
		}
	}
}

func GetUsersBySubid(token string, subscription string) []*SubscriptionUserDto {
	var url = SSoUrl+"subscriptions/"+subscription+"/users"
	var header map[string]string
	header = make(map[string]string)
	authString := "Bearer " + token
	header["Content-type"] = "application/json"
	header["Authorization"] = authString
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		return nil
	}
	var users = new(SubscriptionUsers)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	ret := json_iterator.Unmarshal(bResp, users)
	//fmt.Printf("userInfo:%+v\n", userInfo)
	if ret != nil {
		return nil
	} else {
		if users.Total <= 0 {
			return nil
		} else {
			return users.Resources
		}
	}
}

func CreateSrp(token string, srpForm CreateSrpForm) (*SsoSrp, error){
	var url = SSoUrl+"clients"
	var header map[string]string
	header = make(map[string]string)
	authString := token
	header["Content-type"] = "application/json"
	header["X-Auth-SRPToken"] = authString
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload, _:= json_iterator.MarshalToString(srpForm)
	response,status := httputil.HttpPost(url, header, payload)
	if status != 200 {
		error := errors.New("create srp fail")
		return nil, error
	}
	var srpInfo = new(SsoSrp)
	bResp := []byte(response)
	ret := json_iterator.Unmarshal(bResp, srpInfo)
	//fmt.Printf("userInfo:%+v\n", userInfo)
	if ret != nil {
		error := errors.New("json unmarshal fail")
		return nil, error
	} else {
		return srpInfo, nil
	}
}

func GetDeviceToken(client_id string, client_secret string, grant_type string) string {
	var url = SSoUrl+"oauth/token"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	userPayload := make(map[string]string)
	userPayload["grant_type"] = grant_type
	userPayload["client_id"] = client_id
	userPayload["client_secret"] = client_secret
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload, _:= json_iterator.MarshalToString(userPayload)
	//LOG.Info("payload: ", payload)
	response,status := httputil.HttpPost(url, header, payload)
	if status == 200 {
		//LOG.Info("resp: ", response)
		var tokenPackage = new(TokenPackage)
		bResp := []byte(response)
		error := json_iterator.Unmarshal(bResp, tokenPackage)
		//fmt.Printf("value:%+v\n", tokenPackage)
		if error != nil {
			return ""
		} else {
			return tokenPackage.AccessToken
		}

	} else {
		return ""
	}
}
