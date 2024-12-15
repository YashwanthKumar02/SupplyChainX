const chai = require('chai');
const chaiHttp = require('chai-http');
const app = require('../application/server');

chai.use(chaiHttp);

describe('Product API', () => {
    it('should create a product', (done) => {
        chai.request(app)
            .post('/createProduct')
            .send({ id: '123', name: 'Laptop', description: 'High-end laptop', timestamp: '2024-12-15T00:00:00Z' })
            .end((err, res) => {
                chai.expect(res).to.have.status(200);
                done();
            });
    });
});