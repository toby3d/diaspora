package diaspora

import "testing"

const podEndpoint = "https://joindiaspora.com"

var client *Client

func TestNewClient(t *testing.T) {
	var err error
	client, err = NewClient(podEndpoint)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
}

func TestDiscoveryEndpoint(t *testing.T) {
	provider, err := client.DiscoveryEndpoint()
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if provider.RegistrationEndpoint == "" {
		t.Error("RegistrationEndpoint is empty")
	}
}
