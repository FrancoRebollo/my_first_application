package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NotificationsContactMe(w http.ResponseWriter, r *http.Request) error {
	notificationContact := new(NotificationContact)

	if err := json.NewDecoder(r.Body).Decode(notificationContact); err != nil {
		return err
	}

	commonID := r.Context().Value("commonIdentification").(string)

	if commonID != notificationContact.NotificationSender {
		return fmt.Errorf("who the fuck are u ? ")
	}

	return nil
}
