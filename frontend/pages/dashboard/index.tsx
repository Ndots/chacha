import type { NextPage } from 'next'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import Link from 'next/link'
import { useQuery } from '@tanstack/react-query'
import axios from 'axios'

interface Business {
  id: string
  name: string
  status: string
}

const fetchUserBusinesses = async (): Promise<Business[]> => {
  const res = await axios.get('http://localhost:8080/api/dashboard')
  return res.data
}

const Dashboard: NextPage = () => {
  const { data, isLoading, error } = useQuery<Business[], Error>({
    queryKey: ['businesses'],
    queryFn: fetchUserBusinesses
  })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error loading businesses: {error.message}</div>
  if (!data) return <div>No businesses found</div>

  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-3xl font-bold mb-4">User Dashboard</h1>
        <Link 
          href="/dashboard/business" 
          className="bg-blue-600 text-white p-2 rounded"
        >
          Register a New Business
        </Link>
        <div className="mt-6">
          <h2 className="text-2xl mb-2">Your Business Applications</h2>
          {data.length > 0 ? (
            data.map((business) => (
              <div key={business.id} className="p-4 border rounded mb-2">
                <p className="font-bold">{business.name}</p>
                <p>Status: {business.status}</p>
              </div>
            ))
          ) : (
            <p>No business applications found</p>
          )}
        </div>
      </main>
      <Footer />
    </div>
  )
}

export default Dashboard
