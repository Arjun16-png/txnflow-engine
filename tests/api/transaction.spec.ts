import { test,expect } from "@playwright/test"

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

test('same idempotency return same transaction',async({ request }) => {
    const payload = {
        amount: 167263,
        idempotency_key: `idem-${Date.now()}`
    }

    const firstRes = await request.post('http://localhost:8080/transactions', {
        data: payload
    })

    expect(firstRes.status()).toBe(201)

    const firstBody = await firstRes.json()
    const secondRes = await request.post('http://localhost:8080/transactions',{
        data: payload
    })

    expect(secondRes.ok()).toBeTruthy()
    const secondBody = await secondRes.json()

    expect(secondBody.id).toBe(firstBody.id)
    expect(secondBody.amount).toBe(firstBody.amount)
    expect(secondBody.idempotency_key).toBe(firstBody.idempotency_key)
}
)