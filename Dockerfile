FROM alpine

WORKDIR /usr/src/maps

COPY . .

EXPOSE 5000

CMD [ "./tcp/server/server" ]

