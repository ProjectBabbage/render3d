# render3d, Project Babbage first edition.

The idea for this edition was to build a 3D renderer to understand the important concepts in this field.
We only used CPU power so the rendering is not expected to be time efficient: it was not the goal anyway.
Look at that beautiful red and glass material below !

## Result example

![img](assets/img/FinalScene.png)

## Coordinate system used

Same system as SDL.

    o -- + X
    |
    +
    Y

Z follows the right hand rule (goes up from your eyes to the screen)

Origin (o) is where the Eye is : (0, 0, 0).

## Scene configuration

Check the default config at config/config.go.
The screen is at D (+Z axis) from the eye (D=2 in the default config).
Hence you need to put the object after Z=2 to see it on the rendered image.
SaveAsPNG is at false by default. Set to true to save the rendered image to output/rendered.png.

## Installation

Install golang, then:
(Ubuntu working example)
```
sudo apt install libsdl2-2.0-0 libsdl2-dev

go install
```
then do `go run main.go`

## Testing

Run the tests with `go test ./...`

Run the benchs (and tests) with `go test ./... -bench .`

Run the tests without display tests with `go test -short ./...`

Run a specific test (here called TestNewScreen):

`go test -run TestNewScreen broengine/render`

you must specify the package it's in (broegine/render)

## Documentation

To generate a package doc (render package here), do:
```
go doc -all render
```
