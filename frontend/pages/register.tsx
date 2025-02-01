import { useState } from 'react'
import Header from '../components/Header'
import Footer from '../components/Footer'
import axios from 'axios'
import { useRouter } from 'next/router'

const Register = () => {
  const router = useRouter()
  const [form, setForm] = useState({ name: '', email: '', password: '', address: '' })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await axios.post('/api/register', form)
      router.push('/login')
    } catch (error) {
      console.error('Error registering user:', error)
    }
  }

  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-3xl font-bold mb-4">Register</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label>Name</label>
            <input name="name" type="text" value={form.name} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Email</label>
            <input name="email" type="email" value={form.email} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Password</label>
            <input name="password" type="password" value={form.password} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Address</label>
            <input name="address" type="text" value={form.address} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <button type="submit" className="bg-blue-600 text-white p-2 rounded">Register</button>
        </form>
      </main>
      <Footer />
    </div>
  )
}

export default Register
