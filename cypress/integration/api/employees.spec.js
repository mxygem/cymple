const allEmployees = [
    { "id": 1, "gender": "Male" },
    { "id": 2, "gender": "Female" },
    { "id": 3, "gender": "Nonbinary" },
    { "id": 4, "gender": "Female" }
]

describe('Employees API', () => {
    it('can returns all employees', () => {
        getURL("/employees").then(({ body }) => {
            expect(body).to.deep.equal(allEmployees)
        })
    })

    it('can return individual employees', () => {
        allEmployees.forEach(employee => {
            getURL("/employees/" + employee.id)
                .then(({ body }) => {
                    expect(body).to.deep.equal(employee)
                })
        })
    })
})

function getURL(url) {
    return cy.request({
        method: "GET",
        url: url,
        failOnStatusCode: false
    })
}
