import { useState } from 'react'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import axios from 'axios'
import { useRouter } from 'next/router'

interface Director {
  name: string
  email: string
  position: string
}

const BusinessRegistration = () => {
  const router = useRouter()
  const [directors, setDirectors] = useState<Director[]>([{ name: '', email: '', position: '' }])
  const [formData, setFormData] = useState({
    name: '',
    address: '',
    type: '',
    email: '',
    phone: '',
  })

  const addDirector = () => {
    setDirectors([...directors, { name: '', email: '', position: '' }])
  }

  const updateDirector = (index: number, field: keyof Director, value: string) => {
    const newDirectors = [...directors]
    newDirectors[index][field] = value
    setDirectors(newDirectors)
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await axios.post('http://localhost:8080/api/business', {
        ...formData,
        directors,
      })
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
            <input
              name="name"
              type="text"
              value={formData.name}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              className="border p-2 w-full"
              required
            />
          </div>
          <div>
            <label>Address</label>
            <input
              name="address"
              type="text"
              value={formData.address}
              onChange={(e) => setFormData({ ...formData, address: e.target.value })}
              className="border p-2 w-full"
              required
            />
          </div>
          <div>
            <label>Type</label>
            <input
              name="type"
              type="text"
              value={formData.type}
              onChange={(e) => setFormData({ ...formData, type: e.target.value })}
              className="border p-2 w-full"
              required
            />
          </div>
          <div>
            <label>Email</label>
            <input
              name="email"
              type="email"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              className="border p-2 w-full"
              required
            />
          </div>
          <div>
            <label>Phone</label>
            <input
              name="phone"
              type="text"
              value={formData.phone}
              onChange={(e) => setFormData({ ...formData, phone: e.target.value })}
              className="border p-2 w-full"
              required
            />
          </div>
          <div className="space-y-4">
            <h3>Directors</h3>
            {directors.map((director, index) => (
              <div key={index} className="space-y-2">
                <input
                  type="text"
                  placeholder="Director Name"
                  value={director.name}
                  onChange={(e) => updateDirector(index, 'name', e.target.value)}
                  className="w-full p-2 border rounded"
                />
                <input
                  type="email"
                  placeholder="Director Email"
                  value={director.email}
                  onChange={(e) => updateDirector(index, 'email', e.target.value)}
                  className="w-full p-2 border rounded"
                />
                <input
                  type="text"
                  placeholder="Position"
                  value={director.position}
                  onChange={(e) => updateDirector(index, 'position', e.target.value)}
                  className="w-full p-2 border rounded"
                />
              </div>
            ))}
            <button
              type="button"
              onClick={addDirector}
              className="bg-gray-200 px-4 py-2 rounded"
            >
              Add Another Director
            </button>
          </div>
          <button type="submit" className="bg-blue-600 text-white p-2 rounded">Register Business</button>
        </form>
      </main>
      <Footer />
    </div>
  )
}

export default BusinessRegistration
