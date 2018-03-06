## Consul Check Alerting

* This tool periodically checks the for failing service check in consul and send an slack message
* To set the consul url set `CONSUL_HTTP_ADDR` environment variable
### Help 
```
./consul-checks-slack -h:
     -interval int
           Periodic check interval in seconds (default 60)
     -slack-url string
           Use Slack incoming integration url (default "https://hooks.slack.com/services/xxxx/yyyyy/zzzzzzz")
```

