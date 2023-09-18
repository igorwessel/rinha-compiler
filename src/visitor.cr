  abstract struct Visitor
    abstract def visit(node : Rinha::Parser::Call)
    abstract def visit(node : Rinha::Parser::Function)
    abstract def visit(node : Rinha::Parser::Print)
    abstract def visit(node : Rinha::Parser::Var)
    abstract def visit(node : Rinha::Parser::If)
    abstract def visit(node : Rinha::Parser::Let)
    abstract def visit(node : Rinha::Parser::Binary)
    abstract def visit(node : Rinha::Parser::Int)
    abstract def visit(node : Rinha::Parser::Str)
    abstract def visit(node : Rinha::Parser::Boolean)
    abstract def visit(node : Rinha::Parser::Program)
  end
