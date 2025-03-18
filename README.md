# Overview
Pretty straightforward app to get info from [vesselfinder.com](https://www.vesselfinder.com/)

For now supports only getting vessels information from defined map box and saving them to csv file periodically

## Config file example

Here is example of config file used by application

```
#Essential fields
Host=www.vesselfinder.com
BoxBot=33.521887
BoxTop=40.672535
BoxLeft=-92.697437
BoxRight=-63.154660
Zoom=18
OutputType=csv
OutputDirectory=/home/User/KaraOutput
TimerValueSecs=5
#Optional fields
UserAgent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36
```

## Docker image

Additionally you can pull a docker image:
```
docker pull greedygreen44/kara
```

Note, that running inside a container requires sharing local directories with it, like this:

```bash
docker run -t -v local/directory/with/configs:/Kara/configs:ro -v local/output/directory:Kara/output:rw greedygreen44/ailrun /Kara/configs/config.txt
```
Also you have to change OutputDirectory parameter in config file to match container internal output directory:

```
OutputDirectory=/Kara/output
```
