# Overview

Uses Google's Natural Language processing API to analyze tweets mentioning **Trump**.

# Sample Sentiment Output
Formatted to be readable:
```
2017/03/09 14:46:55 RT @JoyAnnReid: 
In today's episode of #TheKleptocrats ... https://t.co/8YXN65lAFA is POSITIVE with score 0.300000

2017/03/09 14:47:00 RT @blakehounshell: 
If the Obamacare repeal/replace fails, Trump tells conservatives that Plan B is to just blame Democrats https://t.co/XX… is NEGATIVE with score -0.500000

2017/03/09 14:47:05 RT @jk_rowling: 
'I have tremendous respect for women and the many roles they serve that are vital to the fabric of our society &amp; ou…  is POSITIVE with score 0.100000

2017/03/09 14:47:11 
This is the leader of the free world, everybody. https://t.co/iUi726WJuR is POSITIVE with score 0.300000

2017/03/09 14:47:16 RT @saletan: 
But please don't take this as an indication that anything might turn up if we investigate Trump associates' busines…  is POSITIVE with score 0.400000
```

# What happens when we combine it with entites / subject matter

```
2017/03/09 17:03:45 RT @immigrant4trump: Tim Kaine’s Son Arrested For Rioting At Pro-Trump Rally #maga #Trump https://t.co/SaDyJkIOMP is POSITIVE with score 0.600000
2017/03/09 17:03:51 Entity name:  RT @hrkbenowen
2017/03/09 17:03:51 Entity name:  Obama
2017/03/09 17:03:51 Entity name:  Jobs
2017/03/09 17:03:51 Entity name:  Jobs Are Just Not Coming Back
2017/03/09 17:03:51 Entity name:  First Month https://t.co/IPyNpnHlK6
2017/03/09 17:03:51 RT @hrkbenowen: Flashback: Obama Says "Jobs Are Just Not Coming Back" ...Then Trump Adds 298,000 Jobs in First Month https://t.co/IPyNpnHlK6 is NEGATIVE with score -0.600000
2017/03/09 17:03:56 Entity name:  RT @JohnWDean
2017/03/09 17:03:56 Entity name:  Trump
2017/03/09 17:03:56 Entity name:  moles
2017/03/09 17:03:56 Entity name:  posts
2017/03/09 17:03:56 Entity name:  Deep State
2017/03/09 17:03:56 Entity name:  Senate
2017/03/09 17:03:56 Entity name:  executive bureaucracy w
2017/03/09 17:03:56 RT @JohnWDean: Trump's "Deep State:" 100s of Trump moles infiltrate the executive bureaucracy w/o filling Senate confirmable posts: https:/… is NEGATIVE with score -0.100000
```

