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
// gen_go_data -package beta -var YAML_server_tls_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networksecurity/beta/server_tls_policy.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networksecurity/beta/server_tls_policy.yaml
var YAML_server_tls_policy = []byte("info:\n  title: NetworkSecurity/ServerTlsPolicy\n  description: The NetworkSecurity ServerTlsPolicy resource\n  x-dcl-struct-name: ServerTlsPolicy\n  x-dcl-has-create: true\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a ServerTlsPolicy\n    parameters:\n    - name: ServerTlsPolicy\n      required: true\n      description: A full instance of a ServerTlsPolicy\n  apply:\n    description: The function used to apply information about a ServerTlsPolicy\n    parameters:\n    - name: ServerTlsPolicy\n      required: true\n      description: A full instance of a ServerTlsPolicy\n  delete:\n    description: The function used to delete a ServerTlsPolicy\n    parameters:\n    - name: ServerTlsPolicy\n      required: true\n      description: A full instance of a ServerTlsPolicy\n  deleteAll:\n    description: The function used to delete all ServerTlsPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ServerTlsPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ServerTlsPolicy:\n      title: ServerTlsPolicy\n      x-dcl-id: projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: true\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        allowOpen:\n          type: boolean\n          x-dcl-go-name: AllowOpen\n          description: Optional. Determines if server allows plaintext connections.\n            If set to true, server allows plain text connections. By default, it is\n            set to false. This setting is not exclusive of other encryption modes.\n            For example, if allow_open and mtls_policy are set, server allows both\n            plain text and mTLS connections. See documentation of other encryption\n            modes to confirm compatibility.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Free-text description of the resource.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        mtlsPolicy:\n          type: object\n          x-dcl-go-name: MtlsPolicy\n          x-dcl-go-type: ServerTlsPolicyMtlsPolicy\n          description: Optional. Defines a mechanism to provision peer validation\n            certificates for peer to peer authentication (Mutual TLS - mTLS). If not\n            specified, client certificate will not be requested. The connection is\n            treated as TLS and not mTLS. If allow_open and mtls_policy are set, server\n            allows both plain text and mTLS connections.\n          required:\n          - clientValidationCa\n          properties:\n            clientValidationCa:\n              type: array\n              x-dcl-go-name: ClientValidationCa\n              description: Required. Defines the mechanism to obtain the Certificate\n                Authority certificate to validate the client certificate.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ServerTlsPolicyMtlsPolicyClientValidationCa\n                properties:\n                  certificateProviderInstance:\n                    type: object\n                    x-dcl-go-name: CertificateProviderInstance\n                    x-dcl-go-type: ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance\n                    description: The certificate provider instance specification that\n                      will be passed to the data plane, which will be used to load\n                      necessary credential information.\n                    x-dcl-conflicts:\n                    - grpcEndpoint\n                    required:\n                    - pluginInstance\n                    properties:\n                      pluginInstance:\n                        type: string\n                        x-dcl-go-name: PluginInstance\n                        description: Required. Plugin instance name, used to locate\n                          and load CertificateProvider instance configuration. Set\n                          to \"google_cloud_private_spiffe\" to use Certificate Authority\n                          Service certificate provider instance.\n                  grpcEndpoint:\n                    type: object\n                    x-dcl-go-name: GrpcEndpoint\n                    x-dcl-go-type: ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint\n                    description: gRPC specific configuration to access the gRPC server\n                      to obtain the CA certificate.\n                    x-dcl-conflicts:\n                    - certificateProviderInstance\n                    required:\n                    - targetUri\n                    properties:\n                      targetUri:\n                        type: string\n                        x-dcl-go-name: TargetUri\n                        description: Required. The target URI of the gRPC endpoint.\n                          Only UDS path is supported, and should start with “unix:”.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the ServerTlsPolicy resource.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        serverCertificate:\n          type: object\n          x-dcl-go-name: ServerCertificate\n          x-dcl-go-type: ServerTlsPolicyServerCertificate\n          description: Optional. Defines a mechanism to provision server identity\n            (public and private keys). Cannot be combined with allow_open as a permissive\n            mode that allows both plain text and TLS is not supported.\n          properties:\n            certificateProviderInstance:\n              type: object\n              x-dcl-go-name: CertificateProviderInstance\n              x-dcl-go-type: ServerTlsPolicyServerCertificateCertificateProviderInstance\n              description: The certificate provider instance specification that will\n                be passed to the data plane, which will be used to load necessary\n                credential information.\n              x-dcl-conflicts:\n              - grpcEndpoint\n              required:\n              - pluginInstance\n              properties:\n                pluginInstance:\n                  type: string\n                  x-dcl-go-name: PluginInstance\n                  description: Required. Plugin instance name, used to locate and\n                    load CertificateProvider instance configuration. Set to \"google_cloud_private_spiffe\"\n                    to use Certificate Authority Service certificate provider instance.\n            grpcEndpoint:\n              type: object\n              x-dcl-go-name: GrpcEndpoint\n              x-dcl-go-type: ServerTlsPolicyServerCertificateGrpcEndpoint\n              description: gRPC specific configuration to access the gRPC server to\n                obtain the cert and private key.\n              x-dcl-conflicts:\n              - certificateProviderInstance\n              required:\n              - targetUri\n              properties:\n                targetUri:\n                  type: string\n                  x-dcl-go-name: TargetUri\n                  description: Required. The target URI of the gRPC endpoint. Only\n                    UDS path is supported, and should start with “unix:”.\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 8783 bytes
// MD5: 490f8c678a55b1a308e235162adac02f
