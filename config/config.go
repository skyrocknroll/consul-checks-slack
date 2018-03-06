package config

import "flag"

var SlackUrl *string = flag.String("slack-url", "https://hooks.slack.com/services/xxxx/yyyy/zzzzzz", "Use Slack incoming integration url")
var CheckInterval *int = flag.Int("check-interval", 60, "Periodic check interval in seconds")
var WaitInterval *int64 = flag.Int64("wait-interval", 300, "Greater than this period is considered as check failure seconds")

func init() {
	flag.Parse()
}
