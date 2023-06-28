FROM golang
EXPOSE 8600
WORKDIR /app
ENV GO111MODULE=on
#Copiar arquivos para dependencias
COPY go.mod . 
COPY go.sum .
RUN go mod download
#Copiar código
COPY . .
#RUN go get ./...
RUN go build -o alunos-api
CMD [ "./alunos-api" ]
