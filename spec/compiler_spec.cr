require "./spec_helper"

describe Rinha::Compiler do

  it "Variable" do
    ast = get_program "let"
    compiler = Rinha::Compiler.new

    compiler.generate ast
  end
end

