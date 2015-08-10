FROM       scratch
MAINTAINER Kelsey Hightower <kelsey.hightower@gmail.com>
ADD        inspector inspector
ADD        css       css
ENTRYPOINT ["/inspector"]
