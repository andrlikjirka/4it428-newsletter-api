package utils

import "os"

func GenerateUnsubscribeLink(subscriptionID string) string {
	appUrl := os.Getenv("APP_URL")
	link := appUrl + "/api/v1/subscriptions/" + subscriptionID + "/_unsubscribe"
	return "<br><br><a href=\"" + link + "\">Unsubscribe</a>"
}
