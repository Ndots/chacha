import type { NextPage } from 'next'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import { useQuery } from '@tanstack/react-query'
import axios from 'axios'

interface PartnerApplication {
  id: string
  business_id: string
  status: string
  reason?: string
}

const fetchPartnerApplications = async (): Promise<PartnerApplication[]> => {
  // This endpoint should return applications for partner review
  const res = await axios.get('/api/partner/applications')
  return res.data
}

const PartnerDashboard: NextPage = () => {
  const { data, isLoading, error } = useQuery<PartnerApplication[], Error>({
    queryKey: ['partnerApplications'],
    queryFn: fetchPartnerApplications
  })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error loading partner applications</div>

  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-3xl font-bold mb-4">Partner Dashboard</h1>
        {data?.map((app) => (
          <div key={app.id} className="p-4 border rounded mb-2">
            <p>Business ID: {app.business_id}</p>
            <p>Status: {app.status}</p>
            {app.reason && <p>Reason: {app.reason}</p>}
            {/* Here, you could add buttons to approve/reject each application */}
          </div>
        ))}
      </main>
      <Footer />
    </div>
  )
}

export default PartnerDashboard
