package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec,omitempty"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

type WorkspaceStatus struct {
	Status        string `json:"phase,omitempty"`
	OldGeneration int64  `json:"oldgeneration,omitempty"`
	Message       string `json:"message,omitempty"`
	Updatetime    string `json:"updatetime,omitempty"`
}

type WorkspaceSpec struct {
	Description    string `json:"description,omitempty"`
	WorkspaceQuota string `json:"workspacequota,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceList is a list of Workspace resources
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Workspace `json:"items"`
}
