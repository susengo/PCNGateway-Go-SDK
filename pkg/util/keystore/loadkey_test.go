package keystore

import (
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/userstore"
	"testing"
)

func TestLoadKey(t *testing.T) {

	ks, _ := NewFileBasedKeyStore(nil, "./test/msp/keystore", false)
	us := userstore.NewUserStore("./test/msp")

	user := &msp.UserData{
		UserName: "sdktest",
		AppCode:  "app0006202004071529586812466",
	}

	us.Load(user)

	LoadKey(user, ks)

	fmt.Println(string(user.EnrollmentCertificate))
	fmt.Println(user.PrivateKey)
}