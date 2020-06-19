# moth

```
Take the 'oof' out of twofactor

Usage:
  moth [command]

Available Commands:
  add         Add a provider
  delete      Delete a provider
  get         Generate a token
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.moth.yaml)
  -h, --help            help for moth

Use "moth [command] --help" for more information about a command.
```

## Example

```
$ moth add outlook <SECRET>
853611 | outlook

$ moth add okta <SECRET>
333326 | okta

$ moth get all
550386 | okta
858854 | outlook

$ moth get out -c
858854 | outlook
<code copied to clipboard>
```