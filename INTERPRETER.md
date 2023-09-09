# O que é interpretador?

Interpretador nada mais do que é traduzir um código para outro sem que esse código se transforme em código de maquina
Basicamente um interpretador não precisa de um código binario/executável que a maquina possa ser executar.

## Etapas do interpretador

1. Ler uma string e quebrar ela em tokens, mais conhecido como `Lexical Analysis` (Análise Léxica)
Também conhecida como lexer/scanner/tokenizer, mas basicamente vamos ler uma string e transform em uma stream de tokens.
Quero enfatizar `stream` pois é obrigatório que a leitura dos tokens seja sequencial, dado a entrada por exemplo:
"1 + 1" podemos quebrar em uma stream de tokens de -> 
"Token(Inteiro, 1) -> Token(Whitespace, ' ') -> Token(Plus, "+") -> Token(Whitespace, ' ') -> Token(Inteiro, 1)"

### Lexemes

Lexemes nada mais do que é uma sequências de caracteres que formam um token, por exemplo

| Token | Lexemes |
|-------|---------|
| Integer | 2,4,8,10,15,20,18 |
| Plus | + |
| Minus | - |

### Parsing -  Syntax analysis

Parsear é ato de identificar que tipo de estrutura temos em nosso tokens, e quando identificamos essa estrutura
conseguimos identificar o que precisamos realizar, por exemplo:

Token(Integer, 1) - Token(Plus, '+') - Token(Integer, 1)

Nosso parser (Syntax Analyser) consegue identificar que é uma operacao aritmetica de soma, com isso podemos interpretar e retornar o resultado.

# O que é compilador ?

Compilador é o inverso do interpretador, vamos fazer as mesma etapas que o interpretador porém vamos precisar
criar um arquivo que a máquina consiga executar esse código.

## Syntax Diagrama

Diagrama feita para explicar a sintaxe de uma linguagem, e serve também para nos ajudarmos com regras de forma simples.

# Grammars

A anotacao usada para especificar a sintáxe de uma lingaguem de programacao é chamada de `context-free grammars (grammars)`, ou `BNF (Backus-Naur-Form)`
Vamos entender como utilizar uma notacao modificada como a EBNF:
Basicamente um GRAMMAR consiste em sequência de `regras (rules)`, também conhecido como `productions`, como podemos ver no grammar abaixo tmos duas regras:


```EBNF
expr: factor((SUM | MINUS) factor)* <- rule/production 1
factor: integer <- rule/production 2
```

E uma `regra (production)` consiste em um `não terminal (non-terminal)`, que é chamado de `head (ou left-hand side)` de uma `regra (production)`, uma `colon (conhecida também como : 2 pontos)`,
e uma sequências de terminais/não terminal que é conhecido como `body (ou right-hand side)` de uma regra.


```EBNF
[left-hand side]: [right-hand side]
expr            : factor((SUM | MINUS) factor)* 
factor          : integer
```

Os tokens são considerados `terminal` e as variáveis `não terminal`, normalmente `não terminal` consiste em uma sequência de terminal ou/e não terminal.

```EBNF
======== RULE =================
expr: factor((SUM | MINUS) factor)* 
===============================
expr = não terminal 
factor = não terminal
sum = terminal
minus = terminal
factor = não terminal

======== RULE =================
factor: integer 
===============================
factor = não terminal
integer = terminal
```

O não terminal na primeira regra é conhecido como `start symbol`, então no caso abaixo seria `expr`:

```EBNF
======== RULE =================
expr: factor((SUM | MINUS) factor)* 
===============================
expr = start symbol
```

O grammar define uma linguagem explicando que tipo de sentencas a mesma pode formar, por exemplo nesse caso vamos tentar derivar uma expressão artimetica usando grammar:
Você comeca pelo `start symbol` nesse caso expr, e então repetidamente troca um `nào terminal` pela parte `right-hand side (body)` até que você tenha uma sentenca que consiste
apenas em terminal. Essas sentencas forma a linguagem definida pelo grammar.

expr
factor((SUM | MINUS) factor)*
factor SUM factor
integer SUM integer
3       +   3

expr
factor((SUM | MINUS) factor)*
factor SUM factor((SUM | MINUS) factor)*
factor SUM factor MINUS factor
integer SUM integer MIN integer
10      +   10      -   5

expr
factor((SUM | MINUS) factor)*
factor SUM factor((SUM | MINUS) factor)*
factor SUM factor SUM factor((SUM | MINUS) factor)*
factor SUM factor SUM factor MINUS factor((SUM | MINUS) factor)*
factor SUM factor SUM factor MINUS factor SUM factor((SUM | MINUS) factor)*
factor SUM factor SUM factor MINUS factor SUM factor
integer SUM integer MINUS integer SUM integer
10      +   10      -     5       +   5

Seguindo algumas regras conseguimos transformar o grammar criado acima em código:
1. Cada `regra (R/production/rule)` definida no grammar, se torna um metódo com o mesmo nome, e a referência para a regra se tonra o metódo sendo chamado. O corpo desse metódo segue o mesmo fluxo que definimos para regra.
2. As alternativas se torna `if-elif-else`
3. Um grupo opcional se torna um `while` que pode ser iterado 0 ou mais vezes.
4. Cada referência do token se uma chamada para o metódo `eat (T)`, como ele funciona é simplesmente checar se o TIPO do token passado bate com o token atual, se sim pega um novo token do nosso lexer e atribui esse token para o token atual.


# Parser Tree (Concrete Syntax Tree) e AST (Abstract Syntax Tree)

## Parser Tree 

É uma arvore que representa a estrutura de uma linguagem de acordo com uma definicão de grammar.
Basicamente mostra como o "start symbol" de um grammar deriva/transforma em uma string de uma linguagem de programacao. exemplo:

Grammar:
```EBNF
<expr> ::= <digit> "+" <digit>
<digit> ::= [0-9]
```

Parser Tree:
```
        expr
         |
        /|\
    digit+digit
     |      |
     2      2
```

Representa a string: "2+2"

Podemos reparar que:
1. A arvore grava uma sequência de regras que se aplica ao input reconhecido.
2. O nó inicial tem o nome do `start symbol`, que no caso do nosso grammar é `<expr>`
3. Cada nó interior representa um não-terminal, o que significa que essencialmente estamos gravando uma regra.
4. E o nó externo (leaf node/external node) representa o token, que no nosso caso é `<digit>`


## AST

É uma representacao intermediaria (IR) que muitos compiladores/interpretadores utilizam. 
Ela representa uma abstracao sintatica de uma estrutura de linguagem onde cada nó interior e o nó inicial (root node) representa um operator, e 
o filho de cada nó representa o operando desse operador

As diferencas da AST para a Parser Tree:
```
        AST

        "+"
        / \
       2   2
        
    Parser Tree
        expr
         |
        /|\
    digit+digit
     |      |
     2      2
```

Como pode a AST é muito menor do que a Parser Tree, basicamente apenas capturamos a essência da entrada por isso é pequeno

- AST usa como nó inicial operadores, e os nó interiores também utilizam operadores, só que os filhos dos nó interiores usa seus operandos.
- AST não utiliza o nó interior para representar uma regra gramática, como o parser tree faz.
- AST não grava todos os detalhes de uma sintaxe, nenhuma regra/nó/parentheses por exemplo.
- AST são densas comparadas a Parser Tree para mesma linguagem construtora.

