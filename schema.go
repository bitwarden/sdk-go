// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    clientSettings, err := UnmarshalClientSettings(bytes)
//    bytes, err = clientSettings.Marshal()
//
//    deviceType, err := UnmarshalDeviceType(bytes)
//    bytes, err = deviceType.Marshal()
//
//    command, err := UnmarshalCommand(bytes)
//    bytes, err = command.Marshal()
//
//    accessTokenLoginRequest, err := UnmarshalAccessTokenLoginRequest(bytes)
//    bytes, err = accessTokenLoginRequest.Marshal()
//
//    secretsCommand, err := UnmarshalSecretsCommand(bytes)
//    bytes, err = secretsCommand.Marshal()
//
//    secretGetRequest, err := UnmarshalSecretGetRequest(bytes)
//    bytes, err = secretGetRequest.Marshal()
//
//    secretsGetRequest, err := UnmarshalSecretsGetRequest(bytes)
//    bytes, err = secretsGetRequest.Marshal()
//
//    secretCreateRequest, err := UnmarshalSecretCreateRequest(bytes)
//    bytes, err = secretCreateRequest.Marshal()
//
//    secretIdentifiersRequest, err := UnmarshalSecretIdentifiersRequest(bytes)
//    bytes, err = secretIdentifiersRequest.Marshal()
//
//    secretPutRequest, err := UnmarshalSecretPutRequest(bytes)
//    bytes, err = secretPutRequest.Marshal()
//
//    secretsDeleteRequest, err := UnmarshalSecretsDeleteRequest(bytes)
//    bytes, err = secretsDeleteRequest.Marshal()
//
//    secretsSyncRequest, err := UnmarshalSecretsSyncRequest(bytes)
//    bytes, err = secretsSyncRequest.Marshal()
//
//    projectsCommand, err := UnmarshalProjectsCommand(bytes)
//    bytes, err = projectsCommand.Marshal()
//
//    projectGetRequest, err := UnmarshalProjectGetRequest(bytes)
//    bytes, err = projectGetRequest.Marshal()
//
//    projectCreateRequest, err := UnmarshalProjectCreateRequest(bytes)
//    bytes, err = projectCreateRequest.Marshal()
//
//    projectsListRequest, err := UnmarshalProjectsListRequest(bytes)
//    bytes, err = projectsListRequest.Marshal()
//
//    projectPutRequest, err := UnmarshalProjectPutRequest(bytes)
//    bytes, err = projectPutRequest.Marshal()
//
//    projectsDeleteRequest, err := UnmarshalProjectsDeleteRequest(bytes)
//    bytes, err = projectsDeleteRequest.Marshal()
//
//    generatorsCommand, err := UnmarshalGeneratorsCommand(bytes)
//    bytes, err = generatorsCommand.Marshal()
//
//    passwordGeneratorRequest, err := UnmarshalPasswordGeneratorRequest(bytes)
//    bytes, err = passwordGeneratorRequest.Marshal()
//
//    debugCommand, err := UnmarshalDebugCommand(bytes)
//    bytes, err = debugCommand.Marshal()
//
//    responseForAccessTokenLoginResponse, err := UnmarshalResponseForAccessTokenLoginResponse(bytes)
//    bytes, err = responseForAccessTokenLoginResponse.Marshal()
//
//    accessTokenLoginResponse, err := UnmarshalAccessTokenLoginResponse(bytes)
//    bytes, err = accessTokenLoginResponse.Marshal()
//
//    twoFactorProviders, err := UnmarshalTwoFactorProviders(bytes)
//    bytes, err = twoFactorProviders.Marshal()
//
//    authenticator, err := UnmarshalAuthenticator(bytes)
//    bytes, err = authenticator.Marshal()
//
//    email, err := UnmarshalEmail(bytes)
//    bytes, err = email.Marshal()
//
//    duo, err := UnmarshalDuo(bytes)
//    bytes, err = duo.Marshal()
//
//    yubiKey, err := UnmarshalYubiKey(bytes)
//    bytes, err = yubiKey.Marshal()
//
//    remember, err := UnmarshalRemember(bytes)
//    bytes, err = remember.Marshal()
//
//    webAuthn, err := UnmarshalWebAuthn(bytes)
//    bytes, err = webAuthn.Marshal()
//
//    responseForSecretIdentifiersResponse, err := UnmarshalResponseForSecretIdentifiersResponse(bytes)
//    bytes, err = responseForSecretIdentifiersResponse.Marshal()
//
//    secretIdentifiersResponse, err := UnmarshalSecretIdentifiersResponse(bytes)
//    bytes, err = secretIdentifiersResponse.Marshal()
//
//    secretIdentifierResponse, err := UnmarshalSecretIdentifierResponse(bytes)
//    bytes, err = secretIdentifierResponse.Marshal()
//
//    responseForSecretResponse, err := UnmarshalResponseForSecretResponse(bytes)
//    bytes, err = responseForSecretResponse.Marshal()
//
//    secretResponse, err := UnmarshalSecretResponse(bytes)
//    bytes, err = secretResponse.Marshal()
//
//    responseForSecretsResponse, err := UnmarshalResponseForSecretsResponse(bytes)
//    bytes, err = responseForSecretsResponse.Marshal()
//
//    secretsResponse, err := UnmarshalSecretsResponse(bytes)
//    bytes, err = secretsResponse.Marshal()
//
//    responseForSecretsDeleteResponse, err := UnmarshalResponseForSecretsDeleteResponse(bytes)
//    bytes, err = responseForSecretsDeleteResponse.Marshal()
//
//    secretsDeleteResponse, err := UnmarshalSecretsDeleteResponse(bytes)
//    bytes, err = secretsDeleteResponse.Marshal()
//
//    secretDeleteResponse, err := UnmarshalSecretDeleteResponse(bytes)
//    bytes, err = secretDeleteResponse.Marshal()
//
//    responseForSecretsSyncResponse, err := UnmarshalResponseForSecretsSyncResponse(bytes)
//    bytes, err = responseForSecretsSyncResponse.Marshal()
//
//    secretsSyncResponse, err := UnmarshalSecretsSyncResponse(bytes)
//    bytes, err = secretsSyncResponse.Marshal()
//
//    responseForProjectResponse, err := UnmarshalResponseForProjectResponse(bytes)
//    bytes, err = responseForProjectResponse.Marshal()
//
//    projectResponse, err := UnmarshalProjectResponse(bytes)
//    bytes, err = projectResponse.Marshal()
//
//    responseForProjectsResponse, err := UnmarshalResponseForProjectsResponse(bytes)
//    bytes, err = responseForProjectsResponse.Marshal()
//
//    projectsResponse, err := UnmarshalProjectsResponse(bytes)
//    bytes, err = projectsResponse.Marshal()
//
//    responseForProjectsDeleteResponse, err := UnmarshalResponseForProjectsDeleteResponse(bytes)
//    bytes, err = responseForProjectsDeleteResponse.Marshal()
//
//    projectsDeleteResponse, err := UnmarshalProjectsDeleteResponse(bytes)
//    bytes, err = projectsDeleteResponse.Marshal()
//
//    projectDeleteResponse, err := UnmarshalProjectDeleteResponse(bytes)
//    bytes, err = projectDeleteResponse.Marshal()
//
//    responseForString, err := UnmarshalResponseForString(bytes)
//    bytes, err = responseForString.Marshal()

