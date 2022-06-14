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
// gen_go_data -package alpha -var YAML_tag_key blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/alpha/tag_key.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/alpha/tag_key.yaml
var YAML_tag_key = []byte("info:\n  title: CloudResourceManager/TagKey\n  description: The CloudResourceManager TagKey resource\n  x-dcl-struct-name: TagKey\n  x-dcl-has-create: true\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a TagKey\n    parameters:\n    - name: TagKey\n      required: true\n      description: A full instance of a TagKey\n  apply:\n    description: The function used to apply information about a TagKey\n    parameters:\n    - name: TagKey\n      required: true\n      description: A full instance of a TagKey\n  delete:\n    description: The function used to delete a TagKey\n    parameters:\n    - name: TagKey\n      required: true\n      description: A full instance of a TagKey\ncomponents:\n  schemas:\n    TagKey:\n      title: TagKey\n      x-dcl-id: tagKeys/{{name}}\n      x-dcl-has-iam: true\n      type: object\n      required:\n      - parent\n      - shortName\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Creation time.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. User-assigned description of the TagKey. Must not\n            exceed 256 characters. Read-write.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Optional. Entity tag which users can pass to prevent race conditions.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Immutable. The generated numeric id for the TagKey.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        namespacedName:\n          type: string\n          x-dcl-go-name: NamespacedName\n          readOnly: true\n          description: Output only. Immutable. Namespaced name of the TagKey.\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: Immutable. The resource name of the new TagKey's parent. Must\n            be of the form `organizations/{org_id}`.\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n        purpose:\n          type: string\n          x-dcl-go-name: Purpose\n          x-dcl-go-type: TagKeyPurposeEnum\n          description: 'Optional. A purpose denotes that this Tag is intended for\n            use in policies of a specific policy engine, and will involve that policy\n            engine in management operations involving this Tag. A purpose does not\n            grant a policy engine exclusive rights to the Tag, and it may be referenced\n            by other policy engines. A purpose cannot be changed once set. Possible\n            values: PURPOSE_UNSPECIFIED, GCE_FIREWALL'\n          x-kubernetes-immutable: true\n          enum:\n          - PURPOSE_UNSPECIFIED\n          - GCE_FIREWALL\n        purposeData:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: PurposeData\n          description: Optional. Purpose data corresponds to the policy system that\n            the tag is intended for. See documentation for `Purpose` for formatting\n            of this field. Purpose data cannot be changed once set.\n          x-kubernetes-immutable: true\n        shortName:\n          type: string\n          x-dcl-go-name: ShortName\n          description: Required. Immutable. The user friendly name for a TagKey. The\n            short name should be unique for TagKeys within the same tag namespace.\n            The short name must be 1-63 characters, beginning and ending with an alphanumeric\n            character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and\n            alphanumerics between.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Update time.\n          x-kubernetes-immutable: true\n")

// 4154 bytes
// MD5: ffaf39e64d05719b7e7067c789c540b0
