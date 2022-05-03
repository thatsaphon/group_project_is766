package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var ClientAuth *auth.Client
var ctx context.Context

func InitFirebase() (*firebase.App, error) {

	opt := option.WithCredentialsFile("config/is766-project-firebase-adminsdk-wpaoi-73ea0d7c6a.json")
	ctx = context.Background()
	// config := &firebase.Config{ProjectID: "is767-2021-thatsaphon"}
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	FirebaseApp = app

	return app, nil
}

func InitClientAuth(app *firebase.App) (*auth.Client, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v\n", err)
	}
	ClientAuth = client

	return client, nil
}

func GetUserByUid(uid string) (*auth.UserRecord, error) {
	u, err := ClientAuth.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("error getting user %s: %v\n", uid, err)

	}
	log.Printf("Successfully fetched user data: %v\n", u)
	return u, nil
}

func GetUserByEmail(email string) (*auth.UserRecord, error) {
	u, err := ClientAuth.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error getting user by email %s: %v\n", email, err)

	}
	log.Printf("Successfully fetched user data: %v\n", u)
	return u, nil
}

func CreateUser(email string,
	password string, displayName string) error {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		DisplayName(displayName).
		Disabled(false)

	u, err := ClientAuth.CreateUser(ctx, params)
	if err != nil {
		log.Printf("error creating user: %v\n", err)
		return err
	}
	log.Printf("Successfully created user: %v\n", u)
	return nil
}

func ListAllUsers() {
	iter := ClientAuth.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}
		log.Printf("read user user: %#v\n", user)
	}
}

func CreateCustomToken(uid string) (string, error) {
	token, err := ClientAuth.CustomToken(ctx, uid)
	if err != nil {
		// log.Fatalf("error minting custom token: %v\n", err)
		return "", err
	}

	log.Printf("Got custom token: %v\n", token)
	// fmt.Println(token)
	return token, nil
}

func VerifyIdToken(idToken string) {
	token, err := ClientAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
}

func VerifyIdTokenAndCheckRevoke(idToken string) {
	token, err := ClientAuth.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		if err.Error() == "ID token has been revoked" {
			// Token is revoked. Inform the user to reauthenticate or signOut() the user.
		} else {
			// Token is invalid
		}
	}
	log.Printf("Verified ID token: %v\n", token)

}

func EmailSignInLink(email string) {
	result, err := ClientAuth.EmailSignInLink(ctx, email, &auth.ActionCodeSettings{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
