# Kubefs

This project is intended to manage files on a Kubernetes node. It is cloud native and built to run in Kubernetes but can run on  a linux VM outside of Kubernetes. 

- [Original Use case](#original-use-case)
- [CLI](#cli)
- [Functionalities](#functionalities)
- [Usage](#usage)
- [Contribute](CONTRIBUTING.md#label-commits)

## Original Use Case

Depending on the IP Camera and the sensitivity level, the motion detection alarm can generate a lot irrelevent captures, making it difficult to identify the false positives from the true positives.

This project is built to manage images generated from motion detection alarm of an IP-Camera(s).

## CLI

- Serve - serves the webserver
- port - port from which to serve
- cert - server certificate
- key - server private key
- dir - image directory

**_example_**:

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout cert.key -out cert.crt -subj "/CN=kubefs.earth.milkyway/O=kubefs.earth.milkyway"


go run cmd/kubefs/*.go serve -p 8443 --cert cert.crt --key cert.key --dir /tmp
```

**_output_**:
```bash
Starting server at 8443 watching directory /tmp.
```

**_check server cert_**:  
while running the webserver in another terminal...
```bash
openssl s_client -showcerts -connect localhost:8443 | egrep "subject=CN"
```

**_output_**:
```bash
depth=0 CN = kubefs.earth.milkyway, O = kubefs.earth.milkyway
verify error:num=18:self-signed certificate
verify return:1
depth=0 CN = kubefs.earth.milkyway, O = kubefs.earth.milkyway
verify return:1
subject=CN = kubefs.earth.milkyway, O = kubefs.earth.milkyway
```


## Functionalities

* Deletion logic
    * Delete images more than X days old
    * Delete images captured between Date/Time and Date/Time
    * Retain max size of X in directory 
      * delete oldest first
    * Retain X number of files
      * delete oldest first
* View Logic
  * View images by day
  * View images by day and time
  * View files by camera
  * View files by camera, day, and time
  * View files from given camera on given date
* Backup Logic
  * Backup files for X days, then delete

## Components

* Frontend (React) 
* Backend (Go 1.19)
  * CLI