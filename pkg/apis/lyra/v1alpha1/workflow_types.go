package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

////////////////////////////////////////////////////////////////
//
// NOTE
//
// If you make changes you might need to regenerate the controller code:
//
// operator-sdk generate k8s
//
////////////////////////////////////////////////////////////////

// WorkflowSpec defines the desired state of Workflow
type WorkflowSpec struct {
	WorkflowName string            `json:"workflowName"`
	Data         map[string]string `json:"data"`
	RefreshTime  int               `json:"refreshTime"`
}

// WorkflowStatus defines the observed state of Workflow
type WorkflowStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Workflow is the Schema for the workflows API
// +k8s:openapi-gen=true
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkflowSpec   `json:"spec,omitempty"`
	Status WorkflowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkflowList contains a list of Workflow
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}
