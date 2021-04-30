# basiccalc
Basic Prefix and Infix expression calculator

## Prefix Calculator
See ```calc/prefix.go```

## Infix Calculator
See ```calc/infix.go```

## Test
Various test cases have been created for both the prefix and infix calculators.
```
cd calc && go test .
```

## Cavet
The Infix expression needs to be fully parenthesised. For instance

```
( ( 8 + ( 1 + 9 ) ) / 2 ) * ( 3 + 8 )
```
will give the correct result of 99.

However,
```
 ( 8 + ( 1 + 9 ) ) / 2 * ( 3 + 8 )
```
despite semantically being identical, gives an incorrect result of 9. 

The reason is because the current solution uses the openning parenthesis to indicator of the expression's depth level. As the exercise requirement states "full-parenthesized", the limitation seems acceptable. Otherwise, the solution can be patched by checking if the expression stream has reached the end, and by continuing if not.

## Web Server
Source code: ```web/server.go```

To run
```
cd web && go run .
```

To deploy on Google App Engine:
```
gcloud app deploy
```

A live version has been deployed here
https://basiccalc.nw.r.appspot.com

To test
```
curl -H 'Content-Type: application/json' -X POST -d '{"notion": "infix", "expression": "( ( 11 * ( 1 + 9 ) ) / 2 ) * ( 3 + ( 8 / ( 2 * 2 ) ) )"}' https://basiccalc.nw.r.appspot.com/calculator
```

Or
```
curl -H 'Content-Type: application/json' -X POST -d '{"notion": "prefix", "expression": "- / 18 + 1 1 - 19 9 "}' https://basiccalc.nw.r.appspot.com/calculator
```

To test divide by zero
```
curl -H 'Content-Type: application/json' -X POST -d '{"notion": "prefix", "expression": "/ 1 0"}' https://basiccalc.nw.r.appspot.com/calculator
```