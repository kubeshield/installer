/*
Copyright The Kubeshield Authors.

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

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindIdentityServer = "IdentityServer"
	ResourceIdentityServer     = "identityserver"
	ResourceIdentityServers    = "identityservers"
)

// IdentityServer defines the schama for identity server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=identityservers,singular=kubeshieldoperator,categories={kubeshield,appscode}
type IdentityServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              IdentityServerSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type ContianerRef struct {
	Registry   string `json:"registry" protobuf:"bytes,1,opt,name=registry"`
	Repository string `json:"repository" protobuf:"bytes,2,opt,name=repository"`
	Tag        string `json:"tag" protobuf:"bytes,3,opt,name=tag"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources" protobuf:"bytes,4,opt,name=resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext" protobuf:"bytes,5,opt,name=securityContext"`
}

// IdentityServerSpec is the schema for Identity Server values file
type IdentityServerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string       `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	ReplicaCount     int32        `json:"replicaCount" protobuf:"varint,3,opt,name=replicaCount"`
	Image            ContianerRef `json:"image" protobuf:"bytes,4,opt,name=image"`
	ImagePullPolicy  string       `json:"imagePullPolicy" protobuf:"bytes,5,opt,name=imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets" protobuf:"bytes,6,rep,name=imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon" protobuf:"varint,7,opt,name=criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel" protobuf:"varint,8,opt,name=logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,9,rep,name=annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations" protobuf:"bytes,10,rep,name=podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,11,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,12,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,13,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext" protobuf:"bytes,14,opt,name=podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount" protobuf:"bytes,15,opt,name=serviceAccount"`
	Apiserver          WebHookSpec              `json:"apiserver" protobuf:"bytes,16,opt,name=apiserver"`
	//+optional
	EnableAnalytics bool       `json:"enableAnalytics" protobuf:"varint,17,opt,name=enableAnalytics"`
	Monitoring      Monitoring `json:"monitoring" protobuf:"bytes,18,opt,name=monitoring"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create" protobuf:"varint,1,opt,name=create"`
	//+optional
	Name *string `json:"name" protobuf:"bytes,2,opt,name=name"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,3,rep,name=annotations"`
}

type WebHookSpec struct {
	GroupPriorityMinimum       int32           `json:"groupPriorityMinimum" protobuf:"varint,1,opt,name=groupPriorityMinimum"`
	VersionPriority            int32           `json:"versionPriority" protobuf:"varint,2,opt,name=versionPriority"`
	UseKubeapiserverFqdnForAks bool            `json:"useKubeapiserverFqdnForAks" protobuf:"varint,3,opt,name=useKubeapiserverFqdnForAks"`
	Healthcheck                HealthcheckSpec `json:"healthcheck" protobuf:"bytes,4,opt,name=healthcheck"`
	ServingCerts               ServingCerts    `json:"servingCerts" protobuf:"bytes,5,opt,name=servingCerts"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled" protobuf:"varint,1,opt,name=enabled"`
}

type ServingCerts struct {
	Generate bool `json:"generate" protobuf:"varint,1,opt,name=generate"`
	//+optional
	CaCrt string `json:"caCrt" protobuf:"bytes,2,opt,name=caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt" protobuf:"bytes,3,opt,name=serverCrt"`
	//+optional
	ServerKey string `json:"serverKey" protobuf:"bytes,4,opt,name=serverKey"`
}

type Monitoring struct {
	Agent string `json:"agent" protobuf:"bytes,1,opt,name=agent"`
	//+optional
	Server         bool                  `json:"server" protobuf:"varint,2,opt,name=server"`
	Prometheus     *PrometheusSpec       `json:"prometheus" protobuf:"bytes,3,opt,name=prometheus"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor" protobuf:"bytes,4,opt,name=serviceMonitor"`
}

type PrometheusSpec struct {
	//+optional
	Namespace string `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
}

type ServiceMonitorLabels struct {
	//+optional
	Labels map[string]string `json:"labels" protobuf:"bytes,1,rep,name=labels"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityServerList is a list of IdentityServers
type IdentityServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of IdentityServer CRD objects
	Items []IdentityServer `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
