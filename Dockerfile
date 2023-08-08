FROM alpine

RUN apk add tzdata curl bash

ENV GO111MODULE=on CGO_ENABLED=0  GOOS=linux GOARCH=amd64

WORKDIR /app

RUN mkdir data

COPY ./admin/main main
COPY ./data/data/* ./data/

EXPOSE 32325

CMD ["/app/main"]
#CMD ["/bin/bash"]
#CMD ["tail -f /dev/null"]