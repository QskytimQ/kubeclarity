// Copyright © 2021 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
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

package common

const (
	KubeClarityImage                          = "ghcr.io/openclarity/kubeclarity"
	KubeClaritySbomDBImage                    = "ghcr.io/openclarity/kubeclarity-sbom-db"
	KubeClarityRuntimeK8sScannerImage         = "ghcr.io/openclarity/kubeclarity-runtime-k8s-scanner"
	KubeClarityCISDockerBenchmarkScannerImage = "ghcr.io/openclarity/kubeclarity-cis-docker-benchmark-scanner"

	KubeClarityPodLabels             = "app=kubeclarity-kubeclarity"
	KubeClarityServiceName           = "kubeclarity-kubeclarity"
	KubeClarityDeploymentName        = "kubeclarity-kubeclarity"
	KubeClarityNamespace             = "kubeclarity"
	KubeClarityHelmReleaseName       = "kubeclarity"
	KubeClarityPortForwardHostPort   = "8089"
	KubeClarityPortForwardTargetPort = "8080"
)
