type BinaryOperator =
    "Add" |
    "Sub" |
    "Mul" |
    "Div" |
    "Rem" |
    "Eq" |
    "Neq" |
    "Lt" |
    "Gt" |
    "Lte" |
    "Gte" |
    "And" |
    "Or"

type Str = {
    kind: "Str"
    value: string
}

type Int = {
    kind: "Int"
    value: number
}

type Print = {
    kind: "Print"
    value: Term
}

type Bool = {
    kind: "Bool"
    value: boolean
}

type Binary = {
    kind: "Binary",
    lhs: Term
    op: BinaryOperator
    rhs: Term
}

type Term = Str | Int | Bool | Binary | Print

export type File = {
    name: string
    expression: Term
}

function toInt32(x: number) {
    return x | 0
}

//@ts-ignore
export function evaluate(program: Term) {
    switch (program.kind) {
        case "Int": {
            return toInt32(program.value)
        }
        case "Bool": {
            return program.value
        }
        case "Str": {
            return program.value
        }
        case "Binary": {
            switch (program.op) {
                case "Add":
                    return evaluate(program.lhs) + evaluate(program.rhs)
                case "Sub":
                    return evaluate(program.lhs) - evaluate(program.rhs)
                case "Mul":
                    return evaluate(program.lhs) * evaluate(program.rhs)
                case "Div":
                    return evaluate(program.lhs) / evaluate(program.rhs)
                case "Rem":
                    return evaluate(program.lhs) % evaluate(program.rhs)
                case "Eq":
                    return evaluate(program.lhs) == evaluate(program.rhs)
                case "Neq":
                    return evaluate(program.lhs) != evaluate(program.rhs)
                case "Lt":
                    return evaluate(program.lhs) < evaluate(program.rhs)
                case "Gt":
                    return evaluate(program.lhs) > evaluate(program.rhs)
                case "Lte":
                    return evaluate(program.lhs) <= evaluate(program.rhs)
                case "Gte":
                    return evaluate(program.lhs) >= evaluate(program.rhs)
                case "And":
                    return evaluate(program.lhs) && evaluate(program.rhs)
                case "Or":
                    return evaluate(program.lhs) || evaluate(program.rhs)
            }
        }
        case "Print": {
            //@ts-ignore
            return console.log(evaluate(program.value))
        }
    }
}