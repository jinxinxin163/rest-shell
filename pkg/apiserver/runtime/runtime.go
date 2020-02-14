/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package runtime

import (
	"rest-shell/module/ensaas"
	"github.com/emicklei/go-restful"
	"net/http"
	"strings"
)

const (
	ApiRootPath = "/v1"
	ApiTag = "restshell"
)

// container holds all webservice of apiserver
var Container = restful.NewContainer()

type ContainerBuilder []func(c *restful.Container) error

const MimeMergePatchJson = "application/merge-patch+json"
const MimeJsonPatchJson = "application/json-patch+json"

func init() {
	restful.RegisterEntityAccessor(MimeMergePatchJson, restful.NewEntityAccessorJSON(restful.MIME_JSON))
	restful.RegisterEntityAccessor(MimeJsonPatchJson, restful.NewEntityAccessorJSON(restful.MIME_JSON))
}

func NewWebService() *restful.WebService {
	webservice := new(restful.WebService)
	webservice.Path(ApiRootPath + "/" + ApiTag).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	//do not do jwt validate
	webservice.Filter(jwtAuthentication)
	return webservice
}

func jwtAuthentication(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	tokenHeader := req.HeaderParameter("Authorization")
	//LOG.Info("token is:", tokenHeader)
	if tokenHeader == "" {
		resp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		resp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	ensaas.MpToken = tokenHeader
	chain.ProcessFilter(req, resp)
}