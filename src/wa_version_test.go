package main

import (
	"testing"

	"go.mau.fi/whatsmeow/store"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func TestWhatsAppWebVersionUpdatesTheLoginPayload(t *testing.T) {
	previous := store.GetWAVersion()
	t.Cleanup(func() { store.SetWAVersion(previous) })
	for _, test := range []struct{ override, expected string }{
		{"", defaultWhatsAppWebVersion},
		{"2.3000.1054751087", "2.3000.1054751087"},
	} {
		t.Setenv("WAHA_GOWS_WA_VERSION", test.override)
		if err := configureWhatsAppVersion(waLog.Noop); err != nil {
			t.Fatal(err)
		}
		if actual := store.GetWAVersion().String(); actual != test.expected {
			t.Fatalf("client version %s, expected %s", actual, test.expected)
		}
		payload := store.BaseClientPayload.GetUserAgent().GetAppVersion()
		version := store.GetWAVersion()
		if payload.GetPrimary() != version[0] || payload.GetSecondary() != version[1] || payload.GetTertiary() != version[2] {
			t.Fatal("login payload still uses the outdated version")
		}
	}
}

func TestWhatsAppWebVersionRejectsInvalidOverrides(t *testing.T) {
	previous := store.GetWAVersion()
	t.Cleanup(func() { store.SetWAVersion(previous) })
	for _, invalid := range []string{"invalid", "2.3000.-1", "2.3000.9999999999999", "0.0.0"} {
		t.Setenv("WAHA_GOWS_WA_VERSION", invalid)
		if err := configureWhatsAppVersion(waLog.Noop); err == nil {
			t.Fatalf("accepted invalid version %q", invalid)
		}
		if store.GetWAVersion() != previous {
			t.Fatal("invalid override changed the client version")
		}
	}
}
