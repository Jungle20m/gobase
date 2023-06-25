FROM mgo:1.20.5

WORKDIR /otp/source/gobase

COPY . .

CMD ["go", "run", "main.go"]


