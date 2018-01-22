package diaspora

import (
	"fmt"
	"os"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

var client = NewClient(
	"http://joindiaspora.com/",
	&oauth2.Token{
		AccessToken:  os.Getenv("DIASPORA_TOKEN"),
		TokenType:    "bearer",
		RefreshToken: "",
		Expiry:       time.Date(2018, 01, 23, 11, 24, 02, 831168686, time.UTC),
	},
)

func TestGetAspects(t *testing.T) {
	aspects, err := client.GetAspects()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("%#+v", aspects)
	}

	for _, aspect := range aspects {
		t.Run(fmt.Sprintln("get", aspect.Name, "aspect"), func(t *testing.T) {
			aspect, err = client.GetAspect(aspect.ID)
			if err != nil {
				t.Error(err.Error())
				t.Fail()
			}
			t.Log("order:", aspect.Order)
		})
	}
}
