package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//pb "clustergarage.io/fim-proto/golang"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FimWatcher is a specification for a FimWatcher resource
type FimWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FimWatcherSpec   `json:"spec"`
	Status FimWatcherStatus `json:"status"`
}

// FimWatcherSpec is the spec for a FimWatcher resource
type FimWatcherSpec struct {
	Selector  *metav1.LabelSelector `json:"selector" protobuf:"bytes,1,opt,name=selector"`
	Subjects  []*FimWatcherSubject  `json:"subjects" protobuf:"bytes,2,opt,name=subjects"`
	LogFormat string                `json:"logFormat,omitempty" protobuf:"bytes,3,opt,name=logFormat"`
}

// FimWatcherSubject is the spec for a FimWatcherSubject resource
type FimWatcherSubject struct {
	Paths     []string `json:"paths" protobuf:"bytes,1,opt,name=paths"`
	Events    []string `json:"events" protobuf:"bytes,2,opt,name=events"`
	OnlyDir   bool     `json:"onlyDir,omitempty" protobuf:"bytes,3,opt,name=onlyDir"`
	Recursive bool     `json:"recursive,omitempty" protobuf:"bytes,4,opt,name=recursive"`
}

// FimWatcherStatus is the status for a FimWatcher resource
type FimWatcherStatus struct {
	ObservablePods int32 `json:"observablePods"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FimWatcherList is a list of FimWatcher resources
type FimWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []FimWatcher `json:"items"`
}
