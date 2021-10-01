# Cookie strip
Cookie strip is a middleware plugin for [Traefik](https://github.com/traefik/traefik) which strips cookies by name from a request

### Configuration

### Static

```yaml
pilot:
  token: xxxxx

experimental:
  plugins:
    stripcookies:
      moduleName: github.com/nilskohrs/traefik-stripcookies
      version: v0.0.2
```

### Dynamic

```yaml
http:
  middlewares:
    strip-foo:
      stripcookies:
        - cookieName
        - otherCookieName
```
