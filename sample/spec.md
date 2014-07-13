
# Broccoli Language Specification

identifier: `[\w]+`
number: `[\d+.?\d+]`


## Expressions

var reference: `identifier`

number literal: `number`

arithmetic: `expr [+-/*] expr`

function call: `identifier(expression(,expression)*)`
