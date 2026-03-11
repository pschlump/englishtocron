# Convert English to CRON syntax

This project is inspired by the library `natural-cron`, which converts natural language into cron expressions.   It uses a similar syntax.


## Features

- Convert English text descriptions into CRON syntax.
- Use complex patterns to recognize common time scheduling like time ranges.
- Support both AM/PM and 24-hour time formats.

## Supported English Patterns

| English Phrase | CronJob Syntax |
|------------------------------------------------------------------	|---------------------------- |
| every 15 seconds | 0/15 * * * * ? * |
| run every minute | 0 * * * * ? * |
| fire every day at 4:00 pm | 0 0 16 */1 * ? * |
| at 10:00 am | 0 0 10 * * ? * |
| run at midnight on the 1st and 15th of the month | 0 0 0 1,15 * ? * |
| On Sunday at 12:00 | 0 0 12 ? * SUN * |
| 7pm every Thursday | 0 0 19 ? * THU * |
| midnight on Tuesdays | 0 0 ? * TUE * |


