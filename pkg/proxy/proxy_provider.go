/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package proxy

import (
	"k8s.io/kubernetes/pkg/api"
)

// ProxyProvider is the interface provided by proxier implementations.
type ProxyProvider interface {
	// OnUpdate manages the active set of service proxies.
	// Active service proxies are reinitialized if found in the update set or
	// removed if missing from the update set.
	OnUpdate(services []api.Service)
	// SyncLoop runs periodic work.
	// This is expected to run as a goroutine or as the main loop of the app.
	// It does not return.
	SyncLoop()
}
