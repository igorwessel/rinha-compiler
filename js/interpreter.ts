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

type Var = {
    kind: 'Var',
    text: string
}

type If = {
    kind: "If"
    condition: Term,
    then: Term,
    otherwise: Term
}

type Parameter = {
    kind: "Parameter",
    text: string
}

type Function = {
    kind: "Function",
    parameters: Parameter[]
    value: Term
}

type Call = {
    kind: "Call",
    callee: Term,
    arguments: Term[]
}

type Let = {
    kind: "Let",
    name: Parameter,
    value: Term,
    next: Term
}

type Term = Str | Int | Bool | Binary | If | Var | Function | Call | Let | Print

export type File = {
    name: string
    expression: Term
}

function toInt32(x: number) {
    return x | 0
}

const Environment = new Map()

//@ts-ignore
export function evaluate(program: Term, environment = Environment) {
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
        case "Var": {
            let value = environment.get(program.text)

            if (value.kind == "Int" || value.kind == "Bool" || value.kind == "Str") {
                value = evaluate(value, environment)
            }

            return value
        }
        case "Binary": {
            switch (program.op) {
                case "Add":
                    return evaluate(program.lhs, environment) + evaluate(program.rhs, environment)
                case "Sub":
                    return evaluate(program.lhs, environment) - evaluate(program.rhs, environment)
                case "Mul":
                    return evaluate(program.lhs, environment) * evaluate(program.rhs, environment)
                case "Div":
                    return evaluate(program.lhs, environment) / evaluate(program.rhs, environment)
                case "Rem":
                    return evaluate(program.lhs, environment) % evaluate(program.rhs, environment)
                case "Eq":
                    return evaluate(program.lhs, environment) == evaluate(program.rhs, environment)
                case "Neq":
                    return evaluate(program.lhs, environment) != evaluate(program.rhs, environment)
                case "Lt":
                    return evaluate(program.lhs, environment) < evaluate(program.rhs, environment)
                case "Gt":
                    return evaluate(program.lhs, environment) > evaluate(program.rhs, environment)
                case "Lte":
                    return evaluate(program.lhs, environment) <= evaluate(program.rhs, environment)
                case "Gte":
                    return evaluate(program.lhs, environment) >= evaluate(program.rhs, environment)
                case "And":
                    return evaluate(program.lhs, environment) && evaluate(program.rhs, environment)
                case "Or":
                    return evaluate(program.lhs, environment) || evaluate(program.rhs, environment)
            }
        }
        case "If": {
            return evaluate(program.condition, environment) ? evaluate(program.then, environment) : evaluate(program.otherwise, environment)
        }
        case "Function": {
            return evaluate(program.value, environment)
        }
        case "Let": {
            environment.set(program.name.text, program.value)

            return evaluate(program.next, environment)
        }
        case "Call": {
            const body: Function = evaluate(program.callee, environment)

            if (body?.parameters?.length != program?.arguments?.length) {
                throw new Error('wtf')
            }

            for (let i = 0; i < body.parameters.length; i++) {
                environment.set(body.parameters[i].text, evaluate(program.arguments[i], environment))
            }

            return evaluate(body, environment)
        }
        case "Print": {
            const node: Term = evaluate(program.value, environment)

            console.log(node)

            return node
        }
    }

}