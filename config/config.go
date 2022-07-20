package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"okki.hu/garric/ppnext/store"
)

var Repository store.Repository = store.NewMongoRepository() // Repository for the application

var TlsEnabled = configValue("TLS_ENABLED", false, strconv.ParseBool)
var Port = configValue("PORT", 38080, strconv.Atoi)
var PublicPort = configValue("PUBLIC_PORT", Port, strconv.Atoi)

var Domain = configString("DOMAIN", "localhost")
var Support = configString("SUPPORT_EMAIL", "email@example.com")

var CleanupFrequencyMinutes = configValue("CLEANUP_FREQUENCY_MINUTES", 10, strconv.Atoi)
var CleanupFrequency = time.Duration(CleanupFrequencyMinutes) * time.Minute

var MaximumRoomAgeHours = configValue("CLEANUP_MAX_ROOM_AGE_HOURS", 12, strconv.Atoi)
var MaximumRoomAge = time.Duration(MaximumRoomAgeHours) * time.Hour

var AuthCookieName = configString("AUTH_COOKIE_NAME", "ppnext-user")

var AuthCookieExpiryHours = configValue("AUTH_COOKIE_EXPIRY_HOURS", 6, strconv.Atoi)
var AuthCookieExpiry = 60 * 60 * AuthCookieExpiryHours

var ShareUrlBase = func() string {
	scheme := "http"
	if TlsEnabled {
		scheme = "https"
	}
	portStr := fmt.Sprintf(":%d", PublicPort)
	if PublicPort == 80 || PublicPort == 443 {
		portStr = ""
	}
	return fmt.Sprintf("%s://%s%s/rooms/", scheme, Domain, portStr)
}()

func configValue[T any](envVarName string, defaultValue T, convertFn func(string) (T, error)) T {
	env := os.Getenv(envVarName)
	if env == "" {
		return defaultValue
	}
	val, err := convertFn(env)
	if err != nil {
		log.Fatalf("invalid value: %s=%s", envVarName, env)
	}
	return val
}

func configString(envVarName, defaultValue string) string {
	noop := func(s string) (string, error) {
		return s, nil
	}
	return configValue(envVarName, defaultValue, noop)
}
