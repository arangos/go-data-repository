To build local image 

'docker build -t local_postgres .'

To run local image

´docker run --name postgres -p 5432:5432 -d local_postgres´

To connect to local postgres

user: postgres
password: postgres
port: 5432
database: postgres