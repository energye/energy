//go:build windows

package notification

import (
	"testing"

	. "github.com/energye/energy/v3/platform/notification/types"
)

func TestNew(t *testing.T) {
	n := New()
	if n == nil {
		t.Fatal("New() returned nil")
	}
}

func TestInitialize(t *testing.T) {
	n := New()
	err := n.Initialize()
	if err != nil {
		t.Logf("Initialize error (may be expected in test environment): %v", err)
	}
}

func TestRequestNotificationAuthorization(t *testing.T) {
	n := New()
	authorized, err := n.RequestNotificationAuthorization()
	if err != nil {
		t.Fatalf("RequestNotificationAuthorization failed: %v", err)
	}
	if !authorized {
		t.Fatal("Expected authorization to be true on Windows")
	}
}

func TestCheckNotificationAuthorization(t *testing.T) {
	n := New()
	authorized, err := n.CheckNotificationAuthorization()
	if err != nil {
		t.Fatalf("CheckNotificationAuthorization failed: %v", err)
	}
	if !authorized {
		t.Fatal("Expected authorization to be true on Windows")
	}
}

func TestSendNotification(t *testing.T) {
	n := New()
	err := n.SendNotification(Options{
		ID:    "test-1",
		Title: "Test Notification",
		Body:  "This is a test notification",
	})
	if err != nil {
		t.Logf("SendNotification error (may be expected in test environment): %v", err)
	}
}

func TestRegisterNotificationCategory(t *testing.T) {
	n := New()
	err := n.RegisterNotificationCategory(Category{
		ID: "test-category",
		Actions: []Action{
			{ID: "action1", Title: "Action 1"},
			{ID: "action2", Title: "Action 2"},
		},
	})
	if err != nil {
		t.Fatalf("RegisterNotificationCategory failed: %v", err)
	}
}

func TestRemoveNotificationCategory(t *testing.T) {
	n := New()
	err := n.RemoveNotificationCategory("test-category")
	if err != nil {
		t.Fatalf("RemoveNotificationCategory failed: %v", err)
	}
}

func TestRemoveAllPendingNotifications(t *testing.T) {
	n := New()
	err := n.RemoveAllPendingNotifications()
	if err != nil {
		t.Fatalf("RemoveAllPendingNotifications failed: %v", err)
	}
}

func TestRemovePendingNotification(t *testing.T) {
	n := New()
	err := n.RemovePendingNotification("test-id")
	if err != nil {
		t.Fatalf("RemovePendingNotification failed: %v", err)
	}
}

func TestRemoveAllDeliveredNotifications(t *testing.T) {
	n := New()
	err := n.RemoveAllDeliveredNotifications()
	if err != nil {
		t.Fatalf("RemoveAllDeliveredNotifications failed: %v", err)
	}
}

func TestRemoveDeliveredNotification(t *testing.T) {
	n := New()
	err := n.RemoveDeliveredNotification("test-id")
	if err != nil {
		t.Fatalf("RemoveDeliveredNotification failed: %v", err)
	}
}

func TestRemoveNotification(t *testing.T) {
	n := New()
	err := n.RemoveNotification("test-id")
	if err != nil {
		t.Fatalf("RemoveNotification failed: %v", err)
	}
}
