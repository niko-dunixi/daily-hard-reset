# Daily Hard Reset

Originally this was a utility to make sure my applications stayed in a
stable state, but I found that it was a good way to mark boundaries of the
work day and signal to myself that it was time to disconnect from work and
do the things I enjoy.

## Installation

1. First install GoLang to your machine by-way-of your preferred method for
your system.
1. `go get github.com/paul-nelson-baker/daily-hard-reset`
1. Optionally create the associated cron task to execute at the start/end
    of your working day.


If you wish to run this as part of your cron tasks, you need only download
the repository and run `go generate .`, here is an example:

```bash
# Initialize your crontab, this will remove any other entries. Please
# be mindful!
$ go generate . | crontab -

# Validate that everything worked, but fear not! GOPATH is automatically templated in for you!
$ crontab -l
> # ┌───────────── minute (0 - 59)
> # │ ┌───────────── hour (0 - 23)
> # │ │ ┌───────────── day of the month (1 - 31)
> # │ │ │ ┌───────────── month (1 - 12)
> # │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
> # │ │ │ │ │                                   7 is also Sunday on some systems)
> # │ │ │ │ │
> # │ │ │ │ │
> # * * * * * command to execute
>   0 8 * * 1-5 $GOPATH/bin/daily-hard-reset
>  05 17 * * 1-5 $GOPATH/bin/daily-hard-reset
```

### Usage

If you'd like to configure what applications are cycled automatically
all you need to do is create a yaml list of strings and place it in your
home directory. Here is an example of the default `~/.daily-hard-reset.yaml`

```yaml
- "Microsoft Outlook"
- "Slack"
- "zoom.us"
```