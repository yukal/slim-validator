## Benchmark Test

| Strategy                         | Executions | ns/op   | B/op | allocs/op
| :------------------------------- | :--------: | ------: | :--: | :-------:
| Validate (min): int              |   100000   |  594.6  |  432 |  6
| Validate (min): string           |   100000   |  602.1  |  456 |  6
| Validate (min): slice            |   100000   |  610.8  |  440 |  6
| Validate (min): map              |   100000   |  646.1  |  440 |  6
| Validate (date : min): int64     |   100000   |  496.6  |  408 |  5
| Validate (date : min): string    |   100000   |  569.8  |  408 |  5
| Validate (date : min): time      |   100000   |  497.1  |  408 |  5
| Validate (time : min): int64     |   100000   |  489.6  |  408 |  5
| Validate (time : min): string    |   100000   |  523.6  |  408 |  5
| Validate (time : min): time      |   100000   |  508.0  |  408 |  5
| Validate (range): int            |   100000   |  697.3  |  440 |  6
| Validate (range): string         |   100000   |  716.1  |  440 |  6
| Validate (range): slice          |   100000   |  715.6  |  432 |  6
| Validate (range): map            |   100000   |  764.5  |  432 |  6
| Validate (match): string         |   100000   |  670.8  |  504 |  7
| Validate (each : min): int       |   100000   |  663.9  |  440 |  7
| Validate (each : min): string    |   100000   |  635.9  |  456 |  6
| Validate (each : min): slice     |   100000   |  642.9  |  440 |  6
| Validate (each : min): map       |   100000   |  651.3  |  440 |  6
| Validate (each : range): int     |   100000   |  861.6  |  472 | 12
| Validate (each : range): string  |   100000   |  823.7  |  448 | 10
| Validate (each : range): slice   |   100000   |  847.2  |  440 | 10
| Validate (each : range): map     |   100000   |  837.7  |  440 | 10
| Validate (each : match): slice   |   100000   | 1913.0  | 1786 | 22
| Validate (each : match): map     |   100000   | 2107.0  | 1798 | 23
