package protocol

import (
	"context"
	"encoding/base64"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func FirebaseInit() {
	var err error
	sEnc := ""
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	opt := option.WithCredentialsJSON(sDec)
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("firebasev4: %v", err))
	}
}
