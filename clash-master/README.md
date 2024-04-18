<h1 align="center">
  <img src="https://github.com/zhaofenghao/clash_go/raw/master/docs/logo.png" alt="Clash" width="200">
  <br>Clash<br>
</h1>

<h4 align="center">A rule-based tunnel in Go.</h4>

<p align="center">
  <a href="https://github.com/zhaofenghao/clash_go/actions">
    <img src="https://img.shields.io/github/actions/workflow/status/zhaofenghao/clash/release.yml?branch=master&style=flat-square" alt="Github Actions">
  </a>
  <a href="https://goreportcard.com/report/github.com/zhaofenghao/clash_go">
    <img src="https://goreportcard.com/badge/github.com/zhaofenghao/clash_go?style=flat-square">
  </a>
  <img src="https://img.shields.io/github/go-mod/go-version/zhaofenghao/clash?style=flat-square">
  <a href="https://github.com/zhaofenghao/clash_go/releases">
    <img src="https://img.shields.io/github/release/zhaofenghao/clash/all.svg?style=flat-square">
  </a>
  <a href="https://github.com/zhaofenghao/clash_go/releases/tag/premium">
    <img src="https://img.shields.io/badge/release-Premium-00b4f0?style=flat-square">
  </a>
</p>

## Features

This is a general overview of the features that comes with Clash.  

- Inbound: HTTP, HTTPS, SOCKS5 server, TUN device
- Outbound: Shadowsocks(R), VMess, Trojan, Snell, SOCKS5, HTTP(S), Wireguard
- Rule-based Routing: dynamic scripting, domain, IP addresses, process name and more
- Fake-IP DNS: minimises impact on DNS pollution and improves network performance
- Transparent Proxy: Redirect TCP and TProxy TCP/UDP with automatic route table/rule management
- Proxy Groups: automatic fallback, load balancing or latency testing
- Remote Providers: load remote proxy lists dynamically
- RESTful API: update configuration in-place via a comprehensive API

*Some of the features may only be available in the [Premium core](https://zhaofenghao.github.io/clash/premium/introduction.html).*

## Documentation

You can find the latest documentation at [https://zhaofenghao.github.io/clash/](https://zhaofenghao.github.io/clash/).

## Credits

- [riobard/go-shadowsocks2](https://github.com/riobard/go-shadowsocks2)
- [v2ray/v2ray-core](https://github.com/v2ray/v2ray-core)
- [WireGuard/wireguard-go](https://github.com/WireGuard/wireguard-go)

## License

This software is released under the GPL-3.0 license.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fzhaofenghao%2Fclash.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fzhaofenghao%2Fclash?ref=badge_large)
