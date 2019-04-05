# jobcan-fe

To press the jobcan's attendance button.
Web scraping with selenium via ChromeDriver made internal inside the container.
So we use GAE FE custom env.

# deploy

```
gcloud app deploy
```

# run local

```
docker build .
docker run -d -p 8080:8080 XXXXXX 
```

# endpoint

```http
POST http://localhost:8080/jobcan/touch
Accept: application/json

{"email": "your.jobcan.email@domain.com", "password": "your password"}
```

# requirements
- Go version 1.11 (go mod)

*Go 1.12 has not been verified*
