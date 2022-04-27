/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// UsagePlanKeyParameters defines the desired state of UsagePlanKey
type UsagePlanKeyParameters struct {
	// Region is which region the UsagePlanKey will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// [Required] The identifier of a UsagePlanKey resource for a plan customer.
	// +kubebuilder:validation:Required
	KeyID *string `json:"keyID"`
	// [Required] The type of a UsagePlanKey resource for a plan customer.
	// +kubebuilder:validation:Required
	KeyType                      *string `json:"keyType"`
	CustomUsagePlanKeyParameters `json:",inline"`
}

// UsagePlanKeySpec defines the desired state of UsagePlanKey
type UsagePlanKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       UsagePlanKeyParameters `json:"forProvider"`
}

// UsagePlanKeyObservation defines the observed state of UsagePlanKey
type UsagePlanKeyObservation struct {
	// The Id of a usage plan key.
	ID *string `json:"id,omitempty"`
	// The name of a usage plan key.
	Name *string `json:"name,omitempty"`
	// The type of a usage plan key. Currently, the valid key type is API_KEY.
	Type *string `json:"type_,omitempty"`
	// The value of a usage plan key.
	Value *string `json:"value,omitempty"`
}

// UsagePlanKeyStatus defines the observed state of UsagePlanKey.
type UsagePlanKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          UsagePlanKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanKey is the Schema for the UsagePlanKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UsagePlanKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UsagePlanKeySpec   `json:"spec"`
	Status            UsagePlanKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsagePlanKeyList contains a list of UsagePlanKeys
type UsagePlanKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsagePlanKey `json:"items"`
}

// Repository type metadata.
var (
	UsagePlanKeyKind             = "UsagePlanKey"
	UsagePlanKeyGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsagePlanKeyKind}.String()
	UsagePlanKeyKindAPIVersion   = UsagePlanKeyKind + "." + GroupVersion.String()
	UsagePlanKeyGroupVersionKind = GroupVersion.WithKind(UsagePlanKeyKind)
)

func init() {
	SchemeBuilder.Register(&UsagePlanKey{}, &UsagePlanKeyList{})
}