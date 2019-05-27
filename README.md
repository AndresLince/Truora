# Truora
Test for truora

#Comandos para ejecutar adentro de la carpeta de frontend

npm install

#Comandos para ejecutar en la consola adentro de la carpeta de cockroach:
cockroach start --insecure --listen-addr=localhost
cockroach sql --insecure
CREATE USER IF NOT EXISTS maxroach;
CREATE DATABASE test;
GRANT ALL ON DATABASE test TO maxroach;

#Comando para ejecutar adentro de la carpeta backend
go get
go run *.go