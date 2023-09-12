import { readFile } from 'fs/promises'
import { File, evaluate } from './interpreter'

async function run() {
    const json = await readFile(`examples/${process.argv.at(2)}.json`, { encoding: 'utf-8' })
    const ast: File = JSON.parse(json)

    return evaluate(ast.expression)
}

run()