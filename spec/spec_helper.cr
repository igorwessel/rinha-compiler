require "spec"
require "../src/compiler"
require "../src/interpreter"
require "../src/parser"

def get_program(name : String)
  File.open("./examples/#{name}.json") do |file|
    Rinha::Parser::Program.from_json(file)
  end
end

