/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package v1alpha2

import (
	"rest-shell/pkg/apiserver/restshell"
	"rest-shell/pkg/apiserver/runtime"
	"rest-shell/pkg/constants"
	out "rest-shell/pkg/models/manager"
	"rest-shell/pkg/utils/syslog"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"net/http"
)

const (
	RespOK    = "ok"
)

func AddWebService() error {
	ws := runtime.NewWebService()

	//status
	ws.Route(ws.GET("/status").To(restshell.GetStatus)).Doc("Get the restshell server status")

	//run bash shell command
	ws.Route(ws.POST("/bash").To(restshell.RunBashCmd).
		Doc("Run bash shell command.").
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.ManagerSetting}).
		Reads(restshell.ShellBody{}).
		Writes(out.OutputsResult{}).
		Returns(http.StatusOK, RespOK, out.OutputsResult{})).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON)

	//run sh shell command
	ws.Route(ws.POST("/sh").To(restshell.RunShCmd).
		Doc("Run sh shell command.").
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.ManagerSetting}).
		Reads(restshell.ShellBody{}).
		Writes(out.OutputsResult{}).
		Returns(http.StatusOK, RespOK, out.OutputsResult{})).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON)

	LOG.Info("add rest router end")
	restful.DefaultContainer.Add(ws)

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	LOG.Info("Get the API using http://ip:port/apidocs.json")

	//path := os.Getenv("GOPATH") + "/src/rest-shell/docs/swagger"
	path := "/restshellservice/app/swagger"
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir(path))))
	http.ListenAndServe(":8080", nil)
	return nil
}

// refer to  https://github.com/emicklei/go-restful/blob/master/examples/restful-user-resource.go
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "MpbuyService",
			Description: "Resource for Manage Cluster/Workspace/Subscription",
			Contact: &spec.ContactInfo{
				Name:  "jin.xin",
				Email: "jin.xin@advantech.com.cn",
				URL:   "http://www.advantech.com.cn",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "restshell",
		Description: "buy source"}}}
}
