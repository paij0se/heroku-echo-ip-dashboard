package main

import (
	"testing"

	HerokuEchoIpDashboard "github.com/paij0se/heroku-echo-ip-dashboard/src"
)

func Test(testing *testing.T) {
	HerokuEchoIpDashboard.HerokuEchoIpDashboard("8080")
}
