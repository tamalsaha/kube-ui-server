package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WhoAmISpec defines the desired state of WhoAmI
type WhoAmISpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of WhoAmI. Edit whoami_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// WhoAmIStatus defines the observed state of WhoAmI
type WhoAmIStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WhoAmI is the Schema for the whoamis API
type WhoAmI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WhoAmISpec   `json:"spec,omitempty"`
	Status WhoAmIStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WhoAmIList contains a list of WhoAmI
type WhoAmIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WhoAmI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WhoAmI{}, &WhoAmIList{})
}
