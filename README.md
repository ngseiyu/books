# NG SEI YU

Backend stack: **Golang**, **Fiber** web framework, **GORM** with **sqlite** for storage
Frontend stack: **Vue JS** with **Tailwind**

The frontend app has been built, and its dist contents included in the backend code.

Can optionally deploy the backend as a docker container using the commands below in the backend directory:

```
docker build -t books .
docker run -d -p 8080:8080 books
```
