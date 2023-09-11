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


describe("Binary", () => {
    it("Sum", () => {
        const ast = getAST('sum')

        const result = evaluate(ast.expression)

        expect(result).toBe(2)
    })

    it("Minus", () => {
        const ast = getAST('minus')

        const result = evaluate(ast.expression)

        expect(result).toBe(0)
    })

    it("Mul", () => {
        const ast = getAST("mul")

        const result = evaluate(ast.expression)

        expect(result).toBe(4)
    })

    it("Div", () => {
        const ast = getAST("div")

        const result = evaluate(ast.expression)

        expect(result).toBe(5)
    })

    it("Rem", () => {
        const ast = getAST("rem")

        const result = evaluate(ast.expression)

        expect(result).toEqual(0)
    })
})
