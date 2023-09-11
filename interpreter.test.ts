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

        expect(result).toBe(0)
    })

    it("Eq", () => {
        const ast = getAST("eq")

        const result = evaluate(ast.expression)

        expect(result).toBe(true)
    })

    it("Neq", () => {
        const ast = getAST("neq")

        const result = evaluate(ast.expression)

        expect(result).toBe(true)
    })

    it("Lt", () => {
        const ast = getAST("lt")

        const result = evaluate(ast.expression)

        expect(result).toBe(true)
    })
    it("Gt", () => {
        const ast = getAST('gt')

        const result = evaluate(ast.expression)

        expect(result).toBe(false)
    })

    it("Lte", () => {
        const ast = getAST('lte')

        const result = evaluate(ast.expression)

        expect(result).toBe(true)
    })

    it("Gte", () => {
        const ast = getAST('gte')

        const result = evaluate(ast.expression)

        expect(result).toBe(false)
    })

    it("And", () => {
        const ast = getAST('and')

        const result = evaluate(ast.expression)

        expect(result).toBe(false)
    })

    it("Or", () => {
        const ast = getAST('or')

        const result = evaluate(ast.expression)

        expect(result).toBe(true)
    })
})


it("If-Else", () => {
    const ast = getAST("if")

    const result = evaluate(ast.expression)

    expect(result).toBe("a")
})