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
// gen_go_data -package alpha -var YAML_managed_service blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/servicemanagement/alpha/managed_service.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/servicemanagement/alpha/managed_service.yaml
var YAML_managed_service = []byte("info:\n  title: Servicemanagement/ManagedService\n  description: The Servicemanagement ManagedService resource\n  x-dcl-struct-name: ManagedService\n  x-dcl-has-create: true\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a ManagedService\n    parameters:\n    - name: ManagedService\n      required: true\n      description: A full instance of a ManagedService\n  apply:\n    description: The function used to apply information about a ManagedService\n    parameters:\n    - name: ManagedService\n      required: true\n      description: A full instance of a ManagedService\n  delete:\n    description: The function used to delete a ManagedService\n    parameters:\n    - name: ManagedService\n      required: true\n      description: A full instance of a ManagedService\n  deleteAll:\n    description: The function used to delete all ManagedService\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ManagedService\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ManagedService:\n      title: ManagedService\n      x-dcl-id: services/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the service. See the overview for naming requirements.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: ID of the project that produces and owns this service.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n")

// 1917 bytes
// MD5: ffd37f7c0ba66497376ba1ea4cf181ea
