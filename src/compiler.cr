require "./visitor"

class Rinha::Compiler

  struct AssemblyVisitor < Visitor
    getter code
    getter stack_size
    getter vars


    struct Variable
      property stack_lock

      def initialize(@stack_lock : Int32)
      end
    end

    def initialize
      @code = String::Builder.new
      @stack_size = 0
      @vars = {} of String => Variable
    end


    def push(reg : String)
      @code << "    push #{reg}\n"
      @stack_size += 1
    end

    def pop(reg : String)
      @code << "    pop #{reg}\n"
      @stack_size -= 1
    end

    def visit(node : Rinha::Parser::Program)
      @code << "global _start\n_start:\n"

      node.expression.accept self
    end

    def visit(node : Rinha::Parser::Call)
      puts "call"
    end

    def visit(node : Rinha::Parser::Function)
      puts "function"
    end

    def visit(node : Rinha::Parser::If)
      puts "if"
    end

    def visit(node : Rinha::Parser::Print)
      node.value.accept self

      @code << "    mov rax, 60\n"
      self.pop("rdi")
      @code << "    syscall\n"
    end

    def visit(node : Rinha::Parser::Var)
      var = @vars[node.text]
      unless var
        raise "undeclared variable: #{node.text}"
      end
      self.push("QWORD [rsp + #{(@stack_size - var.stack_lock - 1) * 8}]\n")
    end

    def visit(node : Rinha::Parser::Let)
      if @vars.has_value?(node.name.text)
        raise "not ok"
      end

      @vars[node.name.text] = Variable.new @stack_size
      @code << "    mov rax, #{node.value.accept self}\n"
      self.push("rax")

      node.next.accept self
    end

    def visit(node : Rinha::Parser::Binary)
      puts "binary"
    end

    def visit(node : Rinha::Parser::Int)
      node.value
    end

    def visit(node : Rinha::Parser::Str)
      puts "str"
    end

    def visit(node : Rinha::Parser::Boolean)
      puts "boolean"
    end
  end

  def generate(ast : Rinha::Parser::Program)
    visitor = AssemblyVisitor.new

    ast.accept visitor

    visitor.code
  end
end
