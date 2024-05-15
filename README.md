# NEOs-Detecting

#### Script to interact with NASA's Near Earth Object Web Service (NeoWs) to retrieve data about near-Earth objects (NEOs) detected within the last 7 days.

## Table of contents
* [Technologies Used](#Technologies-Used)
* [How to run](#How-to-run)

  

## Technologies Used
* Go 1.22.2
* Docker


## How to run:
Run the following commands:
1. To create a Docker image: 
```docker build -f build Dockerfile -t neos .```
2. Go to https://api.nasa.gov/ and generate your own API key.
3. Enter a generated API key:
```docker run -e API_KEY='DEMO_KEY' neos ```
