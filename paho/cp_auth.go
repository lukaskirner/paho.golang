package paho

import "github.com/eclipse/paho.golang/packets"

type (
	// Auth is a representation of the MQTT Auth packet
	Auth struct {
		ReasonCode byte
		Properties *AuthProperties
	}

	// AuthProperties is a struct of the properties that can be set
	// for a Auth packet
	AuthProperties struct {
		AuthMethod   string
		AuthData     []byte
		ReasonString string
		User         map[string]string
	}
)

// InitProperties is a function that takes a lower level
// Properties struct and completes the properties of the Auth on
// which it is called
func (a *Auth) InitProperties(p *packets.Properties) {
	a.Properties = &AuthProperties{
		AuthMethod:   p.AuthMethod,
		AuthData:     p.AuthData,
		ReasonString: p.ReasonString,
		User:         p.User,
	}
}

// AuthFromPacketAuth takes a packets library Auth and
// returns a paho library Auth
func AuthFromPacketAuth(a *packets.Auth) *Auth {
	v := &Auth{ReasonCode: a.ReasonCode}
	v.InitProperties(a.Properties)

	return v
}

// Packet returns a packets library Auth from the paho Auth
// on which it is called
func (a *Auth) Packet() *packets.Auth {
	v := &packets.Auth{ReasonCode: a.ReasonCode}

	if a.Properties != nil {
		v.Properties = &packets.Properties{
			AuthMethod:   a.Properties.AuthMethod,
			AuthData:     a.Properties.AuthData,
			ReasonString: a.Properties.ReasonString,
			User:         a.Properties.User,
		}
	}

	return v
}

// AuthResponse is a represenation of the response to an Auth
// packet
type AuthResponse struct {
	Success    bool
	ReasonCode byte
	Properties *AuthProperties
}

// AuthResponseFromPacketAuth takes a packets library Auth and
// returns a paho library AuthResponse
func AuthResponseFromPacketAuth(a *packets.Auth) *AuthResponse {
	return &AuthResponse{
		Success:    true,
		ReasonCode: a.ReasonCode,
		Properties: &AuthProperties{
			ReasonString: a.Properties.ReasonString,
			User:         a.Properties.User,
		},
	}
}

// AuthResponseFromPacketDisconnect takes a packets library Disconnect and
// returns a paho library AuthResponse
func AuthResponseFromPacketDisconnect(d *packets.Disconnect) *AuthResponse {
	return &AuthResponse{
		Success:    true,
		ReasonCode: d.ReasonCode,
		Properties: &AuthProperties{
			ReasonString: d.Properties.ReasonString,
			User:         d.Properties.User,
		},
	}
}
