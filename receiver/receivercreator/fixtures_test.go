// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package receivercreator

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

var pod = observer.Pod{
	UID:       "uid-1",
	Namespace: "default",
	Name:      "pod-1",
	Labels: map[string]string{
		"app":    "redis",
		"region": "west-1",
	},
	Annotations: map[string]string{
		"scrape": "true",
	},
}

var podEndpoint = observer.Endpoint{
	ID:      "pod-1",
	Target:  "localhost",
	Details: &pod,
}

var portEndpoint = observer.Endpoint{
	ID:     "port-1",
	Target: "localhost:1234",
	Details: &observer.Port{
		Name:      "http",
		Pod:       pod,
		Port:      1234,
		Transport: observer.ProtocolTCP,
	},
}

var hostportEndpoint = observer.Endpoint{
	ID:     "port-1",
	Target: "localhost:1234",
	Details: &observer.HostPort{
		ProcessName: "splunk",
		Command:     "./splunk",
		Port:        1234,
		Transport:   observer.ProtocolTCP,
	},
}

var container = observer.Container{
	Name:          "otel-agent",
	Image:         "otelcol",
	Port:          8080,
	AlternatePort: 80,
	Transport:     observer.ProtocolTCP,
	Command:       "otelcol -c",
	Host:          "localhost",
	ContainerID:   "abc123",
	Labels: map[string]string{
		"env":    "prod",
		"region": "east-1",
	},
}

var containerEndpoint = observer.Endpoint{
	ID:      "container-1",
	Target:  "localhost:1234",
	Details: &container,
}

var k8sNodeEndpoint = observer.Endpoint{
	ID:     "k8s.node-1",
	Target: "2.3.4.5",
	Details: &observer.K8sNode{
		Annotations: map[string]string{
			"node.alpha.kubernetes.io/ttl":                           "0",
			"volumes.kubernetes.io/controller-managed-attach-detach": "true",
		},
		ExternalDNS:         "an.external.dns",
		ExternalIP:          "1.2.3.4",
		Hostname:            "a.hostname",
		InternalDNS:         "an.internal.dns",
		InternalIP:          "2.3.4.5",
		KubeletEndpointPort: 10250,
		Labels: map[string]string{
			"kubernetes.io/arch": "amd64",
			"kubernetes.io/os":   "linux",
		},
		Name: "a.name",
		UID:  "b344f2a7-1ec1-40f0-8557-8a9bfd8b6f99",
	},
}

var unsupportedEndpoint = observer.Endpoint{
	ID:      "endpoint-1",
	Target:  "localhost:1234",
	Details: nil,
}
