# inventory_pot



Honeypot implements a vulnerable fake inventory system.

1. SQL Injection.
login: `admin'--`
password: `any`

2. BROKER ACCESS. `/admin` route doesn't protect.

3. MISCONFIG. `/debug` route is open.

4. VULN JQUERY - `<script src="https://code.jquery.com/jquery-1.6.1.min.js"></script>`.

### Build

Docker:
```bash
docker built -t inventory_app -f docker/Dockerfile
```

From src:
```bash
go build -o inventory_app
```

### Usage

Any place on you infrastacture. In static file `admin.html` you can add mock data with another honeypot host. In `/debug` route add some info for another honeypot in you infra.
