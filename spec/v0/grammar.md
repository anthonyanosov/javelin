# Language Grammar v0

program      := statement*
statement    := decl_stmt
decl_stmt    := IDENT ":=" INT
IDENT        := letter (letter | digit)*
INT          := digit+