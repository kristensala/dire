package lexer

type Keyword string

const (
    KEYWORD_VAR Keyword = "var"
    KEYWORD_FOR Keyword = "for"
    KEYWORD_FUNCTION Keyword = "fn"
    KEYWORD_IF Keyword = "if"
    KEYWORD_ELSE Keyword = "else"
    KEYWORD_RETURN Keyword = "return"
    KEYWORD_TYPE Keyword = "type"
    KEYWORD_IMPORT Keyword = "import"
    KEYWORD_MUTABLE Keyword = "mut"
    KEYWORD_PUBLIC Keyword = "public"
)
