FROM golang

WORKDIR /app/news-api

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go build -o ./bin/news-api .

CMD [ "./bin/news-api" ]