# D5
Lisp like homoiconic language.
D5 is an interpreted languaged that is influenced by Lisp heavily.

## Why

## Example
```json
{
    "type": "module",
    "name": "testlib",
    "sum_fn": {
        "type": "function",
        "doc": "some docsstring in here",
        "body": {
           "type": "expr",
           "expr": "$args.a + $args.b" // access function arguments using $args.
        } 
    }, 
    "inc_fn": {
        "type": "fn_call",
        "name": "sum_fn",
        "args": {
            "a": 1,
            "b": 2
        }
    }
}
```

## Syntax
D5 is build around blocks. Blocks are just simple Go maps when stored in memory. Each Block can be evaluated if the Block type is defined for D5.
Builtin Table types are<br>

### function
### fn_call
### expr
### if 
### loop
### async
