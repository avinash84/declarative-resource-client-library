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
// gen_go_data -package alpha -var YAML_private_cloud blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/private_cloud.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vmware/alpha/private_cloud.yaml
var YAML_private_cloud = []byte("info:\n  title: Vmware/PrivateCloud\n  description: The Vmware PrivateCloud resource\n  x-dcl-struct-name: PrivateCloud\n  x-dcl-has-create: true\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a PrivateCloud\n    parameters:\n    - name: PrivateCloud\n      required: true\n      description: A full instance of a PrivateCloud\n  apply:\n    description: The function used to apply information about a PrivateCloud\n    parameters:\n    - name: PrivateCloud\n      required: true\n      description: A full instance of a PrivateCloud\n  delete:\n    description: The function used to delete a PrivateCloud\n    parameters:\n    - name: PrivateCloud\n      required: true\n      description: A full instance of a PrivateCloud\n  deleteAll:\n    description: The function used to delete all PrivateCloud\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many PrivateCloud\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    PrivateCloud:\n      title: PrivateCloud\n      x-dcl-id: projects/{{project}}/locations/{{location}}/privateClouds/{{name}}\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-has-iam: true\n      type: object\n      required:\n      - name\n      - networkConfig\n      - managementCluster\n      - project\n      - location\n      properties:\n        conditions:\n          type: array\n          x-dcl-go-name: Conditions\n          readOnly: true\n          description: Output only. The conditions that caused the current private\n            cloud state. For example, private cloud provisioning failure description.\n          x-kubernetes-immutable: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: PrivateCloudConditions\n            properties:\n              code:\n                type: string\n                x-dcl-go-name: Code\n                readOnly: true\n                description: Output only. Machine-readable representation of the condition.\n                x-kubernetes-immutable: true\n              message:\n                type: string\n                x-dcl-go-name: Message\n                readOnly: true\n                description: Output only. Human-readable description of the condition.\n                x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Creation time of this resource in RFC3339 text\n            format.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Output only. Time the resource was marked as deleted, in RFC3339\n            text format.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: User-provided description for this private cloud.\n        expireTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: ExpireTime\n          readOnly: true\n          description: Output only. Planned deletion time of this resource in RFC3339\n            text format.\n          x-kubernetes-immutable: true\n        hcx:\n          type: object\n          x-dcl-go-name: Hcx\n          x-dcl-go-type: PrivateCloudHcx\n          readOnly: true\n          description: Output only. HCX appliance.\n          x-kubernetes-immutable: true\n          properties:\n            externalIP:\n              type: string\n              x-dcl-go-name: ExternalIP\n              description: External IP address of the appliance.\n              x-kubernetes-immutable: true\n            fdqn:\n              type: string\n              x-dcl-go-name: Fdqn\n              description: Fully qualified domain name of the appliance.\n              x-kubernetes-immutable: true\n            internalIP:\n              type: string\n              x-dcl-go-name: InternalIP\n              description: Internal IP address of the appliance.\n              x-kubernetes-immutable: true\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Version of the appliance.\n              x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        managementCluster:\n          type: object\n          x-dcl-go-name: ManagementCluster\n          x-dcl-go-type: PrivateCloudManagementCluster\n          description: Input only. The management cluster for this private cloud.\n            This parameter is required during creation of private cloud to provide\n            details for the default cluster.\n          x-dcl-mutable-unreadable: true\n          required:\n          - clusterId\n          - nodeTypeId\n          - nodeCount\n          properties:\n            clusterId:\n              type: string\n              x-dcl-go-name: ClusterId\n              description: Required. The user-provided identifier of the new `Cluster`.\n              x-kubernetes-immutable: true\n            nodeCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: NodeCount\n              description: Required. Number of nodes in this cluster.\n            nodeTypeId:\n              type: string\n              x-dcl-go-name: NodeTypeId\n              description: 'Required. The canonical identifier of node types (`NodeType`)\n                in this cluster. For example: standard-72.'\n              x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The resource name of this private cloud. Resource names are\n            schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\n            For example: `projects/my-project/locations/us-west1-a/privateClouds/my-cloud`'\n          x-kubernetes-immutable: true\n        networkConfig:\n          type: object\n          x-dcl-go-name: NetworkConfig\n          x-dcl-go-type: PrivateCloudNetworkConfig\n          description: Required. Network configuration.\n          x-kubernetes-immutable: true\n          required:\n          - network\n          - managementCidr\n          properties:\n            managementCidr:\n              type: string\n              x-dcl-go-name: ManagementCidr\n              description: Required. Management CIDR used by VMWare management appliances.\n              x-kubernetes-immutable: true\n            network:\n              type: string\n              x-dcl-go-name: Network\n              description: 'Required. The relative resource name of the consumer VPC\n                network this private cloud is attached to. Specify the name in the\n                following form: `projects/{project}/global/networks/{network_id}`\n                where `{project}` can either be a project number or a project ID.'\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Compute/Network\n                field: name\n            serviceNetwork:\n              type: string\n              x-dcl-go-name: ServiceNetwork\n              readOnly: true\n              description: 'Output only. The relative resource name of the service\n                VPC network this private cloud is attached to. The name is specified\n                in the following form: `projects/{service_project_number}/global/networks/{network_id}`.'\n              x-kubernetes-immutable: true\n        nsx:\n          type: object\n          x-dcl-go-name: Nsx\n          x-dcl-go-type: PrivateCloudNsx\n          readOnly: true\n          description: Output only. NSX appliance.\n          x-kubernetes-immutable: true\n          properties:\n            externalIP:\n              type: string\n              x-dcl-go-name: ExternalIP\n              description: External IP address of the appliance.\n              x-kubernetes-immutable: true\n            fdqn:\n              type: string\n              x-dcl-go-name: Fdqn\n              description: Fully qualified domain name of the appliance.\n              x-kubernetes-immutable: true\n            internalIP:\n              type: string\n              x-dcl-go-name: InternalIP\n              description: Internal IP address of the appliance.\n              x-kubernetes-immutable: true\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Version of the appliance.\n              x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: PrivateCloudStateEnum\n          readOnly: true\n          description: 'Output only. State of the resource. Possible values: ACTIVE,\n            CREATING, UPDATING, FAILED, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - ACTIVE\n          - CREATING\n          - UPDATING\n          - FAILED\n          - DELETED\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Last update time of this resource in RFC3339 text\n            format.\n          x-kubernetes-immutable: true\n        vcenter:\n          type: object\n          x-dcl-go-name: Vcenter\n          x-dcl-go-type: PrivateCloudVcenter\n          readOnly: true\n          description: Output only. Vcenter appliance.\n          x-kubernetes-immutable: true\n          properties:\n            externalIP:\n              type: string\n              x-dcl-go-name: ExternalIP\n              description: External IP address of the appliance.\n              x-kubernetes-immutable: true\n            fdqn:\n              type: string\n              x-dcl-go-name: Fdqn\n              description: Fully qualified domain name of the appliance.\n              x-kubernetes-immutable: true\n            internalIP:\n              type: string\n              x-dcl-go-name: InternalIP\n              description: Internal IP address of the appliance.\n              x-kubernetes-immutable: true\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Version of the appliance.\n              x-kubernetes-immutable: true\n")

// 10997 bytes
// MD5: 2ad5b4b44f27e3a0f75cf1318bd9c249
