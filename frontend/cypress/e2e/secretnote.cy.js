describe('End-to-End Test for Secret Note Application', () => {

    it('should sign up, sign in, create a note, and verify the note', () => {
    
        const uuid = () => Cypress._.random(0, 1e6)
        const id = uuid()
        const username = `testname${id}`
    
        // Sign Up
        cy.visit('http://localhost:3000/signup')
        cy.get('input[type="text"]').type(username)
        cy.get('input[type="password"]').type('password')
        cy.get('button[type="submit"]').click()
        cy.url().should('include', '/signin')
    
        // Sign In
        cy.get('input[id="name"]').type(username)
        cy.get('input[id="password"]').type('password')
        cy.get('button[type="submit"]').click()
        cy.url().should('include', '/create')

        // Create Note
        cy.get('textarea[id="note_content"]').type('Hello world')
        cy.get('button[type="submit"]').click()
        cy.get('.note-created').should('be.visible')

        // View Note
        cy.get('.note-created a').invoke('text').then((noteUrl) => {
            cy.visit(`http://localhost:3000${noteUrl}`)
            cy.contains('Hello world').should('be.visible')
            cy.reload()
            cy.contains('this note has expired or reached its view limit.').should('be.visible')
        })
    })
})