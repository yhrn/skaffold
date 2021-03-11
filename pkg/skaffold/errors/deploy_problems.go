/*
Copyright 2020 The Skaffold Authors

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

package errors

import (
	"fmt"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/cluster"
	"github.com/GoogleContainerTools/skaffold/proto/v1"
)

var (
	// for testing
	isMinikube = cluster.GetClient().IsMinikube
)

func suggestDeployFailedAction(cfg Config) []*proto.Suggestion {
	if isMinikube(cfg.GetKubeContext()) {
		command := "minikube status"
		if cfg.GetKubeContext() != "minikube" {
			command = fmt.Sprintf("minikube status -p %s", cfg.GetKubeContext())
		}
		return []*proto.Suggestion{{
			SuggestionCode: proto.SuggestionCode_CHECK_MINIKUBE_STATUS,
			Action:         fmt.Sprintf("Check if minikube is running using %q and try again.", command),
		}}
	}

	return []*proto.Suggestion{{
		SuggestionCode: proto.SuggestionCode_CHECK_CLUSTER_CONNECTION,
		Action:         fmt.Sprintf("Check your connection to the %q cluster", cfg.GetKubeContext()),
	}}
}
