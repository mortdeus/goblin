package token

const{
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	literal_end
	
	operator_beg

	REF		// $
	HASH		// #
	BSLASH 		// \
	SEMI		// ;
	CARET		// ^
	AND		// &
	PIPE		// |
	ASSIGN		// =
	QUOTE		// '
	BQUOTE		// `

	LBRACE		// {
	RBRACE		// }
	LPAREN		// (
	RPAREN		// )
	LANGLE		// <
	RANGLE		// >


	SHEBANG		// #!
	REF_HASH	// $#
	REF_QUOTE	// $'
	
	operator_end

}