package sdk

import "time"

import "encoding/json"

func UnmarshalClientSettings(data []byte) (ClientSettings, error) {
	var r ClientSettings
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ClientSettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeviceType(data []byte) (DeviceType, error) {
	var r DeviceType
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeviceType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCommand(data []byte) (Command, error) {
	var r Command
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Command) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAccessTokenLoginRequest(data []byte) (AccessTokenLoginRequest, error) {
	var r AccessTokenLoginRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccessTokenLoginRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsCommand(data []byte) (SecretsCommand, error) {
	var r SecretsCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretGetRequest(data []byte) (SecretGetRequest, error) {
	var r SecretGetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretGetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsGetRequest(data []byte) (SecretsGetRequest, error) {
	var r SecretsGetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsGetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretCreateRequest(data []byte) (SecretCreateRequest, error) {
	var r SecretCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretIdentifiersRequest(data []byte) (SecretIdentifiersRequest, error) {
	var r SecretIdentifiersRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretIdentifiersRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretPutRequest(data []byte) (SecretPutRequest, error) {
	var r SecretPutRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretPutRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsDeleteRequest(data []byte) (SecretsDeleteRequest, error) {
	var r SecretsDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsSyncRequest(data []byte) (SecretsSyncRequest, error) {
	var r SecretsSyncRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsSyncRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectsCommand(data []byte) (ProjectsCommand, error) {
	var r ProjectsCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectsCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectGetRequest(data []byte) (ProjectGetRequest, error) {
	var r ProjectGetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectGetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectCreateRequest(data []byte) (ProjectCreateRequest, error) {
	var r ProjectCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectsListRequest(data []byte) (ProjectsListRequest, error) {
	var r ProjectsListRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectsListRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectPutRequest(data []byte) (ProjectPutRequest, error) {
	var r ProjectPutRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectPutRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectsDeleteRequest(data []byte) (ProjectsDeleteRequest, error) {
	var r ProjectsDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectsDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalGeneratorsCommand(data []byte) (GeneratorsCommand, error) {
	var r GeneratorsCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GeneratorsCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPasswordGeneratorRequest(data []byte) (PasswordGeneratorRequest, error) {
	var r PasswordGeneratorRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PasswordGeneratorRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDebugCommand(data []byte) (DebugCommand, error) {
	var r DebugCommand
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DebugCommand) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForAccessTokenLoginResponse(data []byte) (ResponseForAccessTokenLoginResponse, error) {
	var r ResponseForAccessTokenLoginResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForAccessTokenLoginResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAccessTokenLoginResponse(data []byte) (AccessTokenLoginResponse, error) {
	var r AccessTokenLoginResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccessTokenLoginResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwoFactorProviders(data []byte) (TwoFactorProviders, error) {
	var r TwoFactorProviders
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwoFactorProviders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAuthenticator(data []byte) (Authenticator, error) {
	var r Authenticator
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Authenticator) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmail(data []byte) (Email, error) {
	var r Email
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Email) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDuo(data []byte) (Duo, error) {
	var r Duo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Duo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalYubiKey(data []byte) (YubiKey, error) {
	var r YubiKey
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *YubiKey) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRemember(data []byte) (Remember, error) {
	var r Remember
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Remember) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWebAuthn(data []byte) (WebAuthn, error) {
	var r WebAuthn
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WebAuthn) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForSecretIdentifiersResponse(data []byte) (ResponseForSecretIdentifiersResponse, error) {
	var r ResponseForSecretIdentifiersResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForSecretIdentifiersResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretIdentifiersResponse(data []byte) (SecretIdentifiersResponse, error) {
	var r SecretIdentifiersResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretIdentifiersResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretIdentifierResponse(data []byte) (SecretIdentifierResponse, error) {
	var r SecretIdentifierResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretIdentifierResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForSecretResponse(data []byte) (ResponseForSecretResponse, error) {
	var r ResponseForSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretResponse(data []byte) (SecretResponse, error) {
	var r SecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForSecretsResponse(data []byte) (ResponseForSecretsResponse, error) {
	var r ResponseForSecretsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForSecretsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsResponse(data []byte) (SecretsResponse, error) {
	var r SecretsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForSecretsDeleteResponse(data []byte) (ResponseForSecretsDeleteResponse, error) {
	var r ResponseForSecretsDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForSecretsDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsDeleteResponse(data []byte) (SecretsDeleteResponse, error) {
	var r SecretsDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretDeleteResponse(data []byte) (SecretDeleteResponse, error) {
	var r SecretDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForSecretsSyncResponse(data []byte) (ResponseForSecretsSyncResponse, error) {
	var r ResponseForSecretsSyncResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForSecretsSyncResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsSyncResponse(data []byte) (SecretsSyncResponse, error) {
	var r SecretsSyncResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsSyncResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForProjectResponse(data []byte) (ResponseForProjectResponse, error) {
	var r ResponseForProjectResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForProjectResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectResponse(data []byte) (ProjectResponse, error) {
	var r ProjectResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForProjectsResponse(data []byte) (ResponseForProjectsResponse, error) {
	var r ResponseForProjectsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForProjectsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectsResponse(data []byte) (ProjectsResponse, error) {
	var r ProjectsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForProjectsDeleteResponse(data []byte) (ResponseForProjectsDeleteResponse, error) {
	var r ResponseForProjectsDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForProjectsDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectsDeleteResponse(data []byte) (ProjectsDeleteResponse, error) {
	var r ProjectsDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectsDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProjectDeleteResponse(data []byte) (ProjectDeleteResponse, error) {
	var r ProjectDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProjectDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalResponseForString(data []byte) (ResponseForString, error) {
	var r ResponseForString
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ResponseForString) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Basic client behavior settings. These settings specify the various targets and behavior
// of the
// Bitwarden Client. They are optional and uneditable once the client is initialized.
//
// Defaults to
//
// ```
// # use bitwarden_core::{ClientSettings, DeviceType};
// let settings = ClientSettings {
// identity_url: "https://identity.bitwarden.com".to_string(),
// api_url: "https://api.bitwarden.com".to_string(),
// user_agent: "Bitwarden Rust-SDK".to_string(),
// device_type: DeviceType::SDK,
// bitwarden_client_version: None,
// bitwarden_package_type: None,
// device_identifier: None,
// };
// let default = ClientSettings::default();
// ```
type ClientSettings struct {
	// The api url of the targeted Bitwarden instance. Defaults to `https://api.bitwarden.com`            
	APIURL                                                                                    *string     `json:"apiUrl,omitempty"`
	// Bitwarden Client Version to send to Bitwarden. Optional for now in transition period.              
	BitwardenClientVersion                                                                    *string     `json:"bitwardenClientVersion,omitempty"`
	// Bitwarden Package Type to send to Bitwarden. We should evaluate this field to see if it            
	// should be optional later.                                                                          
	BitwardenPackageType                                                                      *string     `json:"bitwardenPackageType,omitempty"`
	// Device identifier to send to Bitwarden. Optional for now in transition period.                     
	DeviceIdentifier                                                                          *string     `json:"deviceIdentifier,omitempty"`
	// Device type to send to Bitwarden. Defaults to SDK                                                  
	DeviceType                                                                                *DeviceType `json:"deviceType,omitempty"`
	// The identity url of the targeted Bitwarden instance. Defaults to                                   
	// `https://identity.bitwarden.com`                                                                   
	IdentityURL                                                                               *string     `json:"identityUrl,omitempty"`
	// The user_agent to sent to Bitwarden. Defaults to `Bitwarden Rust-SDK`                              
	UserAgent                                                                                 *string     `json:"userAgent,omitempty"`
}

// Login with Secrets Manager Access Token
//
// This command is for initiating an authentication handshake with Bitwarden.
//
// Returns: [AccessTokenLoginResponse](bitwarden::secrets_manager::AccessTokenLoginResponse)
type Command struct {
	LoginAccessToken *AccessTokenLoginRequest `json:"loginAccessToken,omitempty"`
	Secrets          *SecretsCommand          `json:"secrets,omitempty"`
	Projects         *ProjectsCommand         `json:"projects,omitempty"`
	Generators       *GeneratorsCommand       `json:"generators,omitempty"`
	Debug            *DebugCommand            `json:"debug,omitempty"`
}

type DebugCommand struct {
	CancellationTest *CancellationTest `json:"cancellationTest,omitempty"`
	ErrorTest        *ErrorTest        `json:"errorTest,omitempty"`
}

type CancellationTest struct {
	DurationMillis int64 `json:"duration_millis"`
}

type ErrorTest struct {
}

// Generate a password
//
// Returns: [String]
type GeneratorsCommand struct {
	GeneratePassword PasswordGeneratorRequest `json:"generatePassword"`
}

// Password generator request options.
type PasswordGeneratorRequest struct {
	// When set to true, the generated password will not contain ambiguous characters.                 
	// The ambiguous characters are: I, O, l, 0, 1                                                     
	AvoidAmbiguous                                                                              bool   `json:"avoidAmbiguous"`
	// The length of the generated password.                                                           
	// Note that the password length must be greater than the sum of all the minimums.                 
	Length                                                                                      int64  `json:"length"`
	// Include lowercase characters (a-z).                                                             
	Lowercase                                                                                   bool   `json:"lowercase"`
	// The minimum number of lowercase characters in the generated password.                           
	// When set, the value must be between 1 and 9. This value is ignored if lowercase is false.       
	MinLowercase                                                                                *int64 `json:"minLowercase,omitempty"`
	// The minimum number of numbers in the generated password.                                        
	// When set, the value must be between 1 and 9. This value is ignored if numbers is false.         
	MinNumber                                                                                   *int64 `json:"minNumber,omitempty"`
	// The minimum number of special characters in the generated password.                             
	// When set, the value must be between 1 and 9. This value is ignored if special is false.         
	MinSpecial                                                                                  *int64 `json:"minSpecial,omitempty"`
	// The minimum number of uppercase characters in the generated password.                           
	// When set, the value must be between 1 and 9. This value is ignored if uppercase is false.       
	MinUppercase                                                                                *int64 `json:"minUppercase,omitempty"`
	// Include numbers (0-9).                                                                          
	Numbers                                                                                     bool   `json:"numbers"`
	// Include special characters: ! @ # $ % ^ & *                                                     
	Special                                                                                     bool   `json:"special"`
	// Include uppercase characters (A-Z).                                                             
	Uppercase                                                                                   bool   `json:"uppercase"`
}

// Login to Bitwarden with access token
type AccessTokenLoginRequest struct {
	// Bitwarden service API access token        
	AccessToken                          string  `json:"accessToken"`
	// Path to the state file                    
	StateFile                            *string `json:"stateFile,omitempty"`
}

// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Retrieve a project by the provided identifier
//
// Returns: [ProjectResponse](bitwarden::secrets_manager::projects::ProjectResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Creates a new project in the provided organization using the given data
//
// Returns: [ProjectResponse](bitwarden::secrets_manager::projects::ProjectResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Lists all projects of the given organization
//
// Returns: [ProjectsResponse](bitwarden::secrets_manager::projects::ProjectsResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Updates an existing project with the provided ID using the given data
//
// Returns: [ProjectResponse](bitwarden::secrets_manager::projects::ProjectResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Deletes all the projects whose IDs match the provided ones
//
// Returns:
// [ProjectsDeleteResponse](bitwarden::secrets_manager::projects::ProjectsDeleteResponse)
type ProjectsCommand struct {
	Get    *ProjectGetRequest     `json:"get,omitempty"`
	Create *ProjectCreateRequest  `json:"create,omitempty"`
	List   *ProjectsListRequest   `json:"list,omitempty"`
	Update *ProjectPutRequest     `json:"update,omitempty"`
	Delete *ProjectsDeleteRequest `json:"delete,omitempty"`
}

type ProjectCreateRequest struct {
	Name                                             string `json:"name"`
	// Organization where the project will be created       
	OrganizationID                                   string `json:"organizationId"`
}

type ProjectsDeleteRequest struct {
	// IDs of the projects to delete         
	IDS                             []string `json:"ids"`
}

type ProjectGetRequest struct {
	// ID of the project to retrieve       
	ID                              string `json:"id"`
}

type ProjectsListRequest struct {
	// Organization to retrieve all the projects from       
	OrganizationID                                   string `json:"organizationId"`
}

type ProjectPutRequest struct {
	// ID of the project to modify                    
	ID                                         string `json:"id"`
	Name                                       string `json:"name"`
	// Organization ID of the project to modify       
	OrganizationID                             string `json:"organizationId"`
}

// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Retrieve a secret by the provided identifier
//
// Returns: [SecretResponse](bitwarden::secrets_manager::secrets::SecretResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Retrieve secrets by the provided identifiers
//
// Returns: [SecretsResponse](bitwarden::secrets_manager::secrets::SecretsResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Creates a new secret in the provided organization using the given data
//
// Returns: [SecretResponse](bitwarden::secrets_manager::secrets::SecretResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Lists all secret identifiers of the given organization, to then retrieve each secret, use
// `CreateSecret`
//
// Returns:
// [SecretIdentifiersResponse](bitwarden::secrets_manager::secrets::SecretIdentifiersResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Updates an existing secret with the provided ID using the given data
//
// Returns: [SecretResponse](bitwarden::secrets_manager::secrets::SecretResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login or calling Sync at least once
//
// Deletes all the secrets whose IDs match the provided ones
//
// Returns:
// [SecretsDeleteResponse](bitwarden::secrets_manager::secrets::SecretsDeleteResponse)
//
// * Requires Authentication
// * Requires using an Access Token for login
//
// Retrieve the secrets accessible by the authenticated machine account
// Optionally, provide the last synced date to assess whether any changes have occurred
// If changes are detected, retrieves all the secrets accessible by the authenticated
// machine
// account
//
// Returns: [SecretsSyncResponse](bitwarden::secrets_manager::secrets::SecretsSyncResponse)
type SecretsCommand struct {
	Get      *SecretGetRequest         `json:"get,omitempty"`
	GetByIDS *SecretsGetRequest        `json:"getByIds,omitempty"`
	Create   *SecretCreateRequest      `json:"create,omitempty"`
	List     *SecretIdentifiersRequest `json:"list,omitempty"`
	Update   *SecretPutRequest         `json:"update,omitempty"`
	Delete   *SecretsDeleteRequest     `json:"delete,omitempty"`
	Sync     *SecretsSyncRequest       `json:"sync,omitempty"`
}

type SecretCreateRequest struct {
	Key                                                   string   `json:"key"`
	Note                                                  string   `json:"note"`
	// Organization where the secret will be created               
	OrganizationID                                        string   `json:"organizationId"`
	// IDs of the projects that this secret will belong to         
	ProjectIDS                                            []string `json:"projectIds,omitempty"`
	Value                                                 string   `json:"value"`
}

type SecretsDeleteRequest struct {
	// IDs of the secrets to delete         
	IDS                            []string `json:"ids"`
}

type SecretGetRequest struct {
	// ID of the secret to retrieve       
	ID                             string `json:"id"`
}

type SecretsGetRequest struct {
	// IDs of the secrets to retrieve         
	IDS                              []string `json:"ids"`
}

type SecretIdentifiersRequest struct {
	// Organization to retrieve all the secrets from       
	OrganizationID                                  string `json:"organizationId"`
}

type SecretsSyncRequest struct {
	// Optional date time a sync last occurred           
	LastSyncedDate                            *time.Time `json:"lastSyncedDate,omitempty"`
	// Organization to sync secrets from                 
	OrganizationID                            string     `json:"organizationId"`
}

type SecretPutRequest struct {
	// ID of the secret to modify                      
	ID                                        string   `json:"id"`
	Key                                       string   `json:"key"`
	Note                                      string   `json:"note"`
	// Organization ID of the secret to modify         
	OrganizationID                            string   `json:"organizationId"`
	ProjectIDS                                []string `json:"projectIds,omitempty"`
	Value                                     string   `json:"value"`
}

type ResponseForAccessTokenLoginResponse struct {
	// The response data. Populated if `success` is true.                                                
	Data                                                                       *AccessTokenLoginResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                          
	ErrorMessage                                                               *string                   `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                         
	Success                                                                    bool                      `json:"success"`
}

type AccessTokenLoginResponse struct {
	Authenticated                                                         bool                `json:"authenticated"`
	// Whether or not the user is required to update their master password                    
	ForcePasswordReset                                                    bool                `json:"forcePasswordReset"`
	// TODO: What does this do?                                                               
	ResetMasterPassword                                                   bool                `json:"resetMasterPassword"`
	TwoFactor                                                             *TwoFactorProviders `json:"twoFactor,omitempty"`
}

type TwoFactorProviders struct {
	Authenticator                                                         *Authenticator `json:"authenticator,omitempty"`
	// Duo-backed 2fa                                                                    
	Duo                                                                   *Duo           `json:"duo,omitempty"`
	// Email 2fa                                                                         
	Email                                                                 *Email         `json:"email,omitempty"`
	// Duo-backed 2fa operated by an organization the user is a member of                
	OrganizationDuo                                                       *Duo           `json:"organizationDuo,omitempty"`
	// Presence indicates the user has stored this device as bypassing 2fa               
	Remember                                                              *Remember      `json:"remember,omitempty"`
	// WebAuthn-backed 2fa                                                               
	WebAuthn                                                              *WebAuthn      `json:"webAuthn,omitempty"`
	// Yubikey-backed 2fa                                                                
	YubiKey                                                               *YubiKey       `json:"yubiKey,omitempty"`
}

type Authenticator struct {
}

type Duo struct {
	Host      string `json:"host"`
	Signature string `json:"signature"`
}

type Email struct {
	// The email to request a 2fa TOTP for       
	Email                                 string `json:"email"`
}

type Remember struct {
}

type WebAuthn struct {
}

type YubiKey struct {
	// Whether the stored yubikey supports near field communication     
	NFC                                                            bool `json:"nfc"`
}

type ResponseForSecretIdentifiersResponse struct {
	// The response data. Populated if `success` is true.                                                 
	Data                                                                       *SecretIdentifiersResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                           
	ErrorMessage                                                               *string                    `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                          
	Success                                                                    bool                       `json:"success"`
}

type SecretIdentifiersResponse struct {
	Data []SecretIdentifierResponse `json:"data"`
}

type SecretIdentifierResponse struct {
	ID             string `json:"id"`
	Key            string `json:"key"`
	OrganizationID string `json:"organizationId"`
}

type ResponseForSecretResponse struct {
	// The response data. Populated if `success` is true.                                      
	Data                                                                       *SecretResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                
	ErrorMessage                                                               *string         `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                               
	Success                                                                    bool            `json:"success"`
}

type SecretResponse struct {
	CreationDate   time.Time `json:"creationDate"`
	ID             string    `json:"id"`
	Key            string    `json:"key"`
	Note           string    `json:"note"`
	OrganizationID string    `json:"organizationId"`
	ProjectID      *string   `json:"projectId,omitempty"`
	RevisionDate   time.Time `json:"revisionDate"`
	Value          string    `json:"value"`
}

type ResponseForSecretsResponse struct {
	// The response data. Populated if `success` is true.                                       
	Data                                                                       *SecretsResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                 
	ErrorMessage                                                               *string          `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                
	Success                                                                    bool             `json:"success"`
}

type SecretsResponse struct {
	Data []SecretResponse `json:"data"`
}

type ResponseForSecretsDeleteResponse struct {
	// The response data. Populated if `success` is true.                                             
	Data                                                                       *SecretsDeleteResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                       
	ErrorMessage                                                               *string                `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                      
	Success                                                                    bool                   `json:"success"`
}

type SecretsDeleteResponse struct {
	Data []SecretDeleteResponse `json:"data"`
}

type SecretDeleteResponse struct {
	Error *string `json:"error,omitempty"`
	ID    string  `json:"id"`
}

type ResponseForSecretsSyncResponse struct {
	// The response data. Populated if `success` is true.                                           
	Data                                                                       *SecretsSyncResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                     
	ErrorMessage                                                               *string              `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                    
	Success                                                                    bool                 `json:"success"`
}

type SecretsSyncResponse struct {
	HasChanges bool             `json:"hasChanges"`
	Secrets    []SecretResponse `json:"secrets,omitempty"`
}

type ResponseForProjectResponse struct {
	// The response data. Populated if `success` is true.                                       
	Data                                                                       *ProjectResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                 
	ErrorMessage                                                               *string          `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                
	Success                                                                    bool             `json:"success"`
}

type ProjectResponse struct {
	CreationDate   time.Time `json:"creationDate"`
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	OrganizationID string    `json:"organizationId"`
	RevisionDate   time.Time `json:"revisionDate"`
}

type ResponseForProjectsResponse struct {
	// The response data. Populated if `success` is true.                                        
	Data                                                                       *ProjectsResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                  
	ErrorMessage                                                               *string           `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                 
	Success                                                                    bool              `json:"success"`
}

type ProjectsResponse struct {
	Data []ProjectResponse `json:"data"`
}

type ResponseForProjectsDeleteResponse struct {
	// The response data. Populated if `success` is true.                                              
	Data                                                                       *ProjectsDeleteResponse `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.                        
	ErrorMessage                                                               *string                 `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                                       
	Success                                                                    bool                    `json:"success"`
}

type ProjectsDeleteResponse struct {
	Data []ProjectDeleteResponse `json:"data"`
}

type ProjectDeleteResponse struct {
	Error *string `json:"error,omitempty"`
	ID    string  `json:"id"`
}

type ResponseForString struct {
	// The response data. Populated if `success` is true.                              
	Data                                                                       *string `json:"data,omitempty"`
	// A message for any error that may occur. Populated if `success` is false.        
	ErrorMessage                                                               *string `json:"errorMessage,omitempty"`
	// Whether or not the SDK request succeeded.                                       
	Success                                                                    bool    `json:"success"`
}

// Device type to send to Bitwarden. Defaults to SDK
type DeviceType string

const (
	Android           DeviceType = "Android"
	AndroidAmazon     DeviceType = "AndroidAmazon"
	ChromeBrowser     DeviceType = "ChromeBrowser"
	ChromeExtension   DeviceType = "ChromeExtension"
	DuckDuckGoBrowser DeviceType = "DuckDuckGoBrowser"
	EdgeBrowser       DeviceType = "EdgeBrowser"
	EdgeExtension     DeviceType = "EdgeExtension"
	FirefoxBrowser    DeviceType = "FirefoxBrowser"
	FirefoxExtension  DeviceType = "FirefoxExtension"
	IEBrowser         DeviceType = "IEBrowser"
	IOS               DeviceType = "iOS"
	LinuxCLI          DeviceType = "LinuxCLI"
	LinuxDesktop      DeviceType = "LinuxDesktop"
	MACOSCLI          DeviceType = "MacOsCLI"
	MACOSDesktop      DeviceType = "MacOsDesktop"
	OperaBrowser      DeviceType = "OperaBrowser"
	OperaExtension    DeviceType = "OperaExtension"
	SDK               DeviceType = "SDK"
	SafariBrowser     DeviceType = "SafariBrowser"
	SafariExtension   DeviceType = "SafariExtension"
	Server            DeviceType = "Server"
	UWP               DeviceType = "UWP"
	UnknownBrowser    DeviceType = "UnknownBrowser"
	VivaldiBrowser    DeviceType = "VivaldiBrowser"
	VivaldiExtension  DeviceType = "VivaldiExtension"
	WindowsCLI        DeviceType = "WindowsCLI"
	WindowsDesktop    DeviceType = "WindowsDesktop"
)

