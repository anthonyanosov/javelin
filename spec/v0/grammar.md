# Language Grammar v0

program      := statement*
statement    := var_stmt
var_stmt     := "var" IDENT ":=" INT
IDENT        := letter (letter | digit)*
INT          := digit+