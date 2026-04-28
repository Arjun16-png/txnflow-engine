import { test,expect } from "@playwright/test";

test('create transaction success', async({ request }) => {
    const res = await request.post('http://localhost:8080/transactions', {
        data: {
            amount: 150000
        }
    })

    expect(res.status()).toBe(201)

    const body = await res.json()

    expect(body.status).toBe('PENDING')
})