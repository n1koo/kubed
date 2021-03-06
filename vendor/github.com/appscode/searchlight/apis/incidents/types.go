/*
Copyright 2018 The Searchlight Authors.

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

package incidents

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:skipVerbs=get,list,update,patch,deleteCollection,watch
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Acknowledgement struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Request  AcknowledgementRequest
	Response AcknowledgementResponse
}

type AcknowledgementRequest struct {
	// Comment by user
	Comment string

	// Skip sending notification
	// +optional
	SkipNotify bool
}

type AcknowledgementResponse struct {
	// The time at which the acknowledgement was done.
	// +optional
	Timestamp metav1.Time
}
