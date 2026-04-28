package cmd

import "github.com/jibankumarpanda/gogcli.1/internal/googleapi"

var newCalendarService = googleapi.NewCalendar

const (
	scopeAll    = literalAll
	scopeSingle = "single"
	scopeFuture = "future"
)
