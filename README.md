`nfsn` DNS module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It
can be used to manage DNS records with nearlyfreespeech.net.

## Caddy nfsn

```
dns.providers.nfsn
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy
JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "nfsn",
				"login": "YOUR_NFSN_LOGIN",
				"api_key": "YOUR_NFSN_API_KEY"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns nfsn <YOUR_NFSN_LOGIN> <YOUR_NFSN_API_KEY> ...
}
```

```
# one site
tls {
	dns nfsn <YOUR_NFSN_LOGIN> <YOUR_NFSN_API_KEY> ...
}
```

## Configuration Parameters

`YOUR_NFSN_LOGIN` is the login name used to access the NFSN member control page.

`YOUR_NFSN_API_KEY` can be obtained by logging into NFSN, navigating to the 'profile' tab, and
selecting 'Manage API Key'. More details on obtaining and managing API keys are available in the
[NFSN FAQs](https://members.nearlyfreespeech.net/faq).
