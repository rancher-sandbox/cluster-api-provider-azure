/*
Copyright 2023 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AzureManagedClusterTemplateSpec defines the desired state of AzureManagedClusterTemplate.
type AzureManagedClusterTemplateSpec struct {
	Template AzureManagedClusterTemplateResource `json:"template"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=azuremanagedclustertemplates,scope=Namespaced,categories=cluster-api,shortName=amct
// +kubebuilder:storageversion
// +kubebuilder:deprecatedversion:warning="AzureManagedClusterTemplate and the AzureManaged API are deprecated. Please migrate to infrastructure.cluster.x-k8s.io/v1beta1 AzureASOManagedClusterTemplate and related AzureASOManaged resources instead."

// AzureManagedClusterTemplate is the Schema for the AzureManagedClusterTemplates API.
type AzureManagedClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AzureManagedClusterTemplateSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// AzureManagedClusterTemplateList contains a list of AzureManagedClusterTemplates.
type AzureManagedClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureManagedClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureManagedClusterTemplate{}, &AzureManagedClusterTemplateList{})
}

// AzureManagedClusterTemplateResource describes the data needed to create an AzureManagedCluster from a template.
type AzureManagedClusterTemplateResource struct {
	Spec AzureManagedClusterTemplateResourceSpec `json:"spec"`
}
