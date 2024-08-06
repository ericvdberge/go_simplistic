# Incredibale go + htmx website

## Build and run docker file
```
docker build -f deployment/Dockerfile -t go_test .
docker run --env-file .env -p 8080:8080 go_test 
```

## Docker image size (14MB)
![image](./docs/image_size.png)

## Memory usage (4.7MB)
![image](./docs//memory%20usage.png)