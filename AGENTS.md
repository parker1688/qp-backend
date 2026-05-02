# Backend Agent Notes

## Scope
This subtree is Go service code for the main backend. Favor targeted fixes inside the touched request path before considering shared refactors.

## Build And Test
Prefer narrow validation first, such as targeted go test for the touched package, before wider repository checks.

## Conventions
Keep controller, service, model, and response contract changes aligned.

Preserve existing error-code behavior unless the task explicitly changes the contract.

Be careful with transaction boundaries, partial writes, nil handling, and config-dependent logic.
