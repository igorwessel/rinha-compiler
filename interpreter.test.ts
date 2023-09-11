import { it, expect, spyOn, describe } from "bun:test";

import {File, evaluate} from './interpreter'
import { readFileSync } from "fs";

const getAST = (file: string) => {
    const json = readFileSync(`./examples/${file}.json`, { encoding: 'utf-8'})
    const ast: File = JSON.parse(json)

    return ast
}

describe("Print", () => {
    it("Number", () => {
        const spy = spyOn(console, "log")
        const ast = getAST('print')

        evaluate(ast.expression)

        expect(spy).toHaveBeenCalled()
        //@ts-ignore
        expect(spy.mock.calls[0][0]).toEqual(1)
    })
})
