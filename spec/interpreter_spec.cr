require "./spec_helper"

describe Rinha::Interpreter do

  describe "Print" do
    it "String" do
      program = get_program "string"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq("rinha")
    end

    it "Number" do
      program = get_program "number"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(1)
    end

    it "Boolean" do
      program = get_program "bool"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(true)
    end

    it "Closure" do
      program = get_program "print_fn"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq("<#closure>")
    end
  end

  describe "Binary" do
    it "Add" do
      program = get_program "sum"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(2)
    end

    it "Minus" do
      program = get_program "minus"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(0)
    end

    it "Mul" do
      program = get_program "mul"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(4)
    end

    it "Div" do
      program = get_program "div"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(5)
    end

    it "Rem" do
      program = get_program "rem"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(0)
    end

    it "Eq" do
      program = get_program "eq"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(true)
    end

    it "Neq" do
      program = get_program "neq"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(true)
    end

    it "Lt" do
      program = get_program "lt"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(true)
    end

    it "Gt" do
      program = get_program "gt"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(false)
    end


    it "Lte" do
      program = get_program "lte"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(true)
    end

    it "Gte" do
      program = get_program "gte"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(false)
    end

    it "And" do
      program = get_program "and"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(false)
    end

    it "Or" do
      program = get_program "or"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(true)
    end
  end

  it "If" do
    program = get_program "if"
    result = Rinha::Interpreter.evaluate program.expression

    result.should eq(true)
  end

  it "Let" do
    program = get_program "let"
    result = Rinha::Interpreter.evaluate program.expression

    result.should eq(1)
  end

  describe "Function" do
    it "Not Recursive" do
      program = get_program "function"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(2)
    end

    it "Recursive" do
      program = get_program "fib"
      result = Rinha::Interpreter.evaluate program.expression

      result.should eq(55)
    end
  end
end
