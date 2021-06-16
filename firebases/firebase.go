package firebases

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// GetFirebaseApp return firebase instance
func GetFirebaseApp() (firebase.App, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return firebase.App{}, err
	}
	return *app, nil
}
