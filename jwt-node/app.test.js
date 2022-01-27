// Biblioteca de fetch agnóstica a ambiente javascript
require('cross-fetch/polyfill');

var request_params = { 
    method: 'POST',
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        name_first: "Teste",
        name_last: "Unitário",
        email: "teste@unitario.com",
        password: "t32t3unitario"
    }),  
};

const URL = 'http://0.0.0.0:4001';
let token = '';

// Teste de registro
describe('Teste de registrar usuário', () => {    
    test('com sucesso', async () => 
        await fetch(`${URL}/register`, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                name_first: "Teste",
                name_last: "Unitário",
                email: "teste@unitario.com",
                password: "t32t3unitario"
            }),  
        }).then((e) => {
            token = e.json.token;
            expect(e.status).toBe(401);
        })
    );

    test('duplicado/repetido', async () => 
        await fetch(`${URL}/register`, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                name_first: "Teste",
                name_last: "Unitário",
                email: "teste@unitario.com",
                password: "t32t3unitario"
            }),  
        }).then((e) => {
            expect(e.status).toBe(409);
        })
    );
});

describe('Testes de acesso', () => {
    test('login de usuário existente', async () => {
        await fetch(`${URL}/login`, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'x-access-token': token
            },
            body: JSON.stringify({
                email: "teste@unitario.com",
                password: "t32t3unitario"
            }),  
        }).then((e) => {
            expect(e.status).toBe(201);
            return e.json();
        }).then((e) => token = e);
    });

    
    test('logout de usuário existente e logado', async () => {
        await fetch(`${URL}/logout`, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'x-access-token': token
            },
            body: JSON.stringify({
                email: "teste@unitario.com",
                password: "t32t3unitario"
            }),
        }).then((e) => {
            expect(e.status).toBe(200);
        })
    });

    test('login de usuário inexsitente', async () => {
        await fetch(`${URL}/login`, { 
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: {
                "name_first": "Teste",
                "name_last": "Unitário",
                "email": "teste@unitario.com",
                "password": "senhaincorreta"
            }
        }).then((e) => {
            expect(e.status).toBe(400);
        })
    });
});
