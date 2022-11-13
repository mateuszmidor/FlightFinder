# Flight Finder

Find flight connections between two given airports:  
- Using OpenAPI3 + go-gin-server generator.  
- Using gin-gonic web framework.   
- Using Redis for connections cache
- Using AWS CloudWatch for metrics:
  - this requires `cloudwatch:PutMetricData` IAM permission assigned on your machine/EC2 or no metrics will be sent
  - if configured correctly, generates metrics in `FlightFinder` CloudWatch Merics namespace

## Run locally from source (with connections cache)

```sh
./redis_cache/run_redis_cache.sh &  
go run cmd/finder_web/main.go -port=8080 -flights_data=./assets -web_data=./web -aws_region=us-east-1 -redis_addr=localhost:6379 -redis_pass=CACHE
```

## Run locally with docker (without connections cache)

```sh
docker build . -t mateuszmidor/flight-finder:latest
docker run --rm --name=flight-finder -p=8080:80 -v $HOME/.aws:/root/.aws mateuszmidor/flight-finder:latest
```

## Run locally as systemd service

```sh
sudo cp init/flight-finder.service /etc/systemd/system/
sudo systemctl daemon-reload 
sudo systemctl enable flight-finder
sudo systemctl start flight-finder
sudo systemctl status flight-finder
```

## Run on AWS EC2 or similar

```sh
#!/usr/bin/env sh

# install docker
curl -fsSL https://get.docker.com | sh

# install systemd service
cat << EOF > /etc/systemd/system/flight-finder.service
[Unit] 
Description=Flight Finder Web Server 
After=network.target 

[Service] 
Type=simple 
Restart=always  
ExecStart=docker run --rm --name=flight-finder -p=80:80 -v $HOME/.aws:/root/.aws mateuszmidor/flight-finder:latest
ExecStop=docker stop flight-finder 
                                   
[Install] 
WantedBy=multi-user.target
EOF

# run systemd service
systemctl daemon-reload    
systemctl enable flight-finder
systemctl start flight-finder
systemctl status flight-finder
```

## Run on AWS BeanStalk

- using platform GO  
    What BeanStalk does with your code uploaded as ZIP archive:
    - unzip the archive files
    - `go build application.go && ./application`
    - map host port 80 to application port 5000

    , so you just need to make a ZIP archive with the app and use it for creating BeanStalk application in AWS Console:

    ```sh
    zip -r flight-finder-platform_go-port5000-v1.zip assets/airports.csv.gz assets/nations.csv.gz assets/segments.csv.gz go.mod go.sum web/ pkg/ cmd/ application.go
    ```
- using platform Docker  
    What BeanStalk does with your code uploaded as ZIP archive:
    - unzip the archive files
    - `docker build . -t server`
    - `docker run -p=80:80 server`

    , so you just need to make a ZIP archive with the app and use it for creating BeanStalk application in AWS Console:
    ```sh
    zip -r flight-finder-platform_docker-port80-v1.zip assets/airports.csv.gz assets/nations.csv.gz assets/segments.csv.gz go.mod go.sum web/ scripts/ api/ pkg/ cmd/ Dockerfile
    ``` 

## Deploy using AWS CodeDeploy

- first setup variables in the script, then:
    ```sh
    ./aws_codedeploy/deploy.sh
    ```

## API

- list all airports
    ```sh
    curl 'http://localhost:8080/api/airports'
    ```

- find connections
    ```sh 
    curl 'http://localhost:8080/api/find?from=gdn&to=sez&maxsegmentcount=2'
    ```

## Showcase

### 2-segments connection: Kraków-Las Palmas

![Kraków-Sevilla-Las Palmas](./website/krk-svq-lpa.png)

### 3-segments connection: Kraków-Colombo

![Kraków-Amman-Muscat-Colombo](./website/krk-amm-mct-cmb.png)


## Design

## Structure

![Logo](website/structure.png)

## Sequence

![Logo](website/sequence.png)