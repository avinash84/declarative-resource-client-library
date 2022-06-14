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
// gen_go_data -package beta -var YAML_feature blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/beta/feature.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/beta/feature.yaml
var YAML_feature = []byte("info:\n  title: GkeHub/Feature\n  description: The GkeHub Feature resource\n  x-dcl-struct-name: Feature\n  x-dcl-has-create: true\n  x-dcl-has-iam: false\n  x-dcl-mutex: '{{project}}/{{location}}/{{feature}}'\npaths:\n  get:\n    description: The function used to get information about a Feature\n    parameters:\n    - name: Feature\n      required: true\n      description: A full instance of a Feature\n  apply:\n    description: The function used to apply information about a Feature\n    parameters:\n    - name: Feature\n      required: true\n      description: A full instance of a Feature\n  delete:\n    description: The function used to delete a Feature\n    parameters:\n    - name: Feature\n      required: true\n      description: A full instance of a Feature\n  deleteAll:\n    description: The function used to delete all Feature\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Feature\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Feature:\n      title: Feature\n      x-dcl-id: projects/{{project}}/locations/{{location}}/features/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. When the Feature resource was created.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Output only. When the Feature resource was deleted.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: GCP labels for this Feature.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The full, unique name of this Feature resource\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        resourceState:\n          type: object\n          x-dcl-go-name: ResourceState\n          x-dcl-go-type: FeatureResourceState\n          readOnly: true\n          description: State of the Feature resource itself.\n          x-kubernetes-immutable: true\n          properties:\n            hasResources:\n              type: boolean\n              x-dcl-go-name: HasResources\n              readOnly: true\n              description: Whether this Feature has outstanding resources that need\n                to be cleaned up before it can be disabled.\n              x-kubernetes-immutable: true\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: FeatureResourceStateStateEnum\n              readOnly: true\n              description: 'The current state of the Feature resource in the Hub API.\n                Possible values: STATE_UNSPECIFIED, ENABLING, ACTIVE, DISABLING, UPDATING,\n                SERVICE_UPDATING'\n              x-kubernetes-immutable: true\n              enum:\n              - STATE_UNSPECIFIED\n              - ENABLING\n              - ACTIVE\n              - DISABLING\n              - UPDATING\n              - SERVICE_UPDATING\n        spec:\n          type: object\n          x-dcl-go-name: Spec\n          x-dcl-go-type: FeatureSpec\n          description: Optional. Hub-wide Feature configuration. If this Feature does\n            not support any Hub-wide configuration, this field may be unused.\n          properties:\n            multiclusteringress:\n              type: object\n              x-dcl-go-name: Multiclusteringress\n              x-dcl-go-type: FeatureSpecMulticlusteringress\n              description: Multicluster Ingress-specific spec.\n              required:\n              - configMembership\n              properties:\n                configMembership:\n                  type: string\n                  x-dcl-go-name: ConfigMembership\n                  description: 'Fully-qualified Membership name which hosts the MultiClusterIngress\n                    CRD. Example: `projects/foo-proj/locations/global/memberships/bar`'\n                  x-dcl-references:\n                  - resource: Gkehub/Membership\n                    field: name\n        state:\n          type: object\n          x-dcl-go-name: State\n          x-dcl-go-type: FeatureState\n          readOnly: true\n          description: Output only. The Hub-wide Feature state\n          x-kubernetes-immutable: true\n          properties:\n            state:\n              type: object\n              x-dcl-go-name: State\n              x-dcl-go-type: FeatureStateState\n              readOnly: true\n              description: Output only. The \"running state\" of the Feature in this\n                Hub.\n              x-kubernetes-immutable: true\n              properties:\n                code:\n                  type: string\n                  x-dcl-go-name: Code\n                  x-dcl-go-type: FeatureStateStateCodeEnum\n                  readOnly: true\n                  description: 'The high-level, machine-readable status of this Feature.\n                    Possible values: CODE_UNSPECIFIED, OK, WARNING, ERROR'\n                  x-kubernetes-immutable: true\n                  enum:\n                  - CODE_UNSPECIFIED\n                  - OK\n                  - WARNING\n                  - ERROR\n                description:\n                  type: string\n                  x-dcl-go-name: Description\n                  readOnly: true\n                  description: A human-readable description of the current status.\n                  x-kubernetes-immutable: true\n                updateTime:\n                  type: string\n                  x-dcl-go-name: UpdateTime\n                  readOnly: true\n                  description: 'The time this status and any related Feature-specific\n                    details were updated. A timestamp in RFC3339 UTC \"Zulu\" format,\n                    with nanosecond resolution and up to nine fractional digits. Examples:\n                    \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\"'\n                  x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. When the Feature resource was last updated.\n          x-kubernetes-immutable: true\n")

// 7244 bytes
// MD5: 30c170fb407be835d3caa63f6a1db396
