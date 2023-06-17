package ninjascript
import (
log_syslog "log/syslog"
)
func init() {if _, ok := Api["log/syslog"]; !ok {
   Api["log/syslog"] = map[string]interface{}{}
}
// Api["log/syslog"]["NewLogger"] = log_syslog.NewLogger
// Api["log/syslog"]["LOG_EMERG"] = log_syslog.LOG_EMERG
// Api["log/syslog"]["LOG_ALERT"] = log_syslog.LOG_ALERT
// Api["log/syslog"]["LOG_CRIT"] = log_syslog.LOG_CRIT
// Api["log/syslog"]["LOG_ERR"] = log_syslog.LOG_ERR
// Api["log/syslog"]["LOG_WARNING"] = log_syslog.LOG_WARNING
// Api["log/syslog"]["LOG_NOTICE"] = log_syslog.LOG_NOTICE
// Api["log/syslog"]["LOG_INFO"] = log_syslog.LOG_INFO
// Api["log/syslog"]["LOG_DEBUG"] = log_syslog.LOG_DEBUG
// Api["log/syslog"]["LOG_KERN"] = log_syslog.LOG_KERN
// Api["log/syslog"]["LOG_USER"] = log_syslog.LOG_USER
// Api["log/syslog"]["LOG_MAIL"] = log_syslog.LOG_MAIL
// Api["log/syslog"]["LOG_DAEMON"] = log_syslog.LOG_DAEMON
// Api["log/syslog"]["LOG_AUTH"] = log_syslog.LOG_AUTH
// Api["log/syslog"]["LOG_SYSLOG"] = log_syslog.LOG_SYSLOG
// Api["log/syslog"]["LOG_LPR"] = log_syslog.LOG_LPR
// Api["log/syslog"]["LOG_NEWS"] = log_syslog.LOG_NEWS
// Api["log/syslog"]["LOG_UUCP"] = log_syslog.LOG_UUCP
// Api["log/syslog"]["LOG_CRON"] = log_syslog.LOG_CRON
// Api["log/syslog"]["LOG_AUTHPRIV"] = log_syslog.LOG_AUTHPRIV
// Api["log/syslog"]["LOG_FTP"] = log_syslog.LOG_FTP
// Api["log/syslog"]["LOG_LOCAL0"] = log_syslog.LOG_LOCAL0
// Api["log/syslog"]["LOG_LOCAL1"] = log_syslog.LOG_LOCAL1
// Api["log/syslog"]["LOG_LOCAL2"] = log_syslog.LOG_LOCAL2
// Api["log/syslog"]["LOG_LOCAL3"] = log_syslog.LOG_LOCAL3
// Api["log/syslog"]["LOG_LOCAL4"] = log_syslog.LOG_LOCAL4
// Api["log/syslog"]["LOG_LOCAL5"] = log_syslog.LOG_LOCAL5
// Api["log/syslog"]["LOG_LOCAL6"] = log_syslog.LOG_LOCAL6
// Api["log/syslog"]["LOG_LOCAL7"] = log_syslog.LOG_LOCAL7
Api["log/syslog"]["Writer"] = log_syslog.Writer{}
Api["log/syslog"]["Dial"] = log_syslog.Dial
Api["log/syslog"]["New"] = log_syslog.New
}
