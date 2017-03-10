# Overview

Uses Google's Natural Language processing API to analyze tweets mentioning **Trump**.

# Quickstart
```bash
go get ./...
go build .
./gcp-nlp-example # add .exe for Windows
```
# Sample Output
```
2017/03/10 09:53:21 Starting Stream...

2017/03/10 09:53:23 Entity name:  RT @InxsyS
2017/03/10 09:53:23 Entity name:  Trump
2017/03/10 09:53:23 Entity name:  'T
2017/03/10 09:53:23 Entity name:  Repugnants
2017/03/10 09:53:23 Entity name:  LIE
2017/03/10 09:53:23 Entity name:  Sphincter
2017/03/10 09:53:23 Entity name:  LIE
2017/03/10 09:53:23 Entity name:  ADMINISTRATION
2017/03/10 09:53:23 Entity name:  PUBLIC
2017/03/10 09:53:23 Entity name:  AMERICAN
2017/03/10 09:53:23 RT @InxsyS: Trump,Sphincter &amp;  All Repugnants- YOU DON'T GET TO TELL LIE AFTER LIE TO THE AMERICAN PUBLIC--THEN SAY YOUR ADMINISTRATION IS… is NEGATIVE with score -0.600000


2017/03/10 09:53:26 Entity name:  Alec Baldwin
2017/03/10 09:53:26 Entity name:  This Is the One
2017/03/10 09:53:26 Entity name:  Donald Trump
2017/03/10 09:53:26 Entity name:  https://t.co/XXarEybxTE https://t.co/FnG55iSbvn
2017/03/10 09:53:26 This Is the One Thing Alec Baldwin Likes About Donald Trump https://t.co/XXarEybxTE https://t.co/FnG55iSbvn is POSITIVE with score 0.700000


2017/03/10 09:53:28 Entity name:  RT @CBSThisMorning
2017/03/10 09:53:28 Entity name:  Trump
2017/03/10 09:53:28 Entity name:  Sean Spicer
2017/03/10 09:53:28 Entity name:  White House
2017/03/10 09:53:28 Entity name:  investigation
2017/03/10 09:53:28 Entity name:  Justice Department
2017/03/10 09:53:28 RT @CBSThisMorning: Sean Spicer says White House is "not aware" of any Justice Department investigation into President Trump. https://t.co/… is POSITIVE with score 0.200000


2017/03/10 09:53:34 Entity name:  A Brand Name
2017/03/10 09:53:34 Entity name:  Hedge Fund Happy Hour: Trump
2017/03/10 09:53:34 Entity name:  Mar-a-Lago
2017/03/10 09:53:34 Entity name:  NYT
2017/03/10 09:53:34 Entity name:  The New York Times https://t.co/tBP18mNEb9
2017/03/10 09:53:34 Entity name:  ALEXANDRA STEVENSON
2017/03/10 09:53:34 "A Brand Name for a Hedge Fund Happy Hour: Trump’s Mar-a-Lago" by ALEXANDRA STEVENSON via NYT The New York Times https://t.co/tBP18mNEb9 is POSITIVE with score 0.600000
```

