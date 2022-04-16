## Super simple ToDo built in Go lang

This repo is just for the simplest study case of the Go lang ...

### Run

Just use docker it`s easy!

```bash
$> docker build -t gotodo .
$> docker run -d -p 80:8080 --name gotodo gotodo
```

Or, the simplest way ...

```bash
$> docker-compose -f stack.yml up
```

After that ... just head to your browser (http://localhost) ... and enjoy a simple and dummy ToDo.

Cheers! ;)