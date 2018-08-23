package models
/*
import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)


const (
	HEADER_AUTH               = "Authorization"





	API_URL_SUFFIX			  = "/api/v1"

)


type Client1 struct {
	Url        string       // The location of the server, for example  "http://localhost:8065"
	ApiUrl     string       // The api location of the server, for example "http://localhost:8065/api/v4"
	HttpClient *http.Client // The http client
	AuthToken  string
	AuthType   string
	HttpHeader map[string]string // Headers to be copied over for each request
}



func NewAPIv1Client(url string) *Client1 {
	return &Client1{ url, url + API_URL_SUFFIX, &http.Client{}, "","",map[string]string{}}
}
func (c *Client1) LoginById(gc *gin.Context, id, password string) {

	var userDataForm User
	decoder := json.NewDecoder(gc.Request.Body)

	if err := decoder.Decode(&userDataForm); err != nil {
		WebApiErrorLog("POST_register converting json to structure made error", 8)
	}

	ctx, err := gc.Get("context")
	var ctx0 core.Server = ctx.(core.Server)

	if err != true {
		WebApiErrorLog("POST_register custom context is not defined", 8)
	}

	if rst := <-ctx0.Store.User().IsUserIdExist(userDataForm.Id); rst.Err != nil {


	}

}

func (c *Client1) GetAccountRoute() string {
	return fmt.Sprintf("/account")
}

func (c *Client1) GetUserExistenceRoute(userName string) string {
	return fmt.Sprintf(c.GetAccountRoute()+"/user/"+userName)
}


func closeBody(r *http.Response){
	if r.Body != nil {
		ioutil.ReadAll(r.Body)
		r.Body.Close()
	}

}

func (c *Client1) DoApiGet(url string) (*http.Response, *AppError){
	return c.DoApiRequest(http.MethodGet, c.ApiUrl+url,"")
}

func (c *Client1) DoApiRequest(method, url, data string) (*http.Response, *AppError){
	rq, _ := http.NewRequest(method, url, strings.NewReader(data))

	if len(c.AuthToken) > 0 {
		rq.Header.Set(HEADER_AUTH,"")
	}
	if rp, err := c.HttpClient.Do(rq); err != nil || rp == nil {
		return nil, NewAppError(url,"model.client1.connectig.app_error",nil,err.Error(),0)
	} else if rp.StatusCode == 304 {
		return rp,nil
	} else if rp.StatusCode >= 300 {
		defer closeBody(rp)
		return rp, AppErrorFromJson(rp.Body)
	}else{
		return rp, nil
	}
}

}

*/