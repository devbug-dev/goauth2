package goauth2

//go:generate go run gen/eventgenerator/main.go -package goauth2 -id UserID -methodName CommandType -aggregateType resource-owner -inFile resource_owner_commands.go -outFile resource_owner_commands_gen.go

type OnBoardUser struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type GrantUserAdministratorRole struct {
	UserID         string `json:"userID"`
	GrantingUserID string `json:"grantingUserID"`
}
type AuthorizeUserToOnBoardClientApplications struct {
	UserID            string `json:"userID"`
	AuthorizingUserID string `json:"authorizingUserID"`
}
type RequestAccessTokenViaImplicitGrant struct {
	UserID      string `json:"userID"`
	ClientID    string `json:"clientID"`
	RedirectURI string `json:"redirectURI"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
type RequestAccessTokenViaROPCGrant struct {
	UserID       string `json:"userID"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
}
type RequestAuthorizationCodeViaAuthorizationCodeGrant struct {
	UserID      string `json:"userID"`
	ClientID    string `json:"clientID"`
	RedirectURI string `json:"redirectURI"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Scope       string `json:"scope"`
}
