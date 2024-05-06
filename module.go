package nfsn

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsnfsn "github.com/libdns/nfsn"
)

// Provider lets Caddy read and manipulate DNS records hosted by nearlyfreespeech.net.
type Provider struct{ *libdnsnfsn.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.nfsn",
		New: func() caddy.Module { return &Provider{new(libdnsnfsn.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	replacer := caddy.NewReplacer()
	p.Provider.Login = replacer.ReplaceAll(p.Provider.Login, "")
	p.Provider.APIKey = replacer.ReplaceAll(p.Provider.APIKey, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// nfsn [<login>] [<api_key>] {
//     login <login>
//     api_key <api_key>
// }
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.Login = d.Val()
		}
		if d.NextArg() {
			p.Provider.APIKey = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if p.Provider.APIKey != "" {
					return d.Err("API key already set")
				}
				if d.NextArg() {
					p.Provider.APIKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "login":
				if p.Provider.Login != "" {
					return d.Err("Login already set")
				}
				if d.NextArg() {
					p.Provider.Login = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIKey == "" {
		return d.Err("missing API key")
	}
	if p.Provider.Login == "" {
		return d.Err("missing Login")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
