# uq

Pretty prints URLs if piped.

Can `add` and `remove` from `query`.

## Example

Pretty print:

```bash
$ echo 'https://www.bing.com/search?q=hello' | uq

- scheme: https
  hostname: www.bing.com
  path: /search
  query:
    q: hello
```

Add to query:

```bash
$ echo 'https://www.bing.com/search?q=hello' | uq query add q world

https://www.bing.com/search?q=hello&q=world
```

Combination of both:

```bash
$ echo 'https://www.bing.com/search?q=hello' | uq query add q world | uq
- scheme: https
  hostname: www.bing.com
  path: /search
  query:
    q: hello, world
```
