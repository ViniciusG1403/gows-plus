package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"go.mau.fi/whatsmeow/store"
	waLog "go.mau.fi/whatsmeow/util/log"
)

// Compatibility update for WhatsApp's ClientOutdated (405) rejection.
// Source: tulir/whatsmeow store/clientpayload.go, checked on 2026-11-17.
const defaultWhatsAppWebVersion = "2.3000.1054751086"

func configureWhatsAppVersion(log waLog.Logger) error {
	version := strings.TrimSpace(os.Getenv("WAHA_GOWS_WA_VERSION"))
	if version == "" {
		version = defaultWhatsAppWebVersion
	}
	if !regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`).MatchString(version) {
		return fmt.Errorf("WAHA_GOWS_WA_VERSION must contain three unsigned numeric parts")
	}
	for _, part := range strings.Split(version, ".") {
		if _, err := strconv.ParseUint(part, 10, 32); err != nil {
			return fmt.Errorf("WAHA_GOWS_WA_VERSION contains an out-of-range part")
		}
	}
	parsed, err := store.ParseVersion(version)
	if err != nil || parsed.IsZero() {
		return fmt.Errorf("WAHA_GOWS_WA_VERSION is not a valid client version")
	}
	store.SetWAVersion(parsed)
	log.Infof("WhatsApp Web client version: %s", store.GetWAVersion().String())
	return nil
}
