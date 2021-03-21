# CRON

[wikipedia](https://en.wikipedia.org/wiki/Cron)

## Syntax

```sh
┌───────────── minute (0 - 59)
│ ┌───────────── hour (0 - 23)
│ │ ┌───────────── day of the month (1 - 31)
│ │ │ ┌───────────── month (1 - 12)
│ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
│ │ │ │ │                                   7 is also Sunday on some systems)
│ │ │ │ │
│ │ │ │ │
* * * * * <command to execute>
```

## Expression

| Field        | Required | Allowed values  | Allowed special characters | Remarks                                                          |
| ------------ | -------- | --------------- | -------------------------- | ---------------------------------------------------------------- |
| Minutes      | Yes      | 0-59            | \* , -                     |                                                                  |
| Hours        | Yes      | 0-23            | \* , -                     |
| Day of month | Yes      | 1-31            | \* , - ? L W ?             | L W only in some implementations                                 |
| Month        | Yes      | 1-12 or JAN-DEC | \* , -                     |
| Day of week  | Yes      | 0-6 or SUN-SAT  | \* , - ? L #               | ? L # ? L # only in some implementations                         |
| Year         | No       | 1970-2099       | \* , -                     | This field is not supported in standard/default implementations. |

- Comma ( , )

  Commas are used to separate items of a list. For example, using "MON,WED,FRI" in the 5th field (day of week) means Mondays, Wednesdays and Fridays

- Dash( - )

  Dash defines ranges. For example, 2000-2010 indicates every year between 2000 and 2010, inclusive

- Percent ( % )

  Percent-signs (%) in the command, unless escaped with backslash (\), are changed into newline characters, and all data after the first % are sent to the command as standard input.

- Question mark (?)

  In some implementations, used instead of '\*' for leaving either day-of-month or day-of-week blank. Other cron implementations substitute "?" with the start-up time of the cron daemon.

- Slash (/)

  For example, \*/5 in the minutes field indicates every 5 minutes (see note below about frequencies). It is shorthand for the more verbose POSIX form 5,10,15,20,25,30,35,40,45,50,55,00.

## Examples

- every minute

```sh
* * * * *
```

- every 5 minutes

```sh
*/5 * * * *
```

- every 2 hours

```sh
* */2 * * *
```

- 2am daily

```sh
0 2 * * *
```

- every Sunday at 5 PM

```sh
0 17 * * sun
```
