docker build . -t go-containerized:latest (build docker)
docker run -e PORT=4000 -p 4000:4000 go-containerized:latest  (run docker with port)
docker compose up (start)
make pro (use make file)