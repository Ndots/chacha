import type { NextPage } from 'next'
import Header from '../../components/Header'
import Footer from '../../components/Footer'
import DashboardCard from '../../components/DashboardCard'
import { useQuery } from '@tanstack/react-query'
import axios from 'axios'

interface DashboardStats {
  total_users: number
  total_businesses: number
  pending_applications: number
  approved_applications: number
  rejected_applications: number
}

const fetchDashboardStats = async (): Promise<DashboardStats> => {
  const res = await axios.get('/api/dashboard')
  return res.data
}

const AdminDashboard: NextPage = () => {
  const { data, isLoading, error } = useQuery<DashboardStats, Error>({
    queryKey: ['dashboardStats'],
    queryFn: fetchDashboardStats
  })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error loading dashboard stats</div>

  return (
    <div>
      <Header />
      <main className="container mx-auto p-6">
        <h1 className="text-3xl font-bold mb-4">Admin Dashboard</h1>
        <div className="grid grid-cols-2 gap-4">
          <DashboardCard title="Total Users" value={data?.total_users} />
          <DashboardCard title="Total Businesses" value={data?.total_businesses} />
          <DashboardCard title="Pending Applications" value={data?.pending_applications} />
          <DashboardCard title="Approved Applications" value={data?.approved_applications} />
          <DashboardCard title="Rejected Applications" value={data?.rejected_applications} />
        </div>
      </main>
      <Footer />
    </div>
  )
}

export default AdminDashboard
