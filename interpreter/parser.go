package interpreter

import "fmt"

// Parser represents parser
type Parser struct {
	lexer   *Lexer
	current *Token
}

// NewParser returns Parser
func NewParser(lexer *Lexer) (*Parser, error) {
	current, err := lexer.Next()

	if err != nil {
		return nil, err
	}

	return &Parser{
		lexer:   lexer,
		current: current,
	}, nil
}

// Parse parses, sun is shining, water is wet
func (p *Parser) Parse() (AST, error) {
	node, err := p.statement()

	if err != nil {
		return nil, err
	}

	if p.current.Type != EOF {
		return nil, fmt.Errorf("%w on %v", ErrUnexpectedEOF, p.current)
	}

	return node, nil
}

func (p *Parser) eat(tokenType int) error {
	if p.current.Type != tokenType {
		return fmt.Errorf("%w on %v", ErrSyntaxError, p.current)
	}

	next, err := p.lexer.Next()
	p.current = next

	return err
}

/*
statement:  SYMBOL ASSIGN statement | expr
expr:       term ((PLUS | MINUS) term)*
term:       factor ((MUL | DIV) factor)*
factor:     (NUMBER | MONEY | SYMBOL) | LPAREN expr RPAREN
*/
func (p *Parser) statement() (AST, error) {
	node, err := p.expr()

	if err != nil {
		return nil, err
	}

	if symbol, ok := node.(*Symbol); ok && p.current.Type == ASSIGN {
		p.eat(ASSIGN)

		statement, err := p.statement()

		if err != nil {
			return nil, err
		}

		node = &Assign{
			Symbol: symbol,
			Value:  statement,
		}
	}

	return node, nil
}

func (p *Parser) expr() (AST, error) {
	node, err := p.term()

	if err != nil {
		return nil, err
	}

	for p.current.Type == PLUS || p.current.Type == MINUS {
		token := p.current

		err := p.eat(token.Type)

		if err != nil {
			return nil, err
		}

		right, err := p.term()

		if err != nil {
			return nil, err
		}

		node = &BinaryOperator{
			Left:     node,
			Operator: token,
			Right:    right,
		}
	}

	return node, nil
}

func (p *Parser) term() (AST, error) {
	node, err := p.factor()

	if err != nil {
		return nil, err
	}

	for p.current.Type == MUL || p.current.Type == DIV {
		token := p.current

		err := p.eat(token.Type)

		if err != nil {
			return nil, err
		}

		right, err := p.factor()

		if err != nil {
			return nil, err
		}

		node = &BinaryOperator{
			Left:     node,
			Operator: token,
			Right:    right,
		}
	}

	return node, nil
}

func (p *Parser) factor() (AST, error) {
	token := p.current

	if token.Type == SYMBOL {
		p.eat(SYMBOL)

		symbol := new(Symbol)
		symbol.Token = token

		return symbol, nil
	}

	if token.Type == NUMBER {
		p.eat(NUMBER)

		number := new(Number)
		number.Token = token

		return number, nil
	}

	if token.Type == MONEY {
		p.eat(MONEY)

		money := new(Money)
		money.Token = token

		return money, nil
	}

	if token.Type == LPAREN {
		p.eat(LPAREN)

		node, err := p.expr()

		if err != nil {
			return nil, err
		}

		err = p.eat(RPAREN)

		if err != nil {
			return nil, err
		}

		return node, nil
	}

	return nil, nil
}
