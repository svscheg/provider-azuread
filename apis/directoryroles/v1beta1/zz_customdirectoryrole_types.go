/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomDirectoryRoleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of the custom directory role.
	// The object ID of the directory role
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`
}

type CustomDirectoryRoleParameters struct {

	// The description of the custom directory role.
	// The description of the custom directory role
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the custom directory role.
	// The display name of the custom directory role
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Indicates whether the role is enabled for assignment.
	// Indicates whether the role is enabled for assignment
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// A collection of permissions blocks as documented below.
	// List of permissions that are included in the custom directory role
	// +kubebuilder:validation:Required
	Permissions []PermissionsParameters `json:"permissions" tf:"permissions,omitempty"`

	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// - The version of the role definition. This can be any arbitrary string between 1-128 characters.
	// The version of the role definition.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type PermissionsObservation struct {
}

type PermissionsParameters struct {

	// A set of tasks that can be performed on a resource. For more information, see the Permissions Reference documentation.
	// Set of tasks that can be performed on a resource
	// +kubebuilder:validation:Required
	AllowedResourceActions []*string `json:"allowedResourceActions" tf:"allowed_resource_actions,omitempty"`
}

// CustomDirectoryRoleSpec defines the desired state of CustomDirectoryRole
type CustomDirectoryRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomDirectoryRoleParameters `json:"forProvider"`
}

// CustomDirectoryRoleStatus defines the observed state of CustomDirectoryRole.
type CustomDirectoryRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomDirectoryRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDirectoryRole is the Schema for the CustomDirectoryRoles API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type CustomDirectoryRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDirectoryRoleSpec   `json:"spec"`
	Status            CustomDirectoryRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDirectoryRoleList contains a list of CustomDirectoryRoles
type CustomDirectoryRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomDirectoryRole `json:"items"`
}

// Repository type metadata.
var (
	CustomDirectoryRole_Kind             = "CustomDirectoryRole"
	CustomDirectoryRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomDirectoryRole_Kind}.String()
	CustomDirectoryRole_KindAPIVersion   = CustomDirectoryRole_Kind + "." + CRDGroupVersion.String()
	CustomDirectoryRole_GroupVersionKind = CRDGroupVersion.WithKind(CustomDirectoryRole_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomDirectoryRole{}, &CustomDirectoryRoleList{})
}
