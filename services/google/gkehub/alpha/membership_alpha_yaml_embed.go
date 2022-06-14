// Copyright 2022 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_membership blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/alpha/membership.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/alpha/membership.yaml
var YAML_membership = []byte("info:\n  title: GkeHub/Membership\n  description: The GkeHub Membership resource\n  x-dcl-struct-name: Membership\n  x-dcl-has-create: true\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Membership\n    parameters:\n    - name: Membership\n      required: true\n      description: A full instance of a Membership\n  apply:\n    description: The function used to apply information about a Membership\n    parameters:\n    - name: Membership\n      required: true\n      description: A full instance of a Membership\n  delete:\n    description: The function used to delete a Membership\n    parameters:\n    - name: Membership\n      required: true\n      description: A full instance of a Membership\n  deleteAll:\n    description: The function used to delete all Membership\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Membership\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Membership:\n      title: Membership\n      x-dcl-id: projects/{{project}}/locations/{{location}}/memberships/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        authority:\n          type: object\n          x-dcl-go-name: Authority\n          x-dcl-go-type: MembershipAuthority\n          description: 'Optional. How to identify workloads from this Membership.\n            See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity'\n          properties:\n            identityProvider:\n              type: string\n              x-dcl-go-name: IdentityProvider\n              readOnly: true\n              description: Output only. An identity provider that reflects the `issuer`\n                in the workload identity pool.\n            issuer:\n              type: string\n              x-dcl-go-name: Issuer\n              description: Optional. A JSON Web Token (JWT) issuer URI. `issuer` must\n                start with `https://` and be a valid URL with length <2000 characters.\n                If set, then Google will allow valid OIDC tokens from this issuer\n                to authenticate within the workload_identity_pool. OIDC discovery\n                will be performed on this URI to validate tokens from the issuer.\n                Clearing `issuer` disables Workload Identity. `issuer` cannot be directly\n                modified; it must be cleared (and Workload Identity disabled) before\n                using a new issuer (and re-enabling Workload Identity).\n            workloadIdentityPool:\n              type: string\n              x-dcl-go-name: WorkloadIdentityPool\n              readOnly: true\n              description: 'Output only. The name of the workload identity pool in\n                which `issuer` will be recognized. There is a single Workload Identity\n                Pool per Hub that is shared between all Memberships that belong to\n                that Hub. For a Hub hosted in: {PROJECT_ID}, the workload pool format\n                is `{PROJECT_ID}.hub.id.goog`, although this is subject to change\n                in newer versions of this API.'\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. When the Membership was created.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Output only. When the Membership was deleted.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: 'Description of this membership, limited to 63 characters.\n            Must match the regex: `*` This field is present for legacy purposes.'\n        endpoint:\n          type: object\n          x-dcl-go-name: Endpoint\n          x-dcl-go-type: MembershipEndpoint\n          description: Optional. Endpoint information to reach this member.\n          properties:\n            gkeCluster:\n              type: object\n              x-dcl-go-name: GkeCluster\n              x-dcl-go-type: MembershipEndpointGkeCluster\n              description: Optional. GKE-specific information. Only present if this\n                Membership is a GKE cluster.\n              properties:\n                resourceLink:\n                  type: string\n                  x-dcl-go-name: ResourceLink\n                  description: 'Immutable. Self-link of the GCP resource for the GKE\n                    cluster. For example: //container.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster\n                    Zonal clusters are also supported.'\n                  x-dcl-references:\n                  - resource: Container/Cluster\n                    field: selfLink\n            kubernetesMetadata:\n              type: object\n              x-dcl-go-name: KubernetesMetadata\n              x-dcl-go-type: MembershipEndpointKubernetesMetadata\n              readOnly: true\n              description: Output only. Useful Kubernetes-specific metadata.\n              properties:\n                kubernetesApiServerVersion:\n                  type: string\n                  x-dcl-go-name: KubernetesApiServerVersion\n                  readOnly: true\n                  description: Output only. Kubernetes API server version string as\n                    reported by `/version`.\n                memoryMb:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: MemoryMb\n                  readOnly: true\n                  description: Output only. The total memory capacity as reported\n                    by the sum of all Kubernetes nodes resources, defined in MB.\n                nodeCount:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: NodeCount\n                  readOnly: true\n                  description: Output only. Node count as reported by Kubernetes nodes\n                    resources.\n                nodeProviderId:\n                  type: string\n                  x-dcl-go-name: NodeProviderId\n                  readOnly: true\n                  description: Output only. Node providerID as reported by the first\n                    node in the list of nodes on the Kubernetes endpoint. On Kubernetes\n                    platforms that support zero-node clusters (like GKE-on-GCP), the\n                    node_count will be zero and the node_provider_id will be empty.\n                updateTime:\n                  type: string\n                  format: date-time\n                  x-dcl-go-name: UpdateTime\n                  readOnly: true\n                  description: Output only. The time at which these details were last\n                    updated. This update_time is different from the Membership-level\n                    update_time since EndpointDetails are updated internally for API\n                    consumers.\n                vcpuCount:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: VcpuCount\n                  readOnly: true\n                  description: Output only. vCPU count as reported by Kubernetes nodes\n                    resources.\n            kubernetesResource:\n              type: object\n              x-dcl-go-name: KubernetesResource\n              x-dcl-go-type: MembershipEndpointKubernetesResource\n              description: 'Optional. The in-cluster Kubernetes Resources that should\n                be applied for a correctly registered cluster, in the steady state.\n                These resources: * Ensure that the cluster is exclusively registered\n                to one and only one Hub Membership. * Propagate Workload Pool Information\n                available in the Membership Authority field. * Ensure proper initial\n                configuration of default Hub Features.'\n              properties:\n                connectResources:\n                  type: array\n                  x-dcl-go-name: ConnectResources\n                  readOnly: true\n                  description: Output only. The Kubernetes resources for installing\n                    the GKE Connect agent This field is only populated in the Membership\n                    returned from a successful long-running operation from CreateMembership\n                    or UpdateMembership. It is not populated during normal GetMembership\n                    or ListMemberships requests. To get the resource manifest after\n                    the initial registration, the caller should make a UpdateMembership\n                    call with an empty field mask.\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: MembershipEndpointKubernetesResourceConnectResources\n                    properties:\n                      clusterScoped:\n                        type: boolean\n                        x-dcl-go-name: ClusterScoped\n                        description: Whether the resource provided in the manifest\n                          is `cluster_scoped`. If unset, the manifest is assumed to\n                          be namespace scoped. This field is used for REST mapping\n                          when applying the resource in a cluster.\n                      manifest:\n                        type: string\n                        x-dcl-go-name: Manifest\n                        description: YAML manifest of the resource.\n                membershipCrManifest:\n                  type: string\n                  x-dcl-go-name: MembershipCrManifest\n                  description: Input only. The YAML representation of the Membership\n                    CR. This field is ignored for GKE clusters where Hub can read\n                    the CR directly. Callers should provide the CR that is currently\n                    present in the cluster during CreateMembership or UpdateMembership,\n                    or leave this field empty if none exists. The CR manifest is used\n                    to validate the cluster has not been registered with another Membership.\n                  x-dcl-mutable-unreadable: true\n                membershipResources:\n                  type: array\n                  x-dcl-go-name: MembershipResources\n                  readOnly: true\n                  description: Output only. Additional Kubernetes resources that need\n                    to be applied to the cluster after Membership creation, and after\n                    every update. This field is only populated in the Membership returned\n                    from a successful long-running operation from CreateMembership\n                    or UpdateMembership. It is not populated during normal GetMembership\n                    or ListMemberships requests. To get the resource manifest after\n                    the initial registration, the caller should make a UpdateMembership\n                    call with an empty field mask.\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: MembershipEndpointKubernetesResourceMembershipResources\n                    properties:\n                      clusterScoped:\n                        type: boolean\n                        x-dcl-go-name: ClusterScoped\n                        description: Whether the resource provided in the manifest\n                          is `cluster_scoped`. If unset, the manifest is assumed to\n                          be namespace scoped. This field is used for REST mapping\n                          when applying the resource in a cluster.\n                      manifest:\n                        type: string\n                        x-dcl-go-name: Manifest\n                        description: YAML manifest of the resource.\n                resourceOptions:\n                  type: object\n                  x-dcl-go-name: ResourceOptions\n                  x-dcl-go-type: MembershipEndpointKubernetesResourceResourceOptions\n                  description: Optional. Options for Kubernetes resource generation.\n                  properties:\n                    connectVersion:\n                      type: string\n                      x-dcl-go-name: ConnectVersion\n                      description: Optional. The Connect agent version to use for\n                        connect_resources. Defaults to the latest GKE Connect version.\n                        The version must be a currently supported version, obsolete\n                        versions will be rejected.\n                    v1beta1Crd:\n                      type: boolean\n                      x-dcl-go-name: V1Beta1Crd\n                      description: Optional. Use `apiextensions/v1beta1` instead of\n                        `apiextensions/v1` for CustomResourceDefinition resources.\n                        This option should be set for clusters with Kubernetes apiserver\n                        versions <1.16.\n        externalId:\n          type: string\n          x-dcl-go-name: ExternalId\n          description: 'Optional. An externally-generated and managed ID for this\n            Membership. This ID may be modified after creation, but this is not recommended.\n            The ID must match the regex: `*` If this Membership represents a Kubernetes\n            cluster, this value should be set to the UID of the `kube-system` namespace\n            object.'\n        infrastructureType:\n          type: string\n          x-dcl-go-name: InfrastructureType\n          x-dcl-go-type: MembershipInfrastructureTypeEnum\n          description: 'Optional. The infrastructure type this Membership is running\n            on. Possible values: INFRASTRUCTURE_TYPE_UNSPECIFIED, ON_PREM, MULTI_CLOUD'\n          enum:\n          - INFRASTRUCTURE_TYPE_UNSPECIFIED\n          - ON_PREM\n          - MULTI_CLOUD\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. GCP labels for this membership.\n        lastConnectionTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: LastConnectionTime\n          readOnly: true\n          description: Output only. For clusters using Connect, the timestamp of the\n            most recent connection established with Google Cloud. This time is updated\n            every several minutes, not continuously. For clusters that do not use\n            GKE Connect, or that have never connected successfully, this field will\n            be unset.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. The full, unique name of this Membership resource\n            in the format `projects/*/locations/*/memberships/{membership_id}`, set\n            during creation. `membership_id` must be a valid RFC 1123 compliant DNS\n            label: 1. At most 63 characters in length 2. It must consist of lower\n            case alphanumeric characters or `-` 3. It must start and end with an alphanumeric\n            character Which can be expressed as the regex: `)?`, with a maximum length\n            of 63 characters.'\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: object\n          x-dcl-go-name: State\n          x-dcl-go-type: MembershipState\n          readOnly: true\n          description: Output only. State of the Membership resource.\n          x-kubernetes-immutable: true\n          properties:\n            code:\n              type: string\n              x-dcl-go-name: Code\n              x-dcl-go-type: MembershipStateCodeEnum\n              readOnly: true\n              description: 'Output only. The current state of the Membership resource.\n                Possible values: CODE_UNSPECIFIED, CREATING, READY, DELETING, UPDATING,\n                SERVICE_UPDATING'\n              x-kubernetes-immutable: true\n              enum:\n              - CODE_UNSPECIFIED\n              - CREATING\n              - READY\n              - DELETING\n              - UPDATING\n              - SERVICE_UPDATING\n        uniqueId:\n          type: string\n          x-dcl-go-name: UniqueId\n          readOnly: true\n          description: Output only. Google-generated UUID for this resource. This\n            is unique across all Membership resources. If a Membership resource is\n            deleted and another resource with the same name is created, it gets a\n            different unique_id.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. When the Membership was last updated.\n          x-kubernetes-immutable: true\n")

// 17804 bytes
// MD5: ea1d3a0ee12c016d2fc490f7b8c850fa
