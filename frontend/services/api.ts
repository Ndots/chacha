import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080', // Adjust if needed
})
//import api from '../services/api'

// Example POST request to register a user
api.post('/api/register', { 
    email: 'test@example.com',
    password: 'password',
    firstName: 'Test',
    lastName: 'User'
})
.then(response => console.log(response.data))
.catch(error => console.error(error))



export default api
