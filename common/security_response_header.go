package common

type ResponseHeader int

const (
	HeadXFrameOptions ResponseHeader = iota
	HeadStrictTransportSecurity
	HeadExpectCT
	HeadContentSecurityPolicy
	HeadXXSSProtection
	HeadXContentTypeOptions
	HeadCorsAllowOrigin
	HeadCorsAllowHeaders
	HeadCorsAllowMethods
)

func (r ResponseHeader) String() string {
	return [...]string{
		"x-frame-options",
		"strict-transport-security",
		"expect-ct",
		"content-security-policy",
		"x-xss-protection",
		"x-content-type-options",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Headers",
	}[r]
}

const (
	headXFrameOptions           = "deny"
	headStrictTransportSecurity = "max-age=31536000 ; includeSubDomains"
	headExpectCT                = "max-age=31536000, enforce"
	headContentSecurityPolicy   = "default-src 'none'; script-src 'none'; object-src 'none'; base-uri 'none';"
	headXXSSProtection          = "1; mode=block"
	headXContentTypeOptions     = "nosniff"
	headAllowCors               = "*"
)

func SetResponseHeader(r ResponseHeader) (key, value string) {
	messages := map[ResponseHeader]string{
		HeadXFrameOptions:           headXFrameOptions,
		HeadStrictTransportSecurity: headStrictTransportSecurity,
		HeadExpectCT:                headExpectCT,
		HeadContentSecurityPolicy:   headContentSecurityPolicy,
		HeadXXSSProtection:          headXXSSProtection,
		HeadXContentTypeOptions:     headXContentTypeOptions,
		HeadCorsAllowOrigin:         headAllowCors,
		HeadCorsAllowHeaders:        headAllowCors,
		HeadCorsAllowMethods:        headAllowCors,
	}

	if messages[r] == "" {
		return HeadContentSecurityPolicy.String(), headContentSecurityPolicy
	}

	return r.String(), messages[r]
}
