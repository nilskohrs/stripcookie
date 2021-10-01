# Stripcookie
Stripcookie is a middleware plugin for [Traefik](https://github.com/traefik/traefik) which strips cookies by name from a request

### Configuration

### Static

```yaml
pilot:
  token: xxxxx

experimental:
  plugins:
    stripcookie:
      moduleName: github.com/nilskohrs/stripcookie
      version: v0.0.4
```

### Dynamic

```yaml
http:
  middlewares:
    strip-foo:
      stripcookie:
        cookies:
          - cookieName
          - otherCookieName
```
