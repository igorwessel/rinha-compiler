require "json"

module Rinha::Parser
  extend self

  alias Value = Bool | Int32 | String | Float64 | Nil

  class Location
    include JSON::Serializable

    property start : Int32
    property end : Int32
    property filename : String
  end

  class Parameter
    include JSON::Serializable

    property text : String
    property location : Location
  end

  abstract class Term
    include JSON::Serializable

    use_json_discriminator "kind", {Str: Str, Int: Int, Print: Print, Bool: Boolean, Binary: Binary, If: If, Let: Let, Var: Var, Function: Function, Call: Call}

    property kind : String
    property location : Location

  end

  class Str < Term
    property value : String

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Int < Term
    property value : Int32

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Boolean < Term
    property value : Bool

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Binary < Term
    property lhs : Term
    property op : String
    property rhs : Term

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class If < Term
    property condition : Term
    property then : Term
    property otherwise : Term

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Let < Term
    property name : Parameter
    property value : Term
    property next : Term

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Var < Term
    property text : String

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Print < Term
    property value : Term

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Function < Term
    property parameters : Array(Parameter)
    property value : Term

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Call < Term
    property callee : Term
    property arguments : Array(Term)

    def accept(visitor)
      visitor.visit(self)
    end
  end

  class Program
    include JSON::Serializable

    property expression : Term
    property location : Location

    def accept(visitor)
      visitor.visit(self)
    end
  end


  def parse(file : String)
    File.open("#{file}.json") do |file|
      Program.from_json file
    end
  end
end
