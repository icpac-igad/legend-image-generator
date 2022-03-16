# Map Legend Image Generator

Generate a legend image for a map from a json configuration.

This was developed to generate static legend images for some layers on [EAHW](https://eahazardswatch.icpac.net) as an alternative to generating directly with HTML.

## Sample Output

Sample output

![Alt text](sample/legend_sample.png "Legend Sample")

## Dependencies

The Legend Image Generator service is built using [Go](https://go.dev/) and  can be executed either natively or using Docker, each of which has its own set of requirements.

Native execution requires:
- [Go](https://go.dev/)

Execution using Docker requires:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)


## Getting started

Start by cloning the repository from github to your execution environment

```
git clone https://github.com/icpac-igad/legend-image-generator.git && cd legend-image-generator
```

After that, follow one of the instructions below:

### Using native execution

1 - Install go dependencies using go:
```
go get

```

2 - Start the application server:
```
go run main.go
```

The endpoints provided by this microservice should now be available on [http:localhost:9000](http:localhost:9000)

Only the root endpoint is currently exposed i.e `/`


### Using Docker
1 - Create and complete your `.env`. You can find an example `.env.sample` file in the project root.

2 - Build the the image

`docker-compose build`

3 - Run the container

`docker-compose up`

### Environment variables

- PORT => Target TCP port in which the service will run
- RESTART_POLICY => Docker container restart policy. You will want this set to `always` in production


## usage

The `/` endpoint only accepts `POST` requests with a `json` payload that should look like below:

```
{
    "legend_type": "discrete",
    "width": 600,
    "item_height": 30,
    "label_size": 20,
    "label_space": 26,
    "transparent": true,
    "items": [
        {
            "color": "#b4b4b4",
            "value": "10%"
        },
        {
            "color": "#BB3426",
            "value": "20%"
        },
        {
            "color": "#D44F24",
            "value": "30%"
        },
        {
            "color": "#EF7424",
            "value": "40%"
        },
        {
            "color": "#FFA229",
            "value": "50%"
        },
        {
            "color": "#FFBD2D",
            "value": "60%"
        },
        {
            "color": "#FFD931",
            "value": "70%"
        },
        {
            "color": "#FFF636",
            "value": "80%"
        },
        {
            "color": "#C6FF35",
            "value": "90%"
        },
        {
            "color": "#75FF32",
            "value": "100%"
        }
    ]
}
```

- legend_type - `string` -  The type of legend you wish to generate. Currently not implemented. We have it here for future addition of other legends types.
- items - `list` - a list of key value pairs of the legend items. The required keys are :
    - color - `string` - the hex color
    - value - `string` - the legend item label
- transparent - `boolean` - Set `true` to generate the image with a transparent background. White background by default
- width - `int` - width of the legend. If not provided, 500 is used by default
- item_height - `int` - height of the color bar. 20 by default
- label_size -`int` - font size of the labes. 12 by default
- label_space - `int` - space between color bars and labels. 5 by default.

This sample configuration will give you an output similar to the example shown above.

Adjust the configuration as necessary to achieve the desired legend images.