require "json"
require "./parser"

module Rinha::Interpreter
  extend self
  VERSION = "0.1.0"

  @@stack = Stack.new

  struct Closure
    property body, env, parameters

    def initialize(@body : Rinha::Parser::Term, @parameters : Array(Rinha::Parser::Parameter), @env : Hash(String, Bool | Int32 | Float64 | Closure | String | Nil))
    end

    def to_s(io)
      io << "<#closure>"
    end
  end

  struct Stack
    property frames = [] of Closure

    def push(closure : Closure)
      @frames.push(closure)
    end

    def pop
      @frames.pop
    end

    def empty
      @frames.empty?
    end
  end

  def evaluate(expression : Rinha::Parser::Term, env = {} of String => Bool | Int32 | Closure | String | Float64 | Nil)
    case expression
    in Rinha::Parser::Str, Rinha::Parser::Int, Rinha::Parser::Boolean
      expression.value
    in Rinha::Parser::If
      evaluate(expression.condition, env) ? evaluate expression.then, env : evaluate expression.otherwise, env
    in Rinha::Parser::Var
      value = env[expression.text]
      if value
        value
      else
        raise "Not defined variable"
      end
    in Rinha::Parser::Let
      env[expression.name.text] = evaluate expression.value, env

      evaluate expression.next, env
    in Rinha::Parser::Binary
      left = evaluate expression.lhs, env
      right = evaluate expression.rhs, env

      case expression.op
      when "Add"
        case {left, right}
        when {Int32, Int32}
          left + right
        when {Int32, String}, {String, Int32}
          "#{left}#{right}"
        else
          raise "Types mismatch for #{left} + #{right}.\nOnly accept Number | String + Number | String"
        end
      when "Sub"
        case {left, right}
        when {Int32, Int32}
          left - right
        else
          raise "Types mismatch for #{left} - #{right}.\nOnly accept number - number"
        end
      when "Mul"
        case {left, right}
        when {Int32, Int32}
          left * right
        else
          raise "Types mismatch for #{left} * #{right}.\nOnly accept number * number"
        end
      when "Div"
        case {left, right}
        when {Int32, Int32}
          left / right
        else
          raise "Types mismatch for #{left} / #{right}.\nOnly accept number / number"
        end
      when "Rem"
        case {left, right}
        when {Int32, Int32}
          left % right
        else
          raise "Types mismatch for #{left} % #{right}.\nOnly accept number + number"
        end
      when "Eq"
        left == right
      when "Neq"
        left != right
      when "Lt"
        case {left, right}
        when {Int32, Int32}
          left < right
        else
          raise "Types mismatch for #{left} < #{right}.\nOnly accept number < number"
        end
      when "Gt"
        case {left, right}
        when {Int32, Int32}
          left > right
        else
          raise "Types mismatch for #{left} > #{right}.\nOnly accept number > number"
        end
      when "Lte"
        case {left, right}
        when {Int32, Int32}
          left <= right
        else
          raise "Types mismatch for #{left} <= #{right}.\nOnly accept number <= number"
        end
      when "Gte"
        case {left, right}
        when {Int32, Int32}
          left >= right
        else
          raise "Types mismatch for #{left} >= #{right}.\nOnly accept number >= number"
        end
      when "And"
        left && right
      when "Or"
        left || right
      else
        raise "Unknown Operation for #{expression.op}"
      end
    in Rinha::Parser::Function
      Closure.new(expression.value, expression.parameters, env)
    in Rinha::Parser::Call
      func = evaluate expression.callee, env

      unless func.is_a?(Closure)
        raise "Not a function"
      end

      if expression.arguments.size != func.parameters.size
        raise "Arguments must matched the size of functions parameters"
      end

      local = env.dup
      func.parameters.zip(expression.arguments) { |parameter, value|
        local[parameter.text] = evaluate(value, env)
      }

      @@stack.push(func)

      until @@stack.empty
        stack_fn = @@stack.pop
        break evaluate stack_fn.body, local
      end
    in Rinha::Parser::Print
      result = evaluate expression.value, env

      puts result
      result
    in Rinha::Parser::Term
      raise "Unexpected error"
    end
  end
end
