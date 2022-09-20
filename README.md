# Timesheet Tools

A simple command-line application to calculate statistics from CSV timesheets. Written in Go.


## Why?

I was sick of manually tracking and managing my working hours with Excel / LibreOffice Calc. Of course there are a lot
of "professional" time tracking tools out there. I don't care. I wanted something simple that works exclusively on the
command-line and offline.


## Expected CSV Timesheet Format

### Example
```
2022-09-01	10:00	10:35	A	Working on Ticket 1
2022-09-01	10:45	12:10	A	Working on Ticket 1
2022-09-01	13:00	14:55	B	Working on the second Ticket
```
### Columns
1. **Date** in YYYY-MM-DD format
2. **Start Time** in HH:MM format
3. **End Time** in HH:MM format
4. **Bucket** duration sums will be calculated per bucket
5. **Comment**

### Delimiters
* columns must be delimited by a tabulator `\t` (not by comma `,` or semicolon `;`)
* line endings must be Linux line endings `\n` (no Windows `\r\n` or Mac `\r`)
* blank lines will be ignored (they might be added between different days for better readability)


## Build Instructions

You can use Docker to build the tool:
```
$ docker run -it --name timesheet_tools -v "$PWD":/usr/src/timesheet_tools -w /usr/src/timesheet_tools golang:1.18 bash

root@e35954335070:/usr/src/timesheet_tools# cd timesheet-analyzer
root@e35954335070:/usr/src/timesheet_tools/timesheet-analyzer# go build -o ../bin/tms-analyze main.go
```

## Usage

The CSV content should be piped via STDIN into the analyzer:
```
$ cat examples/timesheet1.csv | bin/tms-analyze

0: {2022-09-01 10:00:00 +0000 UTC 2022-09-01 10:35:00 +0000 UTC A #1 Bucket Duration}
1: {2022-09-01 10:45:00 +0000 UTC 2022-09-01 12:10:00 +0000 UTC A #1 Bucket Duration}
2: {2022-09-01 13:00:00 +0000 UTC 2022-09-01 14:55:00 +0000 UTC B #2 Overall Duration}

--------------- Durations per Bucket ---------------
B: 01:55
A: 02:00

--------------- Overall Duration ---------------
03:55
```


## Author

* Ralf Henze
