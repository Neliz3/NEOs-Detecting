# NEOs-Detecting

#### Script to interact with NASA's Near Earth Object Web Service (NeoWs) to retrieve data about near-Earth objects (NEOs) detected within the last 7 days.

## Table of contents
* [Technologies Used](#Technologies-Used)
* [Installation](#Installation)
* [How to use it](#How-to-use-it)

  

## Technologies Used
* Go 1.22.2
* Docker


## How to run:

1. Using docker container:
    - Create your own .env file with keys listed in sample.env, then you can fill the values.
    - Set the specified values of build_args in docker-compose file.
    - Run docker-compose up to fetch and build images need for this project.
