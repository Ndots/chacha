import { useState } from 'react'
import Header from '../components/Header'
import Footer from '../components/Footer'
import axios from 'axios'
import { useRouter } from 'next/router'

const Login = () => {
  const router = useRouter()
  const [form, setForm] = useState({ email: '', password: '' })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      const res = await axios.post('/api/login', form)
      localStorage.setItem('token', res.data.token)
      router.push('/dashboard')
    } catch (error) {
      console.error('Error logging in:', error)
    }
  }

  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-3xl font-bold mb-4">Login</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label>Email</label>
            <input name="email" type="email" value={form.email} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Password</label>
            <input name="password" type="password" value={form.password} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <button type="submit" className="bg-blue-600 text-white p-2 rounded">Login</button>
        </form>
      </main>
      <Footer />
    </div>
  )
}

export default Login
