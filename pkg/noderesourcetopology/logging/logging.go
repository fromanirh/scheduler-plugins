/*
Copyright 2021 The Kubernetes Authors.

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

package logging

import (
	"github.com/go-logr/logr"

	corev1 "k8s.io/api/core/v1"
)

// before to replace with FromContext(), at least in filter and score,
// we would need a way to inject a logger instance (preferably a
// per-plugin logger instance) when we create the Scheduler
// (with app.NewSchedulerCommand)

var logh logr.Logger

func SetLogger(lh logr.Logger) {
	logh = lh
}

func Log() logr.Logger {
	return logh
}

func PodLogID(pod *corev1.Pod) string {
	return pod.Namespace + "/" + pod.Name
}
