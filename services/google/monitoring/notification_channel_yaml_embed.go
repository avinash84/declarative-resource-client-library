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
// gen_go_data -package monitoring -var YAML_notification_channel blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/notification_channel.yaml

package monitoring

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/notification_channel.yaml
var YAML_notification_channel = []byte("info:\n  title: Monitoring/NotificationChannel\n  description: The Monitoring NotificationChannel resource\n  x-dcl-struct-name: NotificationChannel\n  x-dcl-has-create: true\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a NotificationChannel\n    parameters:\n    - name: NotificationChannel\n      required: true\n      description: A full instance of a NotificationChannel\n  apply:\n    description: The function used to apply information about a NotificationChannel\n    parameters:\n    - name: NotificationChannel\n      required: true\n      description: A full instance of a NotificationChannel\n  delete:\n    description: The function used to delete a NotificationChannel\n    parameters:\n    - name: NotificationChannel\n      required: true\n      description: A full instance of a NotificationChannel\n  deleteAll:\n    description: The function used to delete all NotificationChannel\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many NotificationChannel\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    NotificationChannel:\n      title: NotificationChannel\n      x-dcl-id: projects/{{project}}/notificationChannels/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional human-readable description of this notification\n            channel. This description may provide additional details, beyond the display\n            name, for the channel. This may not exceed 1024 Unicode characters.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: An optional human-readable name for this notification channel.\n            It is recommended that you specify a non-empty and unique name in order\n            to make it easier to identify the channels in your project, though this\n            is not enforced. The display name is limited to 512 Unicode characters.\n        enabled:\n          type: boolean\n          x-dcl-go-name: Enabled\n          description: Whether notifications are forwarded to the described channel.\n            This makes it possible to disable delivery of notifications to a particular\n            channel without removing the channel from all alerting policies that reference\n            the channel. This is a more convenient approach when the change is temporary\n            and you want to receive notifications from the same set of alerting policies\n            on the channel at some point in the future.\n          default: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Configuration fields that define the channel and its behavior.\n            The permissible and required labels are specified in the [NotificationChannelDescriptor.labels][google.monitoring.v3.NotificationChannelDescriptor.labels]\n            of the `NotificationChannelDescriptor` corresponding to the `type` field.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The full REST resource name for this channel. The format is:\n            projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID] The\n            `[CHANNEL_ID]` is automatically assigned by the server on creation.'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for this notification channel.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        type:\n          type: string\n          x-dcl-go-name: Type\n          description: The type of the notification channel. This field matches the\n            value of the [NotificationChannelDescriptor.type][google.monitoring.v3.NotificationChannelDescriptor.type]\n            field.\n        userLabels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: UserLabels\n          description: User-supplied key/value data that does not need to conform\n            to the corresponding `NotificationChannelDescriptor`'s schema, unlike\n            the `labels` field. This field is intended to be used for orv3nizing and\n            identifying the `NotificationChannel` objects. The field can contain up\n            to 64 entries. Each key and value is limited to 63 Unicode characters\n            or 128 bytes, whichever is smaller. Labels and values can contain only\n            lowercase letters, numerals, underscores, and dashes. Keys must begin\n            with a letter.\n        verificationStatus:\n          type: string\n          x-dcl-go-name: VerificationStatus\n          x-dcl-go-type: NotificationChannelVerificationStatusEnum\n          readOnly: true\n          description: 'Indicates whether this channel has been verified or not. On\n            a [`ListNotificationChannels`][google.monitoring.v3.NotificationChannelService.ListNotificationChannels]\n            or [`GetNotificationChannel`][google.monitoring.v3.NotificationChannelService.GetNotificationChannel]\n            operation, this field is expected to be populated. If the value is `UNVERIFIED`,\n            then it indicates that the channel is non-functioning (it both requires\n            verification and lacks verification); otherwise, it is assumed that the\n            channel works. If the channel is neither `VERIFIED` nor `UNVERIFIED`,\n            it implies that the channel is of a type that does not require verification\n            or that this specific channel has been exempted from verification because\n            it was created prior to verification being required for channels of this\n            type. This field cannot be modified using a standard [`UpdateNotificationChannel`][google.monitoring.v3.NotificationChannelService.UpdateNotificationChannel]\n            operation. To change the value of this field, you must call [`VerifyNotificationChannel`][google.monitoring.v3.NotificationChannelService.VerifyNotificationChannel].\n            Possible values: VERIFICATION_STATUS_UNSPECIFIED, UNVERIFIED, VERIFIED'\n          x-kubernetes-immutable: true\n          enum:\n          - VERIFICATION_STATUS_UNSPECIFIED\n          - UNVERIFIED\n          - VERIFIED\n")

// 6692 bytes
// MD5: 30cb09e30dff4dd556921d6de9e0070f
