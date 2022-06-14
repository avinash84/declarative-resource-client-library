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
// gen_go_data -package alpha -var YAML_metadata_schema blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/alpha/metadata_schema.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/vertexai/alpha/metadata_schema.yaml
var YAML_metadata_schema = []byte("info:\n  title: VertexAI/MetadataSchema\n  description: The VertexAI MetadataSchema resource\n  x-dcl-struct-name: MetadataSchema\n  x-dcl-has-create: true\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a MetadataSchema\n    parameters:\n    - name: MetadataSchema\n      required: true\n      description: A full instance of a MetadataSchema\n  apply:\n    description: The function used to apply information about a MetadataSchema\n    parameters:\n    - name: MetadataSchema\n      required: true\n      description: A full instance of a MetadataSchema\n  list:\n    description: The function used to list information about many MetadataSchema\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: metadatastore\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    MetadataSchema:\n      title: MetadataSchema\n      x-dcl-id: projects/{{project}}/locations/{{location}}/metadataStores/{{metadata_store}}/metadataSchemas/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - schemaVersion\n      - schema\n      - schemaType\n      - project\n      - location\n      - metadataStore\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Timestamp when this MetadataSchema was created.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        metadataStore:\n          type: string\n          x-dcl-go-name: MetadataStore\n          description: The metadata store for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Vertex/MetadataStore\n            field: name\n            parent: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the MetadataSchema.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        schema:\n          type: string\n          x-dcl-go-name: Schema\n          description: Required. The raw YAML string representation of the MetadataSchema.\n            The combination of [MetadataSchema.version] and the schema name given\n            by `title` in [MetadataSchema.schema] must be unique within a MetadataStore.\n            The schema is defined as an OpenAPI 3.0.2 [MetadataSchema Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)\n          x-kubernetes-immutable: true\n        schemaType:\n          type: string\n          x-dcl-go-name: SchemaType\n          x-dcl-go-type: MetadataSchemaSchemaTypeEnum\n          description: 'The type of the MetadataSchema. This is a property that identifies\n            which metadata types will use the MetadataSchema. Possible values: METADATA_SCHEMA_TYPE_UNSPECIFIED,\n            ARTIFACT_TYPE, EXECUTION_TYPE, CONTEXT_TYPE'\n          x-kubernetes-immutable: true\n          enum:\n          - METADATA_SCHEMA_TYPE_UNSPECIFIED\n          - ARTIFACT_TYPE\n          - EXECUTION_TYPE\n          - CONTEXT_TYPE\n        schemaVersion:\n          type: string\n          x-dcl-go-name: SchemaVersion\n          description: 'The version of the MetadataSchema. The version''s format must\n            match the following regular expression: `^[0-9]+.+.+$`, which would allow\n            to order/compare different versions. Example: 1.0.0, 1.0.1, etc.'\n          x-kubernetes-immutable: true\n")

// 4049 bytes
// MD5: 5dc4a217896221f6b052d215c5cdcc81
