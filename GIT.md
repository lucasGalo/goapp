### List tags
```bash
$ git tag
```

*You can also search for tags with a specific pattern.*

```bash
$ git tag -l "v1.8.5*"
```


### Create Tags
```bash
$ git tag -a v1.4 -m "my version 1.4"
```

*You can also create a tag after you have already committed.*
```bash
$ git tag -a v1.2 9fceb02
```

### Push tags

*specific tag*
```bash
$ git push origin v1.5
```

*upload all*
```bash
$ git push origin --tags
```