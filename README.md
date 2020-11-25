```
# server certificate
openssl req -newkey rsa:4096 -nodes -sha256 -x509 -days 30 -out cert.crt -keyout cert.key -subj "/C=PL/ST=DOL/L=Wroclaw/O=server/OU=GS PTF/CN=localhost"

# intermediate certificate
openssl req -newkey rsa:4096 -nodes -sha256 -x509 -days 30 -out inter.crt -keyout inter.key -subj "/C=PL/ST=DOL/L=Wroclaw/O=intermediate/OU=GS PTF/CN=localhost"

# root certificate
openssl req -newkey rsa:4096 -nodes -sha256 -x509 -days 90 -out ca.crt -keyout ca.key -subj "/C=PL/ST=DOL/L=Wroclaw/O=rootCA/OU=GS PTF/CN=localhost"

# ordering certificate when using intermediates:
- server certificate
- intermediate certificate
- root certificate

$ cat cert.crt > server.pem && \
cat inter.crt >> server.pem && \
cat ca.crt >> server.pem


# check ssl certificate chain in the Firefox:
https://localhost:5000/
```