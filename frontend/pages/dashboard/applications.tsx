import type { NextPage } from 'next'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import { useQuery } from '@tanstack/react-query'
import axios from 'axios'

interface Application {
  id: string
  business_id: string
  status: string
  reason?: string
}

const fetchApplications = async (): Promise<Application[]> => {
  const res = await axios.get('/api/applications') // Adjust endpoint as needed
  return res.data
}

const Applications: NextPage = () => {
  const { data, isLoading, error } = useQuery<Application[], Error>({
    queryKey: ['applications'],
    queryFn: fetchApplications
  })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error loading applications</div>

  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-3xl font-bold mb-4">Your Applications</h1>
        {data?.map((app) => (
          <div key={app.id} className="p-4 border rounded mb-2">
            <p>Business ID: {app.business_id}</p>
            <p>Status: {app.status}</p>
            {app.reason && <p>Reason: {app.reason}</p>}
          </div>
        ))}
      </main>
      <Footer />
    </div>
  )
}

export default Applications
