package main 

import "fmt"
import "time"

var p = fmt.Println
func main() {

	now := time.Now()
	// 秒
	secs := now.Unix()
	// 纳秒 
	nanos := now.UnixNano()
	p(now)
	p(secs)
	p(nanos)
	p(now.Format(time.RFC3339))

	t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)
    p(now.Format("3:04PM"))
    p(now.Format("Mon Jan _2 15:04:05 2006"))
    p(now.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    if(e != nil) {
    	p(e)
    } else {
    	p(t2)
    }
	p("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        now.Year(), now.Month(), now.Day(),
        now.Hour(), now.Minute(), now.Second())

	then := time.Date(2009,11,17,20,34,58,651387237,time.UTC)
	p(then)
	p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
    p(then.Weekday())

    // 这几个方法用来比较两个时间,精确到秒
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // 方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
    diff := now.Sub(then)
    p(diff)
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // Add 将时间后移一个时间间隔，或者使用一个 - 来将时间前移一个时间间隔
    p(then.Add(diff))
    p(then.Add(-diff))

}

/*
2018-03-21 14:23:05.25266291 +0800 CST m=+0.000281443
1521613385
1521613385252662910
2018-03-21T14:23:05+08:00
2012-11-01 22:08:41 +0000 +0000
2:23PM
Wed Mar 21 14:23:05 2018
2018-03-21T14:23:05.252662+08:00
0000-01-01 20:41:00 +0000 UTC
%d-%02d-%02dT%02d:%02d:%02d-00:00
 2018 March 21 14 23 5
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
73089h48m6.601275673s
73089.80183368768
4.3853881100212615e+06
2.6312328660127568e+08
263123286601275673
2018-03-21 06:23:05.25266291 +0000 UTC
2001-07-17 10:46:52.050111564 +0000 UTC
*/