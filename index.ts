import { readFile } from 'fs/promises'

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

type Term = Str | Int | Bool | Print

type File = {
    name: string
    expression: Term
}

function toInt32(x: number) {
    return x | 0
}

//@ts-ignore
function evaluate(program: Term) {
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
        case "Print": {
            return console.log(evaluate(program.value))
        }
    }
}


async function run() {
    const json = await readFile(`examples/${process.argv.at(2)}.json`, { encoding: 'utf-8'})
    const ast: File = JSON.parse(json)

    return evaluate(ast.expression)
}


run()