package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
                        "k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MustHaveConfigSpec defines the desired state of MustHaveConfig
type MustHaveConfigSpec struct {
	Namespace string `json:"namespace"`

    // Key inside the ConfigMap data
    Key string `json:"key"`

    // Value to set for the key
    Value string `json:"value"`
}

// MustHaveConfigStatus defines the observed state of MustHaveConfig
type MustHaveConfigStatus struct {
	// Whether the ConfigMap is currently in the desired state
    Synced bool `json:"synced,omitempty"`

    // Optional message (e.g., error or status info)
    Message string `json:"message,omitempty"`

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MustHaveConfig is the Schema for the musthaveconfigs API
type MustHaveConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MustHaveConfigSpec   `json:"spec,omitempty"`
	Status MustHaveConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MustHaveConfigList contains a list of MustHaveConfig
type MustHaveConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MustHaveConfig `json:"items"`
}

func (in *MustHaveConfig) DeepCopyObject() runtime.Object {
    if in == nil {
        return nil
    }
    out := new(MustHaveConfig)
    *out = *in
    return out
}

// DeepCopyObject implements the runtime.Object interface for MustHaveConfigList.
func (in *MustHaveConfigList) DeepCopyObject() runtime.Object {
    if in == nil {
        return nil
    }
    out := new(MustHaveConfigList)
    *out = *in
    return out
}

func init() {
	SchemeBuilder.Register(&MustHaveConfig{}, &MustHaveConfigList{})
}
