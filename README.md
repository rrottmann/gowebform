# gowebform

This is a simple Go web server that serves a HTML web form
and stores the values that it receives via HTTP POST as JSON.

# Why?
The idea behind this is to spawn an access point via hostapd and run
this server on reboot when there is no internet connection available.

After the user connects to the temporary hostapd Wifi and configures
the parameters, the webserver stores the values as JSON data.

The JSON data can then be read by Ansible to reconfigure settings.

As an example, the HTTP form asks for interface, SSID and PSK that
can be used to reconfigure a Raspberry Pi to another Wifi.

This could be expanded with static IP configuration options.

So this webserver alone just handles the HTTP Form and plays a small but
crucial part in the reconfiguration process of a SBC.

## Usage

```
$ ./gowebform  -h
Usage of gowebform:
  -cert string
        Specify TLS certificate file. (default "server.crt")
  -html string
        Specify HTML form to serve. (default "gowebform.html")
  -json string
        Specify JSON file to store received POST form data. (default "gowebform.json")
  -key string
        Specify TLS key file. (default "server.key")
  -port string
        Specify port for HTTPS enabled webserver. (default ":8000")
```

## Generate a new TLS certificate and key

You may run `tlskeygen.sh` in order to create a new self-signed TLS certificate and key.