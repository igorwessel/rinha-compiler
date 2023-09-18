require "option_parser"
require "./parser"
require "./compiler"

outfile = "out"

OptionParser.parse do |parser|
  parser.banner = "Rinha-Compiler in Crystal!"

  parser.on "-v", "--version", "Show Version" do
    puts "v0.1.0"
    exit
  end

  parser.on "-h", "--help", "Show Help" do
    puts parser
    exit
  end

  parser.on "-r FILE", "--run", "Compile a Rinha JSON" do |file|
    rinha_parser = Rinha::Parser.parse(file)
    compiler = Rinha::Compiler.new

    code = compiler.generate rinha_parser

    File.open("#{outfile}.asm", "w") do |file|
      file.print code.to_s
    end

    `nasm -f elf64 #{outfile}.asm && ld -s -o hello #{outfile}.o`
    exit
  end

  parser.missing_option do |option_flag|
    STDERR.puts "ERROR: #{option_flag} is missing something."
    STDERR.puts ""
    STDERR.puts parser
    exit(1)
  end

  parser.invalid_option do |option_flag|
    STDERR.puts "ERROR: #{option_flag} is not a valid option."
    STDERR.puts parser
    exit(1)
  end
end
