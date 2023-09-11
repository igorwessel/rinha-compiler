import { it, expect, spyOn, describe, afterEach } from "bun:test";

import { File, evaluate } from './interpreter'
import { readFileSync } from "fs";

const getAST = (file: string) => {
    const json = readFileSync(`./examples/${file}.json`, { encoding: 'utf-8' })
    const ast: File = JSON.parse(json)

    return ast
}

const spy = spyOn(console, "log")

afterEach(() => {
    spy.mockClear()
})

describe("Print", () => {

    it("String", () => {
        const ast = getAST("string")

        evaluate(ast.expression)
        expect(spy).toHaveBeenCalled()
        //@ts-ignore
        expect(spy.mock.calls[0][0]).toEqual("rinha")
    })

    it("Bool", () => {
        const ast = getAST('bool')

        evaluate(ast.expression)

        expect(spy).toHaveBeenCalled()
        //@ts-ignore
        expect(spy.mock.calls[0][0]).toEqual(true)
    })

    it.todo("Tuple")
    it.todo("Closure")

    it("Number", () => {
        const ast = getAST('number')

        evaluate(ast.expression)

        expect(spy).toHaveBeenCalled()
        //@ts-ignore
        expect(spy.mock.calls[0][0]).toEqual(1)
    })
})
