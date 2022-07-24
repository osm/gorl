# gorl

CLI for interacting with Reolink cameras.

## Features

* List recordings
* Download recordings
* Snap pictures
* Stream

## Example

```sh
# Prepare the environment with address and credentials
$ export GORL_ADDRESS=192.168.1.1
$ export GORL_USERNAME=<username>
$ export GORL_PASSWORD=<password>

# List recordings for todays date
$ gorl ls
Time                    Duration        Name
2021-12-30 08:39:01     44s             Rec_20211230_073901_411_M.mp4

# List recordings for a specific date
$ gorl ls -d 2021-12-29
Time                    Duration        Name
2021-12-29 08:40:29     44s             Rec_20211229_074029_411_M.mp4
2021-12-29 10:54:20     32s             Rec_20211229_095420_411_M.mp4
2021-12-29 11:10:10     32s             Rec_20211229_101010_411_M.mp4

# List recordings for the given date range
$ gorl ls -d 2021-12-29 -e 2021-12-30
Time                    Duration        Name
2021-12-29 08:40:29     44s             Rec_20211229_074029_411_M.mp4
2021-12-29 10:54:20     32s             Rec_20211229_095420_411_M.mp4
2021-12-29 11:10:10     32s             Rec_20211229_101010_411_M.mp4
2021-12-30 08:39:01     44s             Rec_20211230_073901_411_M.mp4

# Download a recording
$ gorl dl -f Rec_20211229_095420_411_M.mp4 >Rec_20211229_095420_411_M.mp4

# Snap a picture
$ gorl snap >current.jpg

# Start a stream
$ gorl stream | mplayer -
```

## Support

The program has been tested with the following models

* RLC-410W with firmware v3.0.0.136_20121102
* RLC-410W with firmware v3.1.0.739_22042505
