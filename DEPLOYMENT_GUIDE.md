# Deployment & Production Checklist

## Pre-Deployment Verification

### Backend Checks
- [ ] Go version 1.19+ installed: `go version`
- [ ] All dependencies installed: `go mod download`
- [ ] Code builds without errors: `go build -o server.exe ./cmd/server/main.go`
- [ ] Routes registered correctly (check `server_out.log`)
- [ ] Database migrations run successfully
- [ ] Admin user created on startup
- [ ] CORS configured for frontend domain
- [ ] Error logs are clean (check `server_err.log`)

### Frontend Checks
- [ ] Node.js 16+ installed: `node -v`
- [ ] npm dependencies installed: `npm install`
- [ ] No build errors: `npm run build`
- [ ] Production build generates: `frontend/dist/`
- [ ] `.gitignore` excludes `node_modules` and `dist`
- [ ] Environment variables documented
- [ ] API base URL configured correctly
- [ ] No console errors or warnings

### Configuration
- [ ] `.env` file created with all required variables:
  ```
  GIN_MODE=release
  SERVER_PORT=8080
  JWT_SECRET=<generated-secure-secret>
  JWT_EXPIRY=72h
  DATABASE_PATH=./school.db
  LOG_LEVEL=info
  ```
- [ ] JWT_SECRET is strong (at least 32 characters, random)
- [ ] Server port configured (default 8080)
- [ ] Database path is writable directory
- [ ] Log level set appropriately (release mode)

### Database
- [ ] SQLite database file is writable
- [ ] Backup strategy in place
- [ ] Schema version tracked
- [ ] Admin user credentials documented
- [ ] Test data cleaned up (optional, for production)

### Security
- [ ] HTTPS/TLS enabled (in production)
- [ ] JWT secret stored securely (not in code)
- [ ] Password hashing verified
- [ ] SQL injection prevention (GORM handles this)
- [ ] CORS whitelist configured
- [ ] Rate limiting considered (optional but recommended)
- [ ] Sensitive endpoints protected by role middleware

### Testing
- [ ] Admin user login works
- [ ] Teacher user login works
- [ ] Student user login works
- [ ] Role-based access working (admin can see admin panel, etc.)
- [ ] Course creation by teacher works
- [ ] Student enrollment works
- [ ] Grade recording works
- [ ] Attendance recording works
- [ ] **NEW**: Enrollment approval works
- [ ] **NEW**: User search/filter works
- [ ] **NEW**: Enrollment search/filter works

---

## Deployment Steps

### Option 1: Traditional Server Deployment

#### 1. Prepare Server
```bash
# SSH into production server
ssh user@server.com

# Install Go (if not already installed)
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Install Node.js (if using Vite)
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs
```

#### 2. Clone Repository
```bash
cd /opt
git clone https://github.com/yourorg/school-management-system.git
cd school-management-system
```

#### 3. Build Backend
```bash
go build -o server ./cmd/server/main.go
chmod +x server
```

#### 4. Build Frontend
```bash
cd frontend
npm install --production
npm run build
cd ..
```

#### 5. Create systemd Service (Linux)
```bash
sudo tee /etc/systemd/system/sms.service > /dev/null <<EOF
[Unit]
Description=School Management System
After=network.target

[Service]
User=sms
WorkingDirectory=/opt/school-management-system
ExecStart=/opt/school-management-system/server
Restart=always
RestartSec=10
Environment="GIN_MODE=release"
Environment="JWT_SECRET=your-secret-key"
Environment="DATABASE_PATH=/opt/school-management-system/school.db"

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable sms
sudo systemctl start sms
```

#### 6. Setup Nginx Reverse Proxy
```nginx
server {
    listen 80;
    server_name example.com;

    # Redirect to HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name example.com;

    # SSL certificates (use Let's Encrypt)
    ssl_certificate /etc/letsencrypt/live/example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/example.com/privkey.pem;

    # Proxy to backend
    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_redirect off;
    }
}
```

#### 7. Verify Deployment
```bash
# Check service status
sudo systemctl status sms

# Check logs
sudo journalctl -u sms -f

# Test API
curl https://example.com/health
```

---

### Option 2: Docker Deployment

#### 1. Create Dockerfile
```dockerfile
# Multi-stage build
FROM golang:1.21-alpine AS backend-builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server/main.go

FROM node:18-alpine AS frontend-builder
WORKDIR /frontend
COPY frontend/package*.json ./
RUN npm ci --omit=dev
COPY frontend .
RUN npm run build

FROM alpine:latest
RUN apk add --no-cache ca-certificates sqlite-libs
WORKDIR /app
COPY --from=backend-builder /app/server .
COPY --from=frontend-builder /frontend/dist ./frontend/dist
COPY --from=frontend-builder /frontend/index.html ./frontend/

EXPOSE 8080
ENV GIN_MODE=release
CMD ["./server"]
```

