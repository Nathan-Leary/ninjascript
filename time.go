package ninjascript
import (
time "time"
)
func init() {if _, ok := Api["time"]; !ok {
   Api["time"] = map[string]interface{}{}
}
Api["time"]["Layout"] = time.Layout
Api["time"]["ANSIC"] = time.ANSIC
Api["time"]["UnixDate"] = time.UnixDate
Api["time"]["RubyDate"] = time.RubyDate
Api["time"]["RFC822"] = time.RFC822
Api["time"]["RFC822Z"] = time.RFC822Z
Api["time"]["RFC850"] = time.RFC850
Api["time"]["RFC1123"] = time.RFC1123
Api["time"]["RFC1123Z"] = time.RFC1123Z
Api["time"]["RFC3339"] = time.RFC3339
Api["time"]["RFC3339Nano"] = time.RFC3339Nano
Api["time"]["Kitchen"] = time.Kitchen
Api["time"]["Stamp"] = time.Stamp
Api["time"]["StampMilli"] = time.StampMilli
Api["time"]["StampMicro"] = time.StampMicro
Api["time"]["StampNano"] = time.StampNano
Api["time"]["DateTime"] = time.DateTime
Api["time"]["DateOnly"] = time.DateOnly
Api["time"]["TimeOnly"] = time.TimeOnly
Api["time"]["Nanosecond"] = time.Nanosecond
Api["time"]["Microsecond"] = time.Microsecond
Api["time"]["Millisecond"] = time.Millisecond
Api["time"]["Second"] = time.Second
Api["time"]["Minute"] = time.Minute
Api["time"]["Hour"] = time.Hour
Api["time"]["After"] = time.After
Api["time"]["Sleep"] = time.Sleep
Api["time"]["Tick"] = time.Tick
Api["time"]["ParseDuration"] = time.ParseDuration
Api["time"]["Since"] = time.Since
Api["time"]["Until"] = time.Until
Api["time"]["Location"] = time.Location{}
Api["time"]["Local"] = time.Local
Api["time"]["UTC"] = time.UTC
Api["time"]["FixedZone"] = time.FixedZone
Api["time"]["LoadLocation"] = time.LoadLocation
Api["time"]["LoadLocationFromTZData"] = time.LoadLocationFromTZData
Api["time"]["January"] = time.January
Api["time"]["February"] = time.February
Api["time"]["March"] = time.March
Api["time"]["April"] = time.April
Api["time"]["May"] = time.May
Api["time"]["June"] = time.June
Api["time"]["July"] = time.July
Api["time"]["August"] = time.August
Api["time"]["September"] = time.September
Api["time"]["October"] = time.October
Api["time"]["November"] = time.November
Api["time"]["December"] = time.December
Api["time"]["ParseError"] = time.ParseError{}
Api["time"]["Ticker"] = time.Ticker{}
Api["time"]["NewTicker"] = time.NewTicker
Api["time"]["Time"] = time.Time{}
Api["time"]["Date"] = time.Date
Api["time"]["Now"] = time.Now
Api["time"]["Parse"] = time.Parse
Api["time"]["ParseInLocation"] = time.ParseInLocation
Api["time"]["Unix"] = time.Unix
Api["time"]["UnixMicro"] = time.UnixMicro
Api["time"]["UnixMilli"] = time.UnixMilli
Api["time"]["Timer"] = time.Timer{}
Api["time"]["AfterFunc"] = time.AfterFunc
Api["time"]["NewTimer"] = time.NewTimer
Api["time"]["Sunday"] = time.Sunday
Api["time"]["Monday"] = time.Monday
Api["time"]["Tuesday"] = time.Tuesday
Api["time"]["Wednesday"] = time.Wednesday
Api["time"]["Thursday"] = time.Thursday
Api["time"]["Friday"] = time.Friday
Api["time"]["Saturday"] = time.Saturday
Api["time"]["After"] = time.After
Api["time"]["Sleep"] = time.Sleep
Api["time"]["Tick"] = time.Tick
Api["time"]["ParseDuration"] = time.ParseDuration
Api["time"]["Since"] = time.Since
Api["time"]["Until"] = time.Until
Api["time"]["FixedZone"] = time.FixedZone
Api["time"]["LoadLocation"] = time.LoadLocation
Api["time"]["LoadLocationFromTZData"] = time.LoadLocationFromTZData
Api["time"]["NewTicker"] = time.NewTicker
Api["time"]["Date"] = time.Date
Api["time"]["Now"] = time.Now
Api["time"]["Parse"] = time.Parse
Api["time"]["ParseInLocation"] = time.ParseInLocation
Api["time"]["Unix"] = time.Unix
Api["time"]["UnixMicro"] = time.UnixMicro
Api["time"]["UnixMilli"] = time.UnixMilli
Api["time"]["AfterFunc"] = time.AfterFunc
Api["time"]["NewTimer"] = time.NewTimer

}