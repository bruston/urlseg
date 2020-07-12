Expands a list of urls into their individual uri elements. Takes input from stdin.

So: https://www.example.com/foo/bar/baz.js

Becomes:

```
https://www.example.com/foo/bar/baz.js
https://www.example.com/
https://www.example.com/foo
https://www.example.com/foo/bar
```

Outputs to stdin.
