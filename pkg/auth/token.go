package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

const host = "https://api.robinhood.com/oauth2/token/"

type MfaRequiredResponse struct {
	MfaRequired bool   `json:"mfa_required"`
	MfaType     string `json:"mfa_type"`
	MfaCode     string `json:"mfa_code,omitempty"`
}

type AuthRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	GrantType   string `json:"grant_type"`
	ClientId    string `json:"client_id"`
	ExpiresIn   int    `json:"expires_in"`
	DeviceToken string `json:"device_token"`
	MfaCode     string `json:"mfa_code,omitempty"`
	Scope       string `json:"scope"`
}
type AuthResponse struct {
	AccessToken  string      `json:"access_token"`
	ExpiresIn    int         `json:"expires_in"`
	TokenType    string      `json:"token_type"`
	Scope        string      `json:"scope"`
	RefreshToken string      `json:"refresh_token"`
	MfaCode      string      `json:"mfa_code"`
	BackupCode   interface{} `json:"backup_code"`
}

func (authRequest *AuthRequest) mfaRequired(ctx context.Context) ([]byte, error) {

	authPayload, err := json.Marshal(authRequest)

	req, err := http.NewRequestWithContext(ctx, "POST", host, bytes.NewBuffer(authPayload))
	client := &http.Client{}

	var mfaResponse *MfaRequiredResponse

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("agent", "Robinhood/823 (iPhone; iOS 7.1.2; Scale/2.00)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &mfaResponse)

	if err != nil {
		return nil, err
	}
	if mfaResponse.MfaRequired {
		fmt.Println("MFA Required")

		_, _ = fmt.Scanf("%s", &authRequest.MfaCode)

		authPayload, err = json.Marshal(&authRequest)

		if err != nil {
			return nil, err
		}
	}
	return authPayload, nil
}

func readCachedToken() (*AuthResponse, bool) {

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, false
	}
	tokenFile := home + "/.robinhood/token.json"
	fBytes, err := os.ReadFile(tokenFile)
	if err != nil {
		return nil, false
	}
	var authResponse *AuthResponse
	err = json.Unmarshal(fBytes, &authResponse)

	if err != nil {

	}

	return authResponse, true

}

func cacheAuthInfo(authResponse *AuthResponse) error {
	homedir, _ := os.UserHomeDir()
	tokenFileName := "token.json"
	err := os.MkdirAll(path.Join(homedir, ".robinhood"), 0755)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(path.Join(homedir, ".robinhood/", tokenFileName))

	if err != nil {
		return err
	}

	fileInfo.ModTime()

	if fileInfo != nil {
		err = os.Remove(path.Join(homedir, ".robinhood/", tokenFileName))
		if err != nil {
			return err
		}

	}

	jsonBytes, err := json.Marshal(authResponse)

	err = os.WriteFile(path.Join(homedir, ".robinhood/", tokenFileName), jsonBytes, 0644)

	if err != nil {
		return err
	}
	return nil
}

// GetToken returns a token from the Robinhood API
func (authRequest *AuthRequest) GetToken(ctx context.Context) (*AuthResponse, error) {

	var authResponse *AuthResponse

	authResponse, ok := readCachedToken()

	if ok {
		return authResponse, nil
	}
	client := &http.Client{
		Timeout: time.Duration(100 * time.Second),
	}

	payloadBytes, err := authRequest.mfaRequired(ctx)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", host, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("agent", "Robinhood/823 (iPhone; iOS 7.1.2; Scale/2.00)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 300 {
		fmt.Println(resp.StatusCode)
		// return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))

	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}

	// cache token
	defer cacheAuthInfo(authResponse)

	return authResponse, nil

}
