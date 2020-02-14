/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package ensaas

import (
	"rest-shell/module/httputil"
	extv1 "rest-shell/pkg/apis/workspace/v1"
	"rest-shell/pkg/utils/syslog"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"os"
)

const (
	DC_PATH = "/datacenter/"
	CLUSTER_PATH = "/cluster/"
)
var MpUrl string = os.Getenv("mp_url")
var MpToken string

func init(){
	LOG.Info("mpurl:", MpUrl)
	LOG.Info("mptoken:", MpToken)
}
type JsonNamespaceList struct {
	TotalCount int `json:totalCount"`
	Items []corev1.Namespace `json:"items"`
}

type JsonWorkspaceList struct {
	TotalCount int `json:totalCount"`
	Items []extv1.Workspace `json:"items"`
}

type JsonDsPodsList struct {
	TotalCount int `json:totalCount"`
	Items []corev1.Pod `json:"items"`
}

type JsonWkspName struct {
	Name string `json:name"`
}

type Source struct {
	Name string `json:name"`
	Id string 	`json:id"`
}
type JsonSourceList struct {
	TotalCount int `json:totalCount"`
	Items []Source `json:"items"`
}

type RoleBody struct {
	Role           string `json:"role,omitempty"`
	DataCenter     string `json:"datacenter,omitempty"`
	Cluster        string `json:"cluster,omitempty"`
	Workspace      string `json:"workspace,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
}

type RolebindingBody struct {
	User       string `json:"user,omitempty"`
	Roles      []*RoleBody `json:"roles,omitempty"`
	IsRollback bool   `json:"isRoleback"`
}

//operate rolebinding
func DeleteRolebinding(dc string, httpBody *RolebindingBody) (err error) {
	var url = MpUrl+DC_PATH+dc+"/rolebinding"
	LOG.Info("DeleteRolebinding url:", url)
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload,err := json_iterator.MarshalToString(httpBody)
	if err != nil {
		return err
	}
	LOG.Info("DeleteRolebinding payload:", payload)
	resp, status := httputil.HttpDelete(url, header, payload)
	fmt.Printf("DeleteRolebinding response::%+v\n", resp)
	if status == 200 {
		LOG.Info("DeleteRolebinding response:", resp)
		return nil
	} else {
		error := errors.New("DeleteRolebinding http operate fail")
		return error
	}
}

// get cluster or workspace by subid
func GetSources(dc string, subid string) (*JsonSourceList, error) {
	var url = MpUrl+DC_PATH+dc+"/clusternames?subscritionId="+subid
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("GetSources http operate fail")
		return nil, error
	}
	var sources = new(JsonSourceList)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, sources)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil, err
	}
	//LOG.Info("get configmap: ", configMap)
	return sources, nil
}

// operate Pod

func DeletePod(dc string, cluster string, LoggingNamespace string, PodName string) (error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+LoggingNamespace+"/pod/"+PodName
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	_, status := httputil.HttpDelete(url, header, "")
	if status == 200 {
		return nil
	} else {
		error := errors.New("DeletePod http operate fail")
		return error
	}
}

//operate Namespace
func ListAllNamespaces(dc string, cluster string) ([]corev1.Namespace) {
	//TODO
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespaces"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,_ := httputil.HttpGet(url, header, "")
	//LOG.Info("ListAllNamespaces, token is:", MpToken)
	//LOG.Info("debug response: ", status, response)
	var animals2 = new(JsonNamespaceList)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, animals2)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil
	}
	return animals2.Items
}

func ListAllNamespacesArray(dc string, cluster string) ([]string) {
	//TODO
	var namespaces []string
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespaces"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,_ := httputil.HttpGet(url, header, "")
	//LOG.Info("debug response: ", status, response)
	var animals2 = new(JsonNamespaceList)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, animals2)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil
	}
	for _, ns := range animals2.Items {
		namespaces = append(namespaces, ns.Name)
	}
	return  namespaces
}

// operate workspace
func ListAllWorkspaces(dc string, cluster string) ([]extv1.Workspace) {
	//TODO
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/workspaces"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,_ := httputil.HttpGet(url, header, "")
	//LOG.Info("debug response: ", status, response)
	var animals2 = new(JsonWorkspaceList)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, animals2)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil
	}
	return animals2.Items
}

func ListAllWorkspacesArray(dc string, cluster string) ([]string) {
	//TODO
	var workspaces []string
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/workspaces"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,_ := httputil.HttpGet(url, header, "")
	//LOG.Info("debug response: ", status, response)
	var animals2 = new(JsonWorkspaceList)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, animals2)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil
	}
	for _, ns := range animals2.Items {
		workspaces = append(workspaces, ns.Name)
	}
	return workspaces
}

func GetWorkspace(dc string, cluster string, workspaceId string) (*extv1.Workspace, error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/workspace/"+workspaceId
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("GetWorkspace http operate fail")
		return nil, error
	}
	var workspace = new(extv1.Workspace)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, workspace)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil, err
	}
	//LOG.Info("get configmap: ", configMap)
	return workspace, nil
}

func GetWorkspaceName(dc string, cluster string, workspaceId string) (string, error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/workspacename/"+workspaceId
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("GetWorkspace http operate fail")
		return "", error
	}
	var wkspWrap = new(JsonWkspName)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, wkspWrap)
	if err != nil {
		LOG.Error("json unmarshal error")
		return "", err
	}
	//LOG.Info("get configmap: ", configMap)
	if wkspWrap.Name == "" {
		error := errors.New("name is null")
		return "", error
	}
	return wkspWrap.Name, nil
}

//operate　ConfigMap
func GetConfigMap(dc string, cluster string, LoggingNamespace string, ConfigMapName string) (*corev1.ConfigMap, error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+LoggingNamespace+"/configmap/"+ConfigMapName
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("GetConfigMap http operate fail")
		return nil, error
	}
	var configMap = new(corev1.ConfigMap)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, configMap)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil, err
	}
	//LOG.Info("get configmap: ", configMap)
	return configMap, nil
}

func DeleteMap(dc string, cluster string, LoggingNamespace string, ConfigMapName string) (error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+LoggingNamespace+"/configmap/"+ConfigMapName
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	_, status := httputil.HttpDelete(url, header, "")
	if status == 200 {
		return nil
	} else {
		error := errors.New("DeleteMap http operate fail")
		return error
	}
}

func CreateConfigMap(dc string, cluster string, LoggingNamespace string, configMap *corev1.ConfigMap) (err error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+LoggingNamespace+"/configmap"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload,err := json_iterator.MarshalToString(configMap)
	if err != nil {
		return err
	}
	_,status := httputil.HttpPost(url, header, payload)
	if status == 200 {
		return nil
	} else {
		error := errors.New("CreateConfigMap http operate fail")
		return error
	}
}

func UpdateConfigMap(dc string, cluster string, LoggingNamespace string, ConfigMapName string, configMap *corev1.ConfigMap) (err error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+LoggingNamespace+"/configmap/"+ConfigMapName
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload,err := json_iterator.MarshalToString(configMap)
	if err != nil {
		return err
	}
	//LOG.Info("update configmap: ", configMap)
	_,status := httputil.HttpPut(url, header, payload)
	if status == 200 {
		return nil
	} else {
		error := errors.New("UpdateConfigMap http operate fail")
		return error
	}
}

// operate Daemonset
func GetDaemonsetPods(dc string, cluster string, LoggingNamespace string, DaemonsetName string) ([]corev1.Pod) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+LoggingNamespace+"/daemonsetpods/"+DaemonsetName
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,_ := httputil.HttpGet(url, header, "")
	//LOG.Info("ListAllNamespaces, token is:", MpToken)
	//LOG.Info("debug response: ", status, response)
	var animals2 = new(JsonDsPodsList)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, animals2)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil
	}
	return animals2.Items
}

func GetDaemonset(dc string, cluster string, namespace string, dsname string) (*v1beta1.DaemonSet, error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+namespace+"/daemonset/"+dsname
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("GetDaemonset http operate fail")
		return nil, error
	}
	var daemonset = new(v1beta1.DaemonSet)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, daemonset)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil, err
	}
	//LOG.Info("get configmap: ", configMap)
	return daemonset, nil
}

func DeleteDaemonset(dc string, cluster string, namespace string, dsname string) (error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+namespace+"/configmap/"+dsname
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	_, status := httputil.HttpDelete(url, header, "")
	if status == 200 {
		return nil
	} else {
		error := errors.New("DeleteDaemonset http operate fail")
		return error
	}
}

func CreateDaemonset(dc string, cluster string, namespace string, daemonst *v1beta1.DaemonSet) (err error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+namespace+"/daemonset"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload,err := json_iterator.MarshalToString(daemonst)
	if err != nil {
		return err
	}
	_,status := httputil.HttpPost(url, header, payload)
	if status == 200 {
		return nil
	} else {
		error := errors.New("CreateDaemonset http operate fail")
		return error
	}
}

// operate cronJob
func GetCronjob(dc string, cluster string, namespace string, cjname string) (*batchv1beta1.CronJob, error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+namespace+"/job/"+cjname
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	response,status := httputil.HttpGet(url, header, "")
	if status != 200 {
		error := errors.New("GetCronjob http operate fail")
		return nil, error
	}
	var cronjob = new(batchv1beta1.CronJob)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bResp := []byte(response)
	err := json_iterator.Unmarshal(bResp, cronjob)
	if err != nil {
		LOG.Error("json unmarshal error")
		return nil, err
	}
	//LOG.Info("get configmap: ", configMap)
	return cronjob, nil
}

func DeleteCronjob(dc string, cluster string, namespace string, cfname string) (error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+namespace+"/job/"+cfname
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	_, status := httputil.HttpDelete(url, header, "")
	if status == 200 {
		return nil
	} else {
		error := errors.New("DeleteCronjob http operate fail")
		return error
	}
}

func CreateCronjob(dc string, cluster string, namespace string, cronjob *batchv1beta1.CronJob) (err error) {
	var url = MpUrl+DC_PATH+dc+CLUSTER_PATH+cluster+"/namespace/"+namespace+"/job"
	var header map[string]string
	header = make(map[string]string)
	header["Content-type"] = "application/json"
	header["Authorization"] = MpToken
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	payload,err := json_iterator.MarshalToString(cronjob)
	if err != nil {
		return err
	}
	_,status := httputil.HttpPost(url, header, payload)
	if status == 200 {
		return nil
	} else {
		error := errors.New("CreateCronjob http operate fail")
		return error
	}
}