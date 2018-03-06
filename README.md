## Consul Check Alerting

* This tool periodically checks the for failing service check in consul and send an slack message
* To set the consul url set `CONSUL_HTTP_ADDR` environment variable
### Help 
```
./consul-checks-slack -h:
    -check-interval int
        	Periodic check interval in seconds (default 60)
      -slack-url string
        	Use Slack incoming integration url (default "https://hooks.slack.com/services/xxxx/yyyy/zzzzzz")
      -wait-interval int
        	Greater than this period is considered as check failure seconds (default 300)

```

