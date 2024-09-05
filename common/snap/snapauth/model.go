package snapauth

type SignatureAccessTokenRequest struct {
}

type SignatureAccessTokenResponse struct {
	Signature string `json:"signature"`
}

type AccessTokenRequest struct {
	GrantType      string      `json:"grantType,omitempty"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
}

type AccessTokenResponse struct {
	ResponseCode    string            `json:"responseCode,omitempty"`
	ResponseMessage string            `json:"responseMessage,omitempty"`
	AccessToken     string            `json:"accessToken,omitempty"`
	TokenType       string            `json:"tokenType,omitempty"`
	ExpiresIn       string            `json:"expiresIn,omitempty"`
	AdditionalInfo  map[string]string `json:"additionalInfo,omitempty"`
}
