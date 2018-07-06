package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the Puffin resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PuffinSpec defines the desired state of Puffin
type PuffinSpec struct {

	// This is the color of the puffin. Set this to any color string you like!
	Color string `json:"color,omitempty"`

}

// PuffinStatus defines the observed state of Puffin
type PuffinStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
	Message string `json:"message,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Puffin
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=puffins
type Puffin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PuffinSpec   `json:"spec,omitempty"`
	Status PuffinStatus `json:"status,omitempty"`
}
