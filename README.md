### GoSpark [![GoDoc](https://godoc.org/github.com/omgnuts/gospark?status.svg)](http://godoc.org/github.com/omgnuts/gospark) [![Build Status](https://travis-ci.org/omgnuts/gospark.svg?branch=master)](https://travis-ci.org/omgnuts/gospark)

##### A small sparkline utility inspired by [holman/spark](https://github.com/holman/spark)

Pretty cool way to display small signals in a heartbeat. Here's a neat
[list of awesome usage ideas](https://github.com/holman/spark/wiki/Wicked-Cool-Usage) for spark.

~~~go
import (
  "github.com/omgnuts/gospark"
)

func main() {
  series := []float64{ 67, 71, 77, 85, 95, 104, 106, 105, 100, 89, 76, 66 }
  signals := gospark.Signals(series)
  fmt.Println(signals)
}

output: "▁▂▃▄▆███▇▅▃▁"
~~~

Another mini example for stock price signals

~~~go
func main() {
  prices := []float64{1398.56,1360.16,1394.46,1409.28,1409.12,1424.97,1424.37,1424.24,1441.72,1411.7,1416.83,1387.12,1389.94,1402.05,1387.67,1388.26,1346.09,1346.09,1352.17,1360.69,1353.43,1333.36,1348.05,1366.42,1379.19,1381.76,1409.17,1391.28,1355.62,1366.7,1401.69,1395.07,1383.62,1359.15,1392.15}
  signals := gospark.Signals(prices)
  fmt.Println(signals)
}

output: "▅▂▅▆▆▇▇▇█▆▇▄▅▆▅▅▁▁▂▃▂▁▂▃▄▄▆▅▂▃▆▅▄▂▅"
~~~

Works from the console too
~~~go
# gospark 10 15 6 23 5 0 0 1 15 0 17 3 0 0

output: "▄▆▃█▂▁▁▁▆▁▆▂▁▁"
~~~

Ok that's all folks. Enjoy!

MIT License
 
