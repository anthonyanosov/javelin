# Semantics v0

Variable Declaration:

decl_stmt := IDENT ":=" INT

Rules:
- INT is evaluated as a 64-bit integer literal
- IDENT introduces a new variable in the current environment
- The value of the variable is the evaluated INT
- If IDENT already exists, it is overwritten