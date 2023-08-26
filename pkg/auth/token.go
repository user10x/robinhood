package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

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

func (authRequest *AuthRequest) mfaRequired(ctx context.Context, host string) ([]byte, error) {

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

// GetToken returns a token from the Robinhood API, reads if the file is there locally
// else queries api
func (authRequest *AuthRequest) GetToken(ctx context.Context, host string) (*AuthResponse, error) {

	var authResponse *AuthResponse

	authResponse, ok := readCachedToken()
	if ok {
		return authResponse, nil
	}
	client := &http.Client{
		Timeout: time.Duration(100 * time.Second),
	}

	payloadBytes, err := authRequest.mfaRequired(ctx, host)

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
		return nil, errors.New(fmt.Sprintf("failed %v", resp.Status))
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
	_, err = writeAuthLocally(authResponse)

	if err != nil {
		return nil, err
	}

	return authResponse, nil

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
		return nil, false
	}

	fileInfo, err := os.Stat(tokenFile)

	if err != nil {
		return nil, false
	}

	if fileInfo.ModTime().Second()-authResponse.ExpiresIn >= 0 {
		return nil, false
	}

	return authResponse, true

}

func writeAuthLocally(authResponse *AuthResponse) (string, error) {
	homedir, _ := os.UserHomeDir()
	tokenFileName := "token.json"
	err := os.MkdirAll(path.Join(homedir, ".robinhood"), 0755)
	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(path.Join(homedir, ".robinhood/", tokenFileName))

	if err != nil {
		return "", err
	}

	if fileInfo != nil {
		err = os.Remove(path.Join(homedir, ".robinhood/", tokenFileName))
		if err != nil {
			return "", err
		}

	}

	jsonBytes, err := json.Marshal(authResponse)
	tokFile := path.Join(homedir, ".robinhood/", tokenFileName)
	err = os.WriteFile(tokFile, jsonBytes, 0644)

	if err != nil {
		return "", err
	}

	return tokFile, nil
}
