import { useState } from 'react'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import axios from 'axios'
import { useRouter } from 'next/router'

const BusinessRegistration = () => {
  const router = useRouter()
  const [form, setForm] = useState({
    name: '',
    address: '',
    type: '',
    proposedNames: ['', '', ''],
    userId: 1 // Replace with actual user ID from auth context
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleProposedNameChange = (index: number, value: string) => {
    const updatedNames = [...form.proposedNames]
    updatedNames[index] = value
    setForm({ ...form, proposedNames: updatedNames })
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await axios.post('/api/business', form)
      router.push('/dashboard')
    } catch (error) {
      console.error('Error registering business:', error)
    }
  }

  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-3xl font-bold mb-4">Register Business</h1>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label>Business Name</label>
            <input name="name" type="text" value={form.name} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Address</label>
            <input name="address" type="text" value={form.address} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Type</label>
            <input name="type" type="text" value={form.type} onChange={handleChange} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Proposed Name 1</label>
            <input type="text" value={form.proposedNames[0]} onChange={(e) => handleProposedNameChange(0, e.target.value)} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Proposed Name 2</label>
            <input type="text" value={form.proposedNames[1]} onChange={(e) => handleProposedNameChange(1, e.target.value)} className="border p-2 w-full" required />
          </div>
          <div>
            <label>Proposed Name 3</label>
            <input type="text" value={form.proposedNames[2]} onChange={(e) => handleProposedNameChange(2, e.target.value)} className="border p-2 w-full" required />
          </div>
          <button type="submit" className="bg-blue-600 text-white p-2 rounded">Register Business</button>
        </form>
      </main>
      <Footer />
    </div>
  )
}

export default BusinessRegistration
