package utils

import (
	"github.com/google/uuid"
	"os"
)

func GenerateUnsubscribeLink(subscriptionID uuid.UUID) string {
	appUrl := os.Getenv("APP_URL")
	link := appUrl + "/api/v1/subscriptions/" + subscriptionID.String() + "/_unsubscribe"
	return "<br><br><a href=\"" + link + "\">Unsubscribe</a>"
}