#### 2. Create docker-compose.yml
```yaml
version: '3.9'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      GIN_MODE: release
      JWT_SECRET: ${JWT_SECRET}
      DATABASE_PATH: /data/school.db
    volumes:
      - data:/data
    restart: unless-stopped

volumes:
  data:
```

#### 3. Deploy with Docker
```bash
docker-compose build
docker-compose up -d

# Check logs
docker-compose logs -f app

# Stop
docker-compose down
```

---

### Option 3: Cloud Platform (Heroku/Railway/Render)

#### 1. Heroku Buildpacks
```bash
heroku create school-management-system
heroku buildpacks:add https://github.com/heroku/heroku-buildpack-go.git
heroku buildpacks:add https://github.com/heroku/heroku-buildpack-nodejs.git

# Set environment variables
heroku config:set JWT_SECRET=your-secret-key
heroku config:set GIN_MODE=release

# Deploy
git push heroku main
```

#### 2. Create Procfile
```
web: server -port=$PORT
```

---

## Post-Deployment

### Health Checks
- [ ] Service is running: `curl https://example.com/health`
- [ ] Frontend loads: `curl https://example.com/`
- [ ] Login works: `curl -X POST https://example.com/api/auth/login`
- [ ] Database is accessible
- [ ] Logs are clean

### Monitoring
- [ ] Setup log aggregation (ELK, Datadog, etc.)
- [ ] Monitor CPU/Memory usage
- [ ] Monitor database size
- [ ] Setup uptime monitoring
- [ ] Configure alerts for errors
- [ ] Track API response times

### Backup Strategy
- [ ] Daily database backups
- [ ] Store backups off-site
- [ ] Test backup restoration
- [ ] Document recovery procedure
- [ ] Version control for code

### Performance Optimization
- [ ] Enable gzip compression
- [ ] Setup CDN for static files
- [ ] Implement caching headers
- [ ] Monitor query performance
- [ ] Index database tables appropriately

---

## Rollback Plan

If deployment issues occur:

### Immediate Rollback
```bash
# Stop current service
systemctl stop sms

# Revert to previous version
git checkout previous-release
go build -o server ./cmd/server/main.go

# Restart
systemctl start sms
```

### Database Rollback
```bash
# Restore from backup
cp backup.db school.db

# Verify data integrity
sqlite3 school.db "SELECT COUNT(*) FROM users;"
```

---

## Maintenance Schedule

### Daily
- [ ] Monitor error logs
- [ ] Check system health metrics
- [ ] Verify backups completed

### Weekly
- [ ] Review performance metrics
- [ ] Check for security updates
- [ ] Test backup restoration

### Monthly
- [ ] Update dependencies
- [ ] Review and optimize slow queries
- [ ] Audit user accounts
- [ ] Review access logs

### Quarterly
- [ ] Security audit
- [ ] Database optimization
- [ ] Capacity planning
- [ ] Update documentation

---

## Common Issues & Solutions

### Issue: High Memory Usage
```bash
# Check process memory
ps aux | grep server

# Restart service to clear memory
systemctl restart sms

# Consider pagination limits if needed
```

### Issue: Slow Database Queries
```bash
# Enable query logging in Go
# Add timing logs to slow handlers
# Create database indexes on frequently used fields
# Consider caching popular data
```

### Issue: SSL Certificate Expiration
```bash
# For Let's Encrypt
certbot renew

# Verify certificate
openssl x509 -in /etc/letsencrypt/live/example.com/fullchain.pem -text -noout
```

### Issue: Disk Space Full
```bash
# Check disk usage
df -h

# Clean logs
find /var/log -type f -name "*.log" -delete

# Archive database if needed
```

---

## Scaling Considerations

### If Traffic Increases
1. **Horizontal Scaling**: Deploy multiple backend instances
2. **Load Balancing**: Use Nginx/HAProxy to distribute traffic
3. **Database**: Migrate from SQLite to PostgreSQL
4. **Caching**: Implement Redis for session/data caching
5. **CDN**: Serve static files from CDN

### If Data Grows Large
1. **Database Indexing**: Add indexes on frequently queried fields
2. **Archiving**: Archive old data to separate storage
3. **Pagination**: Implement pagination for large result sets
4. **Query Optimization**: Profile slow queries and optimize
5. **Sharding**: Partition data across multiple databases (advanced)

---

## Success Criteria

✅ Deployment is successful when:
- All services running without errors
- Health checks passing
- All API endpoints responding
- Frontend loads and is usable
- Login/authentication working
- Database accessible
- Logs are clean
- Performance metrics acceptable
- Backups verified
- Team trained on operations

---

**Last Updated**: December 2024
**Status**: Ready for Deployment
**Version**: 1.0.0
