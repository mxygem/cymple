import allEmployeesImport from '../../fixtures/all-employees.json'

describe('homepage initial information', () => {
    beforeEach(() => {
        cy.visit('/')
    })

    it('displays the expected count of employees', () => {
        cy.get("[data-cy=employees]")
            .find('li').should('have.length', allEmployeesImport.length)
    })

    it('displays the expected details for each employee', () => {
        cy.fixture('all-employees').then((allEmployees) => {

            cy.get("[data-cy=employees]").within(() => {
                allEmployees.forEach((_, i) => {
                    cy.get('li').eq(i)
                        .contains("ID: " + allEmployees[i]["id"])
                        .contains("Gender: " + allEmployees[i]["gender"])
                })
            })

        })
    })
})
