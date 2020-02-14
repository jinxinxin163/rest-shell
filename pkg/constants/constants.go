/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package constants

const (
	APIVersion = "v1alpha1"

	KubeSystemNamespace           = "kube-system"
	OpenPitrixNamespace           = "openpitrix-system"
	KubesphereDevOpsNamespace     = "kubesphere-devops-system"
	IstioNamespace                = "istio-system"
	KubeSphereMonitoringNamespace = "kubesphere-monitoring-system"
	KubeSphereLoggingNamespace    = "kubesphere-logging-system"
	KubeSphereNamespace           = "kubesphere-system"
	KubeSphereControlNamespace    = "kubesphere-controls-system"
	PorterNamespace               = "porter-system"
	IngressControllerNamespace    = KubeSphereControlNamespace
	AdminUserName                 = "admin"
	DataHome                      = "/etc/kubesphere"
	IngressControllerFolder       = DataHome + "/ingress-controller"
	IngressControllerPrefix       = "kubesphere-router-"

	WorkspaceLabelKey              = "workspace"
	System                         = "system"
	OpenPitrixRuntimeAnnotationKey = "openpitrix_runtime"
	WorkspaceAdmin                 = "workspace-admin"
	ClusterAdmin                   = "cluster-admin"
	WorkspaceRegular               = "workspace-regular"
	WorkspaceViewer                = "workspace-viewer"
	DevopsOwner                    = "owner"
	DevopsReporter                 = "reporter"

	UserNameHeader = "X-Token-Username"

	TenantResourcesTag         = "Tenant Resources"
	IdentityManagementTag      = "Identity Management"
	AccessManagementTag        = "Access Management"
	NamespaceResourcesTag      = "Namespace Resources"
	ClusterResourcesTag        = "Cluster Resources"
	ComponentStatusTag         = "Component Status"
	OpenpitrixTag              = "Openpitrix Resources"
	VerificationTag            = "Verification"
	RegistryTag                = "Docker Registry"
	UserResourcesTag           = "User Resources"
	DevOpsProjectTag           = "DevOps Project"
	DevOpsProjectCredentialTag = "DevOps Project Credential"
	DevOpsProjectMemberTag     = "DevOps Project Member"
	DevOpsPipelineTag          = "DevOps Pipeline"
	DevOpsWebhookTag           = "DevOps Webhook"
	DevOpsJenkinsfileTag       = "DevOps Jenkinsfile"
	DevOpsScmTag               = "DevOps Scm"
	ClusterMetricsTag          = "Cluster Metrics"
	NodeMetricsTag             = "Node Metrics"
	NamespaceMetricsTag        = "Namespace Metrics"
	PodMetricsTag              = "Pod Metrics"
	PVCMetricsTag              = "PVC Metrics"
	ContainerMetricsTag        = "Container Metrics"
	WorkloadMetricsTag         = "Workload Metrics"
	WorkspaceMetricsTag        = "Workspace Metrics"
	ComponentMetricsTag        = "Component Metrics"
	ManagerSetting             = "golang rest shell"
)

var (
	WorkSpaceRoles   = []string{WorkspaceAdmin, WorkspaceRegular, WorkspaceViewer}
	SystemNamespaces = []string{KubeSphereNamespace, KubeSphereLoggingNamespace, KubeSphereMonitoringNamespace, OpenPitrixNamespace, KubeSystemNamespace, IstioNamespace, KubesphereDevOpsNamespace, PorterNamespace}
)

