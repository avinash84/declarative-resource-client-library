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
// gen_go_data -package alpha -var YAML_config blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/alpha/config.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/alpha/config.yaml
var YAML_config = []byte("info:\n  title: IdentityToolkit/Config\n  description: The IdentityToolkit Config resource\n  x-dcl-struct-name: Config\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Config\n    parameters:\n    - name: Config\n      required: true\n      description: A full instance of a Config\n  apply:\n    description: The function used to apply information about a Config\n    parameters:\n    - name: Config\n      required: true\n      description: A full instance of a Config\ncomponents:\n  schemas:\n    Config:\n      title: Config\n      x-dcl-id: projects/{{project}}/config\n      x-dcl-uses-state-hint: true\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - project\n      properties:\n        authorizedDomains:\n          type: array\n          x-dcl-go-name: AuthorizedDomains\n          description: List of domains authorized for OAuth redirects\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        blockingFunctions:\n          type: object\n          x-dcl-go-name: BlockingFunctions\n          x-dcl-go-type: ConfigBlockingFunctions\n          description: Configuration related to blocking functions.\n          properties:\n            triggers:\n              type: object\n              additionalProperties:\n                type: object\n                x-dcl-go-type: ConfigBlockingFunctionsTriggers\n                properties:\n                  functionUri:\n                    type: string\n                    x-dcl-go-name: FunctionUri\n                    description: HTTP URI trigger for the Cloud Function.\n                    x-dcl-references:\n                    - resource: Cloudfunctions/Function\n                      field: httpsTrigger.url\n                  updateTime:\n                    type: string\n                    format: date-time\n                    x-dcl-go-name: UpdateTime\n                    readOnly: true\n                    description: When the trigger was changed.\n              x-dcl-go-name: Triggers\n              description: 'Map of Trigger to event type. Key should be one of the\n                supported event types: \"beforeCreate\", \"beforeSignIn\"'\n        client:\n          type: object\n          x-dcl-go-name: Client\n          x-dcl-go-type: ConfigClient\n          description: Options related to how clients making requests on behalf of\n            a project should be configured.\n          properties:\n            apiKey:\n              type: string\n              x-dcl-go-name: ApiKey\n              readOnly: true\n              description: Output only. API key that can be used when making requests\n                for this project.\n              x-dcl-sensitive: true\n            firebaseSubdomain:\n              type: string\n              x-dcl-go-name: FirebaseSubdomain\n              readOnly: true\n              description: Output only. Firebase subdomain.\n            permissions:\n              type: object\n              x-dcl-go-name: Permissions\n              x-dcl-go-type: ConfigClientPermissions\n              description: Configuration related to restricting a user's ability to\n                affect their account.\n              properties:\n                disabledUserDeletion:\n                  type: boolean\n                  x-dcl-go-name: DisabledUserDeletion\n                  description: When true, end users cannot delete their account on\n                    the associated project through any of our API methods\n                disabledUserSignup:\n                  type: boolean\n                  x-dcl-go-name: DisabledUserSignup\n                  description: When true, end users cannot sign up for a new account\n                    on the associated project through any of our API methods\n        mfa:\n          type: object\n          x-dcl-go-name: Mfa\n          x-dcl-go-type: ConfigMfa\n          description: Configuration for this project's multi-factor authentication,\n            including whether it is active and what factors can be used for the second\n            factor\n          properties:\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: ConfigMfaStateEnum\n              description: 'Whether MultiFactor Authentication has been enabled for\n                this project. Possible values: STATE_UNSPECIFIED, DISABLED, ENABLED,\n                MANDATORY'\n              enum:\n              - STATE_UNSPECIFIED\n              - DISABLED\n              - ENABLED\n              - MANDATORY\n        monitoring:\n          type: object\n          x-dcl-go-name: Monitoring\n          x-dcl-go-type: ConfigMonitoring\n          description: Configuration related to monitoring project activity.\n          properties:\n            requestLogging:\n              type: object\n              x-dcl-go-name: RequestLogging\n              x-dcl-go-type: ConfigMonitoringRequestLogging\n              description: Configuration for logging requests made to this project\n                to Stackdriver Logging\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Whether logging is enabled for this project or not.\n        multiTenant:\n          type: object\n          x-dcl-go-name: MultiTenant\n          x-dcl-go-type: ConfigMultiTenant\n          description: Configuration related to multi-tenant functionality.\n          properties:\n            allowTenants:\n              type: boolean\n              x-dcl-go-name: AllowTenants\n              description: Whether this project can have tenants or not.\n            defaultTenantLocation:\n              type: string\n              x-dcl-go-name: DefaultTenantLocation\n              description: The default cloud parent org or folder that the tenant\n                project should be created under. The parent resource name should be\n                in the format of \"<type>/<number>\", such as \"folders/123\" or \"organizations/456\".\n                If the value is not set, the tenant will be created under the same\n                organization or folder as the agent project.\n              x-dcl-references:\n              - resource: Cloudresourcemanager/Folder\n                field: name\n              - resource: Cloudresourcemanager/Organization\n                field: name\n        notification:\n          type: object\n          x-dcl-go-name: Notification\n          x-dcl-go-type: ConfigNotification\n          description: Configuration related to sending notifications to users.\n          properties:\n            defaultLocale:\n              type: string\n              x-dcl-go-name: DefaultLocale\n              description: Default locale used for email and SMS in IETF BCP 47 format.\n            sendEmail:\n              type: object\n              x-dcl-go-name: SendEmail\n              x-dcl-go-type: ConfigNotificationSendEmail\n              description: Options for email sending.\n              properties:\n                callbackUri:\n                  type: string\n                  x-dcl-go-name: CallbackUri\n                  description: action url in email template.\n                changeEmailTemplate:\n                  $ref: '#/components/schemas/EmailTemplate'\n                  x-dcl-go-name: ChangeEmailTemplate\n                dnsInfo:\n                  type: object\n                  x-dcl-go-name: DnsInfo\n                  x-dcl-go-type: ConfigNotificationSendEmailDnsInfo\n                  description: Information of custom domain DNS verification.\n                  properties:\n                    customDomain:\n                      type: string\n                      x-dcl-go-name: CustomDomain\n                      readOnly: true\n                      description: Output only. The applied verified custom domain.\n                      x-kubernetes-immutable: true\n                    customDomainState:\n                      type: string\n                      x-dcl-go-name: CustomDomainState\n                      x-dcl-go-type: ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum\n                      readOnly: true\n                      description: 'Output only. The current verification state of\n                        the custom domain. The custom domain will only be used once\n                        the domain verification is successful. Possible values: VERIFICATION_STATE_UNSPECIFIED,\n                        NOT_STARTED, IN_PROGRESS, FAILED, SUCCEEDED'\n                      x-kubernetes-immutable: true\n                      enum:\n                      - VERIFICATION_STATE_UNSPECIFIED\n                      - NOT_STARTED\n                      - IN_PROGRESS\n                      - FAILED\n                      - SUCCEEDED\n                    domainVerificationRequestTime:\n                      type: string\n                      format: date-time\n                      x-dcl-go-name: DomainVerificationRequestTime\n                      readOnly: true\n                      description: Output only. The timestamp of initial request for\n                        the current domain verification.\n                      x-kubernetes-immutable: true\n                    pendingCustomDomain:\n                      type: string\n                      x-dcl-go-name: PendingCustomDomain\n                      readOnly: true\n                      description: Output only. The custom domain that's to be verified.\n                      x-kubernetes-immutable: true\n                    useCustomDomain:\n                      type: boolean\n                      x-dcl-go-name: UseCustomDomain\n                      description: Whether to use custom domain.\n                method:\n                  type: string\n                  x-dcl-go-name: Method\n                  x-dcl-go-type: ConfigNotificationSendEmailMethodEnum\n                  description: 'The method used for sending an email. Possible values:\n                    METHOD_UNSPECIFIED, DEFAULT, CUSTOM_SMTP'\n                  enum:\n                  - METHOD_UNSPECIFIED\n                  - DEFAULT\n                  - CUSTOM_SMTP\n                resetPasswordTemplate:\n                  $ref: '#/components/schemas/EmailTemplate'\n                  x-dcl-go-name: ResetPasswordTemplate\n                revertSecondFactorAdditionTemplate:\n                  $ref: '#/components/schemas/EmailTemplate'\n                  x-dcl-go-name: RevertSecondFactorAdditionTemplate\n                smtp:\n                  type: object\n                  x-dcl-go-name: Smtp\n                  x-dcl-go-type: ConfigNotificationSendEmailSmtp\n                  description: Use a custom SMTP relay\n                  properties:\n                    host:\n                      type: string\n                      x-dcl-go-name: Host\n                      description: SMTP relay host\n                    password:\n                      type: string\n                      x-dcl-go-name: Password\n                      description: SMTP relay password\n                      x-dcl-sensitive: true\n                      x-dcl-mutable-unreadable: true\n                    port:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Port\n                      description: SMTP relay port\n                    securityMode:\n                      type: string\n                      x-dcl-go-name: SecurityMode\n                      x-dcl-go-type: ConfigNotificationSendEmailSmtpSecurityModeEnum\n                      description: 'SMTP security mode. Possible values: SECURITY_MODE_UNSPECIFIED,\n                        SSL, START_TLS'\n                      enum:\n                      - SECURITY_MODE_UNSPECIFIED\n                      - SSL\n                      - START_TLS\n                    senderEmail:\n                      type: string\n                      x-dcl-go-name: SenderEmail\n                      description: Sender email for the SMTP relay\n                    username:\n                      type: string\n                      x-dcl-go-name: Username\n                      description: SMTP relay username\n                verifyEmailTemplate:\n                  $ref: '#/components/schemas/EmailTemplate'\n                  x-dcl-go-name: VerifyEmailTemplate\n            sendSms:\n              type: object\n              x-dcl-go-name: SendSms\n              x-dcl-go-type: ConfigNotificationSendSms\n              description: Options for SMS sending.\n              properties:\n                smsTemplate:\n                  type: object\n                  x-dcl-go-name: SmsTemplate\n                  x-dcl-go-type: ConfigNotificationSendSmsSmsTemplate\n                  readOnly: true\n                  description: Output only. The template to use when sending an SMS.\n                  x-kubernetes-immutable: true\n                  properties:\n                    content:\n                      type: string\n                      x-dcl-go-name: Content\n                      readOnly: true\n                      description: 'Output only. The SMS''s content. Can contain the\n                        following placeholders which will be replaced with the appropriate\n                        values: %APP_NAME% - For Android or iOS apps, the app''s display\n                        name. For web apps, the domain hosting the application. %LOGIN_CODE%\n                        - The OOB code being sent in the SMS.'\n                      x-kubernetes-immutable: true\n                useDeviceLocale:\n                  type: boolean\n                  x-dcl-go-name: UseDeviceLocale\n                  description: Whether to use the accept_language header for SMS.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project of the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        quota:\n          type: object\n          x-dcl-go-name: Quota\n          x-dcl-go-type: ConfigQuota\n          description: Configuration related to quotas.\n          properties:\n            signUpQuotaConfig:\n              type: object\n              x-dcl-go-name: SignUpQuotaConfig\n              x-dcl-go-type: ConfigQuotaSignUpQuotaConfig\n              description: Quota for the Signup endpoint, if overwritten. Signup quota\n                is measured in sign ups per project per hour per IP.\n              properties:\n                quota:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Quota\n                  description: Corresponds to the 'refill_token_count' field in QuotaServer\n                    config\n                quotaDuration:\n                  type: string\n                  x-dcl-go-name: QuotaDuration\n                  description: How long this quota will be active for\n                startTime:\n                  type: string\n                  format: date-time\n                  x-dcl-go-name: StartTime\n                  description: When this quota will take affect\n        signIn:\n          type: object\n          x-dcl-go-name: SignIn\n          x-dcl-go-type: ConfigSignIn\n          description: Configuration related to local sign in methods.\n          properties:\n            allowDuplicateEmails:\n              type: boolean\n              x-dcl-go-name: AllowDuplicateEmails\n              description: Whether to allow more than one account to have the same\n                email.\n            anonymous:\n              type: object\n              x-dcl-go-name: Anonymous\n              x-dcl-go-type: ConfigSignInAnonymous\n              description: Configuration options related to authenticating an anonymous\n                user.\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Whether anonymous user auth is enabled for the project\n                    or not.\n            email:\n              type: object\n              x-dcl-go-name: Email\n              x-dcl-go-type: ConfigSignInEmail\n              description: Configuration options related to authenticating a user\n                by their email address.\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Whether email auth is enabled for the project or not.\n                hashConfig:\n                  type: object\n                  x-dcl-go-name: HashConfig\n                  x-dcl-go-type: ConfigSignInEmailHashConfig\n                  readOnly: true\n                  description: Output only. Hash config information.\n                  properties:\n                    algorithm:\n                      type: string\n                      x-dcl-go-name: Algorithm\n                      x-dcl-go-type: ConfigSignInEmailHashConfigAlgorithmEnum\n                      readOnly: true\n                      description: 'Output only. Different password hash algorithms\n                        used in Identity Toolkit. Possible values: HASH_ALGORITHM_UNSPECIFIED,\n                        HMAC_SHA256, HMAC_SHA1, HMAC_MD5, SCRYPT, PBKDF_SHA1, MD5,\n                        HMAC_SHA512, SHA1, BCRYPT, PBKDF2_SHA256, SHA256, SHA512,\n                        STANDARD_SCRYPT'\n                      enum:\n                      - HASH_ALGORITHM_UNSPECIFIED\n                      - HMAC_SHA256\n                      - HMAC_SHA1\n                      - HMAC_MD5\n                      - SCRYPT\n                      - PBKDF_SHA1\n                      - MD5\n                      - HMAC_SHA512\n                      - SHA1\n                      - BCRYPT\n                      - PBKDF2_SHA256\n                      - SHA256\n                      - SHA512\n                      - STANDARD_SCRYPT\n                    memoryCost:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: MemoryCost\n                      readOnly: true\n                      description: Output only. Memory cost for hash calculation.\n                        Used by scrypt and other similar password derivation algorithms.\n                        See https://tools.ietf.org/html/rfc7914 for explanation of\n                        field.\n                    rounds:\n                      type: integer\n                      format: int64\n                      x-dcl-go-name: Rounds\n                      readOnly: true\n                      description: Output only. How many rounds for hash calculation.\n                        Used by scrypt and other similar password derivation algorithms.\n                    saltSeparator:\n                      type: string\n                      x-dcl-go-name: SaltSeparator\n                      readOnly: true\n                      description: Output only. Non-printable character to be inserted\n                        between the salt and plain text password in base64.\n                    signerKey:\n                      type: string\n                      x-dcl-go-name: SignerKey\n                      readOnly: true\n                      description: Output only. Signer key in base64.\n                      x-dcl-sensitive: true\n                passwordRequired:\n                  type: boolean\n                  x-dcl-go-name: PasswordRequired\n                  description: Whether a password is required for email auth or not.\n                    If true, both an email and password must be provided to sign in.\n                    If false, a user may sign in via either email/password or email\n                    link.\n            hashConfig:\n              type: object\n              x-dcl-go-name: HashConfig\n              x-dcl-go-type: ConfigSignInHashConfig\n              readOnly: true\n              description: Output only. Hash config information.\n              properties:\n                algorithm:\n                  type: string\n                  x-dcl-go-name: Algorithm\n                  x-dcl-go-type: ConfigSignInHashConfigAlgorithmEnum\n                  readOnly: true\n                  description: 'Output only. Different password hash algorithms used\n                    in Identity Toolkit. Possible values: HASH_ALGORITHM_UNSPECIFIED,\n                    HMAC_SHA256, HMAC_SHA1, HMAC_MD5, SCRYPT, PBKDF_SHA1, MD5, HMAC_SHA512,\n                    SHA1, BCRYPT, PBKDF2_SHA256, SHA256, SHA512, STANDARD_SCRYPT'\n                  enum:\n                  - HASH_ALGORITHM_UNSPECIFIED\n                  - HMAC_SHA256\n                  - HMAC_SHA1\n                  - HMAC_MD5\n                  - SCRYPT\n                  - PBKDF_SHA1\n                  - MD5\n                  - HMAC_SHA512\n                  - SHA1\n                  - BCRYPT\n                  - PBKDF2_SHA256\n                  - SHA256\n                  - SHA512\n                  - STANDARD_SCRYPT\n                memoryCost:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: MemoryCost\n                  readOnly: true\n                  description: Output only. Memory cost for hash calculation. Used\n                    by scrypt and other similar password derivation algorithms. See\n                    https://tools.ietf.org/html/rfc7914 for explanation of field.\n                rounds:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Rounds\n                  readOnly: true\n                  description: Output only. How many rounds for hash calculation.\n                    Used by scrypt and other similar password derivation algorithms.\n                saltSeparator:\n                  type: string\n                  x-dcl-go-name: SaltSeparator\n                  readOnly: true\n                  description: Output only. Non-printable character to be inserted\n                    between the salt and plain text password in base64.\n                signerKey:\n                  type: string\n                  x-dcl-go-name: SignerKey\n                  readOnly: true\n                  description: Output only. Signer key in base64.\n                  x-dcl-sensitive: true\n            phoneNumber:\n              type: object\n              x-dcl-go-name: PhoneNumber\n              x-dcl-go-type: ConfigSignInPhoneNumber\n              description: Configuration options related to authenticated a user by\n                their phone number.\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Whether phone number auth is enabled for the project\n                    or not.\n                testPhoneNumbers:\n                  type: object\n                  additionalProperties:\n                    type: string\n                  x-dcl-go-name: TestPhoneNumbers\n                  description: A map of that can be used for phone auth testing.\n        subtype:\n          type: string\n          x-dcl-go-name: Subtype\n          x-dcl-go-type: ConfigSubtypeEnum\n          readOnly: true\n          description: 'Output only. The subtype of this config. Possible values:\n            SUBTYPE_UNSPECIFIED, IDENTITY_PLATFORM, FIREBASE_AUTH'\n          x-kubernetes-immutable: true\n          enum:\n          - SUBTYPE_UNSPECIFIED\n          - IDENTITY_PLATFORM\n          - FIREBASE_AUTH\n    EmailTemplate:\n      x-dcl-has-iam: false\n      type: object\n      x-dcl-go-name: ResetPasswordTemplate\n      x-dcl-go-type: ConfigEmailTemplate\n      description: Email template for reset password\n      properties:\n        body:\n          type: string\n          x-dcl-go-name: Body\n          description: Email body\n        bodyFormat:\n          type: string\n          x-dcl-go-name: BodyFormat\n          x-dcl-go-type: ConfigEmailTemplateBodyFormatEnum\n          description: 'Email body format Possible values: BODY_FORMAT_UNSPECIFIED,\n            PLAIN_TEXT, HTML'\n          enum:\n          - BODY_FORMAT_UNSPECIFIED\n          - PLAIN_TEXT\n          - HTML\n        customized:\n          type: boolean\n          x-dcl-go-name: Customized\n          readOnly: true\n          description: Output only. Whether the body or subject of the email is customized.\n          x-kubernetes-immutable: true\n        replyTo:\n          type: string\n          x-dcl-go-name: ReplyTo\n          description: Reply-to address\n        senderDisplayName:\n          type: string\n          x-dcl-go-name: SenderDisplayName\n          description: Sender display name\n        senderLocalPart:\n          type: string\n          x-dcl-go-name: SenderLocalPart\n          description: Local part of From address\n        subject:\n          type: string\n          x-dcl-go-name: Subject\n          description: Subject of the email\n")

// 24977 bytes
// MD5: a49d6c65d98f1e49bffd9a8899fe1977
