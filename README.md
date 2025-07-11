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

## Environment

1. APP_PORT - http port. Default `8080`.
2. ADMIN_USER.
3. ADMIN_PASSWORD
4. APP_LOG_FILE. Default - `/var/log/honeypot/inventory_pot.json`.

Usage `.env` file or just inject env in container or process.