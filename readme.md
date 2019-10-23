
#### Running locally
```bash
docker-compose up
```

#### Testing
You can import postman collection from root folder with two requests - for basic ruleset and custom ruleset.
Change path param to apply first or second ruleset


#### The theory
Task fits well into Mealy finite state machine, as has input alphabet (custom_computation_index and A, B, C), output alphabet (the formulas) 
and set of internal states [M, P, T].

FSM approach isn't chosen, because:
1) It's long for develop in terms of test task (with high level of uncertainty ;).
2) It's hard to change and maintain, and such task need flexibility

Instead, approach with decorators around base rules are used. Nevertheless FSM naming is present (CommandToState and StateToExpression), 
because it's just easier to operate on some domain field.

A,B,C and are named "command", because this signals are used for state calculation. 
Commands are represented as binary string because it is easier to read and understand 
(still all system could be implemented using bitmasks for memory optimization, but we all know about evil of premature optimization ;) )

M,P,T are named "state" and ar represented as strings. 

Different calculation logic are named as "rules" and based on route name. Rules incapsulate input->state and state->formula mapping.

Formulas are named "expressions" as in the doc


#### Known issues
No tests as my time is limited and test task is generally small. 
Better error handling/wrapping (with error mechanism from newest golang release) 
Different HTTP codes for different types of errors using gin middleware (e.g. not 500 for every error)
Better validation on HTTP layer and int overflow checks in formulas
Bump versions of alpine and golang in dockerfile
Goimports
Git tree with more than one commit ;)