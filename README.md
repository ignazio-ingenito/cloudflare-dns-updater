# Cloudflare DNS Updater

Cloudflare DNS updater is a tool designed to retrieve your public IP address and update your Cloudflare DNS records automatically. It also provides APIs to fetch update logs in various formats, including JSON, ChartJS, and HTML. Additionally, the tool includes an HTML dashboard for quick and easy inspection of DNS updates.

## Getting Started

You can use docker to run the image issuing the following command.

```
docker run --rm -it volatore74/cloudflare-dns-updater:latest
```

Or you can build it manually with

```
docker build -t cloudflare-dns-updater:latest .
docker run --rm -it cloudflare-dns-updater:latest
```

Or use the docker-compose file issuing the following command

```
docker compose up
```

## Authors

- **Ignazio Ingenito** - _Initial work_ - [ignazio-ingenito](https://github.com/ignazio-ingenito)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
